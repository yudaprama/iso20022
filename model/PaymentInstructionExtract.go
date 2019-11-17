package model

// Details of a payment instruction. The information contained in this component is sufficient to retrieve a payment instruction.
type PaymentInstructionExtract struct {

	// Identification of the payment instruction (eg, field 20 of an MT 103) when meaningful to the case assigner.
	AssignerInstructionIdentification *Max35Text `xml:"AssgnrInstrId,omitempty"`

	// Identification of the payment instruction (eg, field 20 of an MT 103) when meaningful to the case assignee.
	AssigneeInstructionIdentification *Max35Text `xml:"AssgneInstrId,omitempty"`

	// Amount of the payment. Depending on the context it can be either the amount settled (UnableToApply) or the instructed amount (RequestToCancel, RequestToModify, ClaimNonReceipt).
	CurrencyAmount *CurrencyAndAmount `xml:"CcyAmt,omitempty"`

	// Value date of the payment.
	ValueDate *ISODateTime `xml:"ValDt,omitempty"`
}

func (p *PaymentInstructionExtract) SetAssignerInstructionIdentification(value string) {
	p.AssignerInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstructionExtract) SetAssigneeInstructionIdentification(value string) {
	p.AssigneeInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstructionExtract) SetCurrencyAmount(value, currency string) {
	p.CurrencyAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentInstructionExtract) SetValueDate(value string) {
	p.ValueDate = (*ISODateTime)(&value)
}
