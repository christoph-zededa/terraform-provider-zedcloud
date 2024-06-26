package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// ImageConfigurationQueryImageProjectListReader is a Reader for the ImageConfigurationQueryImageProjectList structure.
type ImageConfigurationQueryImageProjectListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageConfigurationQueryImageProjectListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageConfigurationQueryImageProjectListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImageConfigurationQueryImageProjectListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImageConfigurationQueryImageProjectListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImageConfigurationQueryImageProjectListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageConfigurationQueryImageProjectListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewImageConfigurationQueryImageProjectListGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImageConfigurationQueryImageProjectListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImageConfigurationQueryImageProjectListOK creates a ImageConfigurationQueryImageProjectListOK with default headers values
func NewImageConfigurationQueryImageProjectListOK() *ImageConfigurationQueryImageProjectListOK {
	return &ImageConfigurationQueryImageProjectListOK{}
}

/*
ImageConfigurationQueryImageProjectListOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImageConfigurationQueryImageProjectListOK struct {
	Payload *models.ImageProjectList
}

// IsSuccess returns true when this image configuration query image project list o k response has a 2xx status code
func (o *ImageConfigurationQueryImageProjectListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image configuration query image project list o k response has a 3xx status code
func (o *ImageConfigurationQueryImageProjectListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query image project list o k response has a 4xx status code
func (o *ImageConfigurationQueryImageProjectListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration query image project list o k response has a 5xx status code
func (o *ImageConfigurationQueryImageProjectListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration query image project list o k response a status code equal to that given
func (o *ImageConfigurationQueryImageProjectListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the image configuration query image project list o k response
func (o *ImageConfigurationQueryImageProjectListOK) Code() int {
	return 200
}

func (o *ImageConfigurationQueryImageProjectListOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/projects][%d] imageConfigurationQueryImageProjectListOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationQueryImageProjectListOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/projects][%d] imageConfigurationQueryImageProjectListOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationQueryImageProjectListOK) GetPayload() *models.ImageProjectList {
	return o.Payload
}

func (o *ImageConfigurationQueryImageProjectListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageProjectList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImageProjectListBadRequest creates a ImageConfigurationQueryImageProjectListBadRequest with default headers values
func NewImageConfigurationQueryImageProjectListBadRequest() *ImageConfigurationQueryImageProjectListBadRequest {
	return &ImageConfigurationQueryImageProjectListBadRequest{}
}

/*
ImageConfigurationQueryImageProjectListBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type ImageConfigurationQueryImageProjectListBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration query image project list bad request response has a 2xx status code
func (o *ImageConfigurationQueryImageProjectListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration query image project list bad request response has a 3xx status code
func (o *ImageConfigurationQueryImageProjectListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query image project list bad request response has a 4xx status code
func (o *ImageConfigurationQueryImageProjectListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration query image project list bad request response has a 5xx status code
func (o *ImageConfigurationQueryImageProjectListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration query image project list bad request response a status code equal to that given
func (o *ImageConfigurationQueryImageProjectListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the image configuration query image project list bad request response
func (o *ImageConfigurationQueryImageProjectListBadRequest) Code() int {
	return 400
}

func (o *ImageConfigurationQueryImageProjectListBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/projects][%d] imageConfigurationQueryImageProjectListBadRequest  %+v", 400, o.Payload)
}

func (o *ImageConfigurationQueryImageProjectListBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/projects][%d] imageConfigurationQueryImageProjectListBadRequest  %+v", 400, o.Payload)
}

func (o *ImageConfigurationQueryImageProjectListBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryImageProjectListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImageProjectListUnauthorized creates a ImageConfigurationQueryImageProjectListUnauthorized with default headers values
func NewImageConfigurationQueryImageProjectListUnauthorized() *ImageConfigurationQueryImageProjectListUnauthorized {
	return &ImageConfigurationQueryImageProjectListUnauthorized{}
}

/*
ImageConfigurationQueryImageProjectListUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ImageConfigurationQueryImageProjectListUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration query image project list unauthorized response has a 2xx status code
func (o *ImageConfigurationQueryImageProjectListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration query image project list unauthorized response has a 3xx status code
func (o *ImageConfigurationQueryImageProjectListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query image project list unauthorized response has a 4xx status code
func (o *ImageConfigurationQueryImageProjectListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration query image project list unauthorized response has a 5xx status code
func (o *ImageConfigurationQueryImageProjectListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration query image project list unauthorized response a status code equal to that given
func (o *ImageConfigurationQueryImageProjectListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the image configuration query image project list unauthorized response
func (o *ImageConfigurationQueryImageProjectListUnauthorized) Code() int {
	return 401
}

func (o *ImageConfigurationQueryImageProjectListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/projects][%d] imageConfigurationQueryImageProjectListUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationQueryImageProjectListUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/projects][%d] imageConfigurationQueryImageProjectListUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationQueryImageProjectListUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryImageProjectListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImageProjectListForbidden creates a ImageConfigurationQueryImageProjectListForbidden with default headers values
func NewImageConfigurationQueryImageProjectListForbidden() *ImageConfigurationQueryImageProjectListForbidden {
	return &ImageConfigurationQueryImageProjectListForbidden{}
}

/*
ImageConfigurationQueryImageProjectListForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ImageConfigurationQueryImageProjectListForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration query image project list forbidden response has a 2xx status code
func (o *ImageConfigurationQueryImageProjectListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration query image project list forbidden response has a 3xx status code
func (o *ImageConfigurationQueryImageProjectListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query image project list forbidden response has a 4xx status code
func (o *ImageConfigurationQueryImageProjectListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration query image project list forbidden response has a 5xx status code
func (o *ImageConfigurationQueryImageProjectListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration query image project list forbidden response a status code equal to that given
func (o *ImageConfigurationQueryImageProjectListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the image configuration query image project list forbidden response
func (o *ImageConfigurationQueryImageProjectListForbidden) Code() int {
	return 403
}

func (o *ImageConfigurationQueryImageProjectListForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/projects][%d] imageConfigurationQueryImageProjectListForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationQueryImageProjectListForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/projects][%d] imageConfigurationQueryImageProjectListForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationQueryImageProjectListForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryImageProjectListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImageProjectListInternalServerError creates a ImageConfigurationQueryImageProjectListInternalServerError with default headers values
func NewImageConfigurationQueryImageProjectListInternalServerError() *ImageConfigurationQueryImageProjectListInternalServerError {
	return &ImageConfigurationQueryImageProjectListInternalServerError{}
}

/*
ImageConfigurationQueryImageProjectListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ImageConfigurationQueryImageProjectListInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration query image project list internal server error response has a 2xx status code
func (o *ImageConfigurationQueryImageProjectListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration query image project list internal server error response has a 3xx status code
func (o *ImageConfigurationQueryImageProjectListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query image project list internal server error response has a 4xx status code
func (o *ImageConfigurationQueryImageProjectListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration query image project list internal server error response has a 5xx status code
func (o *ImageConfigurationQueryImageProjectListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration query image project list internal server error response a status code equal to that given
func (o *ImageConfigurationQueryImageProjectListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the image configuration query image project list internal server error response
func (o *ImageConfigurationQueryImageProjectListInternalServerError) Code() int {
	return 500
}

func (o *ImageConfigurationQueryImageProjectListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/projects][%d] imageConfigurationQueryImageProjectListInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationQueryImageProjectListInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/projects][%d] imageConfigurationQueryImageProjectListInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationQueryImageProjectListInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryImageProjectListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImageProjectListGatewayTimeout creates a ImageConfigurationQueryImageProjectListGatewayTimeout with default headers values
func NewImageConfigurationQueryImageProjectListGatewayTimeout() *ImageConfigurationQueryImageProjectListGatewayTimeout {
	return &ImageConfigurationQueryImageProjectListGatewayTimeout{}
}

/*
ImageConfigurationQueryImageProjectListGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ImageConfigurationQueryImageProjectListGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration query image project list gateway timeout response has a 2xx status code
func (o *ImageConfigurationQueryImageProjectListGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration query image project list gateway timeout response has a 3xx status code
func (o *ImageConfigurationQueryImageProjectListGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query image project list gateway timeout response has a 4xx status code
func (o *ImageConfigurationQueryImageProjectListGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration query image project list gateway timeout response has a 5xx status code
func (o *ImageConfigurationQueryImageProjectListGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration query image project list gateway timeout response a status code equal to that given
func (o *ImageConfigurationQueryImageProjectListGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the image configuration query image project list gateway timeout response
func (o *ImageConfigurationQueryImageProjectListGatewayTimeout) Code() int {
	return 504
}

func (o *ImageConfigurationQueryImageProjectListGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/projects][%d] imageConfigurationQueryImageProjectListGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationQueryImageProjectListGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/projects][%d] imageConfigurationQueryImageProjectListGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationQueryImageProjectListGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryImageProjectListGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImageProjectListDefault creates a ImageConfigurationQueryImageProjectListDefault with default headers values
func NewImageConfigurationQueryImageProjectListDefault(code int) *ImageConfigurationQueryImageProjectListDefault {
	return &ImageConfigurationQueryImageProjectListDefault{
		_statusCode: code,
	}
}

/*
ImageConfigurationQueryImageProjectListDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImageConfigurationQueryImageProjectListDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this image configuration query image project list default response has a 2xx status code
func (o *ImageConfigurationQueryImageProjectListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this image configuration query image project list default response has a 3xx status code
func (o *ImageConfigurationQueryImageProjectListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this image configuration query image project list default response has a 4xx status code
func (o *ImageConfigurationQueryImageProjectListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this image configuration query image project list default response has a 5xx status code
func (o *ImageConfigurationQueryImageProjectListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this image configuration query image project list default response a status code equal to that given
func (o *ImageConfigurationQueryImageProjectListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the image configuration query image project list default response
func (o *ImageConfigurationQueryImageProjectListDefault) Code() int {
	return o._statusCode
}

func (o *ImageConfigurationQueryImageProjectListDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/projects][%d] ImageConfiguration_QueryImageProjectList default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationQueryImageProjectListDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/projects][%d] ImageConfiguration_QueryImageProjectList default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationQueryImageProjectListDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ImageConfigurationQueryImageProjectListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
