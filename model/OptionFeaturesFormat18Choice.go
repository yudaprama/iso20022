package model

// Choice between a standard code or a proprietary code to specify the features that may apply to a corporate action option.
type OptionFeaturesFormat18Choice struct {

	// Standard code to specify the features that may apply to a corporate action option.
	Code *OptionFeatures6Code `xml:"Cd"`

	// Proprietary identification of the features that may apply to a corporate action option.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (o *OptionFeaturesFormat18Choice) SetCode(value string) {
	o.Code = (*OptionFeatures6Code)(&value)
}

func (o *OptionFeaturesFormat18Choice) AddProprietary() *GenericIdentification30 {
	o.Proprietary = new(GenericIdentification30)
	return o.Proprietary
}
