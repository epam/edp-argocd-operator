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
)

// NewRepositoryServiceValidateAccessParams creates a new RepositoryServiceValidateAccessParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepositoryServiceValidateAccessParams() *RepositoryServiceValidateAccessParams {
	return &RepositoryServiceValidateAccessParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepositoryServiceValidateAccessParamsWithTimeout creates a new RepositoryServiceValidateAccessParams object
// with the ability to set a timeout on a request.
func NewRepositoryServiceValidateAccessParamsWithTimeout(timeout time.Duration) *RepositoryServiceValidateAccessParams {
	return &RepositoryServiceValidateAccessParams{
		timeout: timeout,
	}
}

// NewRepositoryServiceValidateAccessParamsWithContext creates a new RepositoryServiceValidateAccessParams object
// with the ability to set a context for a request.
func NewRepositoryServiceValidateAccessParamsWithContext(ctx context.Context) *RepositoryServiceValidateAccessParams {
	return &RepositoryServiceValidateAccessParams{
		Context: ctx,
	}
}

// NewRepositoryServiceValidateAccessParamsWithHTTPClient creates a new RepositoryServiceValidateAccessParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepositoryServiceValidateAccessParamsWithHTTPClient(client *http.Client) *RepositoryServiceValidateAccessParams {
	return &RepositoryServiceValidateAccessParams{
		HTTPClient: client,
	}
}

/* RepositoryServiceValidateAccessParams contains all the parameters to send to the API endpoint
   for the repository service validate access operation.

   Typically these are written to a http.Request.
*/
type RepositoryServiceValidateAccessParams struct {

	/* Body.

	   The URL to the repo
	*/
	Body string

	/* EnableOci.

	   Whether helm-oci support should be enabled for this repo.
	*/
	EnableOci *bool

	/* GithubAppEnterpriseBaseURL.

	   Github App Enterprise base url if empty will default to https://api.github.com.
	*/
	GithubAppEnterpriseBaseURL *string

	/* GithubAppID.

	   Github App ID of the app used to access the repo.

	   Format: int64
	*/
	GithubAppID *int64

	/* GithubAppInstallationID.

	   Github App Installation ID of the installed GitHub App.

	   Format: int64
	*/
	GithubAppInstallationID *int64

	/* GithubAppPrivateKey.

	   Github App Private Key PEM data.
	*/
	GithubAppPrivateKey *string

	/* Insecure.

	   Whether to skip certificate or host key validation.
	*/
	Insecure *bool

	/* Name.

	   The name of the repo.
	*/
	Name *string

	/* Password.

	   Password for accessing repo.
	*/
	Password *string

	/* Project.

	   Reference between project and repository that allow you automatically to be added as item inside SourceRepos project entity.
	*/
	Project *string

	/* Proxy.

	   HTTP/HTTPS proxy to access the repository.
	*/
	Proxy *string

	/* Repo.

	   The URL to the repo
	*/
	Repo string

	/* SSHPrivateKey.

	   Private key data for accessing SSH repository.
	*/
	SSHPrivateKey *string

	/* TLSClientCertData.

	   TLS client cert data for accessing HTTPS repository.
	*/
	TLSClientCertData *string

	/* TLSClientCertKey.

	   TLS client cert key for accessing HTTPS repository.
	*/
	TLSClientCertKey *string

	/* Type.

	   The type of the repo.
	*/
	Type *string

	/* Username.

	   Username for accessing repo.
	*/
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repository service validate access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepositoryServiceValidateAccessParams) WithDefaults() *RepositoryServiceValidateAccessParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repository service validate access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepositoryServiceValidateAccessParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithTimeout(timeout time.Duration) *RepositoryServiceValidateAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithContext(ctx context.Context) *RepositoryServiceValidateAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithHTTPClient(client *http.Client) *RepositoryServiceValidateAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithBody(body string) *RepositoryServiceValidateAccessParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetBody(body string) {
	o.Body = body
}

// WithEnableOci adds the enableOci to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithEnableOci(enableOci *bool) *RepositoryServiceValidateAccessParams {
	o.SetEnableOci(enableOci)
	return o
}

// SetEnableOci adds the enableOci to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetEnableOci(enableOci *bool) {
	o.EnableOci = enableOci
}

// WithGithubAppEnterpriseBaseURL adds the githubAppEnterpriseBaseURL to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithGithubAppEnterpriseBaseURL(githubAppEnterpriseBaseURL *string) *RepositoryServiceValidateAccessParams {
	o.SetGithubAppEnterpriseBaseURL(githubAppEnterpriseBaseURL)
	return o
}

// SetGithubAppEnterpriseBaseURL adds the githubAppEnterpriseBaseUrl to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetGithubAppEnterpriseBaseURL(githubAppEnterpriseBaseURL *string) {
	o.GithubAppEnterpriseBaseURL = githubAppEnterpriseBaseURL
}

// WithGithubAppID adds the githubAppID to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithGithubAppID(githubAppID *int64) *RepositoryServiceValidateAccessParams {
	o.SetGithubAppID(githubAppID)
	return o
}

// SetGithubAppID adds the githubAppId to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetGithubAppID(githubAppID *int64) {
	o.GithubAppID = githubAppID
}

// WithGithubAppInstallationID adds the githubAppInstallationID to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithGithubAppInstallationID(githubAppInstallationID *int64) *RepositoryServiceValidateAccessParams {
	o.SetGithubAppInstallationID(githubAppInstallationID)
	return o
}

// SetGithubAppInstallationID adds the githubAppInstallationId to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetGithubAppInstallationID(githubAppInstallationID *int64) {
	o.GithubAppInstallationID = githubAppInstallationID
}

// WithGithubAppPrivateKey adds the githubAppPrivateKey to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithGithubAppPrivateKey(githubAppPrivateKey *string) *RepositoryServiceValidateAccessParams {
	o.SetGithubAppPrivateKey(githubAppPrivateKey)
	return o
}

// SetGithubAppPrivateKey adds the githubAppPrivateKey to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetGithubAppPrivateKey(githubAppPrivateKey *string) {
	o.GithubAppPrivateKey = githubAppPrivateKey
}

// WithInsecure adds the insecure to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithInsecure(insecure *bool) *RepositoryServiceValidateAccessParams {
	o.SetInsecure(insecure)
	return o
}

// SetInsecure adds the insecure to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetInsecure(insecure *bool) {
	o.Insecure = insecure
}

// WithName adds the name to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithName(name *string) *RepositoryServiceValidateAccessParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetName(name *string) {
	o.Name = name
}

// WithPassword adds the password to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithPassword(password *string) *RepositoryServiceValidateAccessParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetPassword(password *string) {
	o.Password = password
}

// WithProject adds the project to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithProject(project *string) *RepositoryServiceValidateAccessParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetProject(project *string) {
	o.Project = project
}

// WithProxy adds the proxy to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithProxy(proxy *string) *RepositoryServiceValidateAccessParams {
	o.SetProxy(proxy)
	return o
}

// SetProxy adds the proxy to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetProxy(proxy *string) {
	o.Proxy = proxy
}

// WithRepo adds the repo to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithRepo(repo string) *RepositoryServiceValidateAccessParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithSSHPrivateKey adds the sSHPrivateKey to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithSSHPrivateKey(sSHPrivateKey *string) *RepositoryServiceValidateAccessParams {
	o.SetSSHPrivateKey(sSHPrivateKey)
	return o
}

// SetSSHPrivateKey adds the sshPrivateKey to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetSSHPrivateKey(sSHPrivateKey *string) {
	o.SSHPrivateKey = sSHPrivateKey
}

// WithTLSClientCertData adds the tLSClientCertData to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithTLSClientCertData(tLSClientCertData *string) *RepositoryServiceValidateAccessParams {
	o.SetTLSClientCertData(tLSClientCertData)
	return o
}

// SetTLSClientCertData adds the tlsClientCertData to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetTLSClientCertData(tLSClientCertData *string) {
	o.TLSClientCertData = tLSClientCertData
}

// WithTLSClientCertKey adds the tLSClientCertKey to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithTLSClientCertKey(tLSClientCertKey *string) *RepositoryServiceValidateAccessParams {
	o.SetTLSClientCertKey(tLSClientCertKey)
	return o
}

// SetTLSClientCertKey adds the tlsClientCertKey to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetTLSClientCertKey(tLSClientCertKey *string) {
	o.TLSClientCertKey = tLSClientCertKey
}

// WithType adds the typeVar to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithType(typeVar *string) *RepositoryServiceValidateAccessParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithUsername adds the username to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) WithUsername(username *string) *RepositoryServiceValidateAccessParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the repository service validate access params
func (o *RepositoryServiceValidateAccessParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *RepositoryServiceValidateAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if o.EnableOci != nil {

		// query param enableOci
		var qrEnableOci bool

		if o.EnableOci != nil {
			qrEnableOci = *o.EnableOci
		}
		qEnableOci := swag.FormatBool(qrEnableOci)
		if qEnableOci != "" {

			if err := r.SetQueryParam("enableOci", qEnableOci); err != nil {
				return err
			}
		}
	}

	if o.GithubAppEnterpriseBaseURL != nil {

		// query param githubAppEnterpriseBaseUrl
		var qrGithubAppEnterpriseBaseURL string

		if o.GithubAppEnterpriseBaseURL != nil {
			qrGithubAppEnterpriseBaseURL = *o.GithubAppEnterpriseBaseURL
		}
		qGithubAppEnterpriseBaseURL := qrGithubAppEnterpriseBaseURL
		if qGithubAppEnterpriseBaseURL != "" {

			if err := r.SetQueryParam("githubAppEnterpriseBaseUrl", qGithubAppEnterpriseBaseURL); err != nil {
				return err
			}
		}
	}

	if o.GithubAppID != nil {

		// query param githubAppID
		var qrGithubAppID int64

		if o.GithubAppID != nil {
			qrGithubAppID = *o.GithubAppID
		}
		qGithubAppID := swag.FormatInt64(qrGithubAppID)
		if qGithubAppID != "" {

			if err := r.SetQueryParam("githubAppID", qGithubAppID); err != nil {
				return err
			}
		}
	}

	if o.GithubAppInstallationID != nil {

		// query param githubAppInstallationID
		var qrGithubAppInstallationID int64

		if o.GithubAppInstallationID != nil {
			qrGithubAppInstallationID = *o.GithubAppInstallationID
		}
		qGithubAppInstallationID := swag.FormatInt64(qrGithubAppInstallationID)
		if qGithubAppInstallationID != "" {

			if err := r.SetQueryParam("githubAppInstallationID", qGithubAppInstallationID); err != nil {
				return err
			}
		}
	}

	if o.GithubAppPrivateKey != nil {

		// query param githubAppPrivateKey
		var qrGithubAppPrivateKey string

		if o.GithubAppPrivateKey != nil {
			qrGithubAppPrivateKey = *o.GithubAppPrivateKey
		}
		qGithubAppPrivateKey := qrGithubAppPrivateKey
		if qGithubAppPrivateKey != "" {

			if err := r.SetQueryParam("githubAppPrivateKey", qGithubAppPrivateKey); err != nil {
				return err
			}
		}
	}

	if o.Insecure != nil {

		// query param insecure
		var qrInsecure bool

		if o.Insecure != nil {
			qrInsecure = *o.Insecure
		}
		qInsecure := swag.FormatBool(qrInsecure)
		if qInsecure != "" {

			if err := r.SetQueryParam("insecure", qInsecure); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Password != nil {

		// query param password
		var qrPassword string

		if o.Password != nil {
			qrPassword = *o.Password
		}
		qPassword := qrPassword
		if qPassword != "" {

			if err := r.SetQueryParam("password", qPassword); err != nil {
				return err
			}
		}
	}

	if o.Project != nil {

		// query param project
		var qrProject string

		if o.Project != nil {
			qrProject = *o.Project
		}
		qProject := qrProject
		if qProject != "" {

			if err := r.SetQueryParam("project", qProject); err != nil {
				return err
			}
		}
	}

	if o.Proxy != nil {

		// query param proxy
		var qrProxy string

		if o.Proxy != nil {
			qrProxy = *o.Proxy
		}
		qProxy := qrProxy
		if qProxy != "" {

			if err := r.SetQueryParam("proxy", qProxy); err != nil {
				return err
			}
		}
	}

	// path param repo
	if err := r.SetPathParam("repo", o.Repo); err != nil {
		return err
	}

	if o.SSHPrivateKey != nil {

		// query param sshPrivateKey
		var qrSSHPrivateKey string

		if o.SSHPrivateKey != nil {
			qrSSHPrivateKey = *o.SSHPrivateKey
		}
		qSSHPrivateKey := qrSSHPrivateKey
		if qSSHPrivateKey != "" {

			if err := r.SetQueryParam("sshPrivateKey", qSSHPrivateKey); err != nil {
				return err
			}
		}
	}

	if o.TLSClientCertData != nil {

		// query param tlsClientCertData
		var qrTLSClientCertData string

		if o.TLSClientCertData != nil {
			qrTLSClientCertData = *o.TLSClientCertData
		}
		qTLSClientCertData := qrTLSClientCertData
		if qTLSClientCertData != "" {

			if err := r.SetQueryParam("tlsClientCertData", qTLSClientCertData); err != nil {
				return err
			}
		}
	}

	if o.TLSClientCertKey != nil {

		// query param tlsClientCertKey
		var qrTLSClientCertKey string

		if o.TLSClientCertKey != nil {
			qrTLSClientCertKey = *o.TLSClientCertKey
		}
		qTLSClientCertKey := qrTLSClientCertKey
		if qTLSClientCertKey != "" {

			if err := r.SetQueryParam("tlsClientCertKey", qTLSClientCertKey); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.Username != nil {

		// query param username
		var qrUsername string

		if o.Username != nil {
			qrUsername = *o.Username
		}
		qUsername := qrUsername
		if qUsername != "" {

			if err := r.SetQueryParam("username", qUsername); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
