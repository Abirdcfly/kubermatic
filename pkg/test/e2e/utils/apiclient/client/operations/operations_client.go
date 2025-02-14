// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateOIDCKubeconfig(params *CreateOIDCKubeconfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOIDCKubeconfigOK, error)

	CreateOIDCKubeconfigSecret(params *CreateOIDCKubeconfigSecretParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOIDCKubeconfigSecretOK, *CreateOIDCKubeconfigSecretCreated, error)

	GetAddonConfig(params *GetAddonConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAddonConfigOK, error)

	GetAdmissionPlugins(params *GetAdmissionPluginsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAdmissionPluginsOK, error)

	ListAddonConfigs(params *ListAddonConfigsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAddonConfigsOK, error)

	ListSystemLabels(params *ListSystemLabelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSystemLabelsOK, error)

	MigrateClusterToExternalCCM(params *MigrateClusterToExternalCCMParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MigrateClusterToExternalCCMOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateOIDCKubeconfig Starts OIDC flow and generates kubeconfig, the generated config
contains OIDC provider authentication info
*/
func (a *Client) CreateOIDCKubeconfig(params *CreateOIDCKubeconfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOIDCKubeconfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOIDCKubeconfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createOIDCKubeconfig",
		Method:             "GET",
		PathPattern:        "/api/v1/kubeconfig",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOIDCKubeconfigReader{formats: a.formats},
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
	success, ok := result.(*CreateOIDCKubeconfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateOIDCKubeconfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CreateOIDCKubeconfigSecret Starts OIDC flow and generates kubeconfig, the generated config
contains OIDC provider authentication info. The kubeconfig is stored in the secret.
*/
func (a *Client) CreateOIDCKubeconfigSecret(params *CreateOIDCKubeconfigSecretParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOIDCKubeconfigSecretOK, *CreateOIDCKubeconfigSecretCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOIDCKubeconfigSecretParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createOIDCKubeconfigSecret",
		Method:             "GET",
		PathPattern:        "/api/v2/kubeconfig/secret",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOIDCKubeconfigSecretReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateOIDCKubeconfigSecretOK:
		return value, nil, nil
	case *CreateOIDCKubeconfigSecretCreated:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateOIDCKubeconfigSecretDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetAddonConfig returns specified addon config
*/
func (a *Client) GetAddonConfig(params *GetAddonConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAddonConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAddonConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAddonConfig",
		Method:             "GET",
		PathPattern:        "/api/v1/addonconfigs/{addon_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAddonConfigReader{formats: a.formats},
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
	success, ok := result.(*GetAddonConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAddonConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetAdmissionPlugins returns specified addon config
*/
func (a *Client) GetAdmissionPlugins(params *GetAdmissionPluginsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAdmissionPluginsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdmissionPluginsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAdmissionPlugins",
		Method:             "GET",
		PathPattern:        "/api/v1/admission/plugins/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAdmissionPluginsReader{formats: a.formats},
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
	success, ok := result.(*GetAdmissionPluginsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAdmissionPluginsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListAddonConfigs returns all available addon configs
*/
func (a *Client) ListAddonConfigs(params *ListAddonConfigsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAddonConfigsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAddonConfigsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAddonConfigs",
		Method:             "GET",
		PathPattern:        "/api/v1/addonconfigs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAddonConfigsReader{formats: a.formats},
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
	success, ok := result.(*ListAddonConfigsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAddonConfigsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListSystemLabels List restricted system labels
*/
func (a *Client) ListSystemLabels(params *ListSystemLabelsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSystemLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSystemLabelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listSystemLabels",
		Method:             "GET",
		PathPattern:        "/api/v1/labels/system",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSystemLabelsReader{formats: a.formats},
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
	success, ok := result.(*ListSystemLabelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListSystemLabelsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  MigrateClusterToExternalCCM Enable the migration to the external CCM for the given cluster
*/
func (a *Client) MigrateClusterToExternalCCM(params *MigrateClusterToExternalCCMParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MigrateClusterToExternalCCMOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMigrateClusterToExternalCCMParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "migrateClusterToExternalCCM",
		Method:             "POST",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/externalccmmigration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MigrateClusterToExternalCCMReader{formats: a.formats},
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
	success, ok := result.(*MigrateClusterToExternalCCMOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MigrateClusterToExternalCCMDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
