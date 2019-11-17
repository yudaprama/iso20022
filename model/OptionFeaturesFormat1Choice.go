package model

// Choice between a standard code or a proprietary code to specify the features that may apply to a corporate action option.
type OptionFeaturesFormat1Choice struct {

	// Standard code to specify the features that may apply to a corporate action option.
	Code *OptionFeatures1Code `xml:"Cd"`

	// Proprietary identification of the features that may apply to a corporate action option.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (o *OptionFeaturesFormat1Choice) SetCode(value string) {
	o.Code = (*OptionFeatures1Code)(&value)
}

func (o *OptionFeaturesFormat1Choice) AddProprietary() *GenericIdentification20 {
	o.Proprietary = new(GenericIdentification20)
	return o.Proprietary
}
