package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100102 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.02 Document"`
	Message *ATMPINManagementResponseV02 `xml:"ATMPINMgmtRspn"`
}

func (d *Document01100102) AddMessage() *ATMPINManagementResponseV02 {
	d.Message = new(ATMPINManagementResponseV02)
	return d.Message
}

// The ATMPINManagementResponse message is sent by an ATM manager or its agent to the ATM to provide the information and the outcome of the cardholder PIN operation requested in the ATMPINManagementRequest.
type ATMPINManagementResponseV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMPINManagementResponse *model.ContentInformationType10 `xml:"PrtctdATMPINMgmtRspn,omitempty"`

	// Information related to the response of an ATM PIN Management from an ATM manager.
	ATMPINManagementResponse *model.ATMPINManagementResponse2 `xml:"ATMPINMgmtRspn,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMPINManagementResponseV02) AddHeader() *model.Header31 {
	a.Header = new(model.Header31)
	return a.Header
}

func (a *ATMPINManagementResponseV02) AddProtectedATMPINManagementResponse() *model.ContentInformationType10 {
	a.ProtectedATMPINManagementResponse = new(model.ContentInformationType10)
	return a.ProtectedATMPINManagementResponse
}

func (a *ATMPINManagementResponseV02) AddATMPINManagementResponse() *model.ATMPINManagementResponse2 {
	a.ATMPINManagementResponse = new(model.ATMPINManagementResponse2)
	return a.ATMPINManagementResponse
}

func (a *ATMPINManagementResponseV02) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
