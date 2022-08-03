// Code generated by go-swagger; DO NOT EDIT.

package certificate_service

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

// NewCertificateServiceListCertificatesParams creates a new CertificateServiceListCertificatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCertificateServiceListCertificatesParams() *CertificateServiceListCertificatesParams {
	return &CertificateServiceListCertificatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCertificateServiceListCertificatesParamsWithTimeout creates a new CertificateServiceListCertificatesParams object
// with the ability to set a timeout on a request.
func NewCertificateServiceListCertificatesParamsWithTimeout(timeout time.Duration) *CertificateServiceListCertificatesParams {
	return &CertificateServiceListCertificatesParams{
		timeout: timeout,
	}
}

// NewCertificateServiceListCertificatesParamsWithContext creates a new CertificateServiceListCertificatesParams object
// with the ability to set a context for a request.
func NewCertificateServiceListCertificatesParamsWithContext(ctx context.Context) *CertificateServiceListCertificatesParams {
	return &CertificateServiceListCertificatesParams{
		Context: ctx,
	}
}

// NewCertificateServiceListCertificatesParamsWithHTTPClient creates a new CertificateServiceListCertificatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCertificateServiceListCertificatesParamsWithHTTPClient(client *http.Client) *CertificateServiceListCertificatesParams {
	return &CertificateServiceListCertificatesParams{
		HTTPClient: client,
	}
}

/* CertificateServiceListCertificatesParams contains all the parameters to send to the API endpoint
   for the certificate service list certificates operation.

   Typically these are written to a http.Request.
*/
type CertificateServiceListCertificatesParams struct {

	/* CertSubType.

	   The sub type of the certificate to match (protocol dependent, usually only used for ssh certs).
	*/
	CertSubType *string

	/* CertType.

	   The type of the certificate to match (ssh or https).
	*/
	CertType *string

	/* HostNamePattern.

	   A file-glob pattern (not regular expression) the host name has to match.
	*/
	HostNamePattern *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the certificate service list certificates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CertificateServiceListCertificatesParams) WithDefaults() *CertificateServiceListCertificatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the certificate service list certificates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CertificateServiceListCertificatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the certificate service list certificates params
func (o *CertificateServiceListCertificatesParams) WithTimeout(timeout time.Duration) *CertificateServiceListCertificatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the certificate service list certificates params
func (o *CertificateServiceListCertificatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the certificate service list certificates params
func (o *CertificateServiceListCertificatesParams) WithContext(ctx context.Context) *CertificateServiceListCertificatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the certificate service list certificates params
func (o *CertificateServiceListCertificatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the certificate service list certificates params
func (o *CertificateServiceListCertificatesParams) WithHTTPClient(client *http.Client) *CertificateServiceListCertificatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the certificate service list certificates params
func (o *CertificateServiceListCertificatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCertSubType adds the certSubType to the certificate service list certificates params
func (o *CertificateServiceListCertificatesParams) WithCertSubType(certSubType *string) *CertificateServiceListCertificatesParams {
	o.SetCertSubType(certSubType)
	return o
}

// SetCertSubType adds the certSubType to the certificate service list certificates params
func (o *CertificateServiceListCertificatesParams) SetCertSubType(certSubType *string) {
	o.CertSubType = certSubType
}

// WithCertType adds the certType to the certificate service list certificates params
func (o *CertificateServiceListCertificatesParams) WithCertType(certType *string) *CertificateServiceListCertificatesParams {
	o.SetCertType(certType)
	return o
}

// SetCertType adds the certType to the certificate service list certificates params
func (o *CertificateServiceListCertificatesParams) SetCertType(certType *string) {
	o.CertType = certType
}

// WithHostNamePattern adds the hostNamePattern to the certificate service list certificates params
func (o *CertificateServiceListCertificatesParams) WithHostNamePattern(hostNamePattern *string) *CertificateServiceListCertificatesParams {
	o.SetHostNamePattern(hostNamePattern)
	return o
}

// SetHostNamePattern adds the hostNamePattern to the certificate service list certificates params
func (o *CertificateServiceListCertificatesParams) SetHostNamePattern(hostNamePattern *string) {
	o.HostNamePattern = hostNamePattern
}

// WriteToRequest writes these params to a swagger request
func (o *CertificateServiceListCertificatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CertSubType != nil {

		// query param certSubType
		var qrCertSubType string

		if o.CertSubType != nil {
			qrCertSubType = *o.CertSubType
		}
		qCertSubType := qrCertSubType
		if qCertSubType != "" {

			if err := r.SetQueryParam("certSubType", qCertSubType); err != nil {
				return err
			}
		}
	}

	if o.CertType != nil {

		// query param certType
		var qrCertType string

		if o.CertType != nil {
			qrCertType = *o.CertType
		}
		qCertType := qrCertType
		if qCertType != "" {

			if err := r.SetQueryParam("certType", qCertType); err != nil {
				return err
			}
		}
	}

	if o.HostNamePattern != nil {

		// query param hostNamePattern
		var qrHostNamePattern string

		if o.HostNamePattern != nil {
			qrHostNamePattern = *o.HostNamePattern
		}
		qHostNamePattern := qrHostNamePattern
		if qHostNamePattern != "" {

			if err := r.SetQueryParam("hostNamePattern", qHostNamePattern); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}