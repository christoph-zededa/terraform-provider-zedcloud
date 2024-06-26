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

// HardwareModelCreateHardwareBrandReader is a Reader for the HardwareModelCreateHardwareBrand structure.
type HardwareModelCreateHardwareBrandReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelCreateHardwareBrandReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelCreateHardwareBrandOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHardwareModelCreateHardwareBrandBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHardwareModelCreateHardwareBrandUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelCreateHardwareBrandForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewHardwareModelCreateHardwareBrandConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelCreateHardwareBrandInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelCreateHardwareBrandGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelCreateHardwareBrandDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelCreateHardwareBrandOK creates a HardwareModelCreateHardwareBrandOK with default headers values
func NewHardwareModelCreateHardwareBrandOK() *HardwareModelCreateHardwareBrandOK {
	return &HardwareModelCreateHardwareBrandOK{}
}

/*
HardwareModelCreateHardwareBrandOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelCreateHardwareBrandOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model create hardware brand o k response has a 2xx status code
func (o *HardwareModelCreateHardwareBrandOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hardware model create hardware brand o k response has a 3xx status code
func (o *HardwareModelCreateHardwareBrandOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model create hardware brand o k response has a 4xx status code
func (o *HardwareModelCreateHardwareBrandOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model create hardware brand o k response has a 5xx status code
func (o *HardwareModelCreateHardwareBrandOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model create hardware brand o k response a status code equal to that given
func (o *HardwareModelCreateHardwareBrandOK) IsCode(code int) bool {
	return code == 200
}

func (o *HardwareModelCreateHardwareBrandOK) Error() string {
	return fmt.Sprintf("[POST /v1/brands][%d] hardwareModelCreateHardwareBrandOK  %+v", 200, o.Payload)
}

func (o *HardwareModelCreateHardwareBrandOK) String() string {
	return fmt.Sprintf("[POST /v1/brands][%d] hardwareModelCreateHardwareBrandOK  %+v", 200, o.Payload)
}

func (o *HardwareModelCreateHardwareBrandOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelCreateHardwareBrandOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelCreateHardwareBrandBadRequest creates a HardwareModelCreateHardwareBrandBadRequest with default headers values
func NewHardwareModelCreateHardwareBrandBadRequest() *HardwareModelCreateHardwareBrandBadRequest {
	return &HardwareModelCreateHardwareBrandBadRequest{}
}

/*
HardwareModelCreateHardwareBrandBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type HardwareModelCreateHardwareBrandBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model create hardware brand bad request response has a 2xx status code
func (o *HardwareModelCreateHardwareBrandBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model create hardware brand bad request response has a 3xx status code
func (o *HardwareModelCreateHardwareBrandBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model create hardware brand bad request response has a 4xx status code
func (o *HardwareModelCreateHardwareBrandBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model create hardware brand bad request response has a 5xx status code
func (o *HardwareModelCreateHardwareBrandBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model create hardware brand bad request response a status code equal to that given
func (o *HardwareModelCreateHardwareBrandBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *HardwareModelCreateHardwareBrandBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/brands][%d] hardwareModelCreateHardwareBrandBadRequest  %+v", 400, o.Payload)
}

func (o *HardwareModelCreateHardwareBrandBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/brands][%d] hardwareModelCreateHardwareBrandBadRequest  %+v", 400, o.Payload)
}

func (o *HardwareModelCreateHardwareBrandBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelCreateHardwareBrandBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelCreateHardwareBrandUnauthorized creates a HardwareModelCreateHardwareBrandUnauthorized with default headers values
func NewHardwareModelCreateHardwareBrandUnauthorized() *HardwareModelCreateHardwareBrandUnauthorized {
	return &HardwareModelCreateHardwareBrandUnauthorized{}
}

/*
HardwareModelCreateHardwareBrandUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelCreateHardwareBrandUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model create hardware brand unauthorized response has a 2xx status code
func (o *HardwareModelCreateHardwareBrandUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model create hardware brand unauthorized response has a 3xx status code
func (o *HardwareModelCreateHardwareBrandUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model create hardware brand unauthorized response has a 4xx status code
func (o *HardwareModelCreateHardwareBrandUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model create hardware brand unauthorized response has a 5xx status code
func (o *HardwareModelCreateHardwareBrandUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model create hardware brand unauthorized response a status code equal to that given
func (o *HardwareModelCreateHardwareBrandUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *HardwareModelCreateHardwareBrandUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/brands][%d] hardwareModelCreateHardwareBrandUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelCreateHardwareBrandUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/brands][%d] hardwareModelCreateHardwareBrandUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelCreateHardwareBrandUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelCreateHardwareBrandUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelCreateHardwareBrandForbidden creates a HardwareModelCreateHardwareBrandForbidden with default headers values
func NewHardwareModelCreateHardwareBrandForbidden() *HardwareModelCreateHardwareBrandForbidden {
	return &HardwareModelCreateHardwareBrandForbidden{}
}

/*
HardwareModelCreateHardwareBrandForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelCreateHardwareBrandForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model create hardware brand forbidden response has a 2xx status code
func (o *HardwareModelCreateHardwareBrandForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model create hardware brand forbidden response has a 3xx status code
func (o *HardwareModelCreateHardwareBrandForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model create hardware brand forbidden response has a 4xx status code
func (o *HardwareModelCreateHardwareBrandForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model create hardware brand forbidden response has a 5xx status code
func (o *HardwareModelCreateHardwareBrandForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model create hardware brand forbidden response a status code equal to that given
func (o *HardwareModelCreateHardwareBrandForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *HardwareModelCreateHardwareBrandForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/brands][%d] hardwareModelCreateHardwareBrandForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelCreateHardwareBrandForbidden) String() string {
	return fmt.Sprintf("[POST /v1/brands][%d] hardwareModelCreateHardwareBrandForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelCreateHardwareBrandForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelCreateHardwareBrandForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelCreateHardwareBrandConflict creates a HardwareModelCreateHardwareBrandConflict with default headers values
func NewHardwareModelCreateHardwareBrandConflict() *HardwareModelCreateHardwareBrandConflict {
	return &HardwareModelCreateHardwareBrandConflict{}
}

/*
HardwareModelCreateHardwareBrandConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this hardware brand record will conflict with an already existing hardware brand record.
*/
type HardwareModelCreateHardwareBrandConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model create hardware brand conflict response has a 2xx status code
func (o *HardwareModelCreateHardwareBrandConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model create hardware brand conflict response has a 3xx status code
func (o *HardwareModelCreateHardwareBrandConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model create hardware brand conflict response has a 4xx status code
func (o *HardwareModelCreateHardwareBrandConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model create hardware brand conflict response has a 5xx status code
func (o *HardwareModelCreateHardwareBrandConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model create hardware brand conflict response a status code equal to that given
func (o *HardwareModelCreateHardwareBrandConflict) IsCode(code int) bool {
	return code == 409
}

func (o *HardwareModelCreateHardwareBrandConflict) Error() string {
	return fmt.Sprintf("[POST /v1/brands][%d] hardwareModelCreateHardwareBrandConflict  %+v", 409, o.Payload)
}

func (o *HardwareModelCreateHardwareBrandConflict) String() string {
	return fmt.Sprintf("[POST /v1/brands][%d] hardwareModelCreateHardwareBrandConflict  %+v", 409, o.Payload)
}

func (o *HardwareModelCreateHardwareBrandConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelCreateHardwareBrandConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelCreateHardwareBrandInternalServerError creates a HardwareModelCreateHardwareBrandInternalServerError with default headers values
func NewHardwareModelCreateHardwareBrandInternalServerError() *HardwareModelCreateHardwareBrandInternalServerError {
	return &HardwareModelCreateHardwareBrandInternalServerError{}
}

/*
HardwareModelCreateHardwareBrandInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelCreateHardwareBrandInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model create hardware brand internal server error response has a 2xx status code
func (o *HardwareModelCreateHardwareBrandInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model create hardware brand internal server error response has a 3xx status code
func (o *HardwareModelCreateHardwareBrandInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model create hardware brand internal server error response has a 4xx status code
func (o *HardwareModelCreateHardwareBrandInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model create hardware brand internal server error response has a 5xx status code
func (o *HardwareModelCreateHardwareBrandInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model create hardware brand internal server error response a status code equal to that given
func (o *HardwareModelCreateHardwareBrandInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *HardwareModelCreateHardwareBrandInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/brands][%d] hardwareModelCreateHardwareBrandInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelCreateHardwareBrandInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/brands][%d] hardwareModelCreateHardwareBrandInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelCreateHardwareBrandInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelCreateHardwareBrandInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelCreateHardwareBrandGatewayTimeout creates a HardwareModelCreateHardwareBrandGatewayTimeout with default headers values
func NewHardwareModelCreateHardwareBrandGatewayTimeout() *HardwareModelCreateHardwareBrandGatewayTimeout {
	return &HardwareModelCreateHardwareBrandGatewayTimeout{}
}

/*
HardwareModelCreateHardwareBrandGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelCreateHardwareBrandGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model create hardware brand gateway timeout response has a 2xx status code
func (o *HardwareModelCreateHardwareBrandGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model create hardware brand gateway timeout response has a 3xx status code
func (o *HardwareModelCreateHardwareBrandGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model create hardware brand gateway timeout response has a 4xx status code
func (o *HardwareModelCreateHardwareBrandGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model create hardware brand gateway timeout response has a 5xx status code
func (o *HardwareModelCreateHardwareBrandGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model create hardware brand gateway timeout response a status code equal to that given
func (o *HardwareModelCreateHardwareBrandGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *HardwareModelCreateHardwareBrandGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/brands][%d] hardwareModelCreateHardwareBrandGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelCreateHardwareBrandGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/brands][%d] hardwareModelCreateHardwareBrandGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelCreateHardwareBrandGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelCreateHardwareBrandGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelCreateHardwareBrandDefault creates a HardwareModelCreateHardwareBrandDefault with default headers values
func NewHardwareModelCreateHardwareBrandDefault(code int) *HardwareModelCreateHardwareBrandDefault {
	return &HardwareModelCreateHardwareBrandDefault{
		_statusCode: code,
	}
}

/*
HardwareModelCreateHardwareBrandDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelCreateHardwareBrandDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the hardware model create hardware brand default response
func (o *HardwareModelCreateHardwareBrandDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this hardware model create hardware brand default response has a 2xx status code
func (o *HardwareModelCreateHardwareBrandDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hardware model create hardware brand default response has a 3xx status code
func (o *HardwareModelCreateHardwareBrandDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hardware model create hardware brand default response has a 4xx status code
func (o *HardwareModelCreateHardwareBrandDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hardware model create hardware brand default response has a 5xx status code
func (o *HardwareModelCreateHardwareBrandDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hardware model create hardware brand default response a status code equal to that given
func (o *HardwareModelCreateHardwareBrandDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HardwareModelCreateHardwareBrandDefault) Error() string {
	return fmt.Sprintf("[POST /v1/brands][%d] HardwareModel_CreateHardwareBrand default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelCreateHardwareBrandDefault) String() string {
	return fmt.Sprintf("[POST /v1/brands][%d] HardwareModel_CreateHardwareBrand default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelCreateHardwareBrandDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelCreateHardwareBrandDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
