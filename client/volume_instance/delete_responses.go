package volume_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// VolumeInstanceConfigurationDeleteVolumeInstanceReader is a Reader for the VolumeInstanceConfigurationDeleteVolumeInstance structure.
type VolumeInstanceConfigurationDeleteVolumeInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeInstanceConfigurationDeleteVolumeInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewVolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewVolumeInstanceConfigurationDeleteVolumeInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVolumeInstanceConfigurationDeleteVolumeInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewVolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewVolumeInstanceConfigurationDeleteVolumeInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeInstanceConfigurationDeleteVolumeInstanceOK creates a VolumeInstanceConfigurationDeleteVolumeInstanceOK with default headers values
func NewVolumeInstanceConfigurationDeleteVolumeInstanceOK() *VolumeInstanceConfigurationDeleteVolumeInstanceOK {
	return &VolumeInstanceConfigurationDeleteVolumeInstanceOK{}
}

/*
VolumeInstanceConfigurationDeleteVolumeInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type VolumeInstanceConfigurationDeleteVolumeInstanceOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration delete volume instance o k response has a 2xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume instance configuration delete volume instance o k response has a 3xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration delete volume instance o k response has a 4xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance configuration delete volume instance o k response has a 5xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration delete volume instance o k response a status code equal to that given
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the volume instance configuration delete volume instance o k response
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceOK) Code() int {
	return 200
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationDeleteVolumeInstanceOK  %+v", 200, o.Payload)
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceOK) String() string {
	return fmt.Sprintf("[DELETE /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationDeleteVolumeInstanceOK  %+v", 200, o.Payload)
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized creates a VolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized with default headers values
func NewVolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized() *VolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized {
	return &VolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized{}
}

/*
VolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type VolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration delete volume instance unauthorized response has a 2xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration delete volume instance unauthorized response has a 3xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration delete volume instance unauthorized response has a 4xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance configuration delete volume instance unauthorized response has a 5xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration delete volume instance unauthorized response a status code equal to that given
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the volume instance configuration delete volume instance unauthorized response
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized) Code() int {
	return 401
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationDeleteVolumeInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationDeleteVolumeInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationDeleteVolumeInstanceForbidden creates a VolumeInstanceConfigurationDeleteVolumeInstanceForbidden with default headers values
func NewVolumeInstanceConfigurationDeleteVolumeInstanceForbidden() *VolumeInstanceConfigurationDeleteVolumeInstanceForbidden {
	return &VolumeInstanceConfigurationDeleteVolumeInstanceForbidden{}
}

/*
VolumeInstanceConfigurationDeleteVolumeInstanceForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type VolumeInstanceConfigurationDeleteVolumeInstanceForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration delete volume instance forbidden response has a 2xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration delete volume instance forbidden response has a 3xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration delete volume instance forbidden response has a 4xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance configuration delete volume instance forbidden response has a 5xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration delete volume instance forbidden response a status code equal to that given
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the volume instance configuration delete volume instance forbidden response
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceForbidden) Code() int {
	return 403
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationDeleteVolumeInstanceForbidden  %+v", 403, o.Payload)
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationDeleteVolumeInstanceForbidden  %+v", 403, o.Payload)
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationDeleteVolumeInstanceNotFound creates a VolumeInstanceConfigurationDeleteVolumeInstanceNotFound with default headers values
func NewVolumeInstanceConfigurationDeleteVolumeInstanceNotFound() *NotFound {
	return &NotFound{}
}

/*
NotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type NotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration delete volume instance not found response has a 2xx status code
func (o *NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration delete volume instance not found response has a 3xx status code
func (o *NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration delete volume instance not found response has a 4xx status code
func (o *NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance configuration delete volume instance not found response has a 5xx status code
func (o *NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration delete volume instance not found response a status code equal to that given
func (o *NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the volume instance configuration delete volume instance not found response
func (o *NotFound) Code() int {
	return 404
}

func (o *NotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationDeleteVolumeInstanceNotFound  %+v", 404, o.Payload)
}

func (o *NotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationDeleteVolumeInstanceNotFound  %+v", 404, o.Payload)
}

func (o *NotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError creates a VolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError with default headers values
func NewVolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError() *VolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError {
	return &VolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError{}
}

/*
VolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type VolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration delete volume instance internal server error response has a 2xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration delete volume instance internal server error response has a 3xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration delete volume instance internal server error response has a 4xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance configuration delete volume instance internal server error response has a 5xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this volume instance configuration delete volume instance internal server error response a status code equal to that given
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the volume instance configuration delete volume instance internal server error response
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError) Code() int {
	return 500
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationDeleteVolumeInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationDeleteVolumeInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout creates a VolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout with default headers values
func NewVolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout() *VolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout {
	return &VolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout{}
}

/*
VolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type VolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration delete volume instance gateway timeout response has a 2xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration delete volume instance gateway timeout response has a 3xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration delete volume instance gateway timeout response has a 4xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance configuration delete volume instance gateway timeout response has a 5xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this volume instance configuration delete volume instance gateway timeout response a status code equal to that given
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the volume instance configuration delete volume instance gateway timeout response
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout) Code() int {
	return 504
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/volumes/instances/id/{id}][%d] volumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationDeleteVolumeInstanceDefault creates a VolumeInstanceConfigurationDeleteVolumeInstanceDefault with default headers values
func NewVolumeInstanceConfigurationDeleteVolumeInstanceDefault(code int) *VolumeInstanceConfigurationDeleteVolumeInstanceDefault {
	return &VolumeInstanceConfigurationDeleteVolumeInstanceDefault{
		_statusCode: code,
	}
}

/*
VolumeInstanceConfigurationDeleteVolumeInstanceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type VolumeInstanceConfigurationDeleteVolumeInstanceDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this volume instance configuration delete volume instance default response has a 2xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this volume instance configuration delete volume instance default response has a 3xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this volume instance configuration delete volume instance default response has a 4xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this volume instance configuration delete volume instance default response has a 5xx status code
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this volume instance configuration delete volume instance default response a status code equal to that given
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the volume instance configuration delete volume instance default response
func (o *VolumeInstanceConfigurationDeleteVolumeInstanceDefault) Code() int {
	return o._statusCode
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/volumes/instances/id/{id}][%d] VolumeInstanceConfiguration_DeleteVolumeInstance default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/volumes/instances/id/{id}][%d] VolumeInstanceConfiguration_DeleteVolumeInstance default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *VolumeInstanceConfigurationDeleteVolumeInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
