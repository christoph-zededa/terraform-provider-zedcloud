package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// EdgeNodeConfigurationGetEdgeNodeAttestationReader is a Reader for the EdgeNodeConfigurationGetEdgeNodeAttestation structure.
type EdgeNodeConfigurationGetEdgeNodeAttestationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationGetEdgeNodeAttestationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationGetEdgeNodeAttestationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationGetEdgeNodeAttestationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationGetEdgeNodeAttestationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationGetEdgeNodeAttestationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationGetEdgeNodeAttestationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationGetEdgeNodeAttestationOK creates a EdgeNodeConfigurationGetEdgeNodeAttestationOK with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeAttestationOK() *EdgeNodeConfigurationGetEdgeNodeAttestationOK {
	return &EdgeNodeConfigurationGetEdgeNodeAttestationOK{}
}

/*
EdgeNodeConfigurationGetEdgeNodeAttestationOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationGetEdgeNodeAttestationOK struct {
	Payload *models.Node
}

// IsSuccess returns true when this edge node configuration get edge node attestation o k response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node configuration get edge node attestation o k response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node attestation o k response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration get edge node attestation o k response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get edge node attestation o k response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/attestation][%d] edgeNodeConfigurationGetEdgeNodeAttestationOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationOK) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/attestation][%d] edgeNodeConfigurationGetEdgeNodeAttestationOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationOK) GetPayload() *models.Node {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeAttestationUnauthorized creates a EdgeNodeConfigurationGetEdgeNodeAttestationUnauthorized with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeAttestationUnauthorized() *EdgeNodeConfigurationGetEdgeNodeAttestationUnauthorized {
	return &EdgeNodeConfigurationGetEdgeNodeAttestationUnauthorized{}
}

/*
EdgeNodeConfigurationGetEdgeNodeAttestationUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationGetEdgeNodeAttestationUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node attestation unauthorized response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node attestation unauthorized response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node attestation unauthorized response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration get edge node attestation unauthorized response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get edge node attestation unauthorized response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/attestation][%d] edgeNodeConfigurationGetEdgeNodeAttestationUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/attestation][%d] edgeNodeConfigurationGetEdgeNodeAttestationUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeAttestationForbidden creates a EdgeNodeConfigurationGetEdgeNodeAttestationForbidden with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeAttestationForbidden() *EdgeNodeConfigurationGetEdgeNodeAttestationForbidden {
	return &EdgeNodeConfigurationGetEdgeNodeAttestationForbidden{}
}

/*
EdgeNodeConfigurationGetEdgeNodeAttestationForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationGetEdgeNodeAttestationForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node attestation forbidden response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node attestation forbidden response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node attestation forbidden response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration get edge node attestation forbidden response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get edge node attestation forbidden response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/attestation][%d] edgeNodeConfigurationGetEdgeNodeAttestationForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationForbidden) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/attestation][%d] edgeNodeConfigurationGetEdgeNodeAttestationForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeAttestationNotFound creates a EdgeNodeConfigurationGetEdgeNodeAttestationNotFound with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeAttestationNotFound() *EdgeNodeConfigurationGetEdgeNodeAttestationNotFound {
	return &EdgeNodeConfigurationGetEdgeNodeAttestationNotFound{}
}

/*
EdgeNodeConfigurationGetEdgeNodeAttestationNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationGetEdgeNodeAttestationNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node attestation not found response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node attestation not found response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node attestation not found response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration get edge node attestation not found response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get edge node attestation not found response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/attestation][%d] edgeNodeConfigurationGetEdgeNodeAttestationNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationNotFound) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/attestation][%d] edgeNodeConfigurationGetEdgeNodeAttestationNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeAttestationInternalServerError creates a EdgeNodeConfigurationGetEdgeNodeAttestationInternalServerError with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeAttestationInternalServerError() *EdgeNodeConfigurationGetEdgeNodeAttestationInternalServerError {
	return &EdgeNodeConfigurationGetEdgeNodeAttestationInternalServerError{}
}

/*
EdgeNodeConfigurationGetEdgeNodeAttestationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationGetEdgeNodeAttestationInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node attestation internal server error response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node attestation internal server error response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node attestation internal server error response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration get edge node attestation internal server error response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration get edge node attestation internal server error response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/attestation][%d] edgeNodeConfigurationGetEdgeNodeAttestationInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/attestation][%d] edgeNodeConfigurationGetEdgeNodeAttestationInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout creates a EdgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout() *EdgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout {
	return &EdgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout{}
}

/*
EdgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node attestation gateway timeout response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node attestation gateway timeout response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node attestation gateway timeout response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration get edge node attestation gateway timeout response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration get edge node attestation gateway timeout response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/attestation][%d] edgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/attestation][%d] edgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeAttestationDefault creates a EdgeNodeConfigurationGetEdgeNodeAttestationDefault with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeAttestationDefault(code int) *EdgeNodeConfigurationGetEdgeNodeAttestationDefault {
	return &EdgeNodeConfigurationGetEdgeNodeAttestationDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeConfigurationGetEdgeNodeAttestationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationGetEdgeNodeAttestationDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration get edge node attestation default response
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node configuration get edge node attestation default response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node configuration get edge node attestation default response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node configuration get edge node attestation default response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node configuration get edge node attestation default response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node configuration get edge node attestation default response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeAttestationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/attestation][%d] EdgeNodeConfiguration_GetEdgeNodeAttestation default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationDefault) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/attestation][%d] EdgeNodeConfiguration_GetEdgeNodeAttestation default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeAttestationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
