package model

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransactionInformation28 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed transaction.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalIdentification *Max35Text `xml:"RvslId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the reversed transaction.
	ReversedInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RvsdInstdAmt,omitempty"`

	// Specifies if the creditor and/or debtor will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the reversal message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Set of elements used to provide detailed information on the reversal reason.
	ReversalReasonInformation []*ReversalReasonInformation6 `xml:"RvslRsnInf,omitempty"`

	// Set of key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference13 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation28) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation28) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation28) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation28) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation28) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation28) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransactionInformation28) AddReversalReasonInformation() *ReversalReasonInformation6 {
	newValue := new(ReversalReasonInformation6)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation28) AddOriginalTransactionReference() *OriginalTransactionReference13 {
	p.OriginalTransactionReference = new(OriginalTransactionReference13)
	return p.OriginalTransactionReference
}
