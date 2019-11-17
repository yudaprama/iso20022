package model

// Specifies the type of additional information.
type InformationType1Choice struct {

	// Type of additional information in a coded form.
	Code *InformationType1Code `xml:"Cd"`

	// Type of additional information not catered for by the available codes.
	Proprietary *Max140Text `xml:"Prtry"`
}

func (i *InformationType1Choice) SetCode(value string) {
	i.Code = (*InformationType1Code)(&value)
}

func (i *InformationType1Choice) SetProprietary(value string) {
	i.Proprietary = (*Max140Text)(&value)
}
