package caam

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01000101 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.010.001.01 Document"`
	Message *ATMReconciliationAcknowledgementV01 `xml:"ATMRcncltnAck"`
}

func (d *Document01000101) AddMessage() *ATMReconciliationAcknowledgementV01 {
	d.Message = new(ATMReconciliationAcknowledgementV01)
	return d.Message
}

// The ATMReconciliationAcknowledgement message is sent by an acquirer or its agent to an ATM to acknowledge the receipt of an ATMReconciliationAdvice message.
type ATMReconciliationAcknowledgementV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header21 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMReconciliationAcknowledgement *model.ContentInformationType10 `xml:"PrtctdATMRcncltnAck,omitempty"`

	// Information related to the acknowledgement  of an ATM reconciliation from the ATM manager.
	ATMReconciliationAcknowledgement *model.ATMReconciliationAcknowledgement1 `xml:"ATMRcncltnAck,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMReconciliationAcknowledgementV01) AddHeader() *model.Header21 {
	a.Header = new(model.Header21)
	return a.Header
}

func (a *ATMReconciliationAcknowledgementV01) AddProtectedATMReconciliationAcknowledgement() *model.ContentInformationType10 {
	a.ProtectedATMReconciliationAcknowledgement = new(model.ContentInformationType10)
	return a.ProtectedATMReconciliationAcknowledgement
}

func (a *ATMReconciliationAcknowledgementV01) AddATMReconciliationAcknowledgement() *model.ATMReconciliationAcknowledgement1 {
	a.ATMReconciliationAcknowledgement = new(model.ATMReconciliationAcknowledgement1)
	return a.ATMReconciliationAcknowledgement
}

func (a *ATMReconciliationAcknowledgementV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
