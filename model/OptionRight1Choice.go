package model

// Choice of format for the option rights.
type OptionRight1Choice struct {

	// Option rights expressed as an ISO 20022 code.
	Code *OptionRight1Code `xml:"Cd"`

	// Option rights expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (o *OptionRight1Choice) SetCode(value string) {
	o.Code = (*OptionRight1Code)(&value)
}

func (o *OptionRight1Choice) AddProprietary() *GenericIdentification38 {
	o.Proprietary = new(GenericIdentification38)
	return o.Proprietary
}
