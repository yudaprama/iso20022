package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300101 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:catp.013.001.01 Document"`
	Message *ATMDepositResponseV01 `xml:"ATMDpstRspn"`
}

func (d *Document01300101) AddMessage() *ATMDepositResponseV01 {
	d.Message = new(ATMDepositResponseV01)
	return d.Message
}

// The ATMDepositResponse message is sent by an ATM manager or its agent to inform the ATM of the approval or decline of the deposit transaction.
type ATMDepositResponseV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMDepositResponse *model.ContentInformationType10 `xml:"PrtctdATMDpstRspn,omitempty"`

	// Response to a deposit request.
	ATMDepositResponse *model.ATMDepositResponse1 `xml:"ATMDpstRspn,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMDepositResponseV01) AddHeader() *model.Header31 {
	a.Header = new(model.Header31)
	return a.Header
}

func (a *ATMDepositResponseV01) AddProtectedATMDepositResponse() *model.ContentInformationType10 {
	a.ProtectedATMDepositResponse = new(model.ContentInformationType10)
	return a.ProtectedATMDepositResponse
}

func (a *ATMDepositResponseV01) AddATMDepositResponse() *model.ATMDepositResponse1 {
	a.ATMDepositResponse = new(model.ATMDepositResponse1)
	return a.ATMDepositResponse
}

func (a *ATMDepositResponseV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
