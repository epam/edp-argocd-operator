// Code generated by go-swagger; DO NOT EDIT.

package repository_service

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
	"github.com/go-openapi/swag"

	"github.com/epam/edp-argocd-operator/pkg/argoclient/models"
)

// NewRepositoryServiceCreateRepositoryParams creates a new RepositoryServiceCreateRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepositoryServiceCreateRepositoryParams() *RepositoryServiceCreateRepositoryParams {
	return &RepositoryServiceCreateRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepositoryServiceCreateRepositoryParamsWithTimeout creates a new RepositoryServiceCreateRepositoryParams object
// with the ability to set a timeout on a request.
func NewRepositoryServiceCreateRepositoryParamsWithTimeout(timeout time.Duration) *RepositoryServiceCreateRepositoryParams {
	return &RepositoryServiceCreateRepositoryParams{
		timeout: timeout,
	}
}

// NewRepositoryServiceCreateRepositoryParamsWithContext creates a new RepositoryServiceCreateRepositoryParams object
// with the ability to set a context for a request.
func NewRepositoryServiceCreateRepositoryParamsWithContext(ctx context.Context) *RepositoryServiceCreateRepositoryParams {
	return &RepositoryServiceCreateRepositoryParams{
		Context: ctx,
	}
}

// NewRepositoryServiceCreateRepositoryParamsWithHTTPClient creates a new RepositoryServiceCreateRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepositoryServiceCreateRepositoryParamsWithHTTPClient(client *http.Client) *RepositoryServiceCreateRepositoryParams {
	return &RepositoryServiceCreateRepositoryParams{
		HTTPClient: client,
	}
}

/* RepositoryServiceCreateRepositoryParams contains all the parameters to send to the API endpoint
   for the repository service create repository operation.

   Typically these are written to a http.Request.
*/
type RepositoryServiceCreateRepositoryParams struct {

	/* Body.

	   Repository definition
	*/
	Body *models.V1alpha1Repository

	/* CredsOnly.

	   Whether to operate on credential set instead of repository.
	*/
	CredsOnly *bool

	/* Upsert.

	   Whether to create in upsert mode.
	*/
	Upsert *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repository service create repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepositoryServiceCreateRepositoryParams) WithDefaults() *RepositoryServiceCreateRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repository service create repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepositoryServiceCreateRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repository service create repository params
func (o *RepositoryServiceCreateRepositoryParams) WithTimeout(timeout time.Duration) *RepositoryServiceCreateRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repository service create repository params
func (o *RepositoryServiceCreateRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repository service create repository params
func (o *RepositoryServiceCreateRepositoryParams) WithContext(ctx context.Context) *RepositoryServiceCreateRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repository service create repository params
func (o *RepositoryServiceCreateRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repository service create repository params
func (o *RepositoryServiceCreateRepositoryParams) WithHTTPClient(client *http.Client) *RepositoryServiceCreateRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repository service create repository params
func (o *RepositoryServiceCreateRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repository service create repository params
func (o *RepositoryServiceCreateRepositoryParams) WithBody(body *models.V1alpha1Repository) *RepositoryServiceCreateRepositoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repository service create repository params
func (o *RepositoryServiceCreateRepositoryParams) SetBody(body *models.V1alpha1Repository) {
	o.Body = body
}

// WithCredsOnly adds the credsOnly to the repository service create repository params
func (o *RepositoryServiceCreateRepositoryParams) WithCredsOnly(credsOnly *bool) *RepositoryServiceCreateRepositoryParams {
	o.SetCredsOnly(credsOnly)
	return o
}

// SetCredsOnly adds the credsOnly to the repository service create repository params
func (o *RepositoryServiceCreateRepositoryParams) SetCredsOnly(credsOnly *bool) {
	o.CredsOnly = credsOnly
}

// WithUpsert adds the upsert to the repository service create repository params
func (o *RepositoryServiceCreateRepositoryParams) WithUpsert(upsert *bool) *RepositoryServiceCreateRepositoryParams {
	o.SetUpsert(upsert)
	return o
}

// SetUpsert adds the upsert to the repository service create repository params
func (o *RepositoryServiceCreateRepositoryParams) SetUpsert(upsert *bool) {
	o.Upsert = upsert
}

// WriteToRequest writes these params to a swagger request
func (o *RepositoryServiceCreateRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.CredsOnly != nil {

		// query param credsOnly
		var qrCredsOnly bool

		if o.CredsOnly != nil {
			qrCredsOnly = *o.CredsOnly
		}
		qCredsOnly := swag.FormatBool(qrCredsOnly)
		if qCredsOnly != "" {

			if err := r.SetQueryParam("credsOnly", qCredsOnly); err != nil {
				return err
			}
		}
	}

	if o.Upsert != nil {

		// query param upsert
		var qrUpsert bool

		if o.Upsert != nil {
			qrUpsert = *o.Upsert
		}
		qUpsert := swag.FormatBool(qrUpsert)
		if qUpsert != "" {

			if err := r.SetQueryParam("upsert", qUpsert); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}