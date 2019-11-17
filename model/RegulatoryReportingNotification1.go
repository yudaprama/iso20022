package model

// Provides the details on the regulatory notification.
type RegulatoryReportingNotification1 struct {

	// Unique and unambiguous identification of the transaction notification.
	TransactionNotificationIdentification *Max35Text `xml:"TxNtfctnId"`

	// Party that legally owns the cash account.
	AccountOwner *PartyIdentification77 `xml:"AcctOwnr"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcr"`

	// Certificate against which all currency control transactions are registered.
	TransactionCertificate []*TransactionCertificate1 `xml:"TxCert"`
}

func (r *RegulatoryReportingNotification1) SetTransactionNotificationIdentification(value string) {
	r.TransactionNotificationIdentification = (*Max35Text)(&value)
}

func (r *RegulatoryReportingNotification1) AddAccountOwner() *PartyIdentification77 {
	r.AccountOwner = new(PartyIdentification77)
	return r.AccountOwner
}

func (r *RegulatoryReportingNotification1) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	r.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return r.AccountServicer
}

func (r *RegulatoryReportingNotification1) AddTransactionCertificate() *TransactionCertificate1 {
	newValue := new(TransactionCertificate1)
	r.TransactionCertificate = append(r.TransactionCertificate, newValue)
	return newValue
}
