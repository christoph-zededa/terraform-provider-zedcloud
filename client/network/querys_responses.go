package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// EdgeNetworkConfigurationQueryEdgeNetworksReader is a Reader for the EdgeNetworkConfigurationQueryEdgeNetworks structure.
type EdgeNetworkConfigurationQueryEdgeNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNetworkConfigurationQueryEdgeNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNetworkConfigurationQueryEdgeNetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeNetworkConfigurationQueryEdgeNetworksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeNetworkConfigurationQueryEdgeNetworksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNetworkConfigurationQueryEdgeNetworksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNetworkConfigurationQueryEdgeNetworksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNetworkConfigurationQueryEdgeNetworksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNetworkConfigurationQueryEdgeNetworksOK creates a EdgeNetworkConfigurationQueryEdgeNetworksOK with default headers values
func NewEdgeNetworkConfigurationQueryEdgeNetworksOK() *EdgeNetworkConfigurationQueryEdgeNetworksOK {
	return &EdgeNetworkConfigurationQueryEdgeNetworksOK{}
}

/*
EdgeNetworkConfigurationQueryEdgeNetworksOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNetworkConfigurationQueryEdgeNetworksOK struct {
	Payload *models.NetConfigList
}

// IsSuccess returns true when this edge network configuration query edge networks o k response has a 2xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge network configuration query edge networks o k response has a 3xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration query edge networks o k response has a 4xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network configuration query edge networks o k response has a 5xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network configuration query edge networks o k response a status code equal to that given
func (o *EdgeNetworkConfigurationQueryEdgeNetworksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge network configuration query edge networks o k response
func (o *EdgeNetworkConfigurationQueryEdgeNetworksOK) Code() int {
	return 200
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksOK) Error() string {
	return fmt.Sprintf("[GET /v1/networks][%d] edgeNetworkConfigurationQueryEdgeNetworksOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksOK) String() string {
	return fmt.Sprintf("[GET /v1/networks][%d] edgeNetworkConfigurationQueryEdgeNetworksOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksOK) GetPayload() *models.NetConfigList {
	return o.Payload
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetConfigList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationQueryEdgeNetworksBadRequest creates a EdgeNetworkConfigurationQueryEdgeNetworksBadRequest with default headers values
func NewEdgeNetworkConfigurationQueryEdgeNetworksBadRequest() *EdgeNetworkConfigurationQueryEdgeNetworksBadRequest {
	return &EdgeNetworkConfigurationQueryEdgeNetworksBadRequest{}
}

/*
EdgeNetworkConfigurationQueryEdgeNetworksBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type EdgeNetworkConfigurationQueryEdgeNetworksBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration query edge networks bad request response has a 2xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration query edge networks bad request response has a 3xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration query edge networks bad request response has a 4xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network configuration query edge networks bad request response has a 5xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network configuration query edge networks bad request response a status code equal to that given
func (o *EdgeNetworkConfigurationQueryEdgeNetworksBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge network configuration query edge networks bad request response
func (o *EdgeNetworkConfigurationQueryEdgeNetworksBadRequest) Code() int {
	return 400
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/networks][%d] edgeNetworkConfigurationQueryEdgeNetworksBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/networks][%d] edgeNetworkConfigurationQueryEdgeNetworksBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationQueryEdgeNetworksUnauthorized creates a EdgeNetworkConfigurationQueryEdgeNetworksUnauthorized with default headers values
func NewEdgeNetworkConfigurationQueryEdgeNetworksUnauthorized() *EdgeNetworkConfigurationQueryEdgeNetworksUnauthorized {
	return &EdgeNetworkConfigurationQueryEdgeNetworksUnauthorized{}
}

/*
EdgeNetworkConfigurationQueryEdgeNetworksUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNetworkConfigurationQueryEdgeNetworksUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration query edge networks unauthorized response has a 2xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration query edge networks unauthorized response has a 3xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration query edge networks unauthorized response has a 4xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network configuration query edge networks unauthorized response has a 5xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network configuration query edge networks unauthorized response a status code equal to that given
func (o *EdgeNetworkConfigurationQueryEdgeNetworksUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge network configuration query edge networks unauthorized response
func (o *EdgeNetworkConfigurationQueryEdgeNetworksUnauthorized) Code() int {
	return 401
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/networks][%d] edgeNetworkConfigurationQueryEdgeNetworksUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/networks][%d] edgeNetworkConfigurationQueryEdgeNetworksUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationQueryEdgeNetworksForbidden creates a EdgeNetworkConfigurationQueryEdgeNetworksForbidden with default headers values
func NewEdgeNetworkConfigurationQueryEdgeNetworksForbidden() *EdgeNetworkConfigurationQueryEdgeNetworksForbidden {
	return &EdgeNetworkConfigurationQueryEdgeNetworksForbidden{}
}

/*
EdgeNetworkConfigurationQueryEdgeNetworksForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNetworkConfigurationQueryEdgeNetworksForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration query edge networks forbidden response has a 2xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration query edge networks forbidden response has a 3xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration query edge networks forbidden response has a 4xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network configuration query edge networks forbidden response has a 5xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network configuration query edge networks forbidden response a status code equal to that given
func (o *EdgeNetworkConfigurationQueryEdgeNetworksForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge network configuration query edge networks forbidden response
func (o *EdgeNetworkConfigurationQueryEdgeNetworksForbidden) Code() int {
	return 403
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/networks][%d] edgeNetworkConfigurationQueryEdgeNetworksForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksForbidden) String() string {
	return fmt.Sprintf("[GET /v1/networks][%d] edgeNetworkConfigurationQueryEdgeNetworksForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationQueryEdgeNetworksInternalServerError creates a EdgeNetworkConfigurationQueryEdgeNetworksInternalServerError with default headers values
func NewEdgeNetworkConfigurationQueryEdgeNetworksInternalServerError() *EdgeNetworkConfigurationQueryEdgeNetworksInternalServerError {
	return &EdgeNetworkConfigurationQueryEdgeNetworksInternalServerError{}
}

/*
EdgeNetworkConfigurationQueryEdgeNetworksInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNetworkConfigurationQueryEdgeNetworksInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration query edge networks internal server error response has a 2xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration query edge networks internal server error response has a 3xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration query edge networks internal server error response has a 4xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network configuration query edge networks internal server error response has a 5xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network configuration query edge networks internal server error response a status code equal to that given
func (o *EdgeNetworkConfigurationQueryEdgeNetworksInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge network configuration query edge networks internal server error response
func (o *EdgeNetworkConfigurationQueryEdgeNetworksInternalServerError) Code() int {
	return 500
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/networks][%d] edgeNetworkConfigurationQueryEdgeNetworksInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/networks][%d] edgeNetworkConfigurationQueryEdgeNetworksInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout creates a EdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout with default headers values
func NewEdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout() *EdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout {
	return &EdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout{}
}

/*
EdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge network configuration query edge networks gateway timeout response has a 2xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network configuration query edge networks gateway timeout response has a 3xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network configuration query edge networks gateway timeout response has a 4xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network configuration query edge networks gateway timeout response has a 5xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network configuration query edge networks gateway timeout response a status code equal to that given
func (o *EdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge network configuration query edge networks gateway timeout response
func (o *EdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/networks][%d] edgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/networks][%d] edgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkConfigurationQueryEdgeNetworksDefault creates a EdgeNetworkConfigurationQueryEdgeNetworksDefault with default headers values
func NewEdgeNetworkConfigurationQueryEdgeNetworksDefault(code int) *EdgeNetworkConfigurationQueryEdgeNetworksDefault {
	return &EdgeNetworkConfigurationQueryEdgeNetworksDefault{
		_statusCode: code,
	}
}

/*
EdgeNetworkConfigurationQueryEdgeNetworksDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNetworkConfigurationQueryEdgeNetworksDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge network configuration query edge networks default response has a 2xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge network configuration query edge networks default response has a 3xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge network configuration query edge networks default response has a 4xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge network configuration query edge networks default response has a 5xx status code
func (o *EdgeNetworkConfigurationQueryEdgeNetworksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge network configuration query edge networks default response a status code equal to that given
func (o *EdgeNetworkConfigurationQueryEdgeNetworksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge network configuration query edge networks default response
func (o *EdgeNetworkConfigurationQueryEdgeNetworksDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksDefault) Error() string {
	return fmt.Sprintf("[GET /v1/networks][%d] EdgeNetworkConfiguration_QueryEdgeNetworks default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksDefault) String() string {
	return fmt.Sprintf("[GET /v1/networks][%d] EdgeNetworkConfiguration_QueryEdgeNetworks default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNetworkConfigurationQueryEdgeNetworksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
