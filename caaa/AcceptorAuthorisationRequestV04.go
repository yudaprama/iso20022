package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100104 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.04 Document"`
	Message *AcceptorAuthorisationRequestV04 `xml:"AccptrAuthstnReq"`
}

func (d *Document00100104) AddMessage() *AcceptorAuthorisationRequestV04 {
	d.Message = new(AcceptorAuthorisationRequestV04)
	return d.Message
}

// The AcceptorAuthorisationRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to check with the issuer (or its agent) that the account associated to the card has the resources to fund the payment. This checking will include validation of the card data and any additional transaction data provided.
type AcceptorAuthorisationRequestV04 struct {

	// Authorisation request message management information.
	Header *model.Header10 `xml:"Hdr"`

	// Information related to the authorisation request.
	AuthorisationRequest *model.AcceptorAuthorisationRequest4 `xml:"AuthstnReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType11 `xml:"SctyTrlr"`
}

func (a *AcceptorAuthorisationRequestV04) AddHeader() *model.Header10 {
	a.Header = new(model.Header10)
	return a.Header
}

func (a *AcceptorAuthorisationRequestV04) AddAuthorisationRequest() *model.AcceptorAuthorisationRequest4 {
	a.AuthorisationRequest = new(model.AcceptorAuthorisationRequest4)
	return a.AuthorisationRequest
}

func (a *AcceptorAuthorisationRequestV04) AddSecurityTrailer() *model.ContentInformationType11 {
	a.SecurityTrailer = new(model.ContentInformationType11)
	return a.SecurityTrailer
}
