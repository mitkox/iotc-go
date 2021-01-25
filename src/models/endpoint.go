// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Endpoint endpoint
//
// swagger:discriminator Endpoint type
type Endpoint interface {
	runtime.Validatable

	// Information for connecting to the endpoint.
	// Required: true
	ConnectionString() *string
	SetConnectionString(*string)

	// Name of the entity to send data to.
	// Required: true
	Name() *string
	SetName(*string)

	// Type of the endpoint.
	// Required: true
	Type() string
	SetType(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type endpoint struct {
	connectionStringField *string

	nameField *string

	typeField string
}

// ConnectionString gets the connection string of this polymorphic type
func (m *endpoint) ConnectionString() *string {
	return m.connectionStringField
}

// SetConnectionString sets the connection string of this polymorphic type
func (m *endpoint) SetConnectionString(val *string) {
	m.connectionStringField = val
}

// Name gets the name of this polymorphic type
func (m *endpoint) Name() *string {
	return m.nameField
}

// SetName sets the name of this polymorphic type
func (m *endpoint) SetName(val *string) {
	m.nameField = val
}

// Type gets the type of this polymorphic type
func (m *endpoint) Type() string {
	return "Endpoint"
}

// SetType sets the type of this polymorphic type
func (m *endpoint) SetType(val string) {
}

// UnmarshalEndpointSlice unmarshals polymorphic slices of Endpoint
func UnmarshalEndpointSlice(reader io.Reader, consumer runtime.Consumer) ([]Endpoint, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Endpoint
	for _, element := range elements {
		obj, err := unmarshalEndpoint(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalEndpoint unmarshals polymorphic Endpoint
func UnmarshalEndpoint(reader io.Reader, consumer runtime.Consumer) (Endpoint, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalEndpoint(data, consumer)
}

func unmarshalEndpoint(data []byte, consumer runtime.Consumer) (Endpoint, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var getType struct {
		Type string `json:"type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("type", "body", getType.Type); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch getType.Type {
	case "Endpoint":
		var result endpoint
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "EventHubsEndpoint":
		var result EventHubsEndpoint
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "ServiceBusQueueEndpoint":
		var result ServiceBusQueueEndpoint
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "ServiceBusTopicEndpoint":
		var result ServiceBusTopicEndpoint
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "StorageEndpoint":
		var result StorageEndpoint
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid type value: %q", getType.Type)
}

// Validate validates this endpoint
func (m *endpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionString(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *endpoint) validateConnectionString(formats strfmt.Registry) error {

	if err := validate.Required("connectionString", "body", m.ConnectionString()); err != nil {
		return err
	}

	return nil
}

func (m *endpoint) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}