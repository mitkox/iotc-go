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

// NewRolesGetParams creates a new RolesGetParams object
// with the default values initialized.
func NewRolesGetParams() *RolesGetParams {
	var ()
	return &RolesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRolesGetParamsWithTimeout creates a new RolesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRolesGetParamsWithTimeout(timeout time.Duration) *RolesGetParams {
	var ()
	return &RolesGetParams{

		timeout: timeout,
	}
}

// NewRolesGetParamsWithContext creates a new RolesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewRolesGetParamsWithContext(ctx context.Context) *RolesGetParams {
	var ()
	return &RolesGetParams{

		Context: ctx,
	}
}

// NewRolesGetParamsWithHTTPClient creates a new RolesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRolesGetParamsWithHTTPClient(client *http.Client) *RolesGetParams {
	var ()
	return &RolesGetParams{
		HTTPClient: client,
	}
}

/*RolesGetParams contains all the parameters to send to the API endpoint
for the roles get operation typically these are written to a http.Request
*/
type RolesGetParams struct {

	/*RoleID
	  Unique ID for the role.

	*/
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the roles get params
func (o *RolesGetParams) WithTimeout(timeout time.Duration) *RolesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the roles get params
func (o *RolesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the roles get params
func (o *RolesGetParams) WithContext(ctx context.Context) *RolesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the roles get params
func (o *RolesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the roles get params
func (o *RolesGetParams) WithHTTPClient(client *http.Client) *RolesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the roles get params
func (o *RolesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleID adds the roleID to the roles get params
func (o *RolesGetParams) WithRoleID(roleID string) *RolesGetParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the roles get params
func (o *RolesGetParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *RolesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param role_id
	if err := r.SetPathParam("role_id", o.RoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}