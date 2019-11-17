package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300104 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.003.001.04 Document"`
	Message *AcceptorCompletionAdviceV04 `xml:"AccptrCmpltnAdvc"`
}

func (d *Document00300104) AddMessage() *AcceptorCompletionAdviceV04 {
	d.Message = new(AcceptorCompletionAdviceV04)
	return d.Message
}

// The AcceptorCompletionAdvice message is sent by an acceptor (or its agent) to notify the acquirer (or its agent) of the outcome of the payment at the acceptor, and to transfer the  financial data of the transaction to the acquirer (capture).
// A AcceptorCompletionAdvice message is also sent to reverse an approved authorisation and any associated financial transfer (capture), if the card payment transaction could not be completed successfully.
type AcceptorCompletionAdviceV04 struct {

	// Completion advice message management information.
	Header *model.Header11 `xml:"Hdr"`

	// Information related to the completion advice.
	CompletionAdvice *model.AcceptorCompletionAdvice4 `xml:"CmpltnAdvc"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType11 `xml:"SctyTrlr"`
}

func (a *AcceptorCompletionAdviceV04) AddHeader() *model.Header11 {
	a.Header = new(model.Header11)
	return a.Header
}

func (a *AcceptorCompletionAdviceV04) AddCompletionAdvice() *model.AcceptorCompletionAdvice4 {
	a.CompletionAdvice = new(model.AcceptorCompletionAdvice4)
	return a.CompletionAdvice
}

func (a *AcceptorCompletionAdviceV04) AddSecurityTrailer() *model.ContentInformationType11 {
	a.SecurityTrailer = new(model.ContentInformationType11)
	return a.SecurityTrailer
}
