package model

// Specifies the type of update requested. That is addition, deletion or modification.
type UpdateType14Choice struct {

	// Addition of information to the securities transaction.
	Addition *SecuritiesSettlementTransactionDetails21 `xml:"Addtn"`

	// Deletion of information in the securities transaction.
	Deletion *SecuritiesSettlementTransactionDetails20 `xml:"Deltn"`

	// Modification of information in the securities transaction.
	Modification *SecuritiesSettlementTransactionDetails22 `xml:"Mod"`
}

func (u *UpdateType14Choice) AddAddition() *SecuritiesSettlementTransactionDetails21 {
	u.Addition = new(SecuritiesSettlementTransactionDetails21)
	return u.Addition
}

func (u *UpdateType14Choice) AddDeletion() *SecuritiesSettlementTransactionDetails20 {
	u.Deletion = new(SecuritiesSettlementTransactionDetails20)
	return u.Deletion
}

func (u *UpdateType14Choice) AddModification() *SecuritiesSettlementTransactionDetails22 {
	u.Modification = new(SecuritiesSettlementTransactionDetails22)
	return u.Modification
}
