// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package api

import (
	"net/http"
	"net/http/httputil"
	"strings"
)

const (
	contentType         = "Content-Type"
	textHTMLContentType = "text/html"
	jsonContentType     = "application/json"
)

// DefaultTransport can be used by clients which rely on the apierror.Unwrap
// Can obtain the underlying http.Response is returned with a StatusCode not
// defined within the swagger spec from which the models have been generated.
// Meaning this is a small hack which allows http.Response.Body to be accessed.
// See pkg/api/apierror for details on how apierror.Unwrap works. Note that
// using this variable directly won't allow any of the http.Transport settings
// to be overridden. To customize the transport further, please use
// NewTransport()
var DefaultTransport = new(ErrCatchTransport)

// NewErrCatchTransport Deprecated: initializes an ErrCatchTransport.
// See GoDoc for more help on this type.
func NewErrCatchTransport(rt http.RoundTripper) *ErrCatchTransport {
	return &ErrCatchTransport{rt: rt}
}

// ErrCatchTransport Deprecated: is an http.RoundTripper that which allows the
// http.Response to be accessed in certain types of wrapped errors returned by
// autogenerated code.
// See pkg/api/apierror for details on how apierror.Unwrap works.
type ErrCatchTransport struct {
	rt http.RoundTripper
}

// RoundTrip Deprecated: wraps http.DefaultTransport.RoundTrip to keep track
// of the current request.
func (e *ErrCatchTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if e.rt == nil {
		newDefaultTransport(0)
	}

	res, err := e.rt.RoundTrip(req)
	if res != nil {
		_, _ = httputil.DumpResponse(res, res.Body != nil)

		// When the content type is "text/html", a bit of tweaking is required
		// for the response to be marshaled to JSON. Using the standard error
		// definition and populating it with parts of the request so the error
		// can be identified.
		if strings.Contains(res.Header.Get(contentType), textHTMLContentType) {
			res.Header.Set(contentType, jsonContentType)
			res.Body = newProxyBody(req, res.StatusCode)
		}
	}

	return res, err
}
