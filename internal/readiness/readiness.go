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

// Package readiness is a minimal probe aggregator exposing /readyz and /healthz
// HTTP handlers. It replaces the platform-clients readiness package so the open
// (Apache-2.0) execution layer carries no ELv2 dependency (issue #98).
package readiness

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
)

// Probe is implemented by any dependency that can report its own health.
// Name must be stable and unique within an Aggregator; Check returns nil when
// ready, or a descriptive error when not.
type Probe interface {
	Name() string
	Check(ctx context.Context) error
}

type probeResult struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type readyzResponse struct {
	Status string        `json:"status"`
	Probes []probeResult `json:"probes"`
}

// Aggregator collects Probes and exposes /readyz and /healthz handlers.
// Construct with NewAggregator.
type Aggregator struct {
	mu     sync.RWMutex
	probes []Probe
}

// NewAggregator returns an initialised, empty Aggregator.
func NewAggregator() *Aggregator { return &Aggregator{} }

// Register adds p to the set evaluated on every /readyz request.
func (a *Aggregator) Register(p Probe) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.probes = append(a.probes, p)
}

// ReadyHandler evaluates all registered probes; 200 when all pass, 503 when any
// fails, with a JSON body listing per-probe status.
func (a *Aggregator) ReadyHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.mu.RLock()
		probes := append([]Probe(nil), a.probes...)
		a.mu.RUnlock()

		resp := readyzResponse{Status: "ok"}
		ready := true
		for _, p := range probes {
			pr := probeResult{Name: p.Name(), Status: "ok"}
			if err := p.Check(r.Context()); err != nil {
				pr.Status = "unready"
				pr.Error = err.Error()
				ready = false
			}
			resp.Probes = append(resp.Probes, pr)
		}

		w.Header().Set("Content-Type", "application/json")
		if !ready {
			resp.Status = "unready"
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		_ = json.NewEncoder(w).Encode(resp)
	})
}

// LivenessHandler always responds 200 OK.
func (a *Aggregator) LivenessHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
}
