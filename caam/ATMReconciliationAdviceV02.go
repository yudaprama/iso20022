package caam

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00900102 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.009.001.02 Document"`
	Message *ATMReconciliationAdviceV02 `xml:"ATMRcncltnAdvc"`
}

func (d *Document00900102) AddMessage() *ATMReconciliationAdviceV02 {
	d.Message = new(ATMReconciliationAdviceV02)
	return d.Message
}

// The ATMReconciliationAdvice message is sent by an ATM to an acquirer or its agent to send all the counters of the ATM. It can be sent by an operator function or as a response of a command sent by an agent or a server.
type ATMReconciliationAdviceV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header32 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMReconciliationAdvice *model.ContentInformationType10 `xml:"PrtctdATMRcncltnAdvc,omitempty"`

	// Information related to the reconciliation of an ATM.
	ATMReconciliationAdvice *model.ATMReconciliationAdvice2 `xml:"ATMRcncltnAdvc,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMReconciliationAdviceV02) AddHeader() *model.Header32 {
	a.Header = new(model.Header32)
	return a.Header
}

func (a *ATMReconciliationAdviceV02) AddProtectedATMReconciliationAdvice() *model.ContentInformationType10 {
	a.ProtectedATMReconciliationAdvice = new(model.ContentInformationType10)
	return a.ProtectedATMReconciliationAdvice
}

func (a *ATMReconciliationAdviceV02) AddATMReconciliationAdvice() *model.ATMReconciliationAdvice2 {
	a.ATMReconciliationAdvice = new(model.ATMReconciliationAdvice2)
	return a.ATMReconciliationAdvice
}

func (a *ATMReconciliationAdviceV02) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
