package model

// Choice between types of payment instrument, ie, cheque, credit transfer, direct debit, investment account or payment card.
type PaymentInstrument8Choice struct {

	// Electronic money product that provides the cardholder with a portable and specialised computer device, which typically contains a microprocessor.
	PaymentCardDetails *PaymentCard2 `xml:"PmtCardDtls"`

	// Payment instrument between a debtor and a creditor, which flows through one or more financial institutions or systems.
	CreditTransferDetails *CreditTransfer3 `xml:"CdtTrfDtls"`

	// Instruction, initiated by the creditor, to debit a debtor's account in favour of the creditor. A direct debit can be pre-authorised or not. In most countries, authorisation is in the form of a mandate between the debtor and creditor.
	DirectDebitDetails *DirectDebitMandate2 `xml:"DrctDbtDtls"`

	// Written order on which instructions are given to an account holder (a financial institution) to pay a stated sum to a named recipient (the payee).
	ChequeDetails *Cheque3 `xml:"ChqDtls"`

	// Part of the investment account to or from which cash entries are made.
	AccountDetails *InvestmentAccount15 `xml:"AcctDtls"`
}

func (p *PaymentInstrument8Choice) AddPaymentCardDetails() *PaymentCard2 {
	p.PaymentCardDetails = new(PaymentCard2)
	return p.PaymentCardDetails
}

func (p *PaymentInstrument8Choice) AddCreditTransferDetails() *CreditTransfer3 {
	p.CreditTransferDetails = new(CreditTransfer3)
	return p.CreditTransferDetails
}

func (p *PaymentInstrument8Choice) AddDirectDebitDetails() *DirectDebitMandate2 {
	p.DirectDebitDetails = new(DirectDebitMandate2)
	return p.DirectDebitDetails
}

func (p *PaymentInstrument8Choice) AddChequeDetails() *Cheque3 {
	p.ChequeDetails = new(Cheque3)
	return p.ChequeDetails
}

func (p *PaymentInstrument8Choice) AddAccountDetails() *InvestmentAccount15 {
	p.AccountDetails = new(InvestmentAccount15)
	return p.AccountDetails
}
