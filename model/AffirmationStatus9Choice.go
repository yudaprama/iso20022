package model

// Choice of format for the affirmation status.
type AffirmationStatus9Choice struct {

	// Provides the status of the trade at confirmation level at the time the settlement instruction was sent.
	Code *AffirmationStatus1Code `xml:"Cd"`

	// Provides the status of the trade at confirmation level at the time the settlement instruction was sent.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AffirmationStatus9Choice) SetCode(value string) {
	a.Code = (*AffirmationStatus1Code)(&value)
}

func (a *AffirmationStatus9Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
