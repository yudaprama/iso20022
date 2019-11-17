package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300101 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:catp.003.001.01 Document"`
	Message *ATMWithdrawalCompletionAdviceV01 `xml:"ATMWdrwlCmpltnAdvc"`
}

func (d *Document00300101) AddMessage() *ATMWithdrawalCompletionAdviceV01 {
	d.Message = new(ATMWithdrawalCompletionAdviceV01)
	return d.Message
}

// The ATMWithdrawalCompletionAdvice message is sent by an ATM to an acquirer or its agent to inform of the result of a withdrawal transaction at an ATM.
// If the ATM is configured to only send negative completion, a generic completion message should be used instead of ATMCompletionAdvice.
type ATMWithdrawalCompletionAdviceV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header21 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMWithdrawalCompletionAdvice *model.ContentInformationType10 `xml:"PrtctdATMWdrwlCmpltnAdvc,omitempty"`

	// Information related to the completion of a withdrawal transaction on the ATM.
	ATMWithdrawalCompletionAdvice *model.ATMWithdrawalCompletionAdvice1 `xml:"ATMWdrwlCmpltnAdvc,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMWithdrawalCompletionAdviceV01) AddHeader() *model.Header21 {
	a.Header = new(model.Header21)
	return a.Header
}

func (a *ATMWithdrawalCompletionAdviceV01) AddProtectedATMWithdrawalCompletionAdvice() *model.ContentInformationType10 {
	a.ProtectedATMWithdrawalCompletionAdvice = new(model.ContentInformationType10)
	return a.ProtectedATMWithdrawalCompletionAdvice
}

func (a *ATMWithdrawalCompletionAdviceV01) AddATMWithdrawalCompletionAdvice() *model.ATMWithdrawalCompletionAdvice1 {
	a.ATMWithdrawalCompletionAdvice = new(model.ATMWithdrawalCompletionAdvice1)
	return a.ATMWithdrawalCompletionAdvice
}

func (a *ATMWithdrawalCompletionAdviceV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
