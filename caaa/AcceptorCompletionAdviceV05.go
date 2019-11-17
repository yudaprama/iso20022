package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300105 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.003.001.05 Document"`
	Message *AcceptorCompletionAdviceV05 `xml:"AccptrCmpltnAdvc"`
}

func (d *Document00300105) AddMessage() *AcceptorCompletionAdviceV05 {
	d.Message = new(AcceptorCompletionAdviceV05)
	return d.Message
}

// The AcceptorCompletionAdvice message is sent by an acceptor (or its agent) to notify the acquirer (or its agent) of the outcome of the payment at the acceptor, and to transfer the  financial data of the transaction to the acquirer (capture).
// A AcceptorCompletionAdvice message is also sent to reverse an approved authorisation and any associated financial transfer (capture), if the card payment transaction could not be completed successfully.
type AcceptorCompletionAdviceV05 struct {

	// Completion advice message management information.
	Header *model.Header24 `xml:"Hdr"`

	// Information related to the completion advice.
	CompletionAdvice *model.AcceptorCompletionAdvice5 `xml:"CmpltnAdvc"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorCompletionAdviceV05) AddHeader() *model.Header24 {
	a.Header = new(model.Header24)
	return a.Header
}

func (a *AcceptorCompletionAdviceV05) AddCompletionAdvice() *model.AcceptorCompletionAdvice5 {
	a.CompletionAdvice = new(model.AcceptorCompletionAdvice5)
	return a.CompletionAdvice
}

func (a *AcceptorCompletionAdviceV05) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
