// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// DeleteVpcsIDReader is a Reader for the DeleteVpcsID structure.
type DeleteVpcsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVpcsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteVpcsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteVpcsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteVpcsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteVpcsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteVpcsIDNoContent creates a DeleteVpcsIDNoContent with default headers values
func NewDeleteVpcsIDNoContent() *DeleteVpcsIDNoContent {
	return &DeleteVpcsIDNoContent{}
}

/*DeleteVpcsIDNoContent handles this case with default header values.

error
*/
type DeleteVpcsIDNoContent struct {
	Payload *models.Riaaserror
}

func (o *DeleteVpcsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /vpcs/{id}][%d] deleteVpcsIdNoContent  %+v", 204, o.Payload)
}

func (o *DeleteVpcsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVpcsIDBadRequest creates a DeleteVpcsIDBadRequest with default headers values
func NewDeleteVpcsIDBadRequest() *DeleteVpcsIDBadRequest {
	return &DeleteVpcsIDBadRequest{}
}

/*DeleteVpcsIDBadRequest handles this case with default header values.

error
*/
type DeleteVpcsIDBadRequest struct {
	Payload *models.Riaaserror
}

func (o *DeleteVpcsIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /vpcs/{id}][%d] deleteVpcsIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteVpcsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVpcsIDNotFound creates a DeleteVpcsIDNotFound with default headers values
func NewDeleteVpcsIDNotFound() *DeleteVpcsIDNotFound {
	return &DeleteVpcsIDNotFound{}
}

/*DeleteVpcsIDNotFound handles this case with default header values.

error
*/
type DeleteVpcsIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteVpcsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /vpcs/{id}][%d] deleteVpcsIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteVpcsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVpcsIDInternalServerError creates a DeleteVpcsIDInternalServerError with default headers values
func NewDeleteVpcsIDInternalServerError() *DeleteVpcsIDInternalServerError {
	return &DeleteVpcsIDInternalServerError{}
}

/*DeleteVpcsIDInternalServerError handles this case with default header values.

error
*/
type DeleteVpcsIDInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *DeleteVpcsIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /vpcs/{id}][%d] deleteVpcsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteVpcsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
