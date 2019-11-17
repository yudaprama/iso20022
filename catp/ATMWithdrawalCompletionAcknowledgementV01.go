package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400101 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.004.001.01 Document"`
	Message *ATMWithdrawalCompletionAcknowledgementV01 `xml:"ATMWdrwlCmpltnAck"`
}

func (d *Document00400101) AddMessage() *ATMWithdrawalCompletionAcknowledgementV01 {
	d.Message = new(ATMWithdrawalCompletionAcknowledgementV01)
	return d.Message
}

// The ATMWithdrawalCompletionAcknowledgement message is sent by an acquirer or its agent to an ATM to acknowledge the receipt of an ATMWithdrawalCompletionAdvice message.
type ATMWithdrawalCompletionAcknowledgementV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header21 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMWithdrawalCompletionAcknowledgement *model.ContentInformationType10 `xml:"PrtctdATMWdrwlCmpltnAck,omitempty"`

	// Information related to the acknowledgement  of an ATM withdrawal transaction from the ATM manager.
	ATMWithdrawalCompletionAcknowledgement *model.ATMWithdrawalCompletionAcknowledgement1 `xml:"ATMWdrwlCmpltnAck,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMWithdrawalCompletionAcknowledgementV01) AddHeader() *model.Header21 {
	a.Header = new(model.Header21)
	return a.Header
}

func (a *ATMWithdrawalCompletionAcknowledgementV01) AddProtectedATMWithdrawalCompletionAcknowledgement() *model.ContentInformationType10 {
	a.ProtectedATMWithdrawalCompletionAcknowledgement = new(model.ContentInformationType10)
	return a.ProtectedATMWithdrawalCompletionAcknowledgement
}

func (a *ATMWithdrawalCompletionAcknowledgementV01) AddATMWithdrawalCompletionAcknowledgement() *model.ATMWithdrawalCompletionAcknowledgement1 {
	a.ATMWithdrawalCompletionAcknowledgement = new(model.ATMWithdrawalCompletionAcknowledgement1)
	return a.ATMWithdrawalCompletionAcknowledgement
}

func (a *ATMWithdrawalCompletionAcknowledgementV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
