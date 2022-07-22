// Code generated by go-swagger; DO NOT EDIT.

package repository_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/epam/edp-argocd-operator/pkg/argoclient/models"
)

// RepositoryServiceListRefsReader is a Reader for the RepositoryServiceListRefs structure.
type RepositoryServiceListRefsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepositoryServiceListRefsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepositoryServiceListRefsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRepositoryServiceListRefsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRepositoryServiceListRefsOK creates a RepositoryServiceListRefsOK with default headers values
func NewRepositoryServiceListRefsOK() *RepositoryServiceListRefsOK {
	return &RepositoryServiceListRefsOK{}
}

/* RepositoryServiceListRefsOK describes a response with status code 200, with default header values.

A successful response.
*/
type RepositoryServiceListRefsOK struct {
	Payload *models.RepositoryRefs
}

func (o *RepositoryServiceListRefsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repositories/{repo}/refs][%d] repositoryServiceListRefsOK  %+v", 200, o.Payload)
}
func (o *RepositoryServiceListRefsOK) GetPayload() *models.RepositoryRefs {
	return o.Payload
}

func (o *RepositoryServiceListRefsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RepositoryRefs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryServiceListRefsDefault creates a RepositoryServiceListRefsDefault with default headers values
func NewRepositoryServiceListRefsDefault(code int) *RepositoryServiceListRefsDefault {
	return &RepositoryServiceListRefsDefault{
		_statusCode: code,
	}
}

/* RepositoryServiceListRefsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RepositoryServiceListRefsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the repository service list refs default response
func (o *RepositoryServiceListRefsDefault) Code() int {
	return o._statusCode
}

func (o *RepositoryServiceListRefsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/repositories/{repo}/refs][%d] RepositoryService_ListRefs default  %+v", o._statusCode, o.Payload)
}
func (o *RepositoryServiceListRefsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *RepositoryServiceListRefsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
