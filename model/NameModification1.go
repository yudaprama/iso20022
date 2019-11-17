package model

// Specifies the type of change to name.
type NameModification1 struct {

	// Specifies the type of change.
	ModificationCode *Modification1Code `xml:"ModCd,omitempty"`

	// Name of the account.
	Name *Max70Text `xml:"Nm"`
}

func (n *NameModification1) SetModificationCode(value string) {
	n.ModificationCode = (*Modification1Code)(&value)
}

func (n *NameModification1) SetName(value string) {
	n.Name = (*Max70Text)(&value)
}
