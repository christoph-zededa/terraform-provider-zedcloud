package hardware_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// HardwareModelGetGlobalHardwareModelReader is a Reader for the HardwareModelGetGlobalHardwareModel structure.
type HardwareModelGetGlobalHardwareModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelGetGlobalHardwareModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelGetGlobalHardwareModelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHardwareModelGetGlobalHardwareModelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelGetGlobalHardwareModelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHardwareModelGetGlobalHardwareModelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelGetGlobalHardwareModelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelGetGlobalHardwareModelGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelGetGlobalHardwareModelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelGetGlobalHardwareModelOK creates a HardwareModelGetGlobalHardwareModelOK with default headers values
func NewHardwareModelGetGlobalHardwareModelOK() *HardwareModelGetGlobalHardwareModelOK {
	return &HardwareModelGetGlobalHardwareModelOK{}
}

/*
HardwareModelGetGlobalHardwareModelOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelGetGlobalHardwareModelOK struct {
	Payload *models.SysModel
}

// IsSuccess returns true when this hardware model get global hardware model o k response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareModelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hardware model get global hardware model o k response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareModelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware model o k response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareModelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get global hardware model o k response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareModelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get global hardware model o k response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareModelOK) IsCode(code int) bool {
	return code == 200
}

func (o *HardwareModelGetGlobalHardwareModelOK) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] hardwareModelGetGlobalHardwareModelOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelOK) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] hardwareModelGetGlobalHardwareModelOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelOK) GetPayload() *models.SysModel {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareModelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SysModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareModelUnauthorized creates a HardwareModelGetGlobalHardwareModelUnauthorized with default headers values
func NewHardwareModelGetGlobalHardwareModelUnauthorized() *HardwareModelGetGlobalHardwareModelUnauthorized {
	return &HardwareModelGetGlobalHardwareModelUnauthorized{}
}

/*
HardwareModelGetGlobalHardwareModelUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelGetGlobalHardwareModelUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware model unauthorized response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareModelUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware model unauthorized response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareModelUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware model unauthorized response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareModelUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get global hardware model unauthorized response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareModelUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get global hardware model unauthorized response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareModelUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *HardwareModelGetGlobalHardwareModelUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] hardwareModelGetGlobalHardwareModelUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] hardwareModelGetGlobalHardwareModelUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareModelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareModelForbidden creates a HardwareModelGetGlobalHardwareModelForbidden with default headers values
func NewHardwareModelGetGlobalHardwareModelForbidden() *HardwareModelGetGlobalHardwareModelForbidden {
	return &HardwareModelGetGlobalHardwareModelForbidden{}
}

/*
HardwareModelGetGlobalHardwareModelForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelGetGlobalHardwareModelForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware model forbidden response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareModelForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware model forbidden response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareModelForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware model forbidden response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareModelForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get global hardware model forbidden response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareModelForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get global hardware model forbidden response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareModelForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *HardwareModelGetGlobalHardwareModelForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] hardwareModelGetGlobalHardwareModelForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelForbidden) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] hardwareModelGetGlobalHardwareModelForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareModelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareModelNotFound creates a HardwareModelGetGlobalHardwareModelNotFound with default headers values
func NewHardwareModelGetGlobalHardwareModelNotFound() *HardwareModelGetGlobalHardwareModelNotFound {
	return &HardwareModelGetGlobalHardwareModelNotFound{}
}

/*
HardwareModelGetGlobalHardwareModelNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type HardwareModelGetGlobalHardwareModelNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware model not found response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareModelNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware model not found response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareModelNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware model not found response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareModelNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get global hardware model not found response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareModelNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get global hardware model not found response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareModelNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *HardwareModelGetGlobalHardwareModelNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] hardwareModelGetGlobalHardwareModelNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelNotFound) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] hardwareModelGetGlobalHardwareModelNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareModelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareModelInternalServerError creates a HardwareModelGetGlobalHardwareModelInternalServerError with default headers values
func NewHardwareModelGetGlobalHardwareModelInternalServerError() *HardwareModelGetGlobalHardwareModelInternalServerError {
	return &HardwareModelGetGlobalHardwareModelInternalServerError{}
}

/*
HardwareModelGetGlobalHardwareModelInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelGetGlobalHardwareModelInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware model internal server error response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareModelInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware model internal server error response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareModelInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware model internal server error response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareModelInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get global hardware model internal server error response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareModelInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get global hardware model internal server error response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareModelInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *HardwareModelGetGlobalHardwareModelInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] hardwareModelGetGlobalHardwareModelInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] hardwareModelGetGlobalHardwareModelInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareModelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareModelGatewayTimeout creates a HardwareModelGetGlobalHardwareModelGatewayTimeout with default headers values
func NewHardwareModelGetGlobalHardwareModelGatewayTimeout() *HardwareModelGetGlobalHardwareModelGatewayTimeout {
	return &HardwareModelGetGlobalHardwareModelGatewayTimeout{}
}

/*
HardwareModelGetGlobalHardwareModelGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelGetGlobalHardwareModelGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware model gateway timeout response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareModelGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware model gateway timeout response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareModelGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware model gateway timeout response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareModelGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get global hardware model gateway timeout response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareModelGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get global hardware model gateway timeout response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareModelGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *HardwareModelGetGlobalHardwareModelGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] hardwareModelGetGlobalHardwareModelGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] hardwareModelGetGlobalHardwareModelGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareModelGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareModelDefault creates a HardwareModelGetGlobalHardwareModelDefault with default headers values
func NewHardwareModelGetGlobalHardwareModelDefault(code int) *HardwareModelGetGlobalHardwareModelDefault {
	return &HardwareModelGetGlobalHardwareModelDefault{
		_statusCode: code,
	}
}

/*
HardwareModelGetGlobalHardwareModelDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelGetGlobalHardwareModelDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the hardware model get global hardware model default response
func (o *HardwareModelGetGlobalHardwareModelDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this hardware model get global hardware model default response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareModelDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hardware model get global hardware model default response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareModelDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hardware model get global hardware model default response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareModelDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hardware model get global hardware model default response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareModelDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hardware model get global hardware model default response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareModelDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HardwareModelGetGlobalHardwareModelDefault) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] HardwareModel_GetGlobalHardwareModel default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelDefault) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/id/{id}][%d] HardwareModel_GetGlobalHardwareModel default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareModelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
