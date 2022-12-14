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

// ProjectServiceDeleteTokenReader is a Reader for the ProjectServiceDeleteToken structure.
type ProjectServiceDeleteTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectServiceDeleteTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectServiceDeleteTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectServiceDeleteTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectServiceDeleteTokenOK creates a ProjectServiceDeleteTokenOK with default headers values
func NewProjectServiceDeleteTokenOK() *ProjectServiceDeleteTokenOK {
	return &ProjectServiceDeleteTokenOK{}
}

/* ProjectServiceDeleteTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type ProjectServiceDeleteTokenOK struct {
	Payload models.ProjectEmptyResponse
}

func (o *ProjectServiceDeleteTokenOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project}/roles/{role}/token/{iat}][%d] projectServiceDeleteTokenOK  %+v", 200, o.Payload)
}
func (o *ProjectServiceDeleteTokenOK) GetPayload() models.ProjectEmptyResponse {
	return o.Payload
}

func (o *ProjectServiceDeleteTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectServiceDeleteTokenDefault creates a ProjectServiceDeleteTokenDefault with default headers values
func NewProjectServiceDeleteTokenDefault(code int) *ProjectServiceDeleteTokenDefault {
	return &ProjectServiceDeleteTokenDefault{
		_statusCode: code,
	}
}

/* ProjectServiceDeleteTokenDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProjectServiceDeleteTokenDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the project service delete token default response
func (o *ProjectServiceDeleteTokenDefault) Code() int {
	return o._statusCode
}

func (o *ProjectServiceDeleteTokenDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project}/roles/{role}/token/{iat}][%d] ProjectService_DeleteToken default  %+v", o._statusCode, o.Payload)
}
func (o *ProjectServiceDeleteTokenDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ProjectServiceDeleteTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
