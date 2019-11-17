package model

// Choice between a standard code or a proprietary code to specify the features that may apply to a corporate action option.
type OptionFeaturesFormat17Choice struct {

	// Standard code to specify the features that may apply to a corporate action option.
	Code *OptionFeatures7Code `xml:"Cd"`

	// Proprietary identification of the features that may apply to a corporate action option.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (o *OptionFeaturesFormat17Choice) SetCode(value string) {
	o.Code = (*OptionFeatures7Code)(&value)
}

func (o *OptionFeaturesFormat17Choice) AddProprietary() *GenericIdentification30 {
	o.Proprietary = new(GenericIdentification30)
	return o.Proprietary
}
