package model

// Specifies the technical input channel.
type TechnicalInputChannel1Choice struct {

	// Technical input channel, as published in an external technical input channel code list.
	Code *ExternalTechnicalInputChannel1Code `xml:"Cd"`

	// Technical channel used to input the instruction, in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (t *TechnicalInputChannel1Choice) SetCode(value string) {
	t.Code = (*ExternalTechnicalInputChannel1Code)(&value)
}

func (t *TechnicalInputChannel1Choice) SetProprietary(value string) {
	t.Proprietary = (*Max35Text)(&value)
}
