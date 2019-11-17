package model

// Specifies the type of update requested. That is addition, deletion or modification.
type UpdateType25Choice struct {

	// Addition of information to the securities transaction.
	Addition *SecuritiesSettlementTransactionDetails26 `xml:"Addtn"`

	// Deletion of information in the securities transaction.
	Deletion *SecuritiesSettlementTransactionDetails27 `xml:"Deltn"`

	// Modification of information in the securities transaction.
	Modification *SecuritiesSettlementTransactionDetails28 `xml:"Mod"`
}

func (u *UpdateType25Choice) AddAddition() *SecuritiesSettlementTransactionDetails26 {
	u.Addition = new(SecuritiesSettlementTransactionDetails26)
	return u.Addition
}

func (u *UpdateType25Choice) AddDeletion() *SecuritiesSettlementTransactionDetails27 {
	u.Deletion = new(SecuritiesSettlementTransactionDetails27)
	return u.Deletion
}

func (u *UpdateType25Choice) AddModification() *SecuritiesSettlementTransactionDetails28 {
	u.Modification = new(SecuritiesSettlementTransactionDetails28)
	return u.Modification
}
