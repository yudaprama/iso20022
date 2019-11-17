package model

// Characteristics that apply to the credit side of the payment transactions included in the direct debit initiation.
type PaymentInstruction15 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod2Code `xml:"PmtMtd"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation24 `xml:"PmtTpInf,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent in the payment chain.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Account used to process charges associated with a transaction.
	//
	// Usage: Charges account should be used when charges have to be booked to an account different from the account identified in debtor's account.
	ChargesAccount *CashAccount24 `xml:"ChrgsAcct,omitempty"`

	// Agent that services a charges account.
	//
	// Usage: Charges account agent should only be used when the charges account agent is different from the creditor agent.
	ChargesAccountAgent *BranchAndFinancialInstitutionIdentification5 `xml:"ChrgsAcctAgt,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Provides information on the individual transaction(s) included in the message.
	DirectDebitTransactionInformation []*DirectDebitTransactionInformation18 `xml:"DrctDbtTxInf"`
}

func (p *PaymentInstruction15) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction15) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod2Code)(&value)
}

func (p *PaymentInstruction15) SetBatchBooking(value string) {
	p.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (p *PaymentInstruction15) SetNumberOfTransactions(value string) {
	p.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (p *PaymentInstruction15) SetControlSum(value string) {
	p.ControlSum = (*DecimalNumber)(&value)
}

func (p *PaymentInstruction15) AddPaymentTypeInformation() *PaymentTypeInformation24 {
	p.PaymentTypeInformation = new(PaymentTypeInformation24)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction15) SetRequestedCollectionDate(value string) {
	p.RequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentInstruction15) AddCreditor() *PartyIdentification43 {
	p.Creditor = new(PartyIdentification43)
	return p.Creditor
}

func (p *PaymentInstruction15) AddCreditorAccount() *CashAccount24 {
	p.CreditorAccount = new(CashAccount24)
	return p.CreditorAccount
}

func (p *PaymentInstruction15) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.CreditorAgent
}

func (p *PaymentInstruction15) AddCreditorAgentAccount() *CashAccount24 {
	p.CreditorAgentAccount = new(CashAccount24)
	return p.CreditorAgentAccount
}

func (p *PaymentInstruction15) AddUltimateCreditor() *PartyIdentification43 {
	p.UltimateCreditor = new(PartyIdentification43)
	return p.UltimateCreditor
}

func (p *PaymentInstruction15) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction15) AddChargesAccount() *CashAccount24 {
	p.ChargesAccount = new(CashAccount24)
	return p.ChargesAccount
}

func (p *PaymentInstruction15) AddChargesAccountAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.ChargesAccountAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.ChargesAccountAgent
}

func (p *PaymentInstruction15) AddCreditorSchemeIdentification() *PartyIdentification43 {
	p.CreditorSchemeIdentification = new(PartyIdentification43)
	return p.CreditorSchemeIdentification
}

func (p *PaymentInstruction15) AddDirectDebitTransactionInformation() *DirectDebitTransactionInformation18 {
	newValue := new(DirectDebitTransactionInformation18)
	p.DirectDebitTransactionInformation = append(p.DirectDebitTransactionInformation, newValue)
	return newValue
}
