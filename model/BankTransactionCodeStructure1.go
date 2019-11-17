package model

// Set of elements to fully identify the type of the bank transaction entry.
type BankTransactionCodeStructure1 struct {

	// Specifies the domain, the family and the sub-family of the bank transaction code, in a structured and hierarchical format.
	//
	// Usage: If a specific family or subfamily code cannot be provided, the generic family code defined for the domain or the generic subfamily code defined for the family should be provided.
	Domain *BankTransactionCodeStructure2 `xml:"Domn,omitempty"`

	// Proprietary identification of the bank transaction code, as defined by the issuer.
	Proprietary *ProprietaryBankTransactionCodeStructure1 `xml:"Prtry,omitempty"`
}

func (b *BankTransactionCodeStructure1) AddDomain() *BankTransactionCodeStructure2 {
	b.Domain = new(BankTransactionCodeStructure2)
	return b.Domain
}

func (b *BankTransactionCodeStructure1) AddProprietary() *ProprietaryBankTransactionCodeStructure1 {
	b.Proprietary = new(ProprietaryBankTransactionCodeStructure1)
	return b.Proprietary
}
