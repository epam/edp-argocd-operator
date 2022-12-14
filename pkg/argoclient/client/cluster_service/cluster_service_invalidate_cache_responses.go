// Code generated by go-swagger; DO NOT EDIT.

package cluster_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/epam/edp-argocd-operator/pkg/argoclient/models"
)

// ClusterServiceInvalidateCacheReader is a Reader for the ClusterServiceInvalidateCache structure.
type ClusterServiceInvalidateCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterServiceInvalidateCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterServiceInvalidateCacheOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterServiceInvalidateCacheDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterServiceInvalidateCacheOK creates a ClusterServiceInvalidateCacheOK with default headers values
func NewClusterServiceInvalidateCacheOK() *ClusterServiceInvalidateCacheOK {
	return &ClusterServiceInvalidateCacheOK{}
}

/* ClusterServiceInvalidateCacheOK describes a response with status code 200, with default header values.

A successful response.
*/
type ClusterServiceInvalidateCacheOK struct {
	Payload *models.V1alpha1Cluster
}

func (o *ClusterServiceInvalidateCacheOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/clusters/{id.value}/invalidate-cache][%d] clusterServiceInvalidateCacheOK  %+v", 200, o.Payload)
}
func (o *ClusterServiceInvalidateCacheOK) GetPayload() *models.V1alpha1Cluster {
	return o.Payload
}

func (o *ClusterServiceInvalidateCacheOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterServiceInvalidateCacheDefault creates a ClusterServiceInvalidateCacheDefault with default headers values
func NewClusterServiceInvalidateCacheDefault(code int) *ClusterServiceInvalidateCacheDefault {
	return &ClusterServiceInvalidateCacheDefault{
		_statusCode: code,
	}
}

/* ClusterServiceInvalidateCacheDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ClusterServiceInvalidateCacheDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the cluster service invalidate cache default response
func (o *ClusterServiceInvalidateCacheDefault) Code() int {
	return o._statusCode
}

func (o *ClusterServiceInvalidateCacheDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/clusters/{id.value}/invalidate-cache][%d] ClusterService_InvalidateCache default  %+v", o._statusCode, o.Payload)
}
func (o *ClusterServiceInvalidateCacheDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ClusterServiceInvalidateCacheDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
