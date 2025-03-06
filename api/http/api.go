// Copyright (c) 2025 Cloudflare, Inc.
// Licensed under the Apache 2.0 license found in the LICENSE file or at:
//     https://opensource.org/licenses/Apache-2.0

package http

import (
	"net/http"

	"github.com/prometheus/common/route"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/storage"
)

type apiConfig struct {
	queryAPIOpts []QueryAPIOption
}

type APIOption func(*apiConfig)

func QueryOptions(opts ...QueryAPIOption) APIOption {
	return func(cfg *apiConfig) {
		cfg.queryAPIOpts = opts
	}
}

func NewAPI(queryable storage.Queryable, engine promql.QueryEngine, opts ...APIOption) http.Handler {
	cfg := &apiConfig{}
	for i := range opts {
		opts[i](cfg)
	}

	r := route.New()

	api := r.WithPrefix("/api/v1")
	RegisterQueryV1(api, queryable, engine, cfg.queryAPIOpts...)

	return r
}
