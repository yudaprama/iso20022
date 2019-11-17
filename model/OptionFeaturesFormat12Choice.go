package model

// Choice between a standard code or a proprietary code to specify the features that may apply to a corporate action option.
type OptionFeaturesFormat12Choice struct {

	// Standard code to specify the features that may apply to a corporate action option.
	Code *OptionFeatures5Code `xml:"Cd"`

	// Proprietary identification of the features that may apply to a corporate action option.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (o *OptionFeaturesFormat12Choice) SetCode(value string) {
	o.Code = (*OptionFeatures5Code)(&value)
}

func (o *OptionFeaturesFormat12Choice) AddProprietary() *GenericIdentification20 {
	o.Proprietary = new(GenericIdentification20)
	return o.Proprietary
}
