// Code generated by go-swagger; DO NOT EDIT.

package g_p_g_key_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/epam/edp-argocd-operator/pkg/argoclient/models"
)

// GPGKeyServiceGetReader is a Reader for the GPGKeyServiceGet structure.
type GPGKeyServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GPGKeyServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGPGKeyServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGPGKeyServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGPGKeyServiceGetOK creates a GPGKeyServiceGetOK with default headers values
func NewGPGKeyServiceGetOK() *GPGKeyServiceGetOK {
	return &GPGKeyServiceGetOK{}
}

/* GPGKeyServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type GPGKeyServiceGetOK struct {
	Payload *models.V1alpha1GnuPGPublicKey
}

func (o *GPGKeyServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/gpgkeys/{keyID}][%d] gPGKeyServiceGetOK  %+v", 200, o.Payload)
}
func (o *GPGKeyServiceGetOK) GetPayload() *models.V1alpha1GnuPGPublicKey {
	return o.Payload
}

func (o *GPGKeyServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1GnuPGPublicKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGPGKeyServiceGetDefault creates a GPGKeyServiceGetDefault with default headers values
func NewGPGKeyServiceGetDefault(code int) *GPGKeyServiceGetDefault {
	return &GPGKeyServiceGetDefault{
		_statusCode: code,
	}
}

/* GPGKeyServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GPGKeyServiceGetDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the g p g key service get default response
func (o *GPGKeyServiceGetDefault) Code() int {
	return o._statusCode
}

func (o *GPGKeyServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/gpgkeys/{keyID}][%d] GPGKeyService_Get default  %+v", o._statusCode, o.Payload)
}
func (o *GPGKeyServiceGetDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *GPGKeyServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
