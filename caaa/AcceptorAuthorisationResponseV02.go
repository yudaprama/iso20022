package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200102 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.02 Document"`
	Message *AcceptorAuthorisationResponseV02 `xml:"AccptrAuthstnRspn"`
}

func (d *Document00200102) AddMessage() *AcceptorAuthorisationResponseV02 {
	d.Message = new(AcceptorAuthorisationResponseV02)
	return d.Message
}

// The AcceptorAuthorisationResponse message is sent by the acquirer (or its agent) to an acceptor (or its agent), to return the result of the validation made by issuer about the payment transaction.
type AcceptorAuthorisationResponseV02 struct {

	// Authorisation response message management information.
	Header *model.Header1 `xml:"Hdr"`

	// Information related to the response of the authorisation.
	AuthorisationResponse *model.AcceptorAuthorisationResponse2 `xml:"AuthstnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType6 `xml:"SctyTrlr"`
}

func (a *AcceptorAuthorisationResponseV02) AddHeader() *model.Header1 {
	a.Header = new(model.Header1)
	return a.Header
}

func (a *AcceptorAuthorisationResponseV02) AddAuthorisationResponse() *model.AcceptorAuthorisationResponse2 {
	a.AuthorisationResponse = new(model.AcceptorAuthorisationResponse2)
	return a.AuthorisationResponse
}

func (a *AcceptorAuthorisationResponseV02) AddSecurityTrailer() *model.ContentInformationType6 {
	a.SecurityTrailer = new(model.ContentInformationType6)
	return a.SecurityTrailer
}
