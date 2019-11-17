package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500102 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.005.001.02 Document"`
	Message *AcceptorCancellationRequestV02 `xml:"AccptrCxlReq"`
}

func (d *Document00500102) AddMessage() *AcceptorCancellationRequestV02 {
	d.Message = new(AcceptorCancellationRequestV02)
	return d.Message
}

// The AcceptorCancellationRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to request the cancellation of a successfully completed transaction. Cancellation should only occur before the transaction has been cleared.
//
//
type AcceptorCancellationRequestV02 struct {

	// Cancellation request message management information.
	Header *model.Header1 `xml:"Hdr"`

	// Information related to the cancellation request.
	CancellationRequest *model.AcceptorCancellationRequest2 `xml:"CxlReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType6 `xml:"SctyTrlr"`
}

func (a *AcceptorCancellationRequestV02) AddHeader() *model.Header1 {
	a.Header = new(model.Header1)
	return a.Header
}

func (a *AcceptorCancellationRequestV02) AddCancellationRequest() *model.AcceptorCancellationRequest2 {
	a.CancellationRequest = new(model.AcceptorCancellationRequest2)
	return a.CancellationRequest
}

func (a *AcceptorCancellationRequestV02) AddSecurityTrailer() *model.ContentInformationType6 {
	a.SecurityTrailer = new(model.ContentInformationType6)
	return a.SecurityTrailer
}
