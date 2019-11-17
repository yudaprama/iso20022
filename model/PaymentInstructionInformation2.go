package model

// Set of characteristics that apply to the credit side of the payment transactions included in the direct debit transaction initiation.
type PaymentInstructionInformation2 struct {

	// Reference assigned by a sending party to unambiguously identify the payment information block within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId,omitempty"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod2Code `xml:"PmtMtd"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation2 `xml:"PmtTpInf,omitempty"`

	// Date at which the creditor requests the amount of money to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification8 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount7 `xml:"CdtrAcct"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount7 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification8 `xml:"UltmtCdtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Account used to process charges associated with a transaction.
	//
	// Usage : charges account should be used when charges have to be booked to an account different from the account identified in debtor's account.
	ChargesAccount *CashAccount7 `xml:"ChrgsAcct,omitempty"`

	// Agent that services a charges account.
	//
	// Usage : charges account agent should only be used when the charges account agent is different from the creditor agent.
	ChargesAccountAgent *BranchAndFinancialInstitutionIdentification3 `xml:"ChrgsAcctAgt,omitempty"`

	// Set of elements providing information specific to the individual transaction(s) included in the message.
	DirectDebitTransactionInformation []*DirectDebitTransactionInformation1 `xml:"DrctDbtTxInf"`
}

func (p *PaymentInstructionInformation2) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstructionInformation2) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod2Code)(&value)
}

func (p *PaymentInstructionInformation2) AddPaymentTypeInformation() *PaymentTypeInformation2 {
	p.PaymentTypeInformation = new(PaymentTypeInformation2)
	return p.PaymentTypeInformation
}

func (p *PaymentInstructionInformation2) SetRequestedCollectionDate(value string) {
	p.RequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentInstructionInformation2) AddCreditor() *PartyIdentification8 {
	p.Creditor = new(PartyIdentification8)
	return p.Creditor
}

func (p *PaymentInstructionInformation2) AddCreditorAccount() *CashAccount7 {
	p.CreditorAccount = new(CashAccount7)
	return p.CreditorAccount
}

func (p *PaymentInstructionInformation2) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.CreditorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.CreditorAgent
}

func (p *PaymentInstructionInformation2) AddCreditorAgentAccount() *CashAccount7 {
	p.CreditorAgentAccount = new(CashAccount7)
	return p.CreditorAgentAccount
}

func (p *PaymentInstructionInformation2) AddUltimateCreditor() *PartyIdentification8 {
	p.UltimateCreditor = new(PartyIdentification8)
	return p.UltimateCreditor
}

func (p *PaymentInstructionInformation2) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstructionInformation2) AddChargesAccount() *CashAccount7 {
	p.ChargesAccount = new(CashAccount7)
	return p.ChargesAccount
}

func (p *PaymentInstructionInformation2) AddChargesAccountAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.ChargesAccountAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.ChargesAccountAgent
}

func (p *PaymentInstructionInformation2) AddDirectDebitTransactionInformation() *DirectDebitTransactionInformation1 {
	newValue := new(DirectDebitTransactionInformation1)
	p.DirectDebitTransactionInformation = append(p.DirectDebitTransactionInformation, newValue)
	return newValue
}
