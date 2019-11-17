package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100103 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.03 Document"`
	Message *AcceptorAuthorisationRequestV03 `xml:"AccptrAuthstnReq"`
}

func (d *Document00100103) AddMessage() *AcceptorAuthorisationRequestV03 {
	d.Message = new(AcceptorAuthorisationRequestV03)
	return d.Message
}

// The AcceptorAuthorisationRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to check with the issuer (or its agent) that the account associated to the card has the resources to fund the payment. This checking will include validation of the card data and any additional transaction data provided.
type AcceptorAuthorisationRequestV03 struct {

	// Authorisation request message management information.
	Header *model.Header7 `xml:"Hdr"`

	// Information related to the authorisation request.
	AuthorisationRequest *model.AcceptorAuthorisationRequest3 `xml:"AuthstnReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType8 `xml:"SctyTrlr"`
}

func (a *AcceptorAuthorisationRequestV03) AddHeader() *model.Header7 {
	a.Header = new(model.Header7)
	return a.Header
}

func (a *AcceptorAuthorisationRequestV03) AddAuthorisationRequest() *model.AcceptorAuthorisationRequest3 {
	a.AuthorisationRequest = new(model.AcceptorAuthorisationRequest3)
	return a.AuthorisationRequest
}

func (a *AcceptorAuthorisationRequestV03) AddSecurityTrailer() *model.ContentInformationType8 {
	a.SecurityTrailer = new(model.ContentInformationType8)
	return a.SecurityTrailer
}
