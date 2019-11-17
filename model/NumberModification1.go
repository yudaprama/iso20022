package model

// Specifies the type of change to a number.
type NumberModification1 struct {

	// Specifies the type of change.
	ModificationCode *Modification1Code `xml:"ModCd,omitempty"`

	// Number.
	Number *Max5NumericText `xml:"Nb"`
}

func (n *NumberModification1) SetModificationCode(value string) {
	n.ModificationCode = (*Modification1Code)(&value)
}

func (n *NumberModification1) SetNumber(value string) {
	n.Number = (*Max5NumericText)(&value)
}
