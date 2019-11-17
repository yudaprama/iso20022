package model

// Specifies the type of update requested. That is addition, deletion or modification.
type UpdateType11Choice struct {

	// Addition of information to the securities transaction.
	Addition *SecuritiesSettlementTransactionDetails15 `xml:"Addtn"`

	// Deletion of information in the securities transaction.
	Deletion *SecuritiesSettlementTransactionDetails16 `xml:"Deltn"`

	// Modification of information in the securities transaction.
	Modification *SecuritiesSettlementTransactionDetails14 `xml:"Mod"`
}

func (u *UpdateType11Choice) AddAddition() *SecuritiesSettlementTransactionDetails15 {
	u.Addition = new(SecuritiesSettlementTransactionDetails15)
	return u.Addition
}

func (u *UpdateType11Choice) AddDeletion() *SecuritiesSettlementTransactionDetails16 {
	u.Deletion = new(SecuritiesSettlementTransactionDetails16)
	return u.Deletion
}

func (u *UpdateType11Choice) AddModification() *SecuritiesSettlementTransactionDetails14 {
	u.Modification = new(SecuritiesSettlementTransactionDetails14)
	return u.Modification
}
