package model

// Reference and status information concerning the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransactionInformation4 struct {

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the reversed transaction.
	ReversalIdentification *Max35Text `xml:"RvslId,omitempty"`

	// Unique and unambiguous identifier of the original payment information block as assigned by the original sending party.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId,omitempty"`

	// Original unique instruction identification as assigned by an instructing party for an instructed party to unambiguously identify the original instruction.
	//
	// Usage: the original instruction identification is the original point to point reference used between the instructing party and the instructed party to refer to the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Original unique identification assigned by the initiating party to unambiguously identify the original transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency as ordered by the initiating party in the original transaction.
	OriginalInstructedAmount *CurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the reversed transaction.
	ReversedInstructedAmount *CurrencyAndAmount `xml:"RvsdInstdAmt,omitempty"`

	// Specifies if the creditor and/or debtor will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the reversal message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Detailed information on the reversal reason.
	ReversalReasonInformation []*ReversalReasonInformation1 `xml:"RvslRsnInf,omitempty"`

	// Set of key elements of the original transaction being referred to.
	OriginalTransactionReference *OriginalTransactionReference1 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation4) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation4) SetOriginalPaymentInformationIdentification(value string) {
	p.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation4) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation4) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation4) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation4) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation4) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransactionInformation4) AddReversalReasonInformation() *ReversalReasonInformation1 {
	newValue := new(ReversalReasonInformation1)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation4) AddOriginalTransactionReference() *OriginalTransactionReference1 {
	p.OriginalTransactionReference = new(OriginalTransactionReference1)
	return p.OriginalTransactionReference
}
