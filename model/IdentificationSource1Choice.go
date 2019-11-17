package model

// Choice of proprietary or domestic identification scheme that uniquely identifies a security.
type IdentificationSource1Choice struct {

	// Country of the proprietary identification scheme.
	Domestic *CountryCode `xml:"Dmst"`

	// Entity that issues the proprietary identification.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (i *IdentificationSource1Choice) SetDomestic(value string) {
	i.Domestic = (*CountryCode)(&value)
}

func (i *IdentificationSource1Choice) SetProprietary(value string) {
	i.Proprietary = (*Max35Text)(&value)
}
