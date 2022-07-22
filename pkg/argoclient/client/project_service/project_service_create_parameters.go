// Code generated by go-swagger; DO NOT EDIT.

package project_service

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

	"github.com/epam/edp-argocd-operator/pkg/argoclient/models"
)

// NewProjectServiceCreateParams creates a new ProjectServiceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectServiceCreateParams() *ProjectServiceCreateParams {
	return &ProjectServiceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectServiceCreateParamsWithTimeout creates a new ProjectServiceCreateParams object
// with the ability to set a timeout on a request.
func NewProjectServiceCreateParamsWithTimeout(timeout time.Duration) *ProjectServiceCreateParams {
	return &ProjectServiceCreateParams{
		timeout: timeout,
	}
}

// NewProjectServiceCreateParamsWithContext creates a new ProjectServiceCreateParams object
// with the ability to set a context for a request.
func NewProjectServiceCreateParamsWithContext(ctx context.Context) *ProjectServiceCreateParams {
	return &ProjectServiceCreateParams{
		Context: ctx,
	}
}

// NewProjectServiceCreateParamsWithHTTPClient creates a new ProjectServiceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectServiceCreateParamsWithHTTPClient(client *http.Client) *ProjectServiceCreateParams {
	return &ProjectServiceCreateParams{
		HTTPClient: client,
	}
}

/* ProjectServiceCreateParams contains all the parameters to send to the API endpoint
   for the project service create operation.

   Typically these are written to a http.Request.
*/
type ProjectServiceCreateParams struct {

	// Body.
	Body *models.ProjectProjectCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the project service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectServiceCreateParams) WithDefaults() *ProjectServiceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the project service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectServiceCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the project service create params
func (o *ProjectServiceCreateParams) WithTimeout(timeout time.Duration) *ProjectServiceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project service create params
func (o *ProjectServiceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project service create params
func (o *ProjectServiceCreateParams) WithContext(ctx context.Context) *ProjectServiceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project service create params
func (o *ProjectServiceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project service create params
func (o *ProjectServiceCreateParams) WithHTTPClient(client *http.Client) *ProjectServiceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project service create params
func (o *ProjectServiceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the project service create params
func (o *ProjectServiceCreateParams) WithBody(body *models.ProjectProjectCreateRequest) *ProjectServiceCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the project service create params
func (o *ProjectServiceCreateParams) SetBody(body *models.ProjectProjectCreateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectServiceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
