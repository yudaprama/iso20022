package model

// Choice of format for the affirmation status.
type AffirmationStatus1Choice struct {

	// Provides the status of the trade at confirmation level at the time the settlement instruction was sent.
	Code *AffirmationStatus1Code `xml:"Cd"`

	// Provides the status of the trade at confirmation level at the time the settlement instruction was sent.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (a *AffirmationStatus1Choice) SetCode(value string) {
	a.Code = (*AffirmationStatus1Code)(&value)
}

func (a *AffirmationStatus1Choice) AddProprietary() *GenericIdentification20 {
	a.Proprietary = new(GenericIdentification20)
	return a.Proprietary
}
