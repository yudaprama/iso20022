package model

// Set of elements used to identify the type or operations code of a transaction entry.
type BankTransactionCodeStructure4 struct {

	// Set of elements used to provide the domain, the family and the sub-family of the bank transaction code, in a structured and hierarchical format.
	//
	// Usage: If a specific family or sub-family code cannot be provided, the generic family code defined for the domain or the generic sub-family code defined for the family should be provided.
	Domain *BankTransactionCodeStructure5 `xml:"Domn,omitempty"`

	// Bank transaction code in a proprietary form, as defined by the issuer.
	Proprietary *ProprietaryBankTransactionCodeStructure1 `xml:"Prtry,omitempty"`
}

func (b *BankTransactionCodeStructure4) AddDomain() *BankTransactionCodeStructure5 {
	b.Domain = new(BankTransactionCodeStructure5)
	return b.Domain
}

func (b *BankTransactionCodeStructure4) AddProprietary() *ProprietaryBankTransactionCodeStructure1 {
	b.Proprietary = new(ProprietaryBankTransactionCodeStructure1)
	return b.Proprietary
}
