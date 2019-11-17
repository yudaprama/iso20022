package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400102 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:catp.004.001.02 Document"`
	Message *ATMWithdrawalCompletionAcknowledgementV02 `xml:"ATMWdrwlCmpltnAck"`
}

func (d *Document00400102) AddMessage() *ATMWithdrawalCompletionAcknowledgementV02 {
	d.Message = new(ATMWithdrawalCompletionAcknowledgementV02)
	return d.Message
}

// The ATMWithdrawalCompletionAcknowledgement message is sent by an acquirer or its agent to an ATM to acknowledge the receipt of an ATMWithdrawalCompletionAdvice message.
type ATMWithdrawalCompletionAcknowledgementV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header32 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMWithdrawalCompletionAcknowledgement *model.ContentInformationType10 `xml:"PrtctdATMWdrwlCmpltnAck,omitempty"`

	// Information related to the acknowledgement  of an ATM withdrawal transaction from the ATM manager.
	ATMWithdrawalCompletionAcknowledgement *model.ATMWithdrawalCompletionAcknowledgement2 `xml:"ATMWdrwlCmpltnAck,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMWithdrawalCompletionAcknowledgementV02) AddHeader() *model.Header32 {
	a.Header = new(model.Header32)
	return a.Header
}

func (a *ATMWithdrawalCompletionAcknowledgementV02) AddProtectedATMWithdrawalCompletionAcknowledgement() *model.ContentInformationType10 {
	a.ProtectedATMWithdrawalCompletionAcknowledgement = new(model.ContentInformationType10)
	return a.ProtectedATMWithdrawalCompletionAcknowledgement
}

func (a *ATMWithdrawalCompletionAcknowledgementV02) AddATMWithdrawalCompletionAcknowledgement() *model.ATMWithdrawalCompletionAcknowledgement2 {
	a.ATMWithdrawalCompletionAcknowledgement = new(model.ATMWithdrawalCompletionAcknowledgement2)
	return a.ATMWithdrawalCompletionAcknowledgement
}

func (a *ATMWithdrawalCompletionAcknowledgementV02) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
