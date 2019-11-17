package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700104 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.04 Document"`
	Message *AcceptorCancellationAdviceV04 `xml:"AccptrCxlAdvc"`
}

func (d *Document00700104) AddMessage() *AcceptorCancellationAdviceV04 {
	d.Message = new(AcceptorCancellationAdviceV04)
	return d.Message
}

// The AcceptorCancellationAdvice message is sent by an acceptor (or its agent) to notify the acquirer (or its agent) of the cancellation of a successfully completed transaction. The transaction has been completed without financial transfer, or the acceptor is aware that the transaction was not cleared by the acquirer.
type AcceptorCancellationAdviceV04 struct {

	// Cancellation advice message management information.
	Header *model.Header11 `xml:"Hdr"`

	// Information related to the cancellation advice.
	CancellationAdvice *model.AcceptorCancellationAdvice4 `xml:"CxlAdvc"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType11 `xml:"SctyTrlr"`
}

func (a *AcceptorCancellationAdviceV04) AddHeader() *model.Header11 {
	a.Header = new(model.Header11)
	return a.Header
}

func (a *AcceptorCancellationAdviceV04) AddCancellationAdvice() *model.AcceptorCancellationAdvice4 {
	a.CancellationAdvice = new(model.AcceptorCancellationAdvice4)
	return a.CancellationAdvice
}

func (a *AcceptorCancellationAdviceV04) AddSecurityTrailer() *model.ContentInformationType11 {
	a.SecurityTrailer = new(model.ContentInformationType11)
	return a.SecurityTrailer
}
