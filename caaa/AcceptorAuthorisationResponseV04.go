package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200104 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Document"`
	Message *AcceptorAuthorisationResponseV04 `xml:"AccptrAuthstnRspn"`
}

func (d *Document00200104) AddMessage() *AcceptorAuthorisationResponseV04 {
	d.Message = new(AcceptorAuthorisationResponseV04)
	return d.Message
}

// The AcceptorAuthorisationResponse message is sent by the acquirer (or its agent) to an acceptor (or its agent), to return the result of the validation made by issuer about the payment transaction.
type AcceptorAuthorisationResponseV04 struct {

	// Authorisation response message management information.
	Header *model.Header10 `xml:"Hdr"`

	// Information related to the response of the authorisation.
	AuthorisationResponse *model.AcceptorAuthorisationResponse4 `xml:"AuthstnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType11 `xml:"SctyTrlr"`
}

func (a *AcceptorAuthorisationResponseV04) AddHeader() *model.Header10 {
	a.Header = new(model.Header10)
	return a.Header
}

func (a *AcceptorAuthorisationResponseV04) AddAuthorisationResponse() *model.AcceptorAuthorisationResponse4 {
	a.AuthorisationResponse = new(model.AcceptorAuthorisationResponse4)
	return a.AuthorisationResponse
}

func (a *AcceptorAuthorisationResponseV04) AddSecurityTrailer() *model.ContentInformationType11 {
	a.SecurityTrailer = new(model.ContentInformationType11)
	return a.SecurityTrailer
}
