package model

// Specifies the type of update requested. That is addition, deletion or modification.
type UpdateType5Choice struct {

	// Addition of information to the securities transaction.
	Addition *SecuritiesSettlementTransactionDetails3 `xml:"Addtn"`

	// Deletion of information in the securities transaction.
	Deletion *SecuritiesSettlementTransactionDetails4 `xml:"Deltn"`

	// Modification of information in the securities transaction.
	Modification *SecuritiesSettlementTransactionDetails2 `xml:"Mod"`
}

func (u *UpdateType5Choice) AddAddition() *SecuritiesSettlementTransactionDetails3 {
	u.Addition = new(SecuritiesSettlementTransactionDetails3)
	return u.Addition
}

func (u *UpdateType5Choice) AddDeletion() *SecuritiesSettlementTransactionDetails4 {
	u.Deletion = new(SecuritiesSettlementTransactionDetails4)
	return u.Deletion
}

func (u *UpdateType5Choice) AddModification() *SecuritiesSettlementTransactionDetails2 {
	u.Modification = new(SecuritiesSettlementTransactionDetails2)
	return u.Modification
}
