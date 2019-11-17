package model

// Reference and status information concerning the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransactionInformation3 struct {

	// Unique and unambiguous identifier of a cancellation request, assigned by an instructing party.
	//
	// Usage: the cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request transaction.
	CancellationIdentification *Max35Text `xml:"CxlId,omitempty"`

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

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency as ordered by the initiating party in the original transaction.
	OriginalInstructedAmount *CurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`

	// Detailed information on the cancellation reason.
	CancellationReasonInformation []*CancellationReasonInformation1 `xml:"CxlRsnInf,omitempty"`

	// Set of key elements of the original transaction being referred to.
	OriginalTransactionReference *OriginalTransactionReference1 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation3) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation3) SetOriginalPaymentInformationIdentification(value string) {
	p.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation3) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation3) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation3) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation3) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation3) SetOriginalInstructedAmount(value, currency string) {
	p.OriginalInstructedAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation3) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.InstructingAgent
}

func (p *PaymentTransactionInformation3) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.InstructedAgent
}

func (p *PaymentTransactionInformation3) AddCancellationReasonInformation() *CancellationReasonInformation1 {
	newValue := new(CancellationReasonInformation1)
	p.CancellationReasonInformation = append(p.CancellationReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation3) AddOriginalTransactionReference() *OriginalTransactionReference1 {
	p.OriginalTransactionReference = new(OriginalTransactionReference1)
	return p.OriginalTransactionReference
}
