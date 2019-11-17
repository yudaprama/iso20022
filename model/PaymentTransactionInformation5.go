package model

// Reference and status information concerning the original transactions, included in the original instruction, to which the reversal message applies.
type PaymentTransactionInformation5 struct {

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

	// Original identification of a transaction, as assigned by the first instructing agent and passed on, unchanged, throughout the entire interbank chain.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Amount of money transferred between the instructing agent and the instructed agent in the original transaction.
	OriginalInterbankSettlementAmount *CurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent in the reversed transaction.
	ReversedInterbankSettlementAmount *CurrencyAndAmount `xml:"RvsdIntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: the InterbankSettlementDate is the interbank settlement date of the reversal message, and not of the original instruction.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the reversed transaction.
	ReversedInstructedAmount *CurrencyAndAmount `xml:"RvsdInstdAmt,omitempty"`

	// The factor used for conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Amount of money asked or paid as compensation for the processing of the instruction.
	CompensationAmount *CurrencyAndAmount `xml:"CompstnAmt,omitempty"`

	// Specifies if the creditor and/or debtor will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the reversal message, not to the original instruction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Transaction charges to be paid by the charge bearer for the reversal transaction.
	ChargesInformation []*ChargesInformation1 `xml:"ChrgsInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`

	// Detailed information on the reversal reason.
	ReversalReasonInformation []*ReversalReasonInformation1 `xml:"RvslRsnInf,omitempty"`

	// Set of key elements of the original transaction being referred to.
	OriginalTransactionReference *OriginalTransactionReference1 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation5) SetReversalIdentification(value string) {
	p.ReversalIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation5) SetOriginalPaymentInformationIdentification(value string) {
	p.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation5) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation5) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation5) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation5) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation5) SetReversedInterbankSettlementAmount(value, currency string) {
	p.ReversedInterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation5) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransactionInformation5) SetReversedInstructedAmount(value, currency string) {
	p.ReversedInstructedAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation5) SetExchangeRate(value string) {
	p.ExchangeRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransactionInformation5) SetCompensationAmount(value, currency string) {
	p.CompensationAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation5) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransactionInformation5) AddChargesInformation() *ChargesInformation1 {
	newValue := new(ChargesInformation1)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation5) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.InstructingAgent
}

func (p *PaymentTransactionInformation5) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.InstructedAgent
}

func (p *PaymentTransactionInformation5) AddReversalReasonInformation() *ReversalReasonInformation1 {
	newValue := new(ReversalReasonInformation1)
	p.ReversalReasonInformation = append(p.ReversalReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation5) AddOriginalTransactionReference() *OriginalTransactionReference1 {
	p.OriginalTransactionReference = new(OriginalTransactionReference1)
	return p.OriginalTransactionReference
}
