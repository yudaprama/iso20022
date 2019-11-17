package model

// Specifies a service level restricted to the proprietary element.
type RestrictedProprietaryChoice struct {

	// Proprietary identification of a pre-agreed level of service between the parties.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (r *RestrictedProprietaryChoice) SetProprietary(value string) {
	r.Proprietary = (*Max35Text)(&value)
}
