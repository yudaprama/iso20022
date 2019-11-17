package model

// Identifies the response parameters for which there is a match with the search criteria.
type InvestigatedParties1Choice struct {

	// Specifies the investigated parties as a code.
	Code *InvestigatedParties1Code `xml:"Cd"`

	// Specifies the investigated parties as a proprietary code.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (i *InvestigatedParties1Choice) SetCode(value string) {
	i.Code = (*InvestigatedParties1Code)(&value)
}

func (i *InvestigatedParties1Choice) SetProprietary(value string) {
	i.Proprietary = (*Max35Text)(&value)
}
