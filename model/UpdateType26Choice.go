package model

// Specifies the type of update requested. That is addition, deletion or modification.
type UpdateType26Choice struct {

	// Addition of information to the securities transaction.
	Addition *SecuritiesSettlementTransactionDetails29 `xml:"Addtn"`

	// Deletion of information in the securities transaction.
	Deletion *SecuritiesSettlementTransactionDetails30 `xml:"Deltn"`

	// Modification of information in the securities transaction.
	Modification *SecuritiesSettlementTransactionDetails31 `xml:"Mod"`
}

func (u *UpdateType26Choice) AddAddition() *SecuritiesSettlementTransactionDetails29 {
	u.Addition = new(SecuritiesSettlementTransactionDetails29)
	return u.Addition
}

func (u *UpdateType26Choice) AddDeletion() *SecuritiesSettlementTransactionDetails30 {
	u.Deletion = new(SecuritiesSettlementTransactionDetails30)
	return u.Deletion
}

func (u *UpdateType26Choice) AddModification() *SecuritiesSettlementTransactionDetails31 {
	u.Modification = new(SecuritiesSettlementTransactionDetails31)
	return u.Modification
}
