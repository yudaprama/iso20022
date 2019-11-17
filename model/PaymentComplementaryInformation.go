package model

// Additional information from the underlying payment instruction which cannot be transferred in a regular statement message.
type PaymentComplementaryInformation struct {

	// Remittance information.
	RemittanceChoice *RemittanceInformation3Choice `xml:"RmtChc,omitempty"`

	// Debtor or Ordering customer of the payment instruction.
	Debtor *PartyIdentification1 `xml:"Dbtr,omitempty"`

	// Debtor account or Ordering customer account.
	DebtorAccount *CashAccount3 `xml:"DbtrAcct,omitempty"`

	// First Agent or Field 52 in Fin messages.
	FirstAgent *BranchAndFinancialInstitutionIdentification `xml:"FrstAgt,omitempty"`

	// Instructed amount of the payment instruction (Field 33B)
	Amount *AmountType1Choice `xml:"Amt,omitempty"`

	// Indicates the account used to cover a payment.
	NostroVostroAccount *CashAccount3 `xml:"NstrVstrAcct,omitempty"`

	// Intermediary bank.
	Intermediary *Intermediary1 `xml:"Intrmy,omitempty"`

	// Correspondent of the first agent (Field 53 in MT202).
	FirstSettlementAgent *BranchAndFinancialInstitutionIdentification `xml:"FrstSttlmAgt,omitempty"`

	// Correspondent of the Final agent (Field 54 of Mt 202)
	LastSettlementAgent *BranchAndFinancialInstitutionIdentification `xml:"LastSttlmAgt,omitempty"`

	// Equivalent to Field 55 in MT202.
	IntermediarySettlementAgent *BranchAndFinancialInstitutionIdentification `xml:"IntrmySttlmAgt,omitempty"`

	// Creditor or Beneficiary customer of the payment instruction.
	Creditor *PartyIdentification1 `xml:"Cdtr,omitempty"`

	// Creditor account or Beneficiary customer account.
	CreditorAccount *CashAccount3 `xml:"CdtrAcct,omitempty"`

	// Unformatted information from the sender to the receiver.
	SenderToReceiverInformation []*Max35Text `xml:"SndrToRcvrInf,omitempty"`
}

func (p *PaymentComplementaryInformation) AddRemittanceChoice() *RemittanceInformation3Choice {
	p.RemittanceChoice = new(RemittanceInformation3Choice)
	return p.RemittanceChoice
}

func (p *PaymentComplementaryInformation) AddDebtor() *PartyIdentification1 {
	p.Debtor = new(PartyIdentification1)
	return p.Debtor
}

func (p *PaymentComplementaryInformation) AddDebtorAccount() *CashAccount3 {
	p.DebtorAccount = new(CashAccount3)
	return p.DebtorAccount
}

func (p *PaymentComplementaryInformation) AddFirstAgent() *BranchAndFinancialInstitutionIdentification {
	p.FirstAgent = new(BranchAndFinancialInstitutionIdentification)
	return p.FirstAgent
}

func (p *PaymentComplementaryInformation) AddAmount() *AmountType1Choice {
	p.Amount = new(AmountType1Choice)
	return p.Amount
}

func (p *PaymentComplementaryInformation) AddNostroVostroAccount() *CashAccount3 {
	p.NostroVostroAccount = new(CashAccount3)
	return p.NostroVostroAccount
}

func (p *PaymentComplementaryInformation) AddIntermediary() *Intermediary1 {
	p.Intermediary = new(Intermediary1)
	return p.Intermediary
}

func (p *PaymentComplementaryInformation) AddFirstSettlementAgent() *BranchAndFinancialInstitutionIdentification {
	p.FirstSettlementAgent = new(BranchAndFinancialInstitutionIdentification)
	return p.FirstSettlementAgent
}

func (p *PaymentComplementaryInformation) AddLastSettlementAgent() *BranchAndFinancialInstitutionIdentification {
	p.LastSettlementAgent = new(BranchAndFinancialInstitutionIdentification)
	return p.LastSettlementAgent
}

func (p *PaymentComplementaryInformation) AddIntermediarySettlementAgent() *BranchAndFinancialInstitutionIdentification {
	p.IntermediarySettlementAgent = new(BranchAndFinancialInstitutionIdentification)
	return p.IntermediarySettlementAgent
}

func (p *PaymentComplementaryInformation) AddCreditor() *PartyIdentification1 {
	p.Creditor = new(PartyIdentification1)
	return p.Creditor
}

func (p *PaymentComplementaryInformation) AddCreditorAccount() *CashAccount3 {
	p.CreditorAccount = new(CashAccount3)
	return p.CreditorAccount
}

func (p *PaymentComplementaryInformation) AddSenderToReceiverInformation(value string) {
	p.SenderToReceiverInformation = append(p.SenderToReceiverInformation, (*Max35Text)(&value))
}
