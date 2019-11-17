package model

// Provides reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction38 struct {

	// Unique and unambiguous identifier of a cancellation request, as assigned by the assigner.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationIdentification *Max35Text `xml:"CxlId,omitempty"`

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case3 `xml:"Case,omitempty"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OriginalInterbankSettlementDate *ISODate `xml:"OrgnlIntrBkSttlmDt,omitempty"`

	// Party who assigns the case.
	// Usage: This is also the agent that instructs the next party in the chain to carry out the (set of) cancellation request(s).
	Assigner *BranchAndFinancialInstitutionIdentification5 `xml:"Assgnr,omitempty"`

	// Party to which the case is assigned.
	// Usage: This is also the agent that is instructed by the previous party in the chain to carry out the (set of) cancellation request(s).
	Assignee *BranchAndFinancialInstitutionIdentification5 `xml:"Assgne,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason2 `xml:"CxlRsnInf,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransaction38) SetCancellationIdentification(value string) {
	p.CancellationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction38) AddCase() *Case3 {
	p.Case = new(Case3)
	return p.Case
}

func (p *PaymentTransaction38) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction38) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction38) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction38) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction38) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction38) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction38) SetOriginalInterbankSettlementDate(value string) {
	p.OriginalInterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction38) AddAssigner() *BranchAndFinancialInstitutionIdentification5 {
	p.Assigner = new(BranchAndFinancialInstitutionIdentification5)
	return p.Assigner
}

func (p *PaymentTransaction38) AddAssignee() *BranchAndFinancialInstitutionIdentification5 {
	p.Assignee = new(BranchAndFinancialInstitutionIdentification5)
	return p.Assignee
}

func (p *PaymentTransaction38) AddCancellationReasonInformation() *PaymentCancellationReason2 {
	newValue := new(PaymentCancellationReason2)
	p.CancellationReasonInformation = append(p.CancellationReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction38) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}
