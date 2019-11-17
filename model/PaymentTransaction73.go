package model

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction73 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the status request.
	// Usage: The instructing party is the party sending the request message and not the party that sent the original instruction that is being reported on.
	StatusRequestIdentification *Max35Text `xml:"StsReqId,omitempty"`

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OriginalTransactionIdentification *Max35Text `xml:"OrgnlTxId,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	//
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	//
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

	// Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionReference *OriginalTransactionReference24 `xml:"OrgnlTxRef,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (p *PaymentTransaction73) SetStatusRequestIdentification(value string) {
	p.StatusRequestIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction73) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	p.OriginalGroupInformation = new(OriginalGroupInformation3)
	return p.OriginalGroupInformation
}

func (p *PaymentTransaction73) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction73) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction73) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransaction73) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransaction73) SetClearingSystemReference(value string) {
	p.ClearingSystemReference = (*Max35Text)(&value)
}

func (p *PaymentTransaction73) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructingAgent
}

func (p *PaymentTransaction73) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstructedAgent
}

func (p *PaymentTransaction73) AddOriginalTransactionReference() *OriginalTransactionReference24 {
	p.OriginalTransactionReference = new(OriginalTransactionReference24)
	return p.OriginalTransactionReference
}

func (p *PaymentTransaction73) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SupplementaryData = append(p.SupplementaryData, newValue)
	return newValue
}
