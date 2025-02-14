// Code generated by go-swagger; DO NOT EDIT.

package applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new applications API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for applications API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateApplicationInstallation(params *CreateApplicationInstallationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateApplicationInstallationCreated, error)

	DeleteApplicationInstallation(params *DeleteApplicationInstallationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteApplicationInstallationOK, error)

	GetApplicationInstallation(params *GetApplicationInstallationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetApplicationInstallationOK, error)

	ListApplicationDefinitions(params *ListApplicationDefinitionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListApplicationDefinitionsOK, error)

	ListApplicationInstallations(params *ListApplicationInstallationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListApplicationInstallationsOK, error)

	UpdateApplicationInstallation(params *UpdateApplicationInstallationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateApplicationInstallationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateApplicationInstallation Creates ApplicationInstallation into the given cluster
*/
func (a *Client) CreateApplicationInstallation(params *CreateApplicationInstallationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateApplicationInstallationCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateApplicationInstallationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createApplicationInstallation",
		Method:             "POST",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateApplicationInstallationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateApplicationInstallationCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateApplicationInstallationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteApplicationInstallation Deletes the given ApplicationInstallation
*/
func (a *Client) DeleteApplicationInstallation(params *DeleteApplicationInstallationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteApplicationInstallationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteApplicationInstallationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteApplicationInstallation",
		Method:             "DELETE",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations/{namespace}/{appinstall_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteApplicationInstallationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteApplicationInstallationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteApplicationInstallationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetApplicationInstallation Gets the given ApplicationInstallation
*/
func (a *Client) GetApplicationInstallation(params *GetApplicationInstallationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetApplicationInstallationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApplicationInstallationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getApplicationInstallation",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations/{namespace}/{appinstall_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetApplicationInstallationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApplicationInstallationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetApplicationInstallationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListApplicationDefinitions List ApplicationDefinitions which are available in the KKP installation
*/
func (a *Client) ListApplicationDefinitions(params *ListApplicationDefinitionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListApplicationDefinitionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListApplicationDefinitionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listApplicationDefinitions",
		Method:             "GET",
		PathPattern:        "/api/v2/applicationdefinitions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListApplicationDefinitionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListApplicationDefinitionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListApplicationDefinitionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListApplicationInstallations List ApplicationInstallations which belong to the given cluster
*/
func (a *Client) ListApplicationInstallations(params *ListApplicationInstallationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListApplicationInstallationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListApplicationInstallationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listApplicationInstallations",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListApplicationInstallationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListApplicationInstallationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListApplicationInstallationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateApplicationInstallation Updates the given ApplicationInstallation
*/
func (a *Client) UpdateApplicationInstallation(params *UpdateApplicationInstallationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateApplicationInstallationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateApplicationInstallationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateApplicationInstallation",
		Method:             "PUT",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations/{namespace}/{appinstall_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateApplicationInstallationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateApplicationInstallationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateApplicationInstallationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
