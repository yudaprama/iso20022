package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02800107 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.07 Document"`
	Message *AdditionalPaymentInformationV07 `xml:"AddtlPmtInf"`
}

func (d *Document02800107) AddMessage() *AdditionalPaymentInformationV07 {
	d.Message = new(AdditionalPaymentInformationV07)
	return d.Message
}

// Scope
// The AdditionalPaymentInformation message is sent by an account servicing institution to an account owner.
// This message is used to provide additional or corrected information on a payment instruction or statement entry, in order to allow reconciliation.
// Usage
// The AdditionalPaymentInformation message provides elements which are usually not reported in a statement or advice (for example full remittance information or identification of parties other than the account servicing institution and the account owner). It complements information about a payment instruction that has already been received, in the form of one or several entries of the original payment instruction.
// The AdditionalPaymentInformation message covers one and only one original payment instruction. If several payment instructions need further details, multiple AdditionalPaymentInformation messages must be used, one for each of the payment instructions.
// The AdditionalPaymentInformation message may be used as a result of two investigation processes and in a way outlined below.
// - A ClaimNonReceipt workflow raised by the creditor or recipient of the payment: This means that the payment instruction has reached the creditor or beneficiary. The account owner needs further details or correct information for its reconciliation processes. The AdditionalPaymentInformation can be used to provide the missing information.
// - A RequestToModifyPayment workflow raised by the debtor or one of the intermediate agents upstream: When the payment instruction has reached its intended recipient and the modification does not affect the accounting at the account servicing institution, this AdditionalPaymentInformation message allows the account owner to receive further particulars or correct information about a payment instruction or an entry passed to its account.
// The AdditionalPayment Information message cannot be used to trigger a request for modification of a payment instruction activity. A RequestToModifyPayment message must be used. In other words, if a debtor or one of intermediate agent (excluding the account servicing institution of the creditor) realises the some information was missing in the original payment instruction, he should not use an AdditionalPaymentInformation but instead a RequestToModifyPayment message.
// It is assumed that when an account servicing institution sends out an AdditionalPaymentInformation message, the institution is fairly confident that this will resolve the case. Therefore it does not need to wait for a Resolution Of Investigation message. Neither does the account owner, or whoever receives the additional information, need to send back a ResolutionOfInvestigation message. Positive resolution in this case is implicit. Both parties are expected to close the case. In the event that the problem does not go away, a party can re-open the case.
type AdditionalPaymentInformationV07 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *model.CaseAssignment3 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *model.Case3 `xml:"Case"`

	// Identifies the underlying payment instruction.
	Underlying *model.UnderlyingTransaction3Choice `xml:"Undrlyg"`

	// Additional information to the underlying payment instruction.
	Information *model.PaymentComplementaryInformation6 `xml:"Inf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AdditionalPaymentInformationV07) AddAssignment() *model.CaseAssignment3 {
	a.Assignment = new(model.CaseAssignment3)
	return a.Assignment
}

func (a *AdditionalPaymentInformationV07) AddCase() *model.Case3 {
	a.Case = new(model.Case3)
	return a.Case
}

func (a *AdditionalPaymentInformationV07) AddUnderlying() *model.UnderlyingTransaction3Choice {
	a.Underlying = new(model.UnderlyingTransaction3Choice)
	return a.Underlying
}

func (a *AdditionalPaymentInformationV07) AddInformation() *model.PaymentComplementaryInformation6 {
	a.Information = new(model.PaymentComplementaryInformation6)
	return a.Information
}

func (a *AdditionalPaymentInformationV07) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
