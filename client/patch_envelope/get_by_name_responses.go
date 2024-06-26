// Code generated by go-swagger; DO NOT EDIT.

package patch_envelope

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// PatchEnvelopeConfigurationGetPatchEnvelopeByNameReader is a Reader for the PatchEnvelopeConfigurationGetPatchEnvelopeByName structure.
type PatchEnvelopeConfigurationGetPatchEnvelopeByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameOK creates a PatchEnvelopeConfigurationGetPatchEnvelopeByNameOK with default headers values
func NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameOK() *PatchEnvelopeConfigurationGetPatchEnvelopeByNameOK {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeByNameOK{}
}

/*
PatchEnvelopeConfigurationGetPatchEnvelopeByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchEnvelopeConfigurationGetPatchEnvelopeByNameOK struct {
	Payload *models.PatchEnvelope
}

// IsSuccess returns true when this patch envelope configuration get patch envelope by name o k response has a 2xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch envelope configuration get patch envelope by name o k response has a 3xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration get patch envelope by name o k response has a 4xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration get patch envelope by name o k response has a 5xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration get patch envelope by name o k response a status code equal to that given
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch envelope configuration get patch envelope by name o k response
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameOK) Code() int {
	return 200
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/name/{name}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByNameOK  %+v", 200, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/name/{name}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByNameOK  %+v", 200, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameOK) GetPayload() *models.PatchEnvelope {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PatchEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized creates a PatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized with default headers values
func NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized() *PatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized{}
}

/*
PatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type PatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration get patch envelope by name unauthorized response has a 2xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration get patch envelope by name unauthorized response has a 3xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration get patch envelope by name unauthorized response has a 4xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch envelope configuration get patch envelope by name unauthorized response has a 5xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration get patch envelope by name unauthorized response a status code equal to that given
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patch envelope configuration get patch envelope by name unauthorized response
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized) Code() int {
	return 401
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/name/{name}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/name/{name}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden creates a PatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden with default headers values
func NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden() *PatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden{}
}

/*
PatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type PatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration get patch envelope by name forbidden response has a 2xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration get patch envelope by name forbidden response has a 3xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration get patch envelope by name forbidden response has a 4xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch envelope configuration get patch envelope by name forbidden response has a 5xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration get patch envelope by name forbidden response a status code equal to that given
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch envelope configuration get patch envelope by name forbidden response
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden) Code() int {
	return 403
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/name/{name}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden  %+v", 403, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/name/{name}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden  %+v", 403, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound creates a PatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound with default headers values
func NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound() *PatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound{}
}

/*
PatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type PatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration get patch envelope by name not found response has a 2xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration get patch envelope by name not found response has a 3xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration get patch envelope by name not found response has a 4xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch envelope configuration get patch envelope by name not found response has a 5xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration get patch envelope by name not found response a status code equal to that given
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the patch envelope configuration get patch envelope by name not found response
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound) Code() int {
	return 404
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/name/{name}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound  %+v", 404, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/name/{name}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound  %+v", 404, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError creates a PatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError with default headers values
func NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError() *PatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError{}
}

/*
PatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type PatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration get patch envelope by name internal server error response has a 2xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration get patch envelope by name internal server error response has a 3xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration get patch envelope by name internal server error response has a 4xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration get patch envelope by name internal server error response has a 5xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch envelope configuration get patch envelope by name internal server error response a status code equal to that given
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the patch envelope configuration get patch envelope by name internal server error response
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError) Code() int {
	return 500
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/name/{name}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/name/{name}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout creates a PatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout with default headers values
func NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout() *PatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout{}
}

/*
PatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type PatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration get patch envelope by name gateway timeout response has a 2xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration get patch envelope by name gateway timeout response has a 3xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration get patch envelope by name gateway timeout response has a 4xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration get patch envelope by name gateway timeout response has a 5xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch envelope configuration get patch envelope by name gateway timeout response a status code equal to that given
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the patch envelope configuration get patch envelope by name gateway timeout response
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout) Code() int {
	return 504
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/name/{name}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/name/{name}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault creates a PatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault with default headers values
func NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault(code int) *PatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault{
		_statusCode: code,
	}
}

/*
PatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this patch envelope configuration get patch envelope by name default response has a 2xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch envelope configuration get patch envelope by name default response has a 3xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch envelope configuration get patch envelope by name default response has a 4xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch envelope configuration get patch envelope by name default response has a 5xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch envelope configuration get patch envelope by name default response a status code equal to that given
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch envelope configuration get patch envelope by name default response
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault) Code() int {
	return o._statusCode
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/name/{name}][%d] PatchEnvelopeConfiguration_GetPatchEnvelopeByName default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/name/{name}][%d] PatchEnvelopeConfiguration_GetPatchEnvelopeByName default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
