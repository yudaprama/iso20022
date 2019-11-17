package model

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransactionInformation33 struct {

	// Unique and unambiguous identifier of a cancellation request status, as assigned by the assigner.
	//
	// Usage: The cancellation status identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CancellationStatusIdentification *Max35Text `xml:"CxlStsId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case2 `xml:"RslvdCase,omitempty"`

	// Set of elements used to provide information on the original messsage.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OriginalClearingSystemReference *Max35Text `xml:"OrgnlClrSysRef,omitempty"`

	// Specifies the status of the transaction cancellation request.
	TransactionCancellationStatus *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty"`

	// Set of elements used to provide detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReasonInformation1 `xml:"CxlStsRsnInf,omitempty"`

	// Reference of a return or a reversal transaction that is initiated to fix the case under investigation as part of the resolution.
	ResolutionRelatedInformation *ResolutionInformation1 `xml:"RsltnRltdInf,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the instructing agent and the instructed agent.
	OriginalInterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty"`

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OriginalInterbankSettlementDate *ISODate `xml:"OrgnlIntrBkSttlmDt,omitempty"`

	// Party who assigns the case.
	// Usage: This is also the agent that instructs the next party in the chain to carry out the (set of) cancellation request(s).
	Assigner *Party7Choice `xml:"Assgnr,omitempty"`

	// Party to which the case is assigned.
	// Usage: This is also the agent that is instructed by the previous party in the chain to carry out the (set of) cancellation request(s).
	Assignee *Party7Choice `xml:"Assgne,omitempty"`

	// Set of key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference13 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation33) SetCancellationStatusIdentification(value string) {
	p.CancellationStatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation33) AddResolvedCase() *Case2 {
	p.ResolvedCase = new(Case2)
	return p.ResolvedCase
}

func (p *PaymentTransactionInformation33) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransactionInformation33) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation33) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation33) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation33) SetOriginalClearingSystemReference(value string) {
	p.OriginalClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation33) SetTransactionCancellationStatus(value string) {
	p.TransactionCancellationStatus = (*CancellationIndividualStatus1Code)(&value)
}

func (p *PaymentTransactionInformation33) AddCancellationStatusReasonInformation() *CancellationStatusReasonInformation1 {
	newValue := new(CancellationStatusReasonInformation1)
	p.CancellationStatusReasonInformation = append(p.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation33) AddResolutionRelatedInformation() *ResolutionInformation1 {
	p.ResolutionRelatedInformation = new(ResolutionInformation1)
	return p.ResolutionRelatedInformation
}

func (p *PaymentTransactionInformation33) SetOriginalInterbankSettlementAmount(value, currency string) {
	p.OriginalInterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransactionInformation33) SetOriginalInterbankSettlementDate(value string) {
	p.OriginalInterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransactionInformation33) AddAssigner() *Party7Choice {
	p.Assigner = new(Party7Choice)
	return p.Assigner
}

func (p *PaymentTransactionInformation33) AddAssignee() *Party7Choice {
	p.Assignee = new(Party7Choice)
	return p.Assignee
}

func (p *PaymentTransactionInformation33) AddOriginalTransactionReference() *OriginalTransactionReference13 {
	p.OriginalTransactionReference = new(OriginalTransactionReference13)
	return p.OriginalTransactionReference
}
