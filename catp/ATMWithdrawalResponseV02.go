package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200102 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.002.001.02 Document"`
	Message *ATMWithdrawalResponseV02 `xml:"ATMWdrwlRspn"`
}

func (d *Document00200102) AddMessage() *ATMWithdrawalResponseV02 {
	d.Message = new(ATMWithdrawalResponseV02)
	return d.Message
}

// The ATMWithdrawalResponse message is sent by an acquirer or its agent to an ATM in response to the ATMWithdrawalRequest to inform the ATM of the approval or decline of the withdrawal transaction.
type ATMWithdrawalResponseV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMWithdrawalResponse *model.ContentInformationType10 `xml:"PrtctdATMWdrwlRspn,omitempty"`

	// Information related to the response of an ATM withdrawal transaction from an ATM manager.
	ATMWithdrawalResponse *model.ATMWithdrawalResponse2 `xml:"ATMWdrwlRspn,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMWithdrawalResponseV02) AddHeader() *model.Header31 {
	a.Header = new(model.Header31)
	return a.Header
}

func (a *ATMWithdrawalResponseV02) AddProtectedATMWithdrawalResponse() *model.ContentInformationType10 {
	a.ProtectedATMWithdrawalResponse = new(model.ContentInformationType10)
	return a.ProtectedATMWithdrawalResponse
}

func (a *ATMWithdrawalResponseV02) AddATMWithdrawalResponse() *model.ATMWithdrawalResponse2 {
	a.ATMWithdrawalResponse = new(model.ATMWithdrawalResponse2)
	return a.ATMWithdrawalResponse
}

func (a *ATMWithdrawalResponseV02) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
