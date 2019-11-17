package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300102 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.003.001.02 Document"`
	Message *ATMWithdrawalCompletionAdviceV02 `xml:"ATMWdrwlCmpltnAdvc"`
}

func (d *Document00300102) AddMessage() *ATMWithdrawalCompletionAdviceV02 {
	d.Message = new(ATMWithdrawalCompletionAdviceV02)
	return d.Message
}

// The ATMWithdrawalCompletionAdvice message is sent by an ATM to an acquirer or its agent to inform of the result of a withdrawal transaction at an ATM.
// If the ATM is configured to only send negative completion, a generic completion message should be used instead of ATMCompletionAdvice.
type ATMWithdrawalCompletionAdviceV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header32 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMWithdrawalCompletionAdvice *model.ContentInformationType10 `xml:"PrtctdATMWdrwlCmpltnAdvc,omitempty"`

	// Information related to the completion of a withdrawal transaction on the ATM.
	ATMWithdrawalCompletionAdvice *model.ATMWithdrawalCompletionAdvice2 `xml:"ATMWdrwlCmpltnAdvc,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMWithdrawalCompletionAdviceV02) AddHeader() *model.Header32 {
	a.Header = new(model.Header32)
	return a.Header
}

func (a *ATMWithdrawalCompletionAdviceV02) AddProtectedATMWithdrawalCompletionAdvice() *model.ContentInformationType10 {
	a.ProtectedATMWithdrawalCompletionAdvice = new(model.ContentInformationType10)
	return a.ProtectedATMWithdrawalCompletionAdvice
}

func (a *ATMWithdrawalCompletionAdviceV02) AddATMWithdrawalCompletionAdvice() *model.ATMWithdrawalCompletionAdvice2 {
	a.ATMWithdrawalCompletionAdvice = new(model.ATMWithdrawalCompletionAdvice2)
	return a.ATMWithdrawalCompletionAdvice
}

func (a *ATMWithdrawalCompletionAdviceV02) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
