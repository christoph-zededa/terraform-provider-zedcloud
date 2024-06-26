package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameReader is a Reader for the EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByName structure.
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK creates a EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK with default headers values
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK() *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK{}
}

/*
EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK struct {
	Payload *models.Application
}

// IsSuccess returns true when this edge application configuration get global edge application bundle by name o k response has a 2xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge application configuration get global edge application bundle by name o k response has a 3xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get global edge application bundle by name o k response has a 4xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration get global edge application bundle by name o k response has a 5xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration get global edge application bundle by name o k response a status code equal to that given
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge application configuration get global edge application bundle by name o k response
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK) Code() int {
	return 200
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global/name/{name}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/global/name/{name}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK) GetPayload() *models.Application {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized creates a EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized with default headers values
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized() *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized{}
}

/*
EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get global edge application bundle by name unauthorized response has a 2xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get global edge application bundle by name unauthorized response has a 3xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get global edge application bundle by name unauthorized response has a 4xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration get global edge application bundle by name unauthorized response has a 5xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration get global edge application bundle by name unauthorized response a status code equal to that given
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge application configuration get global edge application bundle by name unauthorized response
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized) Code() int {
	return 401
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global/name/{name}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/global/name/{name}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden creates a EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden with default headers values
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden() *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden{}
}

/*
EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get global edge application bundle by name forbidden response has a 2xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get global edge application bundle by name forbidden response has a 3xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get global edge application bundle by name forbidden response has a 4xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration get global edge application bundle by name forbidden response has a 5xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration get global edge application bundle by name forbidden response a status code equal to that given
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge application configuration get global edge application bundle by name forbidden response
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden) Code() int {
	return 403
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global/name/{name}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/global/name/{name}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound creates a EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound with default headers values
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound() *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound{}
}

/*
EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get global edge application bundle by name not found response has a 2xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get global edge application bundle by name not found response has a 3xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get global edge application bundle by name not found response has a 4xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration get global edge application bundle by name not found response has a 5xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration get global edge application bundle by name not found response a status code equal to that given
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge application configuration get global edge application bundle by name not found response
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound) Code() int {
	return 404
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global/name/{name}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/global/name/{name}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError creates a EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError with default headers values
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError() *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError{}
}

/*
EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get global edge application bundle by name internal server error response has a 2xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get global edge application bundle by name internal server error response has a 3xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get global edge application bundle by name internal server error response has a 4xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration get global edge application bundle by name internal server error response has a 5xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application configuration get global edge application bundle by name internal server error response a status code equal to that given
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge application configuration get global edge application bundle by name internal server error response
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError) Code() int {
	return 500
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global/name/{name}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/global/name/{name}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout creates a EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout with default headers values
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout() *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout{}
}

/*
EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get global edge application bundle by name gateway timeout response has a 2xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get global edge application bundle by name gateway timeout response has a 3xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get global edge application bundle by name gateway timeout response has a 4xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration get global edge application bundle by name gateway timeout response has a 5xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application configuration get global edge application bundle by name gateway timeout response a status code equal to that given
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge application configuration get global edge application bundle by name gateway timeout response
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global/name/{name}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps/global/name/{name}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault creates a EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault with default headers values
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault(code int) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault{
		_statusCode: code,
	}
}

/*
EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge application configuration get global edge application bundle by name default response has a 2xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge application configuration get global edge application bundle by name default response has a 3xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge application configuration get global edge application bundle by name default response has a 4xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge application configuration get global edge application bundle by name default response has a 5xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge application configuration get global edge application bundle by name default response a status code equal to that given
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge application configuration get global edge application bundle by name default response
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global/name/{name}][%d] EdgeApplicationConfiguration_GetGlobalEdgeApplicationBundleByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps/global/name/{name}][%d] EdgeApplicationConfiguration_GetGlobalEdgeApplicationBundleByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
