package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200101 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:catp.002.001.01 Document"`
	Message *ATMWithdrawalResponseV01 `xml:"ATMWdrwlRspn"`
}

func (d *Document00200101) AddMessage() *ATMWithdrawalResponseV01 {
	d.Message = new(ATMWithdrawalResponseV01)
	return d.Message
}

// The ATMWithdrawalResponse message is sent by an acquirer or its agent to an ATM in response to the ATMWithdrawalRequest to inform the ATM of the approval or decline of the withdrawal transaction.
type ATMWithdrawalResponseV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header20 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMWithdrawalResponse *model.ContentInformationType10 `xml:"PrtctdATMWdrwlRspn,omitempty"`

	// Information related to the response of an ATM withdrawal transaction from an ATM manager.
	ATMWithdrawalResponse *model.ATMWithdrawalResponse1 `xml:"ATMWdrwlRspn,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMWithdrawalResponseV01) AddHeader() *model.Header20 {
	a.Header = new(model.Header20)
	return a.Header
}

func (a *ATMWithdrawalResponseV01) AddProtectedATMWithdrawalResponse() *model.ContentInformationType10 {
	a.ProtectedATMWithdrawalResponse = new(model.ContentInformationType10)
	return a.ProtectedATMWithdrawalResponse
}

func (a *ATMWithdrawalResponseV01) AddATMWithdrawalResponse() *model.ATMWithdrawalResponse1 {
	a.ATMWithdrawalResponse = new(model.ATMWithdrawalResponse1)
	return a.ATMWithdrawalResponse
}

func (a *ATMWithdrawalResponseV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
