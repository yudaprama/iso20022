package model

// Set of elements used to identify the type or operations code of a transaction entry.
type BankTransactionCodeStructure6 struct {

	// Specifies the family within a domain.
	Code *ExternalBankTransactionFamily1Code `xml:"Cd"`

	// Specifies the sub-product family within a specific family.
	SubFamilyCode *ExternalBankTransactionSubFamily1Code `xml:"SubFmlyCd"`
}

func (b *BankTransactionCodeStructure6) SetCode(value string) {
	b.Code = (*ExternalBankTransactionFamily1Code)(&value)
}

func (b *BankTransactionCodeStructure6) SetSubFamilyCode(value string) {
	b.SubFamilyCode = (*ExternalBankTransactionSubFamily1Code)(&value)
}
