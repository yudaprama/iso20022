package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01500105 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.05 Document"`
	Message *AcceptorRejectionV05 `xml:"AccptrRjctn"`
}

func (d *Document01500105) AddMessage() *AcceptorRejectionV05 {
	d.Message = new(AcceptorRejectionV05)
	return d.Message
}

// The AcceptorRejection message is sent by the acquirer (or its agent) to reject a message request or advice sent by an acceptor (or its agent), to indicate that the received message could not be processed.
type AcceptorRejectionV05 struct {

	// Rejection message management information.
	Header *model.Header26 `xml:"Hdr"`

	// Information related to the reject.
	Reject *model.AcceptorRejection2 `xml:"Rjct"`
}

func (a *AcceptorRejectionV05) AddHeader() *model.Header26 {
	a.Header = new(model.Header26)
	return a.Header
}

func (a *AcceptorRejectionV05) AddReject() *model.AcceptorRejection2 {
	a.Reject = new(model.AcceptorRejection2)
	return a.Reject
}
