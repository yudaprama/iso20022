package model

// Specifies the type of update requested. That is addition, deletion or modification.
type UpdateType9Choice struct {

	// Addition of information to the securities transaction.
	Addition *SecuritiesSettlementTransactionDetails8 `xml:"Addtn"`

	// Deletion of information in the securities transaction.
	Deletion *SecuritiesSettlementTransactionDetails9 `xml:"Deltn"`

	// Modification of information in the securities transaction.
	Modification *SecuritiesSettlementTransactionDetails10 `xml:"Mod"`
}

func (u *UpdateType9Choice) AddAddition() *SecuritiesSettlementTransactionDetails8 {
	u.Addition = new(SecuritiesSettlementTransactionDetails8)
	return u.Addition
}

func (u *UpdateType9Choice) AddDeletion() *SecuritiesSettlementTransactionDetails9 {
	u.Deletion = new(SecuritiesSettlementTransactionDetails9)
	return u.Deletion
}

func (u *UpdateType9Choice) AddModification() *SecuritiesSettlementTransactionDetails10 {
	u.Modification = new(SecuritiesSettlementTransactionDetails10)
	return u.Modification
}
