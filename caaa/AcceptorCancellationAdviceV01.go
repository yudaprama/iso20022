package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700101 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.01 Document"`
	Message *AcceptorCancellationAdviceV01 `xml:"AccptrCxlAdvc"`
}

func (d *Document00700101) AddMessage() *AcceptorCancellationAdviceV01 {
	d.Message = new(AcceptorCancellationAdviceV01)
	return d.Message
}

// Scope
// The AcceptorCancellationAdvice message is sent by a card acceptor to notify the cancellation of a successfully completed card payment transaction. The message can be sent directly to the acquirer or through an agent.
// Usage
// The AcceptorCancellationAdvice message is sent by the card acceptor to an acquirer when the acquirer did not receive a correct response to the AcceptorCompletionAdvice message.
type AcceptorCancellationAdviceV01 struct {

	// Cancellation advice message management information.
	Header *model.Header2 `xml:"Hdr"`

	// Information related to the cancellation advice.
	CancellationAdvice *model.AcceptorCancellationAdvice1 `xml:"CxlAdvc"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType3 `xml:"SctyTrlr"`
}

func (a *AcceptorCancellationAdviceV01) AddHeader() *model.Header2 {
	a.Header = new(model.Header2)
	return a.Header
}

func (a *AcceptorCancellationAdviceV01) AddCancellationAdvice() *model.AcceptorCancellationAdvice1 {
	a.CancellationAdvice = new(model.AcceptorCancellationAdvice1)
	return a.CancellationAdvice
}

func (a *AcceptorCancellationAdviceV01) AddSecurityTrailer() *model.ContentInformationType3 {
	a.SecurityTrailer = new(model.ContentInformationType3)
	return a.SecurityTrailer
}
