// Code generated by go-swagger; DO NOT EDIT.

package project_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/epam/edp-argocd-operator/pkg/argoclient/models"
)

// ProjectServiceListEventsReader is a Reader for the ProjectServiceListEvents structure.
type ProjectServiceListEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectServiceListEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectServiceListEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectServiceListEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectServiceListEventsOK creates a ProjectServiceListEventsOK with default headers values
func NewProjectServiceListEventsOK() *ProjectServiceListEventsOK {
	return &ProjectServiceListEventsOK{}
}

/* ProjectServiceListEventsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ProjectServiceListEventsOK struct {
	Payload *models.V1EventList
}

func (o *ProjectServiceListEventsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{name}/events][%d] projectServiceListEventsOK  %+v", 200, o.Payload)
}
func (o *ProjectServiceListEventsOK) GetPayload() *models.V1EventList {
	return o.Payload
}

func (o *ProjectServiceListEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1EventList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectServiceListEventsDefault creates a ProjectServiceListEventsDefault with default headers values
func NewProjectServiceListEventsDefault(code int) *ProjectServiceListEventsDefault {
	return &ProjectServiceListEventsDefault{
		_statusCode: code,
	}
}

/* ProjectServiceListEventsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProjectServiceListEventsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the project service list events default response
func (o *ProjectServiceListEventsDefault) Code() int {
	return o._statusCode
}

func (o *ProjectServiceListEventsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{name}/events][%d] ProjectService_ListEvents default  %+v", o._statusCode, o.Payload)
}
func (o *ProjectServiceListEventsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ProjectServiceListEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
