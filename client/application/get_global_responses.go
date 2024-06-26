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

// EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleReader is a Reader for the EdgeApplicationConfigurationGetGlobalEdgeApplicationBundle structure.
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK creates a EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK with default headers values
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK() *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK{}
}

/*
EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK struct {
	Payload *models.Application
}

// IsSuccess returns true when this edge application configuration get global edge application bundle o k response has a 2xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge application configuration get global edge application bundle o k response has a 3xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get global edge application bundle o k response has a 4xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration get global edge application bundle o k response has a 5xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration get global edge application bundle o k response a status code equal to that given
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge application configuration get global edge application bundle o k response
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK) Code() int {
	return 200
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global/id/{id}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/global/id/{id}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK) GetPayload() *models.Application {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized creates a EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized with default headers values
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized() *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized{}
}

/*
EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get global edge application bundle unauthorized response has a 2xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get global edge application bundle unauthorized response has a 3xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get global edge application bundle unauthorized response has a 4xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration get global edge application bundle unauthorized response has a 5xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration get global edge application bundle unauthorized response a status code equal to that given
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge application configuration get global edge application bundle unauthorized response
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized) Code() int {
	return 401
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global/id/{id}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/global/id/{id}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden creates a EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden with default headers values
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden() *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden{}
}

/*
EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get global edge application bundle forbidden response has a 2xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get global edge application bundle forbidden response has a 3xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get global edge application bundle forbidden response has a 4xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration get global edge application bundle forbidden response has a 5xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration get global edge application bundle forbidden response a status code equal to that given
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge application configuration get global edge application bundle forbidden response
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden) Code() int {
	return 403
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global/id/{id}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/global/id/{id}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound creates a EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound with default headers values
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound() *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound{}
}

/*
EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get global edge application bundle not found response has a 2xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get global edge application bundle not found response has a 3xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get global edge application bundle not found response has a 4xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration get global edge application bundle not found response has a 5xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration get global edge application bundle not found response a status code equal to that given
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge application configuration get global edge application bundle not found response
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound) Code() int {
	return 404
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global/id/{id}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/global/id/{id}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError creates a EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError with default headers values
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError() *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError{}
}

/*
EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get global edge application bundle internal server error response has a 2xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get global edge application bundle internal server error response has a 3xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get global edge application bundle internal server error response has a 4xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration get global edge application bundle internal server error response has a 5xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application configuration get global edge application bundle internal server error response a status code equal to that given
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge application configuration get global edge application bundle internal server error response
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError) Code() int {
	return 500
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global/id/{id}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/global/id/{id}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout creates a EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout with default headers values
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout() *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout{}
}

/*
EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get global edge application bundle gateway timeout response has a 2xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get global edge application bundle gateway timeout response has a 3xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get global edge application bundle gateway timeout response has a 4xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration get global edge application bundle gateway timeout response has a 5xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application configuration get global edge application bundle gateway timeout response a status code equal to that given
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge application configuration get global edge application bundle gateway timeout response
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global/id/{id}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps/global/id/{id}][%d] edgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault creates a EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault with default headers values
func NewEdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault(code int) *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault {
	return &EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault{
		_statusCode: code,
	}
}

/*
EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge application configuration get global edge application bundle default response has a 2xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge application configuration get global edge application bundle default response has a 3xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge application configuration get global edge application bundle default response has a 4xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge application configuration get global edge application bundle default response has a 5xx status code
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge application configuration get global edge application bundle default response a status code equal to that given
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge application configuration get global edge application bundle default response
func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global/id/{id}][%d] EdgeApplicationConfiguration_GetGlobalEdgeApplicationBundle default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps/global/id/{id}][%d] EdgeApplicationConfiguration_GetGlobalEdgeApplicationBundle default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetGlobalEdgeApplicationBundleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
