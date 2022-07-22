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

// ApplicationServiceRunResourceActionReader is a Reader for the ApplicationServiceRunResourceAction structure.
type ApplicationServiceRunResourceActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationServiceRunResourceActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationServiceRunResourceActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationServiceRunResourceActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationServiceRunResourceActionOK creates a ApplicationServiceRunResourceActionOK with default headers values
func NewApplicationServiceRunResourceActionOK() *ApplicationServiceRunResourceActionOK {
	return &ApplicationServiceRunResourceActionOK{}
}

/* ApplicationServiceRunResourceActionOK describes a response with status code 200, with default header values.

A successful response.
*/
type ApplicationServiceRunResourceActionOK struct {
	Payload models.ApplicationApplicationResponse
}

func (o *ApplicationServiceRunResourceActionOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/applications/{name}/resource/actions][%d] applicationServiceRunResourceActionOK  %+v", 200, o.Payload)
}
func (o *ApplicationServiceRunResourceActionOK) GetPayload() models.ApplicationApplicationResponse {
	return o.Payload
}

func (o *ApplicationServiceRunResourceActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationServiceRunResourceActionDefault creates a ApplicationServiceRunResourceActionDefault with default headers values
func NewApplicationServiceRunResourceActionDefault(code int) *ApplicationServiceRunResourceActionDefault {
	return &ApplicationServiceRunResourceActionDefault{
		_statusCode: code,
	}
}

/* ApplicationServiceRunResourceActionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ApplicationServiceRunResourceActionDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the application service run resource action default response
func (o *ApplicationServiceRunResourceActionDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationServiceRunResourceActionDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/applications/{name}/resource/actions][%d] ApplicationService_RunResourceAction default  %+v", o._statusCode, o.Payload)
}
func (o *ApplicationServiceRunResourceActionDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ApplicationServiceRunResourceActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
