package model

// Choice of formats to  express the  feature of an option.
type OptionFeatures1FormatChoice struct {

	// Standard code to  specify the  feature of an option.
	Code *OptionFeatures1Code `xml:"Cd"`

	// Proprietary code to  express the  feature of an option.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (o *OptionFeatures1FormatChoice) SetCode(value string) {
	o.Code = (*OptionFeatures1Code)(&value)
}

func (o *OptionFeatures1FormatChoice) AddProprietary() *GenericIdentification13 {
	o.Proprietary = new(GenericIdentification13)
	return o.Proprietary
}
