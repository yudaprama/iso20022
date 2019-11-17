package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700102 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.02 Document"`
	Message *AcceptorCancellationAdviceV02 `xml:"AccptrCxlAdvc"`
}

func (d *Document00700102) AddMessage() *AcceptorCancellationAdviceV02 {
	d.Message = new(AcceptorCancellationAdviceV02)
	return d.Message
}

// The AcceptorCancellationAdvice message is sent by an acceptor (or its agent) to notify the acquirer (or its agent) of the cancellation of a successfully completed transaction. The transaction has been completed without financial transfer, or the acceptor is aware that the transaction was not cleared by the acquirer.
type AcceptorCancellationAdviceV02 struct {

	// Cancellation advice message management information.
	Header *model.Header2 `xml:"Hdr"`

	// Information related to the cancellation advice.
	CancellationAdvice *model.AcceptorCancellationAdvice2 `xml:"CxlAdvc"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType6 `xml:"SctyTrlr"`
}

func (a *AcceptorCancellationAdviceV02) AddHeader() *model.Header2 {
	a.Header = new(model.Header2)
	return a.Header
}

func (a *AcceptorCancellationAdviceV02) AddCancellationAdvice() *model.AcceptorCancellationAdvice2 {
	a.CancellationAdvice = new(model.AcceptorCancellationAdvice2)
	return a.CancellationAdvice
}

func (a *AcceptorCancellationAdviceV02) AddSecurityTrailer() *model.ContentInformationType6 {
	a.SecurityTrailer = new(model.ContentInformationType6)
	return a.SecurityTrailer
}
