package model

// Means by which a payment will be or has been made for settlement purposes.
type PaymentMeans1 struct {

	// Type, or nature, of the payment, eg, express payment.
	PaymentType *PaymentTypeInformation19 `xml:"PmtTp"`

	// Transfer method to be used for the transfer.
	PaymentMethodCode *PaymentMethod4Code `xml:"PmtMtdCd"`

	// Creditor financial account of the payee party for this payment means.
	PayeeCreditorAccount *CashAccount16 `xml:"PyeeCdtrAcct"`

	// Creditor financial institution of the payee party specified for this payment means.
	PayeeFinancialInstitution *BranchAndFinancialInstitutionIdentification4 `xml:"PyeeFI"`

	// Debtor financial account of the payer party for this payment means.
	PayerDebtorAccount *CashAccount16 `xml:"PyerDbtrAcct,omitempty"`

	// Debtor financial institution of the payer party specified for this payment means.
	PayerFinancialInstitution *BranchAndFinancialInstitutionIdentification4 `xml:"PyerFI,omitempty"`
}

func (p *PaymentMeans1) AddPaymentType() *PaymentTypeInformation19 {
	p.PaymentType = new(PaymentTypeInformation19)
	return p.PaymentType
}

func (p *PaymentMeans1) SetPaymentMethodCode(value string) {
	p.PaymentMethodCode = (*PaymentMethod4Code)(&value)
}

func (p *PaymentMeans1) AddPayeeCreditorAccount() *CashAccount16 {
	p.PayeeCreditorAccount = new(CashAccount16)
	return p.PayeeCreditorAccount
}

func (p *PaymentMeans1) AddPayeeFinancialInstitution() *BranchAndFinancialInstitutionIdentification4 {
	p.PayeeFinancialInstitution = new(BranchAndFinancialInstitutionIdentification4)
	return p.PayeeFinancialInstitution
}

func (p *PaymentMeans1) AddPayerDebtorAccount() *CashAccount16 {
	p.PayerDebtorAccount = new(CashAccount16)
	return p.PayerDebtorAccount
}

func (p *PaymentMeans1) AddPayerFinancialInstitution() *BranchAndFinancialInstitutionIdentification4 {
	p.PayerFinancialInstitution = new(BranchAndFinancialInstitutionIdentification4)
	return p.PayerFinancialInstitution
}
