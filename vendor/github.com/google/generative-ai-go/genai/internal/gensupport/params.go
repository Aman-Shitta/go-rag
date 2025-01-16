// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gensupport

import (
	"net/http"
	"net/url"

	"google.golang.org/api/googleapi"
)

// URLParams is a simplified replacement for url.Values
// that safely builds up URL parameters for encoding.
type URLParams map[string][]string

// Get returns the first value for the given key, or "".
func (u URLParams) Get(key string) string {
	vs := u[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

// Set sets the key to value.
// It replaces any existing values.
func (u URLParams) Set(key, value string) {
	u[key] = []string{value}
}

// SetMulti sets the key to an array of values.
// It replaces any existing values.
// Note that values must not be modified after calling SetMulti
// so the caller is responsible for making a copy if necessary.
func (u URLParams) SetMulti(key string, values []string) {
	u[key] = values
}

// Encode encodes the values into “URL encoded” form
// ("bar=baz&foo=quux") sorted by key.
func (u URLParams) Encode() string {
	return url.Values(u).Encode()
}

// SetOptions sets the URL params and any additional `CallOption` or
// `MultiCallOption` passed in.
func SetOptions(u URLParams, opts ...googleapi.CallOption) {
	for _, o := range opts {
		m, ok := o.(googleapi.MultiCallOption)
		if ok {
			u.SetMulti(m.GetMulti())
			continue
		}
		u.Set(o.Get())
	}
}

// SetHeaders sets common headers for all requests. The keyvals header pairs
// should have a corresponding value for every key provided. If there is an odd
// number of keyvals this method will panic.
func SetHeaders(userAgent, contentType string, userHeaders http.Header, keyvals ...string) http.Header {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+GoVersion()+" gdcl/"+"0.179.0")
	for i := 0; i < len(keyvals); i = i + 2 {
		reqHeaders.Set(keyvals[i], keyvals[i+1])
	}
	reqHeaders.Set("User-Agent", userAgent)
	if contentType != "" {
		reqHeaders.Set("Content-Type", contentType)
	}
	for k, v := range userHeaders {
		reqHeaders[k] = v
	}
	return reqHeaders
}
