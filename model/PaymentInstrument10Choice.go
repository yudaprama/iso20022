package model

// Choice between types of payment instrument, ie, credit transfer, cheque, payment card, investment cash account or direct debit.
type PaymentInstrument10Choice struct {

	// Electronic money product that provides the cardholder with a portable and specialised computer device, which typically contains a microprocessor.
	PaymentCardDetails *PaymentCard2 `xml:"PmtCardDtls"`

	// Payment instrument between a debtor and a creditor, which flows through one or more financial institutions or systems.
	CreditTransferDetails *CreditTransfer4 `xml:"CdtTrfDtls"`

	// Instruction, initiated by the creditor, to debit a debtor's account in favour of the creditor. A direct debit can be pre-authorised or not. In most countries, authorisation is in the form of a mandate between the debtor and creditor.
	DirectDebitDetails *DirectDebitMandate2 `xml:"DrctDbtDtls"`

	// Written order on which instructions are given to an account holder (a financial institution) to pay a stated sum to a named recipient (the payee).
	ChequeDetails *Cheque3 `xml:"ChqDtls"`

	// Part of the investment account to or from which cash entries are made.
	AccountDetails *InvestmentAccount15 `xml:"AcctDtls"`
}

func (p *PaymentInstrument10Choice) AddPaymentCardDetails() *PaymentCard2 {
	p.PaymentCardDetails = new(PaymentCard2)
	return p.PaymentCardDetails
}

func (p *PaymentInstrument10Choice) AddCreditTransferDetails() *CreditTransfer4 {
	p.CreditTransferDetails = new(CreditTransfer4)
	return p.CreditTransferDetails
}

func (p *PaymentInstrument10Choice) AddDirectDebitDetails() *DirectDebitMandate2 {
	p.DirectDebitDetails = new(DirectDebitMandate2)
	return p.DirectDebitDetails
}

func (p *PaymentInstrument10Choice) AddChequeDetails() *Cheque3 {
	p.ChequeDetails = new(Cheque3)
	return p.ChequeDetails
}

func (p *PaymentInstrument10Choice) AddAccountDetails() *InvestmentAccount15 {
	p.AccountDetails = new(InvestmentAccount15)
	return p.AccountDetails
}
