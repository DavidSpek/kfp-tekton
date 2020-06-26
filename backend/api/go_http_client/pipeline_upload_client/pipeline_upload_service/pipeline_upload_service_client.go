// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package pipeline_upload_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new pipeline upload service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pipeline upload service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
UploadPipeline upload pipeline API
*/
func (a *Client) UploadPipeline(params *UploadPipelineParams, authInfo runtime.ClientAuthInfoWriter) (*UploadPipelineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadPipelineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UploadPipeline",
		Method:             "POST",
		PathPattern:        "/apis/v1beta1/pipelines/upload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UploadPipelineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UploadPipelineOK), nil

}

/*
UploadPipelineVersion upload pipeline version API
*/
func (a *Client) UploadPipelineVersion(params *UploadPipelineVersionParams, authInfo runtime.ClientAuthInfoWriter) (*UploadPipelineVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadPipelineVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UploadPipelineVersion",
		Method:             "POST",
		PathPattern:        "/apis/v1beta1/pipelines/upload_version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UploadPipelineVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UploadPipelineVersionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}