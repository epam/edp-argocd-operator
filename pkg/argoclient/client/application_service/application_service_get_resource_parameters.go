// Code generated by go-swagger; DO NOT EDIT.

package application_service

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

// NewApplicationServiceGetResourceParams creates a new ApplicationServiceGetResourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationServiceGetResourceParams() *ApplicationServiceGetResourceParams {
	return &ApplicationServiceGetResourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationServiceGetResourceParamsWithTimeout creates a new ApplicationServiceGetResourceParams object
// with the ability to set a timeout on a request.
func NewApplicationServiceGetResourceParamsWithTimeout(timeout time.Duration) *ApplicationServiceGetResourceParams {
	return &ApplicationServiceGetResourceParams{
		timeout: timeout,
	}
}

// NewApplicationServiceGetResourceParamsWithContext creates a new ApplicationServiceGetResourceParams object
// with the ability to set a context for a request.
func NewApplicationServiceGetResourceParamsWithContext(ctx context.Context) *ApplicationServiceGetResourceParams {
	return &ApplicationServiceGetResourceParams{
		Context: ctx,
	}
}

// NewApplicationServiceGetResourceParamsWithHTTPClient creates a new ApplicationServiceGetResourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationServiceGetResourceParamsWithHTTPClient(client *http.Client) *ApplicationServiceGetResourceParams {
	return &ApplicationServiceGetResourceParams{
		HTTPClient: client,
	}
}

/* ApplicationServiceGetResourceParams contains all the parameters to send to the API endpoint
   for the application service get resource operation.

   Typically these are written to a http.Request.
*/
type ApplicationServiceGetResourceParams struct {

	// Group.
	Group *string

	// Kind.
	Kind *string

	// Name.
	Name string

	// Namespace.
	Namespace *string

	// ResourceName.
	ResourceName *string

	// Version.
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application service get resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationServiceGetResourceParams) WithDefaults() *ApplicationServiceGetResourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application service get resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationServiceGetResourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the application service get resource params
func (o *ApplicationServiceGetResourceParams) WithTimeout(timeout time.Duration) *ApplicationServiceGetResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application service get resource params
func (o *ApplicationServiceGetResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application service get resource params
func (o *ApplicationServiceGetResourceParams) WithContext(ctx context.Context) *ApplicationServiceGetResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application service get resource params
func (o *ApplicationServiceGetResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application service get resource params
func (o *ApplicationServiceGetResourceParams) WithHTTPClient(client *http.Client) *ApplicationServiceGetResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application service get resource params
func (o *ApplicationServiceGetResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroup adds the group to the application service get resource params
func (o *ApplicationServiceGetResourceParams) WithGroup(group *string) *ApplicationServiceGetResourceParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the application service get resource params
func (o *ApplicationServiceGetResourceParams) SetGroup(group *string) {
	o.Group = group
}

// WithKind adds the kind to the application service get resource params
func (o *ApplicationServiceGetResourceParams) WithKind(kind *string) *ApplicationServiceGetResourceParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the application service get resource params
func (o *ApplicationServiceGetResourceParams) SetKind(kind *string) {
	o.Kind = kind
}

// WithName adds the name to the application service get resource params
func (o *ApplicationServiceGetResourceParams) WithName(name string) *ApplicationServiceGetResourceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the application service get resource params
func (o *ApplicationServiceGetResourceParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the application service get resource params
func (o *ApplicationServiceGetResourceParams) WithNamespace(namespace *string) *ApplicationServiceGetResourceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the application service get resource params
func (o *ApplicationServiceGetResourceParams) SetNamespace(namespace *string) {
	o.Namespace = namespace
}

// WithResourceName adds the resourceName to the application service get resource params
func (o *ApplicationServiceGetResourceParams) WithResourceName(resourceName *string) *ApplicationServiceGetResourceParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the application service get resource params
func (o *ApplicationServiceGetResourceParams) SetResourceName(resourceName *string) {
	o.ResourceName = resourceName
}

// WithVersion adds the version to the application service get resource params
func (o *ApplicationServiceGetResourceParams) WithVersion(version *string) *ApplicationServiceGetResourceParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the application service get resource params
func (o *ApplicationServiceGetResourceParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationServiceGetResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Group != nil {

		// query param group
		var qrGroup string

		if o.Group != nil {
			qrGroup = *o.Group
		}
		qGroup := qrGroup
		if qGroup != "" {

			if err := r.SetQueryParam("group", qGroup); err != nil {
				return err
			}
		}
	}

	if o.Kind != nil {

		// query param kind
		var qrKind string

		if o.Kind != nil {
			qrKind = *o.Kind
		}
		qKind := qrKind
		if qKind != "" {

			if err := r.SetQueryParam("kind", qKind); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Namespace != nil {

		// query param namespace
		var qrNamespace string

		if o.Namespace != nil {
			qrNamespace = *o.Namespace
		}
		qNamespace := qrNamespace
		if qNamespace != "" {

			if err := r.SetQueryParam("namespace", qNamespace); err != nil {
				return err
			}
		}
	}

	if o.ResourceName != nil {

		// query param resourceName
		var qrResourceName string

		if o.ResourceName != nil {
			qrResourceName = *o.ResourceName
		}
		qResourceName := qrResourceName
		if qResourceName != "" {

			if err := r.SetQueryParam("resourceName", qResourceName); err != nil {
				return err
			}
		}
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
