package model

// Code of the underlying bank transaction.
type BankTransactionCodeStructure2 struct {

	// Specifies the business area of the underlying transaction.
	Code *ExternalBankTransactionDomainCode `xml:"Cd"`

	// Specifies the family and the sub-family of the bank transaction code, within a specific domain, in a structured and hierarchical format.
	Family *BankTransactionCodeStructure3 `xml:"Fmly"`
}

func (b *BankTransactionCodeStructure2) SetCode(value string) {
	b.Code = (*ExternalBankTransactionDomainCode)(&value)
}

func (b *BankTransactionCodeStructure2) AddFamily() *BankTransactionCodeStructure3 {
	b.Family = new(BankTransactionCodeStructure3)
	return b.Family
}
