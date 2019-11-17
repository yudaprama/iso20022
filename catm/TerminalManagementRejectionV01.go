package catm

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400101 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.004.001.01 Document"`
	Message *TerminalManagementRejectionV01 `xml:"TermnlMgmtRjctn"`
}

func (d *Document00400101) AddMessage() *TerminalManagementRejectionV01 {
	d.Message = new(TerminalManagementRejectionV01)
	return d.Message
}

// The TerminalManagementRejection message is sent by the terminal manager to reject a message request sent by an acceptor, to indicate that the received message could not be processed.
type TerminalManagementRejectionV01 struct {

	// Rejection message management information.
	Header *model.Header6 `xml:"Hdr"`

	// Information related to the reject.
	Reject *model.AcceptorRejection1 `xml:"Rjct"`
}

func (t *TerminalManagementRejectionV01) AddHeader() *model.Header6 {
	t.Header = new(model.Header6)
	return t.Header
}

func (t *TerminalManagementRejectionV01) AddReject() *model.AcceptorRejection1 {
	t.Reject = new(model.AcceptorRejection1)
	return t.Reject
}
