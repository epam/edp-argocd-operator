// Code generated by go-swagger; DO NOT EDIT.

package application_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/epam/edp-argocd-operator/pkg/argoclient/models"
)

// ApplicationServiceListResourceEventsReader is a Reader for the ApplicationServiceListResourceEvents structure.
type ApplicationServiceListResourceEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationServiceListResourceEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationServiceListResourceEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationServiceListResourceEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationServiceListResourceEventsOK creates a ApplicationServiceListResourceEventsOK with default headers values
func NewApplicationServiceListResourceEventsOK() *ApplicationServiceListResourceEventsOK {
	return &ApplicationServiceListResourceEventsOK{}
}

/* ApplicationServiceListResourceEventsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ApplicationServiceListResourceEventsOK struct {
	Payload *models.V1EventList
}

func (o *ApplicationServiceListResourceEventsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/applications/{name}/events][%d] applicationServiceListResourceEventsOK  %+v", 200, o.Payload)
}
func (o *ApplicationServiceListResourceEventsOK) GetPayload() *models.V1EventList {
	return o.Payload
}

func (o *ApplicationServiceListResourceEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1EventList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationServiceListResourceEventsDefault creates a ApplicationServiceListResourceEventsDefault with default headers values
func NewApplicationServiceListResourceEventsDefault(code int) *ApplicationServiceListResourceEventsDefault {
	return &ApplicationServiceListResourceEventsDefault{
		_statusCode: code,
	}
}

/* ApplicationServiceListResourceEventsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ApplicationServiceListResourceEventsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the application service list resource events default response
func (o *ApplicationServiceListResourceEventsDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationServiceListResourceEventsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/applications/{name}/events][%d] ApplicationService_ListResourceEvents default  %+v", o._statusCode, o.Payload)
}
func (o *ApplicationServiceListResourceEventsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ApplicationServiceListResourceEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}