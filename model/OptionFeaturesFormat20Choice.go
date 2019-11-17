package model

// Choice between a standard code or a proprietary code to specify the features that may apply to a corporate action option.
type OptionFeaturesFormat20Choice struct {

	// Standard code to specify the features that may apply to a corporate action option.
	Code *OptionFeatures7Code `xml:"Cd"`

	// Proprietary identification of the features that may apply to a corporate action option.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (o *OptionFeaturesFormat20Choice) SetCode(value string) {
	o.Code = (*OptionFeatures7Code)(&value)
}

func (o *OptionFeaturesFormat20Choice) AddProprietary() *GenericIdentification47 {
	o.Proprietary = new(GenericIdentification47)
	return o.Proprietary
}
