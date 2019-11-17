package model

// Specifies the type of modification to account type.
type TypeModification1 struct {

	// Specifies the type of change.
	ModificationCode *Modification1Code `xml:"ModCd,omitempty"`

	// Type of the account.
	Type *CashAccountType2Choice `xml:"Tp"`
}

func (t *TypeModification1) SetModificationCode(value string) {
	t.ModificationCode = (*Modification1Code)(&value)
}

func (t *TypeModification1) AddType() *CashAccountType2Choice {
	t.Type = new(CashAccountType2Choice)
	return t.Type
}
