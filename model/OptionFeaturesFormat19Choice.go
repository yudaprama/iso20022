package model

// Choice between a standard code or a proprietary code to specify the features that may apply to a corporate action option.
type OptionFeaturesFormat19Choice struct {

	// Standard code to specify the features that may apply to a corporate action option.
	Code *OptionFeatures6Code `xml:"Cd"`

	// Proprietary identification of the features that may apply to a corporate action option.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (o *OptionFeaturesFormat19Choice) SetCode(value string) {
	o.Code = (*OptionFeatures6Code)(&value)
}

func (o *OptionFeaturesFormat19Choice) AddProprietary() *GenericIdentification47 {
	o.Proprietary = new(GenericIdentification47)
	return o.Proprietary
}
