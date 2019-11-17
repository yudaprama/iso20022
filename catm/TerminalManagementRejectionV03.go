package catm

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400103 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.004.001.03 Document"`
	Message *TerminalManagementRejectionV03 `xml:"TermnlMgmtRjctn"`
}

func (d *Document00400103) AddMessage() *TerminalManagementRejectionV03 {
	d.Message = new(TerminalManagementRejectionV03)
	return d.Message
}

// The TerminalManagementRejection message is sent by the terminal manager to reject a message request sent by an acceptor, to indicate that the received message could not be processed.
type TerminalManagementRejectionV03 struct {

	// Rejection message management information.
	Header *model.Header15 `xml:"Hdr"`

	// Information related to the reject.
	Reject *model.AcceptorRejection3 `xml:"Rjct"`
}

func (t *TerminalManagementRejectionV03) AddHeader() *model.Header15 {
	t.Header = new(model.Header15)
	return t.Header
}

func (t *TerminalManagementRejectionV03) AddReject() *model.AcceptorRejection3 {
	t.Reject = new(model.AcceptorRejection3)
	return t.Reject
}
