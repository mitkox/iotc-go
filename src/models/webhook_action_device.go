// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebhookActionDevice webhook action device
//
// swagger:model WebhookActionDevice
type WebhookActionDevice struct {
	Device

	// Cloud properties that triggered the webhook.
	CloudProperties interface{} `json:"cloudProperties,omitempty"`

	// Device properties that triggered the webhook.
	Properties interface{} `json:"properties,omitempty"`

	// Device telemetry that triggered the webhook.
	Telemetry interface{} `json:"telemetry,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WebhookActionDevice) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Device
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Device = aO0

	// AO1
	var dataAO1 struct {
		CloudProperties interface{} `json:"cloudProperties,omitempty"`

		Properties interface{} `json:"properties,omitempty"`

		Telemetry interface{} `json:"telemetry,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.CloudProperties = dataAO1.CloudProperties

	m.Properties = dataAO1.Properties

	m.Telemetry = dataAO1.Telemetry

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WebhookActionDevice) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Device)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		CloudProperties interface{} `json:"cloudProperties,omitempty"`

		Properties interface{} `json:"properties,omitempty"`

		Telemetry interface{} `json:"telemetry,omitempty"`
	}

	dataAO1.CloudProperties = m.CloudProperties

	dataAO1.Properties = m.Properties

	dataAO1.Telemetry = m.Telemetry

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this webhook action device
func (m *WebhookActionDevice) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Device
	if err := m.Device.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WebhookActionDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhookActionDevice) UnmarshalBinary(b []byte) error {
	var res WebhookActionDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
