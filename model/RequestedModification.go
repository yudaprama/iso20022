package model

// Contains the requested modifications.
type RequestedModification struct {

	// Reference relating to a linked payment instruction or agreement which is meaningful to both parties (eg, the content of field 21 in a cover instruction).
	RelatedReference *Max35Text `xml:"RltdRef,omitempty"`

	// SWIFT defined service level applies to the payment instruction.
	BankOperationCode *SWIFTServiceLevel2Code `xml:"BkOprCd,omitempty"`

	// Further information related to the processing of the payment instruction. The instruction can relate to a level of service between the bank and the customer, or give instructions to and for specific parties in the payment chain.
	InstructionCode *Instruction1Code `xml:"InstrCd,omitempty"`

	// Date and time the debtor requests the clearing agent to process the payment instruction.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettledAmount *CurrencyAndAmount `xml:"IntrBkSttldAmt,omitempty"`

	// Debtor or Ordering customer of the payment instruction.
	Debtor *PartyIdentification1 `xml:"Dbtr,omitempty"`

	// Account to or from which a cash entry is made.
	DebtorAccount *CashAccount3 `xml:"DbtrAcct,omitempty"`

	// Party that executes a cash transfer received via a clearing agent or on request of an agreement party.
	IntermediarySettlementAgent *BranchAndFinancialInstitutionIdentification `xml:"IntrmySttlmAgt,omitempty"`

	// Party that executes a cash transfer received via a clearing agent or on request of an agreement party.
	LastSettlementAgent *BranchAndFinancialInstitutionIdentification `xml:"LastSttlmAgt,omitempty"`

	// Specification of a pre-agreed offering between clearing agents, or the channel through which the payment instruction is to be processed. This payment scheme can point to a specific rulebook governing the rules of clearing and settlement between two parties.
	PaymentScheme *PaymentSchemeChoice `xml:"PmtSchme,omitempty"`

	// Account to or from which a cash entry is made.
	BeneficiaryInstitutionAccount *CashAccount3 `xml:"BnfcryInstnAcct,omitempty"`

	// Entity involved in an activity.
	Creditor *PartyIdentification1 `xml:"Cdtr,omitempty"`

	// Account to or from which a cash entry is made.
	CreditorAccount *CashAccount3 `xml:"CdtrAcct,omitempty"`

	// Structured information that enables the matching, ie, reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation3Choice `xml:"RmtInf,omitempty"`

	// Underlying reason for the payment transaction.
	Purpose *PurposeChoice `xml:"Purp,omitempty"`

	// Further information related to the processing of the payment instruction. The instruction can relate to a level of service between the bank and the customer, or give instructions to and for specific parties in the payment chain.
	InstructionForFinalAgent *InstructionForFinalAgent `xml:"InstrForFnlAgt,omitempty"`

	// Party(ies) liable for a charge associated with the transfer of cash.
	DetailsOfCharges *ChargeBearer1Code `xml:"DtlsOfChrgs,omitempty"`

	// Unformatted information from the sender to the receiver.
	SenderToReceiverInformation []*Max35Text `xml:"SndrToRcvrInf,omitempty"`
}

func (r *RequestedModification) SetRelatedReference(value string) {
	r.RelatedReference = (*Max35Text)(&value)
}

func (r *RequestedModification) SetBankOperationCode(value string) {
	r.BankOperationCode = (*SWIFTServiceLevel2Code)(&value)
}

func (r *RequestedModification) SetInstructionCode(value string) {
	r.InstructionCode = (*Instruction1Code)(&value)
}

func (r *RequestedModification) SetRequestedExecutionDate(value string) {
	r.RequestedExecutionDate = (*ISODate)(&value)
}

func (r *RequestedModification) SetValueDate(value string) {
	r.ValueDate = (*ISODate)(&value)
}

func (r *RequestedModification) SetInterbankSettledAmount(value, currency string) {
	r.InterbankSettledAmount = NewCurrencyAndAmount(value, currency)
}

func (r *RequestedModification) AddDebtor() *PartyIdentification1 {
	r.Debtor = new(PartyIdentification1)
	return r.Debtor
}

func (r *RequestedModification) AddDebtorAccount() *CashAccount3 {
	r.DebtorAccount = new(CashAccount3)
	return r.DebtorAccount
}

func (r *RequestedModification) AddIntermediarySettlementAgent() *BranchAndFinancialInstitutionIdentification {
	r.IntermediarySettlementAgent = new(BranchAndFinancialInstitutionIdentification)
	return r.IntermediarySettlementAgent
}

func (r *RequestedModification) AddLastSettlementAgent() *BranchAndFinancialInstitutionIdentification {
	r.LastSettlementAgent = new(BranchAndFinancialInstitutionIdentification)
	return r.LastSettlementAgent
}

func (r *RequestedModification) AddPaymentScheme() *PaymentSchemeChoice {
	r.PaymentScheme = new(PaymentSchemeChoice)
	return r.PaymentScheme
}

func (r *RequestedModification) AddBeneficiaryInstitutionAccount() *CashAccount3 {
	r.BeneficiaryInstitutionAccount = new(CashAccount3)
	return r.BeneficiaryInstitutionAccount
}

func (r *RequestedModification) AddCreditor() *PartyIdentification1 {
	r.Creditor = new(PartyIdentification1)
	return r.Creditor
}

func (r *RequestedModification) AddCreditorAccount() *CashAccount3 {
	r.CreditorAccount = new(CashAccount3)
	return r.CreditorAccount
}

func (r *RequestedModification) AddRemittanceInformation() *RemittanceInformation3Choice {
	r.RemittanceInformation = new(RemittanceInformation3Choice)
	return r.RemittanceInformation
}

func (r *RequestedModification) AddPurpose() *PurposeChoice {
	r.Purpose = new(PurposeChoice)
	return r.Purpose
}

func (r *RequestedModification) AddInstructionForFinalAgent() *InstructionForFinalAgent {
	r.InstructionForFinalAgent = new(InstructionForFinalAgent)
	return r.InstructionForFinalAgent
}

func (r *RequestedModification) SetDetailsOfCharges(value string) {
	r.DetailsOfCharges = (*ChargeBearer1Code)(&value)
}

func (r *RequestedModification) AddSenderToReceiverInformation(value string) {
	r.SenderToReceiverInformation = append(r.SenderToReceiverInformation, (*Max35Text)(&value))
}
