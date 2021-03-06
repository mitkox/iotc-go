// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewActionsGetParams creates a new ActionsGetParams object
// with the default values initialized.
func NewActionsGetParams() *ActionsGetParams {
	var ()
	return &ActionsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewActionsGetParamsWithTimeout creates a new ActionsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewActionsGetParamsWithTimeout(timeout time.Duration) *ActionsGetParams {
	var ()
	return &ActionsGetParams{

		timeout: timeout,
	}
}

// NewActionsGetParamsWithContext creates a new ActionsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewActionsGetParamsWithContext(ctx context.Context) *ActionsGetParams {
	var ()
	return &ActionsGetParams{

		Context: ctx,
	}
}

// NewActionsGetParamsWithHTTPClient creates a new ActionsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewActionsGetParamsWithHTTPClient(client *http.Client) *ActionsGetParams {
	var ()
	return &ActionsGetParams{
		HTTPClient: client,
	}
}

/*ActionsGetParams contains all the parameters to send to the API endpoint
for the actions get operation typically these are written to a http.Request
*/
type ActionsGetParams struct {

	/*ActionID
	  Unique ID of the action.

	*/
	ActionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the actions get params
func (o *ActionsGetParams) WithTimeout(timeout time.Duration) *ActionsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the actions get params
func (o *ActionsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the actions get params
func (o *ActionsGetParams) WithContext(ctx context.Context) *ActionsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the actions get params
func (o *ActionsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the actions get params
func (o *ActionsGetParams) WithHTTPClient(client *http.Client) *ActionsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the actions get params
func (o *ActionsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the actions get params
func (o *ActionsGetParams) WithActionID(actionID string) *ActionsGetParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the actions get params
func (o *ActionsGetParams) SetActionID(actionID string) {
	o.ActionID = actionID
}

// WriteToRequest writes these params to a swagger request
func (o *ActionsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param action_id
	if err := r.SetPathParam("action_id", o.ActionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
