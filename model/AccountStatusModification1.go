package model

// Specifies the type of change to the status of the account.
type AccountStatusModification1 struct {

	// Specifies the type of change.
	ModificationCode *Modification1Code `xml:"ModCd,omitempty"`

	// Status of the account
	Status *AccountStatus3Code `xml:"Sts"`
}

func (a *AccountStatusModification1) SetModificationCode(value string) {
	a.ModificationCode = (*Modification1Code)(&value)
}

func (a *AccountStatusModification1) SetStatus(value string) {
	a.Status = (*AccountStatus3Code)(&value)
}
