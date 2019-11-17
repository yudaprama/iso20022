package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01500103 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.03 Document"`
	Message *AcceptorRejectionV03 `xml:"AccptrRjctn"`
}

func (d *Document01500103) AddMessage() *AcceptorRejectionV03 {
	d.Message = new(AcceptorRejectionV03)
	return d.Message
}

// The AcceptorRejection message is sent by the acquirer (or its agent) to reject a message request or advice sent by an acceptor (or its agent), to indicate that the received message could not be processed.
type AcceptorRejectionV03 struct {

	// Rejection message management information.
	Header *model.Header9 `xml:"Hdr"`

	// Information related to the reject.
	Reject *model.AcceptorRejection2 `xml:"Rjct"`
}

func (a *AcceptorRejectionV03) AddHeader() *model.Header9 {
	a.Header = new(model.Header9)
	return a.Header
}

func (a *AcceptorRejectionV03) AddReject() *model.AcceptorRejection2 {
	a.Reject = new(model.AcceptorRejection2)
	return a.Reject
}
