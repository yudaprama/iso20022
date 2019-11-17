package model

// Choice between a standard code or a proprietary code to specify the features that may apply to a corporate action option.
type OptionFeaturesFormat16Choice struct {

	// Standard code to specify the features that may apply to a corporate action option.
	Code *OptionFeatures8Code `xml:"Cd"`

	// Proprietary identification of the features that may apply to a corporate action option.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (o *OptionFeaturesFormat16Choice) SetCode(value string) {
	o.Code = (*OptionFeatures8Code)(&value)
}

func (o *OptionFeaturesFormat16Choice) AddProprietary() *GenericIdentification30 {
	o.Proprietary = new(GenericIdentification30)
	return o.Proprietary
}
