// Copyright 2026 zero-day.ai
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package observability provides minimal OpenTelemetry + slog initialisation
// for the runner. It replaces the platform-clients observability package so the
// open (Apache-2.0) execution layer carries no ELv2 dependency (issue #98,
// ADR-0054).
//
// Behaviour matches the prior helper: a composite W3C TraceContext + Baggage
// propagator (so GIBSON_TRACE_ID/SPAN_ID propagate), an OTLP/gRPC trace exporter
// when OTEL_EXPORTER_OTLP_ENDPOINT is set (a plain provider otherwise — never a
// hard failure), and a JSON slog logger pre-loaded with service_name.
package observability

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// Observability holds the initialised OTel + slog stack for the process.
type Observability struct {
	// Logger is a JSON slog.Logger pre-loaded with the service_name attribute.
	Logger *slog.Logger

	shutdown func(context.Context) error
}

// Init wires the global tracer provider + propagator and returns the stack.
// serviceName must be non-empty. The OTLP endpoint is read from
// OTEL_EXPORTER_OTLP_ENDPOINT; when unset, a no-export provider is used so the
// global tracer is still valid and propagation still works.
func Init(ctx context.Context, serviceName string) (*Observability, error) {
	if serviceName == "" {
		return nil, fmt.Errorf("observability.Init: serviceName must not be empty")
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})).With("service_name", serviceName)

	// Composite W3C TraceContext + Baggage propagator.
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	// resource.New can return a partially-constructed resource with a non-fatal
	// error (e.g. schema-URL conflicts); the resource is still usable.
	res, _ := resource.New(ctx, resource.WithAttributes(
		attribute.String("service.name", serviceName),
	))

	var tp *sdktrace.TracerProvider
	if endpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"); endpoint != "" {
		exp, err := otlptracegrpc.New(ctx)
		if err != nil {
			return nil, fmt.Errorf("observability.Init: trace exporter: %w", err)
		}
		tp = sdktrace.NewTracerProvider(
			sdktrace.WithBatcher(exp),
			sdktrace.WithResource(res),
		)
	} else {
		tp = sdktrace.NewTracerProvider(sdktrace.WithResource(res))
	}
	otel.SetTracerProvider(tp)

	return &Observability{Logger: logger, shutdown: tp.Shutdown}, nil
}

// Shutdown flushes and stops the tracer provider.
func (o *Observability) Shutdown(ctx context.Context) error {
	if o == nil || o.shutdown == nil {
		return nil
	}
	return o.shutdown(ctx)
}
