package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500103 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.005.001.03 Document"`
	Message *AcceptorCancellationRequestV03 `xml:"AccptrCxlReq"`
}

func (d *Document00500103) AddMessage() *AcceptorCancellationRequestV03 {
	d.Message = new(AcceptorCancellationRequestV03)
	return d.Message
}

// The AcceptorCancellationRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to request the cancellation of a successfully completed transaction. Cancellation should only occur before the transaction has been cleared.
//
//
type AcceptorCancellationRequestV03 struct {

	// Cancellation request message management information.
	Header *model.Header7 `xml:"Hdr"`

	// Information related to the cancellation request.
	CancellationRequest *model.AcceptorCancellationRequest3 `xml:"CxlReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType8 `xml:"SctyTrlr"`
}

func (a *AcceptorCancellationRequestV03) AddHeader() *model.Header7 {
	a.Header = new(model.Header7)
	return a.Header
}

func (a *AcceptorCancellationRequestV03) AddCancellationRequest() *model.AcceptorCancellationRequest3 {
	a.CancellationRequest = new(model.AcceptorCancellationRequest3)
	return a.CancellationRequest
}

func (a *AcceptorCancellationRequestV03) AddSecurityTrailer() *model.ContentInformationType8 {
	a.SecurityTrailer = new(model.ContentInformationType8)
	return a.SecurityTrailer
}
