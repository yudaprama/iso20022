package catm

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400104 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.004.001.04 Document"`
	Message *TerminalManagementRejectionV04 `xml:"TermnlMgmtRjctn"`
}

func (d *Document00400104) AddMessage() *TerminalManagementRejectionV04 {
	d.Message = new(TerminalManagementRejectionV04)
	return d.Message
}

// The TerminalManagementRejection message is sent by the terminal manager to reject a message request sent by an acceptor, to indicate that the received message could not be processed.
type TerminalManagementRejectionV04 struct {

	// Rejection message management information.
	Header *model.Header28 `xml:"Hdr"`

	// Information related to the reject.
	Reject *model.AcceptorRejection3 `xml:"Rjct"`
}

func (t *TerminalManagementRejectionV04) AddHeader() *model.Header28 {
	t.Header = new(model.Header28)
	return t.Header
}

func (t *TerminalManagementRejectionV04) AddReject() *model.AcceptorRejection3 {
	t.Reject = new(model.AcceptorRejection3)
	return t.Reject
}
