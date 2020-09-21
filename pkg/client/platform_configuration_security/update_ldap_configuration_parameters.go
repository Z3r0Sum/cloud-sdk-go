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

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewUpdateLdapConfigurationParams creates a new UpdateLdapConfigurationParams object
// with the default values initialized.
func NewUpdateLdapConfigurationParams() *UpdateLdapConfigurationParams {
	var ()
	return &UpdateLdapConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLdapConfigurationParamsWithTimeout creates a new UpdateLdapConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateLdapConfigurationParamsWithTimeout(timeout time.Duration) *UpdateLdapConfigurationParams {
	var ()
	return &UpdateLdapConfigurationParams{

		timeout: timeout,
	}
}

// NewUpdateLdapConfigurationParamsWithContext creates a new UpdateLdapConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateLdapConfigurationParamsWithContext(ctx context.Context) *UpdateLdapConfigurationParams {
	var ()
	return &UpdateLdapConfigurationParams{

		Context: ctx,
	}
}

// NewUpdateLdapConfigurationParamsWithHTTPClient creates a new UpdateLdapConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateLdapConfigurationParamsWithHTTPClient(client *http.Client) *UpdateLdapConfigurationParams {
	var ()
	return &UpdateLdapConfigurationParams{
		HTTPClient: client,
	}
}

/*UpdateLdapConfigurationParams contains all the parameters to send to the API endpoint
for the update ldap configuration operation typically these are written to a http.Request
*/
type UpdateLdapConfigurationParams struct {

	/*Body
	  The LDAP configuration

	*/
	Body *models.LdapSettings
	/*RealmID
	  The Elasticsearch Security realm identifier.

	*/
	RealmID string
	/*Version
	  When specified, checks for conflicts against the version stored in the persistent store (returned in 'x-cloud-resource-version' of the GET request)

	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update ldap configuration params
func (o *UpdateLdapConfigurationParams) WithTimeout(timeout time.Duration) *UpdateLdapConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update ldap configuration params
func (o *UpdateLdapConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update ldap configuration params
func (o *UpdateLdapConfigurationParams) WithContext(ctx context.Context) *UpdateLdapConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update ldap configuration params
func (o *UpdateLdapConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update ldap configuration params
func (o *UpdateLdapConfigurationParams) WithHTTPClient(client *http.Client) *UpdateLdapConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update ldap configuration params
func (o *UpdateLdapConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update ldap configuration params
func (o *UpdateLdapConfigurationParams) WithBody(body *models.LdapSettings) *UpdateLdapConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update ldap configuration params
func (o *UpdateLdapConfigurationParams) SetBody(body *models.LdapSettings) {
	o.Body = body
}

// WithRealmID adds the realmID to the update ldap configuration params
func (o *UpdateLdapConfigurationParams) WithRealmID(realmID string) *UpdateLdapConfigurationParams {
	o.SetRealmID(realmID)
	return o
}

// SetRealmID adds the realmId to the update ldap configuration params
func (o *UpdateLdapConfigurationParams) SetRealmID(realmID string) {
	o.RealmID = realmID
}

// WithVersion adds the version to the update ldap configuration params
func (o *UpdateLdapConfigurationParams) WithVersion(version *string) *UpdateLdapConfigurationParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the update ldap configuration params
func (o *UpdateLdapConfigurationParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLdapConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param realm_id
	if err := r.SetPathParam("realm_id", o.RealmID); err != nil {
		return err
	}

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
