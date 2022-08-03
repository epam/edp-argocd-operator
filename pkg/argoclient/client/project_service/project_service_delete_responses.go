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

// ProjectServiceDeleteReader is a Reader for the ProjectServiceDelete structure.
type ProjectServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectServiceDeleteOK creates a ProjectServiceDeleteOK with default headers values
func NewProjectServiceDeleteOK() *ProjectServiceDeleteOK {
	return &ProjectServiceDeleteOK{}
}

/* ProjectServiceDeleteOK describes a response with status code 200, with default header values.

A successful response.
*/
type ProjectServiceDeleteOK struct {
	Payload models.ProjectEmptyResponse
}

func (o *ProjectServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{name}][%d] projectServiceDeleteOK  %+v", 200, o.Payload)
}
func (o *ProjectServiceDeleteOK) GetPayload() models.ProjectEmptyResponse {
	return o.Payload
}

func (o *ProjectServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectServiceDeleteDefault creates a ProjectServiceDeleteDefault with default headers values
func NewProjectServiceDeleteDefault(code int) *ProjectServiceDeleteDefault {
	return &ProjectServiceDeleteDefault{
		_statusCode: code,
	}
}

/* ProjectServiceDeleteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProjectServiceDeleteDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the project service delete default response
func (o *ProjectServiceDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ProjectServiceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{name}][%d] ProjectService_Delete default  %+v", o._statusCode, o.Payload)
}
func (o *ProjectServiceDeleteDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ProjectServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}