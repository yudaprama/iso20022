package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500105 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.005.001.05 Document"`
	Message *AcceptorCancellationRequestV05 `xml:"AccptrCxlReq"`
}

func (d *Document00500105) AddMessage() *AcceptorCancellationRequestV05 {
	d.Message = new(AcceptorCancellationRequestV05)
	return d.Message
}

// The AcceptorCancellationRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to request the cancellation of a successfully completed transaction. Cancellation should only occur before the transaction has been cleared.
//
//
type AcceptorCancellationRequestV05 struct {

	// Cancellation request message management information.
	Header *model.Header30 `xml:"Hdr"`

	// Information related to the cancellation request.
	CancellationRequest *model.AcceptorCancellationRequest5 `xml:"CxlReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorCancellationRequestV05) AddHeader() *model.Header30 {
	a.Header = new(model.Header30)
	return a.Header
}

func (a *AcceptorCancellationRequestV05) AddCancellationRequest() *model.AcceptorCancellationRequest5 {
	a.CancellationRequest = new(model.AcceptorCancellationRequest5)
	return a.CancellationRequest
}

func (a *AcceptorCancellationRequestV05) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
