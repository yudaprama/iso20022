package model

// Certificate against which all currency control transactions are registered.
type TransactionCertificate1 struct {

	// Unique and unambiguous identification of the transaction.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Reference of the transaction certificate.
	Certificate *DocumentIdentification28 `xml:"Cert"`

	// Cash account, linked to the registered contract, on which the cash entry is made.
	Account *CashAccount24 `xml:"Acct,omitempty"`

	// Country in which the bank account is located, when the bank which services the account is located in another country.
	BankAccountDomiciliationCountry *CountryCode `xml:"BkAcctDmcltnCtry,omitempty"`

	// Amendment indicator details.
	Amendment *DocumentAmendment1 `xml:"Amdmnt,omitempty"`

	// Record of the transaction certificate.
	CertificateRecord []*TransactionCertificateRecord1 `xml:"CertRcrd"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TransactionCertificate1) SetTransactionIdentification(value string) {
	t.TransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionCertificate1) AddCertificate() *DocumentIdentification28 {
	t.Certificate = new(DocumentIdentification28)
	return t.Certificate
}

func (t *TransactionCertificate1) AddAccount() *CashAccount24 {
	t.Account = new(CashAccount24)
	return t.Account
}

func (t *TransactionCertificate1) SetBankAccountDomiciliationCountry(value string) {
	t.BankAccountDomiciliationCountry = (*CountryCode)(&value)
}

func (t *TransactionCertificate1) AddAmendment() *DocumentAmendment1 {
	t.Amendment = new(DocumentAmendment1)
	return t.Amendment
}

func (t *TransactionCertificate1) AddCertificateRecord() *TransactionCertificateRecord1 {
	newValue := new(TransactionCertificateRecord1)
	t.CertificateRecord = append(t.CertificateRecord, newValue)
	return newValue
}

func (t *TransactionCertificate1) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}
