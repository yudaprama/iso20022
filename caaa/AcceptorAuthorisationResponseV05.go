package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200105 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.05 Document"`
	Message *AcceptorAuthorisationResponseV05 `xml:"AccptrAuthstnRspn"`
}

func (d *Document00200105) AddMessage() *AcceptorAuthorisationResponseV05 {
	d.Message = new(AcceptorAuthorisationResponseV05)
	return d.Message
}

// The AcceptorAuthorisationResponse message is sent by the acquirer (or its agent) to an acceptor (or its agent), to return the result of the validation made by issuer about the payment transaction.
type AcceptorAuthorisationResponseV05 struct {

	// Authorisation response message management information.
	Header *model.Header30 `xml:"Hdr"`

	// Information related to the response of the authorisation.
	AuthorisationResponse *model.AcceptorAuthorisationResponse5 `xml:"AuthstnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorAuthorisationResponseV05) AddHeader() *model.Header30 {
	a.Header = new(model.Header30)
	return a.Header
}

func (a *AcceptorAuthorisationResponseV05) AddAuthorisationResponse() *model.AcceptorAuthorisationResponse5 {
	a.AuthorisationResponse = new(model.AcceptorAuthorisationResponse5)
	return a.AuthorisationResponse
}

func (a *AcceptorAuthorisationResponseV05) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
