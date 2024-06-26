package application_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceReader is a Reader for the EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstance structure.
type EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK creates a EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK with default headers values
func NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK() *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK {
	return &EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK{}
}

/*
EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration de activate edge application instance o k response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge application instance configuration de activate edge application instance o k response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration de activate edge application instance o k response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration de activate edge application instance o k response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration de activate edge application instance o k response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge application instance configuration de activate edge application instance o k response
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK) Code() int {
	return 200
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/deactivate][%d] edgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK) String() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/deactivate][%d] edgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized creates a EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized with default headers values
func NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized() *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized {
	return &EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized{}
}

/*
EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration de activate edge application instance unauthorized response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration de activate edge application instance unauthorized response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration de activate edge application instance unauthorized response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration de activate edge application instance unauthorized response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration de activate edge application instance unauthorized response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge application instance configuration de activate edge application instance unauthorized response
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized) Code() int {
	return 401
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/deactivate][%d] edgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/deactivate][%d] edgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden creates a EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden with default headers values
func NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden() *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden {
	return &EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden{}
}

/*
EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration de activate edge application instance forbidden response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration de activate edge application instance forbidden response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration de activate edge application instance forbidden response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration de activate edge application instance forbidden response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration de activate edge application instance forbidden response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge application instance configuration de activate edge application instance forbidden response
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden) Code() int {
	return 403
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/deactivate][%d] edgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/deactivate][%d] edgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound creates a EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound with default headers values
func NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound() *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound {
	return &EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound{}
}

/*
EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration de activate edge application instance not found response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration de activate edge application instance not found response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration de activate edge application instance not found response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration de activate edge application instance not found response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration de activate edge application instance not found response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge application instance configuration de activate edge application instance not found response
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound) Code() int {
	return 404
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/deactivate][%d] edgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/deactivate][%d] edgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict creates a EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict with default headers values
func NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict() *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict {
	return &EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict{}
}

/*
EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge network record.
*/
type EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration de activate edge application instance conflict response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration de activate edge application instance conflict response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration de activate edge application instance conflict response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration de activate edge application instance conflict response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration de activate edge application instance conflict response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the edge application instance configuration de activate edge application instance conflict response
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict) Code() int {
	return 409
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/deactivate][%d] edgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict  %+v", 409, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict) String() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/deactivate][%d] edgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict  %+v", 409, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError creates a EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError with default headers values
func NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError() *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError {
	return &EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError{}
}

/*
EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration de activate edge application instance internal server error response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration de activate edge application instance internal server error response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration de activate edge application instance internal server error response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration de activate edge application instance internal server error response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance configuration de activate edge application instance internal server error response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge application instance configuration de activate edge application instance internal server error response
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError) Code() int {
	return 500
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/deactivate][%d] edgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/deactivate][%d] edgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout creates a EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout with default headers values
func NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout() *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout {
	return &EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout{}
}

/*
EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration de activate edge application instance gateway timeout response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration de activate edge application instance gateway timeout response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration de activate edge application instance gateway timeout response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration de activate edge application instance gateway timeout response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance configuration de activate edge application instance gateway timeout response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge application instance configuration de activate edge application instance gateway timeout response
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/deactivate][%d] edgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/deactivate][%d] edgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault creates a EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault with default headers values
func NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault(code int) *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault {
	return &EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault{
		_statusCode: code,
	}
}

/*
EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge application instance configuration de activate edge application instance default response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge application instance configuration de activate edge application instance default response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge application instance configuration de activate edge application instance default response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge application instance configuration de activate edge application instance default response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge application instance configuration de activate edge application instance default response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge application instance configuration de activate edge application instance default response
func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/deactivate][%d] EdgeApplicationInstanceConfiguration_DeActivateEdgeApplicationInstance default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault) String() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}/deactivate][%d] EdgeApplicationInstanceConfiguration_DeActivateEdgeApplicationInstance default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
