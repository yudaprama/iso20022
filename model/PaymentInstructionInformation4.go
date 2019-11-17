package model

// Set of characteristics that apply to the credit side of the payment transactions included in the direct debit initiation.
type PaymentInstructionInformation4 struct {

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
	PaymentTypeInformation *PaymentTypeInformation20 `xml:"PmtTpInf,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification32 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount16 `xml:"CdtrAcct"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent in the payment chain.
	CreditorAgentAccount *CashAccount16 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification32 `xml:"UltmtCdtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Account used to process charges associated with a transaction.
	//
	// Usage: Charges account should be used when charges have to be booked to an account different from the account identified in debtor's account.
	ChargesAccount *CashAccount16 `xml:"ChrgsAcct,omitempty"`

	// Agent that services a charges account.
	//
	// Usage: Charges account agent should only be used when the charges account agent is different from the creditor agent.
	ChargesAccountAgent *BranchAndFinancialInstitutionIdentification4 `xml:"ChrgsAcctAgt,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification32 `xml:"CdtrSchmeId,omitempty"`

	// Set of elements used to provide information on the individual transaction(s) included in the message.
	DirectDebitTransactionInformation []*DirectDebitTransactionInformation9 `xml:"DrctDbtTxInf"`
}

func (p *PaymentInstructionInformation4) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstructionInformation4) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod2Code)(&value)
}

func (p *PaymentInstructionInformation4) SetBatchBooking(value string) {
	p.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (p *PaymentInstructionInformation4) SetNumberOfTransactions(value string) {
	p.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (p *PaymentInstructionInformation4) SetControlSum(value string) {
	p.ControlSum = (*DecimalNumber)(&value)
}

func (p *PaymentInstructionInformation4) AddPaymentTypeInformation() *PaymentTypeInformation20 {
	p.PaymentTypeInformation = new(PaymentTypeInformation20)
	return p.PaymentTypeInformation
}

func (p *PaymentInstructionInformation4) SetRequestedCollectionDate(value string) {
	p.RequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentInstructionInformation4) AddCreditor() *PartyIdentification32 {
	p.Creditor = new(PartyIdentification32)
	return p.Creditor
}

func (p *PaymentInstructionInformation4) AddCreditorAccount() *CashAccount16 {
	p.CreditorAccount = new(CashAccount16)
	return p.CreditorAccount
}

func (p *PaymentInstructionInformation4) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification4 {
	p.CreditorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return p.CreditorAgent
}

func (p *PaymentInstructionInformation4) AddCreditorAgentAccount() *CashAccount16 {
	p.CreditorAgentAccount = new(CashAccount16)
	return p.CreditorAgentAccount
}

func (p *PaymentInstructionInformation4) AddUltimateCreditor() *PartyIdentification32 {
	p.UltimateCreditor = new(PartyIdentification32)
	return p.UltimateCreditor
}

func (p *PaymentInstructionInformation4) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstructionInformation4) AddChargesAccount() *CashAccount16 {
	p.ChargesAccount = new(CashAccount16)
	return p.ChargesAccount
}

func (p *PaymentInstructionInformation4) AddChargesAccountAgent() *BranchAndFinancialInstitutionIdentification4 {
	p.ChargesAccountAgent = new(BranchAndFinancialInstitutionIdentification4)
	return p.ChargesAccountAgent
}

func (p *PaymentInstructionInformation4) AddCreditorSchemeIdentification() *PartyIdentification32 {
	p.CreditorSchemeIdentification = new(PartyIdentification32)
	return p.CreditorSchemeIdentification
}

func (p *PaymentInstructionInformation4) AddDirectDebitTransactionInformation() *DirectDebitTransactionInformation9 {
	newValue := new(DirectDebitTransactionInformation9)
	p.DirectDebitTransactionInformation = append(p.DirectDebitTransactionInformation, newValue)
	return newValue
}
