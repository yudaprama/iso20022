package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400102 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.004.001.02 Document"`
	Message *AcceptorCompletionAdviceResponseV02 `xml:"AccptrCmpltnAdvcRspn"`
}

func (d *Document00400102) AddMessage() *AcceptorCompletionAdviceResponseV02 {
	d.Message = new(AcceptorCompletionAdviceResponseV02)
	return d.Message
}

// The AcceptorCompletionAdviceResponse message is sent by the acquirer (or its agent) to acknowledge the acceptor (or its agent) of the outcome of the payment transaction, and the transfer the  financial data of the transaction contained in the completion advice.
type AcceptorCompletionAdviceResponseV02 struct {

	// Completion advice response message management information.
	Header *model.Header2 `xml:"Hdr"`

	// Information related to the completion advice response.
	CompletionAdviceResponse *model.AcceptorCompletionAdviceResponse2 `xml:"CmpltnAdvcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType6 `xml:"SctyTrlr"`
}

func (a *AcceptorCompletionAdviceResponseV02) AddHeader() *model.Header2 {
	a.Header = new(model.Header2)
	return a.Header
}

func (a *AcceptorCompletionAdviceResponseV02) AddCompletionAdviceResponse() *model.AcceptorCompletionAdviceResponse2 {
	a.CompletionAdviceResponse = new(model.AcceptorCompletionAdviceResponse2)
	return a.CompletionAdviceResponse
}

func (a *AcceptorCompletionAdviceResponseV02) AddSecurityTrailer() *model.ContentInformationType6 {
	a.SecurityTrailer = new(model.ContentInformationType6)
	return a.SecurityTrailer
}
