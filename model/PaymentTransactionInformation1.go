package model

// Information concerning the original transactions, to which the status report message refers.
type PaymentTransactionInformation1 struct {

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the reported status.
	StatusIdentification *Max35Text `xml:"StsId,omitempty"`

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

	// Specifies the status of a transaction, in a coded form.
	TransactionStatus *TransactionIndividualStatus1Code `xml:"TxSts,omitempty"`

	// Detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation1 `xml:"StsRsnInf,omitempty"`

	// Information on charges related to the processing of the rejection of the instruction.
	//
	// Usage: ChargesInformation is past on for information purposes only. Settlement of the charges will be done separately.
	ChargesInformation []*ChargesInformation1 `xml:"ChrgsInf,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent (debtor's agent in case of a credit transfer, creditor's agent in case of a direct debit). This means - amongst others - that the account servicing agent has received the payment order and has applied checks as eg, authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`

	// Set of key elements of the original transaction being referred to.
	OriginalTransactionReference *OriginalTransactionReference1 `xml:"OrgnlTxRef,omitempty"`
}

func (p *PaymentTransactionInformation1) SetStatusIdentification(value string) {
	p.StatusIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation1) SetOriginalPaymentInformationIdentification(value string) {
	p.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation1) SetOriginalInstructionIdentification(value string) {
	p.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation1) SetOriginalEndToEndIdentification(value string) {
	p.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation1) SetOriginalTransactionIdentification(value string) {
	p.OriginalTransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentTransactionInformation1) SetTransactionStatus(value string) {
	p.TransactionStatus = (*TransactionIndividualStatus1Code)(&value)
}

func (p *PaymentTransactionInformation1) AddStatusReasonInformation() *StatusReasonInformation1 {
	newValue := new(StatusReasonInformation1)
	p.StatusReasonInformation = append(p.StatusReasonInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation1) AddChargesInformation() *ChargesInformation1 {
	newValue := new(ChargesInformation1)
	p.ChargesInformation = append(p.ChargesInformation, newValue)
	return newValue
}

func (p *PaymentTransactionInformation1) SetAcceptanceDateTime(value string) {
	p.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (p *PaymentTransactionInformation1) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.InstructingAgent
}

func (p *PaymentTransactionInformation1) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.InstructedAgent
}

func (p *PaymentTransactionInformation1) AddOriginalTransactionReference() *OriginalTransactionReference1 {
	p.OriginalTransactionReference = new(OriginalTransactionReference1)
	return p.OriginalTransactionReference
}
