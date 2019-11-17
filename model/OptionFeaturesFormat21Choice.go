package model

// Choice between a standard code or a proprietary code to specify the features that may apply to a corporate action option.
type OptionFeaturesFormat21Choice struct {

	// Standard code to specify the features that may apply to a corporate action option.
	Code *OptionFeatures8Code `xml:"Cd"`

	// Proprietary identification of the features that may apply to a corporate action option.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (o *OptionFeaturesFormat21Choice) SetCode(value string) {
	o.Code = (*OptionFeatures8Code)(&value)
}

func (o *OptionFeaturesFormat21Choice) AddProprietary() *GenericIdentification47 {
	o.Proprietary = new(GenericIdentification47)
	return o.Proprietary
}
