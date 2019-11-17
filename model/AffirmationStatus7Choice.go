package model

// Choice of format for the affirmation status.
type AffirmationStatus7Choice struct {

	// Status of affirmation of a trade.
	Code *AffirmationStatus1Code `xml:"Cd"`

	// Trade has been unaffirmed.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (a *AffirmationStatus7Choice) SetCode(value string) {
	a.Code = (*AffirmationStatus1Code)(&value)
}

func (a *AffirmationStatus7Choice) AddProprietary() *GenericIdentification38 {
	a.Proprietary = new(GenericIdentification38)
	return a.Proprietary
}
