// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new actions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for actions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CancelAction cancels action stops an action
*/
func (a *Client) CancelAction(params *CancelActionParams) (*CancelActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CancelAction",
		Method:             "POST",
		PathPattern:        "/v0/management/Actions/Cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CancelActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CancelActionOK), nil

}

/*
GetAction gets action gets an result of given action
*/
func (a *Client) GetAction(params *GetActionParams) (*GetActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAction",
		Method:             "POST",
		PathPattern:        "/v0/management/Actions/Get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetActionOK), nil

}

/*
StartMySQLExplainAction starts my SQL explain action starts my SQL e x p l a i n action
*/
func (a *Client) StartMySQLExplainAction(params *StartMySQLExplainActionParams) (*StartMySQLExplainActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartMySQLExplainActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartMySQLExplainAction",
		Method:             "POST",
		PathPattern:        "/v0/management/Actions/StartMySQLExplain",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartMySQLExplainActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartMySQLExplainActionOK), nil

}

/*
StartMySQLExplainJSONAction starts my SQL explain JSON action starts my SQL JSON e x p l a i n action
*/
func (a *Client) StartMySQLExplainJSONAction(params *StartMySQLExplainJSONActionParams) (*StartMySQLExplainJSONActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartMySQLExplainJSONActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartMySQLExplainJSONAction",
		Method:             "POST",
		PathPattern:        "/v0/management/Actions/StartMySQLExplainJSON",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartMySQLExplainJSONActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartMySQLExplainJSONActionOK), nil

}

/*
StartMySQLShowCreateTableAction starts my SQL show create table action starts my SQL s h o w c r e a t e t a b l e action
*/
func (a *Client) StartMySQLShowCreateTableAction(params *StartMySQLShowCreateTableActionParams) (*StartMySQLShowCreateTableActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartMySQLShowCreateTableActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartMySQLShowCreateTableAction",
		Method:             "POST",
		PathPattern:        "/v0/management/Actions/StartMySQLShowCreateTable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartMySQLShowCreateTableActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartMySQLShowCreateTableActionOK), nil

}

/*
StartMySQLShowIndexAction starts my SQL show index action starts my SQL s h o w i n d e x action
*/
func (a *Client) StartMySQLShowIndexAction(params *StartMySQLShowIndexActionParams) (*StartMySQLShowIndexActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartMySQLShowIndexActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartMySQLShowIndexAction",
		Method:             "POST",
		PathPattern:        "/v0/management/Actions/StartMySQLShowIndex",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartMySQLShowIndexActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartMySQLShowIndexActionOK), nil

}

/*
StartMySQLShowTableStatusAction starts my SQL show table status action starts my SQL s h o w t a b l e s t a t u s action
*/
func (a *Client) StartMySQLShowTableStatusAction(params *StartMySQLShowTableStatusActionParams) (*StartMySQLShowTableStatusActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartMySQLShowTableStatusActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartMySQLShowTableStatusAction",
		Method:             "POST",
		PathPattern:        "/v0/management/Actions/StartMySQLShowTableStatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartMySQLShowTableStatusActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartMySQLShowTableStatusActionOK), nil

}

/*
StartPTMySQLSummaryAction starts p t my SQL summary action starts pt mysql summary action
*/
func (a *Client) StartPTMySQLSummaryAction(params *StartPTMySQLSummaryActionParams) (*StartPTMySQLSummaryActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartPTMySQLSummaryActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartPTMySQLSummaryAction",
		Method:             "POST",
		PathPattern:        "/v0/management/Actions/StartPTMySQLSummary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartPTMySQLSummaryActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartPTMySQLSummaryActionOK), nil

}

/*
StartPTSummaryAction starts p t summary action starts pt summary action
*/
func (a *Client) StartPTSummaryAction(params *StartPTSummaryActionParams) (*StartPTSummaryActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartPTSummaryActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartPTSummaryAction",
		Method:             "POST",
		PathPattern:        "/v0/management/Actions/StartPTSummary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartPTSummaryActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartPTSummaryActionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}