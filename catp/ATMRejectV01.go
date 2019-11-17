package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500101 struct {
	XMLName xml.Name      `xml:"urn:iso:std:iso:20022:tech:xsd:catp.005.001.01 Document"`
	Message *ATMRejectV01 `xml:"ATMRjct"`
}

func (d *Document00500101) AddMessage() *ATMRejectV01 {
	d.Message = new(ATMRejectV01)
	return d.Message
}

// The ATMReject message is sent by any entity to reject a received message.
type ATMRejectV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header22 `xml:"Hdr"`

	// Information related to the reject of a message from an ATM or an ATM manager.
	ATMReject *model.ATMReject1 `xml:"ATMRjct"`
}

func (a *ATMRejectV01) AddHeader() *model.Header22 {
	a.Header = new(model.Header22)
	return a.Header
}

func (a *ATMRejectV01) AddATMReject() *model.ATMReject1 {
	a.ATMReject = new(model.ATMReject1)
	return a.ATMReject
}
