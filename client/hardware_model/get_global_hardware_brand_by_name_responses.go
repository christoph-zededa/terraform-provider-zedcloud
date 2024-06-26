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

// HardwareModelGetGlobalHardwareBrandByNameReader is a Reader for the HardwareModelGetGlobalHardwareBrandByName structure.
type HardwareModelGetGlobalHardwareBrandByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelGetGlobalHardwareBrandByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelGetGlobalHardwareBrandByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHardwareModelGetGlobalHardwareBrandByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelGetGlobalHardwareBrandByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHardwareModelGetGlobalHardwareBrandByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelGetGlobalHardwareBrandByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelGetGlobalHardwareBrandByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelGetGlobalHardwareBrandByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelGetGlobalHardwareBrandByNameOK creates a HardwareModelGetGlobalHardwareBrandByNameOK with default headers values
func NewHardwareModelGetGlobalHardwareBrandByNameOK() *HardwareModelGetGlobalHardwareBrandByNameOK {
	return &HardwareModelGetGlobalHardwareBrandByNameOK{}
}

/*
HardwareModelGetGlobalHardwareBrandByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelGetGlobalHardwareBrandByNameOK struct {
	Payload *models.SysBrand
}

// IsSuccess returns true when this hardware model get global hardware brand by name o k response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hardware model get global hardware brand by name o k response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware brand by name o k response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get global hardware brand by name o k response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get global hardware brand by name o k response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareBrandByNameOK) IsCode(code int) bool {
	return code == 200
}

func (o *HardwareModelGetGlobalHardwareBrandByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] hardwareModelGetGlobalHardwareBrandByNameOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] hardwareModelGetGlobalHardwareBrandByNameOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandByNameOK) GetPayload() *models.SysBrand {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareBrandByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SysBrand)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareBrandByNameUnauthorized creates a HardwareModelGetGlobalHardwareBrandByNameUnauthorized with default headers values
func NewHardwareModelGetGlobalHardwareBrandByNameUnauthorized() *HardwareModelGetGlobalHardwareBrandByNameUnauthorized {
	return &HardwareModelGetGlobalHardwareBrandByNameUnauthorized{}
}

/*
HardwareModelGetGlobalHardwareBrandByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelGetGlobalHardwareBrandByNameUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware brand by name unauthorized response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware brand by name unauthorized response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware brand by name unauthorized response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get global hardware brand by name unauthorized response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get global hardware brand by name unauthorized response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareBrandByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *HardwareModelGetGlobalHardwareBrandByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] hardwareModelGetGlobalHardwareBrandByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] hardwareModelGetGlobalHardwareBrandByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandByNameUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareBrandByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareBrandByNameForbidden creates a HardwareModelGetGlobalHardwareBrandByNameForbidden with default headers values
func NewHardwareModelGetGlobalHardwareBrandByNameForbidden() *HardwareModelGetGlobalHardwareBrandByNameForbidden {
	return &HardwareModelGetGlobalHardwareBrandByNameForbidden{}
}

/*
HardwareModelGetGlobalHardwareBrandByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelGetGlobalHardwareBrandByNameForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware brand by name forbidden response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware brand by name forbidden response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware brand by name forbidden response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get global hardware brand by name forbidden response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get global hardware brand by name forbidden response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareBrandByNameForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *HardwareModelGetGlobalHardwareBrandByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] hardwareModelGetGlobalHardwareBrandByNameForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] hardwareModelGetGlobalHardwareBrandByNameForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandByNameForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareBrandByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareBrandByNameNotFound creates a HardwareModelGetGlobalHardwareBrandByNameNotFound with default headers values
func NewHardwareModelGetGlobalHardwareBrandByNameNotFound() *HardwareModelGetGlobalHardwareBrandByNameNotFound {
	return &HardwareModelGetGlobalHardwareBrandByNameNotFound{}
}

/*
HardwareModelGetGlobalHardwareBrandByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type HardwareModelGetGlobalHardwareBrandByNameNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware brand by name not found response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware brand by name not found response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware brand by name not found response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get global hardware brand by name not found response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get global hardware brand by name not found response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareBrandByNameNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *HardwareModelGetGlobalHardwareBrandByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] hardwareModelGetGlobalHardwareBrandByNameNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] hardwareModelGetGlobalHardwareBrandByNameNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandByNameNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareBrandByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareBrandByNameInternalServerError creates a HardwareModelGetGlobalHardwareBrandByNameInternalServerError with default headers values
func NewHardwareModelGetGlobalHardwareBrandByNameInternalServerError() *HardwareModelGetGlobalHardwareBrandByNameInternalServerError {
	return &HardwareModelGetGlobalHardwareBrandByNameInternalServerError{}
}

/*
HardwareModelGetGlobalHardwareBrandByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelGetGlobalHardwareBrandByNameInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware brand by name internal server error response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware brand by name internal server error response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware brand by name internal server error response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get global hardware brand by name internal server error response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get global hardware brand by name internal server error response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareBrandByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *HardwareModelGetGlobalHardwareBrandByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] hardwareModelGetGlobalHardwareBrandByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] hardwareModelGetGlobalHardwareBrandByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandByNameInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareBrandByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareBrandByNameGatewayTimeout creates a HardwareModelGetGlobalHardwareBrandByNameGatewayTimeout with default headers values
func NewHardwareModelGetGlobalHardwareBrandByNameGatewayTimeout() *HardwareModelGetGlobalHardwareBrandByNameGatewayTimeout {
	return &HardwareModelGetGlobalHardwareBrandByNameGatewayTimeout{}
}

/*
HardwareModelGetGlobalHardwareBrandByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelGetGlobalHardwareBrandByNameGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware brand by name gateway timeout response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware brand by name gateway timeout response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware brand by name gateway timeout response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get global hardware brand by name gateway timeout response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get global hardware brand by name gateway timeout response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareBrandByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *HardwareModelGetGlobalHardwareBrandByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] hardwareModelGetGlobalHardwareBrandByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] hardwareModelGetGlobalHardwareBrandByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandByNameGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareBrandByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareBrandByNameDefault creates a HardwareModelGetGlobalHardwareBrandByNameDefault with default headers values
func NewHardwareModelGetGlobalHardwareBrandByNameDefault(code int) *HardwareModelGetGlobalHardwareBrandByNameDefault {
	return &HardwareModelGetGlobalHardwareBrandByNameDefault{
		_statusCode: code,
	}
}

/*
HardwareModelGetGlobalHardwareBrandByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelGetGlobalHardwareBrandByNameDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the hardware model get global hardware brand by name default response
func (o *HardwareModelGetGlobalHardwareBrandByNameDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this hardware model get global hardware brand by name default response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hardware model get global hardware brand by name default response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hardware model get global hardware brand by name default response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hardware model get global hardware brand by name default response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareBrandByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hardware model get global hardware brand by name default response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareBrandByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HardwareModelGetGlobalHardwareBrandByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] HardwareModel_GetGlobalHardwareBrandByName default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/brands/global/name/{name}][%d] HardwareModel_GetGlobalHardwareBrandByName default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandByNameDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareBrandByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
