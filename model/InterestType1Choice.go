package model

// Interest type is expressed as a code or a text.
type InterestType1Choice struct {

	// Specifies the type of interest.
	Code *InterestType1Code `xml:"Cd"`

	// Specifies the type of interest in uncoded form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (i *InterestType1Choice) SetCode(value string) {
	i.Code = (*InterestType1Code)(&value)
}

func (i *InterestType1Choice) SetProprietary(value string) {
	i.Proprietary = (*Max35Text)(&value)
}
