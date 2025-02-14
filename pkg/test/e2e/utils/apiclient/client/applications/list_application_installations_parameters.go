// Code generated by go-swagger; DO NOT EDIT.

package applications

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

// NewListApplicationInstallationsParams creates a new ListApplicationInstallationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListApplicationInstallationsParams() *ListApplicationInstallationsParams {
	return &ListApplicationInstallationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListApplicationInstallationsParamsWithTimeout creates a new ListApplicationInstallationsParams object
// with the ability to set a timeout on a request.
func NewListApplicationInstallationsParamsWithTimeout(timeout time.Duration) *ListApplicationInstallationsParams {
	return &ListApplicationInstallationsParams{
		timeout: timeout,
	}
}

// NewListApplicationInstallationsParamsWithContext creates a new ListApplicationInstallationsParams object
// with the ability to set a context for a request.
func NewListApplicationInstallationsParamsWithContext(ctx context.Context) *ListApplicationInstallationsParams {
	return &ListApplicationInstallationsParams{
		Context: ctx,
	}
}

// NewListApplicationInstallationsParamsWithHTTPClient creates a new ListApplicationInstallationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListApplicationInstallationsParamsWithHTTPClient(client *http.Client) *ListApplicationInstallationsParams {
	return &ListApplicationInstallationsParams{
		HTTPClient: client,
	}
}

/* ListApplicationInstallationsParams contains all the parameters to send to the API endpoint
   for the list application installations operation.

   Typically these are written to a http.Request.
*/
type ListApplicationInstallationsParams struct {

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list application installations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListApplicationInstallationsParams) WithDefaults() *ListApplicationInstallationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list application installations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListApplicationInstallationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list application installations params
func (o *ListApplicationInstallationsParams) WithTimeout(timeout time.Duration) *ListApplicationInstallationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list application installations params
func (o *ListApplicationInstallationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list application installations params
func (o *ListApplicationInstallationsParams) WithContext(ctx context.Context) *ListApplicationInstallationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list application installations params
func (o *ListApplicationInstallationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list application installations params
func (o *ListApplicationInstallationsParams) WithHTTPClient(client *http.Client) *ListApplicationInstallationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list application installations params
func (o *ListApplicationInstallationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list application installations params
func (o *ListApplicationInstallationsParams) WithClusterID(clusterID string) *ListApplicationInstallationsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list application installations params
func (o *ListApplicationInstallationsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the list application installations params
func (o *ListApplicationInstallationsParams) WithProjectID(projectID string) *ListApplicationInstallationsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list application installations params
func (o *ListApplicationInstallationsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListApplicationInstallationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
