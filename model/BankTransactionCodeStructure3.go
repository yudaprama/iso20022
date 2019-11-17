package model

// Code of the underlying bank transaction.
type BankTransactionCodeStructure3 struct {

	// Specifies the family within a domain.
	Code *ExternalBankTransactionFamilyCode `xml:"Cd"`

	// Specifies the sub-product family within a specific family.
	SubFamilyCode *ExternalBankTransactionSubFamilyCode `xml:"SubFmlyCd"`
}

func (b *BankTransactionCodeStructure3) SetCode(value string) {
	b.Code = (*ExternalBankTransactionFamilyCode)(&value)
}

func (b *BankTransactionCodeStructure3) SetSubFamilyCode(value string) {
	b.SubFamilyCode = (*ExternalBankTransactionSubFamilyCode)(&value)
}
