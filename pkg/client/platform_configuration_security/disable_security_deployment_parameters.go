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

// Code generated by go-swagger; DO NOT EDIT.

package platform_configuration_security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDisableSecurityDeploymentParams creates a new DisableSecurityDeploymentParams object
// with the default values initialized.
func NewDisableSecurityDeploymentParams() *DisableSecurityDeploymentParams {
	var ()
	return &DisableSecurityDeploymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDisableSecurityDeploymentParamsWithTimeout creates a new DisableSecurityDeploymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDisableSecurityDeploymentParamsWithTimeout(timeout time.Duration) *DisableSecurityDeploymentParams {
	var ()
	return &DisableSecurityDeploymentParams{

		timeout: timeout,
	}
}

// NewDisableSecurityDeploymentParamsWithContext creates a new DisableSecurityDeploymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewDisableSecurityDeploymentParamsWithContext(ctx context.Context) *DisableSecurityDeploymentParams {
	var ()
	return &DisableSecurityDeploymentParams{

		Context: ctx,
	}
}

// NewDisableSecurityDeploymentParamsWithHTTPClient creates a new DisableSecurityDeploymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDisableSecurityDeploymentParamsWithHTTPClient(client *http.Client) *DisableSecurityDeploymentParams {
	var ()
	return &DisableSecurityDeploymentParams{
		HTTPClient: client,
	}
}

/*DisableSecurityDeploymentParams contains all the parameters to send to the API endpoint
for the disable security deployment operation typically these are written to a http.Request
*/
type DisableSecurityDeploymentParams struct {

	/*Version
	  When specified, checks for conflicts against the version stored in the persistent store (returned in 'x-cloud-resource-version' of the GET request)

	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the disable security deployment params
func (o *DisableSecurityDeploymentParams) WithTimeout(timeout time.Duration) *DisableSecurityDeploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disable security deployment params
func (o *DisableSecurityDeploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disable security deployment params
func (o *DisableSecurityDeploymentParams) WithContext(ctx context.Context) *DisableSecurityDeploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disable security deployment params
func (o *DisableSecurityDeploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disable security deployment params
func (o *DisableSecurityDeploymentParams) WithHTTPClient(client *http.Client) *DisableSecurityDeploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disable security deployment params
func (o *DisableSecurityDeploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVersion adds the version to the disable security deployment params
func (o *DisableSecurityDeploymentParams) WithVersion(version *string) *DisableSecurityDeploymentParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the disable security deployment params
func (o *DisableSecurityDeploymentParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DisableSecurityDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Version != nil {

		// query param version
		var qrVersion string
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
