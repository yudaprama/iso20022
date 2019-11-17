package model

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction43 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *TransactionIndividualStatus3Code `xml:"TxSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Provides information on the charges related to the processing of the rejection of the instruction.
	//
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the status message and not the party that sent the original instruction that is being reported on.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the status message and not the party that received the original instruction that is being reported on.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference16 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction43) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction43) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction43) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction43) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction43) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (p *PaymentTransaction43) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransaction43) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransaction43) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction43) SetAccountServicerReference(value string) {
	p.AccountServicerReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction43) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction43) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction43) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction43) AddOriginalTransactionReference() *OriginalTransactionReference16 {
	p.OriginalTransactionReference = new(OriginalTransactionReference16)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction43) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}
