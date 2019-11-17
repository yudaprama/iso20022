package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400101 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:catp.014.001.01 Document"`
	Message *ATMDepositCompletionAdviceV01 `xml:"ATMDpstCmpltnAdvc"`
}

func (d *Document01400101) AddMessage() *ATMDepositCompletionAdviceV01 {
	d.Message = new(ATMDepositCompletionAdviceV01)
	return d.Message
}

// The ATMDepositCompletionAdvice message is sent by an ATM to an acquirer or its agent to inform of the result of a deposit transaction at an ATM.
// If the ATM is configured to only send negative completion, a generic completion message should be used instead of ATMCompletionAdvice.
// This message can be used each time a bundle is put in the ATM safe and/or at the end of a multi bundle deposit.
type ATMDepositCompletionAdviceV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header32 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMDepositCompletionAdvice *model.ContentInformationType10 `xml:"PrtctdATMDpstCmpltnAdvc,omitempty"`

	// Information related to the completion of a deposit transaction on the ATM.
	ATMDepositCompletionAdvice *model.ATMDepositCompletionAdvice1 `xml:"ATMDpstCmpltnAdvc,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMDepositCompletionAdviceV01) AddHeader() *model.Header32 {
	a.Header = new(model.Header32)
	return a.Header
}

func (a *ATMDepositCompletionAdviceV01) AddProtectedATMDepositCompletionAdvice() *model.ContentInformationType10 {
	a.ProtectedATMDepositCompletionAdvice = new(model.ContentInformationType10)
	return a.ProtectedATMDepositCompletionAdvice
}

func (a *ATMDepositCompletionAdviceV01) AddATMDepositCompletionAdvice() *model.ATMDepositCompletionAdvice1 {
	a.ATMDepositCompletionAdvice = new(model.ATMDepositCompletionAdvice1)
	return a.ATMDepositCompletionAdvice
}

func (a *ATMDepositCompletionAdviceV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
