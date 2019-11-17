package model

// Specifies the type of update requested. That is addition, deletion or modification.
type UpdateType22Choice struct {

	// Addition of information to the securities transaction.
	Addition *SecuritiesSettlementTransactionDetails23 `xml:"Addtn"`

	// Deletion of information in the securities transaction.
	Deletion *SecuritiesSettlementTransactionDetails24 `xml:"Deltn"`

	// Modification of information in the securities transaction.
	Modification *SecuritiesSettlementTransactionDetails25 `xml:"Mod"`
}

func (u *UpdateType22Choice) AddAddition() *SecuritiesSettlementTransactionDetails23 {
	u.Addition = new(SecuritiesSettlementTransactionDetails23)
	return u.Addition
}

func (u *UpdateType22Choice) AddDeletion() *SecuritiesSettlementTransactionDetails24 {
	u.Deletion = new(SecuritiesSettlementTransactionDetails24)
	return u.Deletion
}

func (u *UpdateType22Choice) AddModification() *SecuritiesSettlementTransactionDetails25 {
	u.Modification = new(SecuritiesSettlementTransactionDetails25)
	return u.Modification
}
