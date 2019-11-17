package model

// Set of elements used to identify the type or operations code of a transaction entry.
type BankTransactionCodeStructure5 struct {

	// Specifies the business area of the underlying transaction.
	Code *ExternalBankTransactionDomain1Code `xml:"Cd"`

	// Specifies the family and the sub-family of the bank transaction code, within a specific domain, in a structured and hierarchical format.
	Family *BankTransactionCodeStructure6 `xml:"Fmly"`
}

func (b *BankTransactionCodeStructure5) SetCode(value string) {
	b.Code = (*ExternalBankTransactionDomain1Code)(&value)
}

func (b *BankTransactionCodeStructure5) AddFamily() *BankTransactionCodeStructure6 {
	b.Family = new(BankTransactionCodeStructure6)
	return b.Family
}
