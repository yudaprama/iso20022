package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01500104 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.04 Document"`
	Message *AcceptorRejectionV04 `xml:"AccptrRjctn"`
}

func (d *Document01500104) AddMessage() *AcceptorRejectionV04 {
	d.Message = new(AcceptorRejectionV04)
	return d.Message
}

// The AcceptorRejection message is sent by the acquirer (or its agent) to reject a message request or advice sent by an acceptor (or its agent), to indicate that the received message could not be processed.
type AcceptorRejectionV04 struct {

	// Rejection message management information.
	Header *model.Header13 `xml:"Hdr"`

	// Information related to the reject.
	Reject *model.AcceptorRejection2 `xml:"Rjct"`
}

func (a *AcceptorRejectionV04) AddHeader() *model.Header13 {
	a.Header = new(model.Header13)
	return a.Header
}

func (a *AcceptorRejectionV04) AddReject() *model.AcceptorRejection2 {
	a.Reject = new(model.AcceptorRejection2)
	return a.Reject
}
