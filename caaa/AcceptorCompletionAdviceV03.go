package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300103 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.003.001.03 Document"`
	Message *AcceptorCompletionAdviceV03 `xml:"AccptrCmpltnAdvc"`
}

func (d *Document00300103) AddMessage() *AcceptorCompletionAdviceV03 {
	d.Message = new(AcceptorCompletionAdviceV03)
	return d.Message
}

// The AcceptorCompletionAdvice message is sent by an acceptor (or its agent) to notify the acquirer (or its agent) of the outcome of the payment at the acceptor, and to transfer the  financial data of the transaction to the acquirer (capture).
// A AcceptorCompletionAdvice message is also sent to reverse an approved authorisation and any associated financial transfer (capture), if the card payment transaction could not be completed successfully.
type AcceptorCompletionAdviceV03 struct {

	// Completion advice message management information.
	Header *model.Header8 `xml:"Hdr"`

	// Information related to the completion advice.
	CompletionAdvice *model.AcceptorCompletionAdvice3 `xml:"CmpltnAdvc"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType8 `xml:"SctyTrlr"`
}

func (a *AcceptorCompletionAdviceV03) AddHeader() *model.Header8 {
	a.Header = new(model.Header8)
	return a.Header
}

func (a *AcceptorCompletionAdviceV03) AddCompletionAdvice() *model.AcceptorCompletionAdvice3 {
	a.CompletionAdvice = new(model.AcceptorCompletionAdvice3)
	return a.CompletionAdvice
}

func (a *AcceptorCompletionAdviceV03) AddSecurityTrailer() *model.ContentInformationType8 {
	a.SecurityTrailer = new(model.ContentInformationType8)
	return a.SecurityTrailer
}
