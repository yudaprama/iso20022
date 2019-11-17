package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01500101 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.015.001.01 Document"`
	Message *AcceptorRejectionV01 `xml:"AccptrRjctn"`
}

func (d *Document01500101) AddMessage() *AcceptorRejectionV01 {
	d.Message = new(AcceptorRejectionV01)
	return d.Message
}

// Scope
// The AcceptorRejection message is used by the acquirer to reject a message received from the card acceptor. The acquirer uses this message as a substitute to a response or an advice response message sent to the card acceptor.
// Usage
// The AcceptorRejection message is used to indicate that the received message could not be processed by the acquirer (for example, unable to read or process the message, security error, duplicate message).
type AcceptorRejectionV01 struct {

	// Rejection message management information.
	Header *model.Header1 `xml:"Hdr"`

	// Information related to the reject.
	Reject *model.AcceptorRejection1 `xml:"Rjct"`
}

func (a *AcceptorRejectionV01) AddHeader() *model.Header1 {
	a.Header = new(model.Header1)
	return a.Header
}

func (a *AcceptorRejectionV01) AddReject() *model.AcceptorRejection1 {
	a.Reject = new(model.AcceptorRejection1)
	return a.Reject
}
