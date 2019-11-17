package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400104 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.004.001.04 Document"`
	Message *AcceptorCompletionAdviceResponseV04 `xml:"AccptrCmpltnAdvcRspn"`
}

func (d *Document00400104) AddMessage() *AcceptorCompletionAdviceResponseV04 {
	d.Message = new(AcceptorCompletionAdviceResponseV04)
	return d.Message
}

// The AcceptorCompletionAdviceResponse message is sent by the acquirer (or its agent) to acknowledge the acceptor (or its agent) of the outcome of the payment transaction, and the transfer the  financial data of the transaction contained in the completion advice.
type AcceptorCompletionAdviceResponseV04 struct {

	// Completion advice response message management information.
	Header *model.Header11 `xml:"Hdr"`

	// Information related to the completion advice response.
	CompletionAdviceResponse *model.AcceptorCompletionAdviceResponse4 `xml:"CmpltnAdvcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType11 `xml:"SctyTrlr"`
}

func (a *AcceptorCompletionAdviceResponseV04) AddHeader() *model.Header11 {
	a.Header = new(model.Header11)
	return a.Header
}

func (a *AcceptorCompletionAdviceResponseV04) AddCompletionAdviceResponse() *model.AcceptorCompletionAdviceResponse4 {
	a.CompletionAdviceResponse = new(model.AcceptorCompletionAdviceResponse4)
	return a.CompletionAdviceResponse
}

func (a *AcceptorCompletionAdviceResponseV04) AddSecurityTrailer() *model.ContentInformationType11 {
	a.SecurityTrailer = new(model.ContentInformationType11)
	return a.SecurityTrailer
}
