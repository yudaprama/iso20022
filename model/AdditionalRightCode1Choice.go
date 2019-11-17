package model

// Choice of additional right code.
type AdditionalRightCode1Choice struct {

	// Identifies the additional right code.
	Code *AdditionalRight1Code `xml:"Cd"`

	// This code can be used in case another reason is required.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (a *AdditionalRightCode1Choice) SetCode(value string) {
	a.Code = (*AdditionalRight1Code)(&value)
}

func (a *AdditionalRightCode1Choice) AddProprietary() *GenericIdentification13 {
	a.Proprietary = new(GenericIdentification13)
	return a.Proprietary
}
