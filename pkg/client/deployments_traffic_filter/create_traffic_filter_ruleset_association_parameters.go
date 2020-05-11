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

package deployments_traffic_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewCreateTrafficFilterRulesetAssociationParams creates a new CreateTrafficFilterRulesetAssociationParams object
// with the default values initialized.
func NewCreateTrafficFilterRulesetAssociationParams() *CreateTrafficFilterRulesetAssociationParams {
	var ()
	return &CreateTrafficFilterRulesetAssociationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTrafficFilterRulesetAssociationParamsWithTimeout creates a new CreateTrafficFilterRulesetAssociationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTrafficFilterRulesetAssociationParamsWithTimeout(timeout time.Duration) *CreateTrafficFilterRulesetAssociationParams {
	var ()
	return &CreateTrafficFilterRulesetAssociationParams{

		timeout: timeout,
	}
}

// NewCreateTrafficFilterRulesetAssociationParamsWithContext creates a new CreateTrafficFilterRulesetAssociationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTrafficFilterRulesetAssociationParamsWithContext(ctx context.Context) *CreateTrafficFilterRulesetAssociationParams {
	var ()
	return &CreateTrafficFilterRulesetAssociationParams{

		Context: ctx,
	}
}

// NewCreateTrafficFilterRulesetAssociationParamsWithHTTPClient creates a new CreateTrafficFilterRulesetAssociationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTrafficFilterRulesetAssociationParamsWithHTTPClient(client *http.Client) *CreateTrafficFilterRulesetAssociationParams {
	var ()
	return &CreateTrafficFilterRulesetAssociationParams{
		HTTPClient: client,
	}
}

/*CreateTrafficFilterRulesetAssociationParams contains all the parameters to send to the API endpoint
for the create traffic filter ruleset association operation typically these are written to a http.Request
*/
type CreateTrafficFilterRulesetAssociationParams struct {

	/*Body
	  Mandatory ruleset association description

	*/
	Body *models.FilterAssociation
	/*RulesetID
	  The mandatory ruleset ID.

	*/
	RulesetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create traffic filter ruleset association params
func (o *CreateTrafficFilterRulesetAssociationParams) WithTimeout(timeout time.Duration) *CreateTrafficFilterRulesetAssociationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create traffic filter ruleset association params
func (o *CreateTrafficFilterRulesetAssociationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create traffic filter ruleset association params
func (o *CreateTrafficFilterRulesetAssociationParams) WithContext(ctx context.Context) *CreateTrafficFilterRulesetAssociationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create traffic filter ruleset association params
func (o *CreateTrafficFilterRulesetAssociationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create traffic filter ruleset association params
func (o *CreateTrafficFilterRulesetAssociationParams) WithHTTPClient(client *http.Client) *CreateTrafficFilterRulesetAssociationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create traffic filter ruleset association params
func (o *CreateTrafficFilterRulesetAssociationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create traffic filter ruleset association params
func (o *CreateTrafficFilterRulesetAssociationParams) WithBody(body *models.FilterAssociation) *CreateTrafficFilterRulesetAssociationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create traffic filter ruleset association params
func (o *CreateTrafficFilterRulesetAssociationParams) SetBody(body *models.FilterAssociation) {
	o.Body = body
}

// WithRulesetID adds the rulesetID to the create traffic filter ruleset association params
func (o *CreateTrafficFilterRulesetAssociationParams) WithRulesetID(rulesetID string) *CreateTrafficFilterRulesetAssociationParams {
	o.SetRulesetID(rulesetID)
	return o
}

// SetRulesetID adds the rulesetId to the create traffic filter ruleset association params
func (o *CreateTrafficFilterRulesetAssociationParams) SetRulesetID(rulesetID string) {
	o.RulesetID = rulesetID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTrafficFilterRulesetAssociationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param ruleset_id
	if err := r.SetPathParam("ruleset_id", o.RulesetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}