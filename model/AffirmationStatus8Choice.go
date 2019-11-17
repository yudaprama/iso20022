package model

// Choice of format for the affirmation status.
type AffirmationStatus8Choice struct {

	// Provides the status of the trade at confirmation level at the time the settlement instruction was sent.
	Code *AffirmationStatus1Code `xml:"Cd"`

	// Provides the status of the trade at confirmation level at the time the settlement instruction was sent.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (a *AffirmationStatus8Choice) SetCode(value string) {
	a.Code = (*AffirmationStatus1Code)(&value)
}

func (a *AffirmationStatus8Choice) AddProprietary() *GenericIdentification30 {
	a.Proprietary = new(GenericIdentification30)
	return a.Proprietary
}
