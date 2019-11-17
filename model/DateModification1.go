package model

// Specifies the type of change to a date.
type DateModification1 struct {

	// Specifies the type of change.
	ModificationCode *Modification1Code `xml:"ModCd,omitempty"`

	// Date.
	Date *ISODate `xml:"Dt"`
}

func (d *DateModification1) SetModificationCode(value string) {
	d.ModificationCode = (*Modification1Code)(&value)
}

func (d *DateModification1) SetDate(value string) {
	d.Date = (*ISODate)(&value)
}
