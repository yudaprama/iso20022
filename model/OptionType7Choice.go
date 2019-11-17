package model

// Choice of format for the option type.
type OptionType7Choice struct {

	// Option type expressed as an ISO 20022 code.
	Code *OptionType1Code `xml:"Cd"`

	// Option type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (o *OptionType7Choice) SetCode(value string) {
	o.Code = (*OptionType1Code)(&value)
}

func (o *OptionType7Choice) AddProprietary() *GenericIdentification47 {
	o.Proprietary = new(GenericIdentification47)
	return o.Proprietary
}
