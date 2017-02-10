package dhcp_reservations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rackn/rocket-skates/models"
)

/*POSTDhcpReservationCreated p o s t dhcp reservation created

swagger:response pOSTDhcpReservationCreated
*/
type POSTDhcpReservationCreated struct {

	/*
	  In: Body
	*/
	Payload *models.DhcpReservationInput `json:"body,omitempty"`
}

// NewPOSTDhcpReservationCreated creates POSTDhcpReservationCreated with default headers values
func NewPOSTDhcpReservationCreated() *POSTDhcpReservationCreated {
	return &POSTDhcpReservationCreated{}
}

// WithPayload adds the payload to the p o s t dhcp reservation created response
func (o *POSTDhcpReservationCreated) WithPayload(payload *models.DhcpReservationInput) *POSTDhcpReservationCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the p o s t dhcp reservation created response
func (o *POSTDhcpReservationCreated) SetPayload(payload *models.DhcpReservationInput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *POSTDhcpReservationCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}