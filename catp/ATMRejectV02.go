package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500102 struct {
	XMLName xml.Name      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.02 Document"`
	Message *ATMRejectV02 `xml:"ATMRjct"`
}

func (d *Document00500102) AddMessage() *ATMRejectV02 {
	d.Message = new(ATMRejectV02)
	return d.Message
}

// The ATMReject message is sent by any entity to reject a received message.
type ATMRejectV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header33 `xml:"Hdr"`

	// Information related to the reject of a message from an ATM or an ATM manager.
	ATMReject *model.ATMReject2 `xml:"ATMRjct"`
}

func (a *ATMRejectV02) AddHeader() *model.Header33 {
	a.Header = new(model.Header33)
	return a.Header
}

func (a *ATMRejectV02) AddATMReject() *model.ATMReject2 {
	a.ATMReject = new(model.ATMReject2)
	return a.ATMReject
}
