package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00800103 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.008.001.03 Document"`
	Message *AcceptorCancellationAdviceResponseV03 `xml:"AccptrCxlAdvcRspn"`
}

func (d *Document00800103) AddMessage() *AcceptorCancellationAdviceResponseV03 {
	d.Message = new(AcceptorCancellationAdviceResponseV03)
	return d.Message
}

// The AcceptorCancellationAdviceResponse message is sent by the acquirer (or its agent) to acknowledge the acceptor (or its agent) about the notification of the payment cancellation.
type AcceptorCancellationAdviceResponseV03 struct {

	// Cancellation advice response message management information.
	Header *model.Header8 `xml:"Hdr"`

	// Information related to the cancellation advice response.
	CancellationAdviceResponse *model.AcceptorCancellationAdviceResponse3 `xml:"CxlAdvcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType8 `xml:"SctyTrlr"`
}

func (a *AcceptorCancellationAdviceResponseV03) AddHeader() *model.Header8 {
	a.Header = new(model.Header8)
	return a.Header
}

func (a *AcceptorCancellationAdviceResponseV03) AddCancellationAdviceResponse() *model.AcceptorCancellationAdviceResponse3 {
	a.CancellationAdviceResponse = new(model.AcceptorCancellationAdviceResponse3)
	return a.CancellationAdviceResponse
}

func (a *AcceptorCancellationAdviceResponseV03) AddSecurityTrailer() *model.ContentInformationType8 {
	a.SecurityTrailer = new(model.ContentInformationType8)
	return a.SecurityTrailer
}
