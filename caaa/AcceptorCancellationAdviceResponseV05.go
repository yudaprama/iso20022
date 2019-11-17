package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00800105 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.008.001.05 Document"`
	Message *AcceptorCancellationAdviceResponseV05 `xml:"AccptrCxlAdvcRspn"`
}

func (d *Document00800105) AddMessage() *AcceptorCancellationAdviceResponseV05 {
	d.Message = new(AcceptorCancellationAdviceResponseV05)
	return d.Message
}

// The AcceptorCancellationAdviceResponse message is sent by the acquirer (or its agent) to acknowledge the acceptor (or its agent) about the notification of the payment cancellation.
type AcceptorCancellationAdviceResponseV05 struct {

	// Cancellation advice response message management information.
	Header *model.Header24 `xml:"Hdr"`

	// Information related to the cancellation advice response.
	CancellationAdviceResponse *model.AcceptorCancellationAdviceResponse5 `xml:"CxlAdvcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorCancellationAdviceResponseV05) AddHeader() *model.Header24 {
	a.Header = new(model.Header24)
	return a.Header
}

func (a *AcceptorCancellationAdviceResponseV05) AddCancellationAdviceResponse() *model.AcceptorCancellationAdviceResponse5 {
	a.CancellationAdviceResponse = new(model.AcceptorCancellationAdviceResponse5)
	return a.CancellationAdviceResponse
}

func (a *AcceptorCancellationAdviceResponseV05) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
