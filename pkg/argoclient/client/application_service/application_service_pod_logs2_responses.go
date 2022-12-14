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

// ApplicationServicePodLogs2Reader is a Reader for the ApplicationServicePodLogs2 structure.
type ApplicationServicePodLogs2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationServicePodLogs2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationServicePodLogs2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationServicePodLogs2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationServicePodLogs2OK creates a ApplicationServicePodLogs2OK with default headers values
func NewApplicationServicePodLogs2OK() *ApplicationServicePodLogs2OK {
	return &ApplicationServicePodLogs2OK{}
}

/* ApplicationServicePodLogs2OK describes a response with status code 200, with default header values.

A successful response.(streaming responses)
*/
type ApplicationServicePodLogs2OK struct {
	Payload *models.ApplicationServicePodLogs2OKBody
}

func (o *ApplicationServicePodLogs2OK) Error() string {
	return fmt.Sprintf("[GET /api/v1/applications/{name}/logs][%d] applicationServicePodLogs2OK  %+v", 200, o.Payload)
}
func (o *ApplicationServicePodLogs2OK) GetPayload() *models.ApplicationServicePodLogs2OKBody {
	return o.Payload
}

func (o *ApplicationServicePodLogs2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationServicePodLogs2OKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationServicePodLogs2Default creates a ApplicationServicePodLogs2Default with default headers values
func NewApplicationServicePodLogs2Default(code int) *ApplicationServicePodLogs2Default {
	return &ApplicationServicePodLogs2Default{
		_statusCode: code,
	}
}

/* ApplicationServicePodLogs2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ApplicationServicePodLogs2Default struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the application service pod logs2 default response
func (o *ApplicationServicePodLogs2Default) Code() int {
	return o._statusCode
}

func (o *ApplicationServicePodLogs2Default) Error() string {
	return fmt.Sprintf("[GET /api/v1/applications/{name}/logs][%d] ApplicationService_PodLogs2 default  %+v", o._statusCode, o.Payload)
}
func (o *ApplicationServicePodLogs2Default) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ApplicationServicePodLogs2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
