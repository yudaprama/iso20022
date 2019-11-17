package model

// Choice between a standard code or a proprietary code to specify the features that may apply to a corporate action option.
type OptionFeaturesFormat9Choice struct {

	// Standard code to specify the features that may apply to a corporate action option.
	Code *OptionFeatures4Code `xml:"Cd"`

	// Proprietary identification of the features that may apply to a corporate action option.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (o *OptionFeaturesFormat9Choice) SetCode(value string) {
	o.Code = (*OptionFeatures4Code)(&value)
}

func (o *OptionFeaturesFormat9Choice) AddProprietary() *GenericIdentification20 {
	o.Proprietary = new(GenericIdentification20)
	return o.Proprietary
}
