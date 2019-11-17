package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100101 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:catp.011.001.01 Document"`
	Message *ATMPINManagementResponseV01 `xml:"ATMPINMgmtRspn"`
}

func (d *Document01100101) AddMessage() *ATMPINManagementResponseV01 {
	d.Message = new(ATMPINManagementResponseV01)
	return d.Message
}

// The ATMPINManagementResponse message is sent by an ATM manager or its agent to the ATM to provide the information and the outcome of the cardholder PIN operation requested in the ATMPINManagementRequest.
type ATMPINManagementResponseV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header20 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMPINManagementResponse *model.ContentInformationType10 `xml:"PrtctdATMPINMgmtRspn,omitempty"`

	// Information related to the response of an ATM PIN Management from an ATM manager.
	ATMPINManagementResponse *model.ATMPINManagementResponse1 `xml:"ATMPINMgmtRspn,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMPINManagementResponseV01) AddHeader() *model.Header20 {
	a.Header = new(model.Header20)
	return a.Header
}

func (a *ATMPINManagementResponseV01) AddProtectedATMPINManagementResponse() *model.ContentInformationType10 {
	a.ProtectedATMPINManagementResponse = new(model.ContentInformationType10)
	return a.ProtectedATMPINManagementResponse
}

func (a *ATMPINManagementResponseV01) AddATMPINManagementResponse() *model.ATMPINManagementResponse1 {
	a.ATMPINManagementResponse = new(model.ATMPINManagementResponse1)
	return a.ATMPINManagementResponse
}

func (a *ATMPINManagementResponseV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
