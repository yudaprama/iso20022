package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00800104 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.008.001.04 Document"`
	Message *AcceptorCancellationAdviceResponseV04 `xml:"AccptrCxlAdvcRspn"`
}

func (d *Document00800104) AddMessage() *AcceptorCancellationAdviceResponseV04 {
	d.Message = new(AcceptorCancellationAdviceResponseV04)
	return d.Message
}

// The AcceptorCancellationAdviceResponse message is sent by the acquirer (or its agent) to acknowledge the acceptor (or its agent) about the notification of the payment cancellation.
type AcceptorCancellationAdviceResponseV04 struct {

	// Cancellation advice response message management information.
	Header *model.Header11 `xml:"Hdr"`

	// Information related to the cancellation advice response.
	CancellationAdviceResponse *model.AcceptorCancellationAdviceResponse4 `xml:"CxlAdvcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType11 `xml:"SctyTrlr"`
}

func (a *AcceptorCancellationAdviceResponseV04) AddHeader() *model.Header11 {
	a.Header = new(model.Header11)
	return a.Header
}

func (a *AcceptorCancellationAdviceResponseV04) AddCancellationAdviceResponse() *model.AcceptorCancellationAdviceResponse4 {
	a.CancellationAdviceResponse = new(model.AcceptorCancellationAdviceResponse4)
	return a.CancellationAdviceResponse
}

func (a *AcceptorCancellationAdviceResponseV04) AddSecurityTrailer() *model.ContentInformationType11 {
	a.SecurityTrailer = new(model.ContentInformationType11)
	return a.SecurityTrailer
}
