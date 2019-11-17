package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400101 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.004.001.01 Document"`
	Message *AcceptorCompletionAdviceResponseV01 `xml:"AccptrCmpltnAdvcRspn"`
}

func (d *Document00400101) AddMessage() *AcceptorCompletionAdviceResponseV01 {
	d.Message = new(AcceptorCompletionAdviceResponseV01)
	return d.Message
}

// Scope
// The AcceptorCompletionAdviceResponse message is sent by the acquirer to acknowledge the proper receipt of an AcceptorCompletionAdvice. The message can be sent directly to the card acceptor or through an agent.
// Usage
// The AcceptorCompletionAdviceResponse message is used to acknowledge the data capture process performed by the acquirer based on the data required to carry out the financial clearing and settlement of the transaction.
type AcceptorCompletionAdviceResponseV01 struct {

	// Completion advice response message management information.
	Header *model.Header2 `xml:"Hdr"`

	// Information related to the completion advice response.
	CompletionAdviceResponse *model.AcceptorCompletionAdviceResponse1 `xml:"CmpltnAdvcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType3 `xml:"SctyTrlr"`
}

func (a *AcceptorCompletionAdviceResponseV01) AddHeader() *model.Header2 {
	a.Header = new(model.Header2)
	return a.Header
}

func (a *AcceptorCompletionAdviceResponseV01) AddCompletionAdviceResponse() *model.AcceptorCompletionAdviceResponse1 {
	a.CompletionAdviceResponse = new(model.AcceptorCompletionAdviceResponse1)
	return a.CompletionAdviceResponse
}

func (a *AcceptorCompletionAdviceResponseV01) AddSecurityTrailer() *model.ContentInformationType3 {
	a.SecurityTrailer = new(model.ContentInformationType3)
	return a.SecurityTrailer
}
