// Code generated by go-swagger; DO NOT EDIT.

package account_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/epam/edp-argocd-operator/pkg/argoclient/models"
)

// AccountServiceCreateTokenReader is a Reader for the AccountServiceCreateToken structure.
type AccountServiceCreateTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountServiceCreateTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountServiceCreateTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAccountServiceCreateTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountServiceCreateTokenOK creates a AccountServiceCreateTokenOK with default headers values
func NewAccountServiceCreateTokenOK() *AccountServiceCreateTokenOK {
	return &AccountServiceCreateTokenOK{}
}

/* AccountServiceCreateTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type AccountServiceCreateTokenOK struct {
	Payload *models.AccountCreateTokenResponse
}

func (o *AccountServiceCreateTokenOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/account/{name}/token][%d] accountServiceCreateTokenOK  %+v", 200, o.Payload)
}
func (o *AccountServiceCreateTokenOK) GetPayload() *models.AccountCreateTokenResponse {
	return o.Payload
}

func (o *AccountServiceCreateTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountCreateTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountServiceCreateTokenDefault creates a AccountServiceCreateTokenDefault with default headers values
func NewAccountServiceCreateTokenDefault(code int) *AccountServiceCreateTokenDefault {
	return &AccountServiceCreateTokenDefault{
		_statusCode: code,
	}
}

/* AccountServiceCreateTokenDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AccountServiceCreateTokenDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the account service create token default response
func (o *AccountServiceCreateTokenDefault) Code() int {
	return o._statusCode
}

func (o *AccountServiceCreateTokenDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/account/{name}/token][%d] AccountService_CreateToken default  %+v", o._statusCode, o.Payload)
}
func (o *AccountServiceCreateTokenDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AccountServiceCreateTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
