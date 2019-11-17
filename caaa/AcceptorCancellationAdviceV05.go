package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700105 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.05 Document"`
	Message *AcceptorCancellationAdviceV05 `xml:"AccptrCxlAdvc"`
}

func (d *Document00700105) AddMessage() *AcceptorCancellationAdviceV05 {
	d.Message = new(AcceptorCancellationAdviceV05)
	return d.Message
}

// The AcceptorCancellationAdvice message is sent by an acceptor (or its agent) to notify the acquirer (or its agent) of the cancellation of a successfully completed transaction. The transaction has been completed without financial transfer, or the acceptor is aware that the transaction was not cleared by the acquirer.
type AcceptorCancellationAdviceV05 struct {

	// Cancellation advice message management information.
	Header *model.Header24 `xml:"Hdr"`

	// Information related to the cancellation advice.
	CancellationAdvice *model.AcceptorCancellationAdvice5 `xml:"CxlAdvc"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorCancellationAdviceV05) AddHeader() *model.Header24 {
	a.Header = new(model.Header24)
	return a.Header
}

func (a *AcceptorCancellationAdviceV05) AddCancellationAdvice() *model.AcceptorCancellationAdvice5 {
	a.CancellationAdvice = new(model.AcceptorCancellationAdvice5)
	return a.CancellationAdvice
}

func (a *AcceptorCancellationAdviceV05) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
