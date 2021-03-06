// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPcloudEventsGetsinceParams creates a new PcloudEventsGetsinceParams object
// with the default values initialized.
func NewPcloudEventsGetsinceParams() *PcloudEventsGetsinceParams {
	var ()
	return &PcloudEventsGetsinceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudEventsGetsinceParamsWithTimeout creates a new PcloudEventsGetsinceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudEventsGetsinceParamsWithTimeout(timeout time.Duration) *PcloudEventsGetsinceParams {
	var ()
	return &PcloudEventsGetsinceParams{

		timeout: timeout,
	}
}

// NewPcloudEventsGetsinceParamsWithContext creates a new PcloudEventsGetsinceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudEventsGetsinceParamsWithContext(ctx context.Context) *PcloudEventsGetsinceParams {
	var ()
	return &PcloudEventsGetsinceParams{

		Context: ctx,
	}
}

// NewPcloudEventsGetsinceParamsWithHTTPClient creates a new PcloudEventsGetsinceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudEventsGetsinceParamsWithHTTPClient(client *http.Client) *PcloudEventsGetsinceParams {
	var ()
	return &PcloudEventsGetsinceParams{
		HTTPClient: client,
	}
}

/*PcloudEventsGetsinceParams contains all the parameters to send to the API endpoint
for the pcloud events getsince operation typically these are written to a http.Request
*/
type PcloudEventsGetsinceParams struct {

	/*AcceptLanguage
	  The language requested for the return document

	*/
	AcceptLanguage *string
	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*Time
	  A time in either ISO 8601 or unix epoch format

	*/
	Time string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud events getsince params
func (o *PcloudEventsGetsinceParams) WithTimeout(timeout time.Duration) *PcloudEventsGetsinceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud events getsince params
func (o *PcloudEventsGetsinceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud events getsince params
func (o *PcloudEventsGetsinceParams) WithContext(ctx context.Context) *PcloudEventsGetsinceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud events getsince params
func (o *PcloudEventsGetsinceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud events getsince params
func (o *PcloudEventsGetsinceParams) WithHTTPClient(client *http.Client) *PcloudEventsGetsinceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud events getsince params
func (o *PcloudEventsGetsinceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the pcloud events getsince params
func (o *PcloudEventsGetsinceParams) WithAcceptLanguage(acceptLanguage *string) *PcloudEventsGetsinceParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the pcloud events getsince params
func (o *PcloudEventsGetsinceParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud events getsince params
func (o *PcloudEventsGetsinceParams) WithCloudInstanceID(cloudInstanceID string) *PcloudEventsGetsinceParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud events getsince params
func (o *PcloudEventsGetsinceParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithTime adds the time to the pcloud events getsince params
func (o *PcloudEventsGetsinceParams) WithTime(time string) *PcloudEventsGetsinceParams {
	o.SetTime(time)
	return o
}

// SetTime adds the time to the pcloud events getsince params
func (o *PcloudEventsGetsinceParams) SetTime(time string) {
	o.Time = time
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudEventsGetsinceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AcceptLanguage != nil {

		// header param Accept-Language
		if err := r.SetHeaderParam("Accept-Language", *o.AcceptLanguage); err != nil {
			return err
		}

	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// query param time
	qrTime := o.Time
	qTime := qrTime
	if qTime != "" {
		if err := r.SetQueryParam("time", qTime); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
