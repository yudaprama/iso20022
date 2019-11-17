package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200103 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.03 Document"`
	Message *AcceptorAuthorisationResponseV03 `xml:"AccptrAuthstnRspn"`
}

func (d *Document00200103) AddMessage() *AcceptorAuthorisationResponseV03 {
	d.Message = new(AcceptorAuthorisationResponseV03)
	return d.Message
}

// The AcceptorAuthorisationResponse message is sent by the acquirer (or its agent) to an acceptor (or its agent), to return the result of the validation made by issuer about the payment transaction.
type AcceptorAuthorisationResponseV03 struct {

	// Authorisation response message management information.
	Header *model.Header7 `xml:"Hdr"`

	// Information related to the response of the authorisation.
	AuthorisationResponse *model.AcceptorAuthorisationResponse3 `xml:"AuthstnRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType8 `xml:"SctyTrlr"`
}

func (a *AcceptorAuthorisationResponseV03) AddHeader() *model.Header7 {
	a.Header = new(model.Header7)
	return a.Header
}

func (a *AcceptorAuthorisationResponseV03) AddAuthorisationResponse() *model.AcceptorAuthorisationResponse3 {
	a.AuthorisationResponse = new(model.AcceptorAuthorisationResponse3)
	return a.AuthorisationResponse
}

func (a *AcceptorAuthorisationResponseV03) AddSecurityTrailer() *model.ContentInformationType8 {
	a.SecurityTrailer = new(model.ContentInformationType8)
	return a.SecurityTrailer
}
