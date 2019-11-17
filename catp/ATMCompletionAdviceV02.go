package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00800102 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.008.001.02 Document"`
	Message *ATMCompletionAdviceV02 `xml:"ATMCmpltnAdvc"`
}

func (d *Document00800102) AddMessage() *ATMCompletionAdviceV02 {
	d.Message = new(ATMCompletionAdviceV02)
	return d.Message
}

// The ATMCompletionAdvice message is sent by an ATM to an acquirer or its agent to inform of the result of a transaction performed on the ATM.
type ATMCompletionAdviceV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header32 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMCompletionAdvice *model.ContentInformationType10 `xml:"PrtctdATMCmpltnAdvc,omitempty"`

	// Information related to the completion of an operation on the ATM.
	ATMCompletionAdvice *model.ATMCompletionAdvice2 `xml:"ATMCmpltnAdvc,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMCompletionAdviceV02) AddHeader() *model.Header32 {
	a.Header = new(model.Header32)
	return a.Header
}

func (a *ATMCompletionAdviceV02) AddProtectedATMCompletionAdvice() *model.ContentInformationType10 {
	a.ProtectedATMCompletionAdvice = new(model.ContentInformationType10)
	return a.ProtectedATMCompletionAdvice
}

func (a *ATMCompletionAdviceV02) AddATMCompletionAdvice() *model.ATMCompletionAdvice2 {
	a.ATMCompletionAdvice = new(model.ATMCompletionAdvice2)
	return a.ATMCompletionAdvice
}

func (a *ATMCompletionAdviceV02) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
