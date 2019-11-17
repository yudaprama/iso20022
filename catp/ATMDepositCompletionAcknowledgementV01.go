package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01500101 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.015.001.01 Document"`
	Message *ATMDepositCompletionAcknowledgementV01 `xml:"ATMDpstCmpltnAck"`
}

func (d *Document01500101) AddMessage() *ATMDepositCompletionAcknowledgementV01 {
	d.Message = new(ATMDepositCompletionAcknowledgementV01)
	return d.Message
}

// The ATMDepositCompletionAcknowledgement message is sent by an acquirer or its agent to an ATM to acknowledge the receipt of an ATMDepositCompletionAdvice message.
type ATMDepositCompletionAcknowledgementV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header32 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMDepositCompletionAcknowledgement *model.ContentInformationType10 `xml:"PrtctdATMDpstCmpltnAck,omitempty"`

	// Information related to the acknowledgement  of an ATM deposit transaction from the ATM manager.
	ATMDepositCompletionAcknowledgement *model.ATMDepositCompletionAcknowledgement1 `xml:"ATMDpstCmpltnAck,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMDepositCompletionAcknowledgementV01) AddHeader() *model.Header32 {
	a.Header = new(model.Header32)
	return a.Header
}

func (a *ATMDepositCompletionAcknowledgementV01) AddProtectedATMDepositCompletionAcknowledgement() *model.ContentInformationType10 {
	a.ProtectedATMDepositCompletionAcknowledgement = new(model.ContentInformationType10)
	return a.ProtectedATMDepositCompletionAcknowledgement
}

func (a *ATMDepositCompletionAcknowledgementV01) AddATMDepositCompletionAcknowledgement() *model.ATMDepositCompletionAcknowledgement1 {
	a.ATMDepositCompletionAcknowledgement = new(model.ATMDepositCompletionAcknowledgement1)
	return a.ATMDepositCompletionAcknowledgement
}

func (a *ATMDepositCompletionAcknowledgementV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
