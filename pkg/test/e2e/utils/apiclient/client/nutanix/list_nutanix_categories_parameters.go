// Code generated by go-swagger; DO NOT EDIT.

package nutanix

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

// NewListNutanixCategoriesParams creates a new ListNutanixCategoriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListNutanixCategoriesParams() *ListNutanixCategoriesParams {
	return &ListNutanixCategoriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListNutanixCategoriesParamsWithTimeout creates a new ListNutanixCategoriesParams object
// with the ability to set a timeout on a request.
func NewListNutanixCategoriesParamsWithTimeout(timeout time.Duration) *ListNutanixCategoriesParams {
	return &ListNutanixCategoriesParams{
		timeout: timeout,
	}
}

// NewListNutanixCategoriesParamsWithContext creates a new ListNutanixCategoriesParams object
// with the ability to set a context for a request.
func NewListNutanixCategoriesParamsWithContext(ctx context.Context) *ListNutanixCategoriesParams {
	return &ListNutanixCategoriesParams{
		Context: ctx,
	}
}

// NewListNutanixCategoriesParamsWithHTTPClient creates a new ListNutanixCategoriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListNutanixCategoriesParamsWithHTTPClient(client *http.Client) *ListNutanixCategoriesParams {
	return &ListNutanixCategoriesParams{
		HTTPClient: client,
	}
}

/* ListNutanixCategoriesParams contains all the parameters to send to the API endpoint
   for the list nutanix categories operation.

   Typically these are written to a http.Request.
*/
type ListNutanixCategoriesParams struct {

	// Credential.
	Credential *string

	// NutanixPassword.
	NutanixPassword *string

	// NutanixProxyURL.
	NutanixProxyURL *string

	// NutanixUsername.
	NutanixUsername *string

	/* Dc.

	   KKP Datacenter to use for endpoint
	*/
	DC string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list nutanix categories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNutanixCategoriesParams) WithDefaults() *ListNutanixCategoriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list nutanix categories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNutanixCategoriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list nutanix categories params
func (o *ListNutanixCategoriesParams) WithTimeout(timeout time.Duration) *ListNutanixCategoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list nutanix categories params
func (o *ListNutanixCategoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list nutanix categories params
func (o *ListNutanixCategoriesParams) WithContext(ctx context.Context) *ListNutanixCategoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list nutanix categories params
func (o *ListNutanixCategoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list nutanix categories params
func (o *ListNutanixCategoriesParams) WithHTTPClient(client *http.Client) *ListNutanixCategoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list nutanix categories params
func (o *ListNutanixCategoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list nutanix categories params
func (o *ListNutanixCategoriesParams) WithCredential(credential *string) *ListNutanixCategoriesParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list nutanix categories params
func (o *ListNutanixCategoriesParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithNutanixPassword adds the nutanixPassword to the list nutanix categories params
func (o *ListNutanixCategoriesParams) WithNutanixPassword(nutanixPassword *string) *ListNutanixCategoriesParams {
	o.SetNutanixPassword(nutanixPassword)
	return o
}

// SetNutanixPassword adds the nutanixPassword to the list nutanix categories params
func (o *ListNutanixCategoriesParams) SetNutanixPassword(nutanixPassword *string) {
	o.NutanixPassword = nutanixPassword
}

// WithNutanixProxyURL adds the nutanixProxyURL to the list nutanix categories params
func (o *ListNutanixCategoriesParams) WithNutanixProxyURL(nutanixProxyURL *string) *ListNutanixCategoriesParams {
	o.SetNutanixProxyURL(nutanixProxyURL)
	return o
}

// SetNutanixProxyURL adds the nutanixProxyUrl to the list nutanix categories params
func (o *ListNutanixCategoriesParams) SetNutanixProxyURL(nutanixProxyURL *string) {
	o.NutanixProxyURL = nutanixProxyURL
}

// WithNutanixUsername adds the nutanixUsername to the list nutanix categories params
func (o *ListNutanixCategoriesParams) WithNutanixUsername(nutanixUsername *string) *ListNutanixCategoriesParams {
	o.SetNutanixUsername(nutanixUsername)
	return o
}

// SetNutanixUsername adds the nutanixUsername to the list nutanix categories params
func (o *ListNutanixCategoriesParams) SetNutanixUsername(nutanixUsername *string) {
	o.NutanixUsername = nutanixUsername
}

// WithDC adds the dc to the list nutanix categories params
func (o *ListNutanixCategoriesParams) WithDC(dc string) *ListNutanixCategoriesParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list nutanix categories params
func (o *ListNutanixCategoriesParams) SetDC(dc string) {
	o.DC = dc
}

// WriteToRequest writes these params to a swagger request
func (o *ListNutanixCategoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.NutanixPassword != nil {

		// header param NutanixPassword
		if err := r.SetHeaderParam("NutanixPassword", *o.NutanixPassword); err != nil {
			return err
		}
	}

	if o.NutanixProxyURL != nil {

		// header param NutanixProxyURL
		if err := r.SetHeaderParam("NutanixProxyURL", *o.NutanixProxyURL); err != nil {
			return err
		}
	}

	if o.NutanixUsername != nil {

		// header param NutanixUsername
		if err := r.SetHeaderParam("NutanixUsername", *o.NutanixUsername); err != nil {
			return err
		}
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
