package model

// Choice between a standard code or a proprietary code to specify the type of fraction disposition.
type FractionDispositionType26Choice struct {

	// Standard code to specify how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	Code *FractionDispositionType8Code `xml:"Cd"`

	// Proprietary identification of the type of fraction disposition.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (f *FractionDispositionType26Choice) SetCode(value string) {
	f.Code = (*FractionDispositionType8Code)(&value)
}

func (f *FractionDispositionType26Choice) AddProprietary() *GenericIdentification30 {
	f.Proprietary = new(GenericIdentification30)
	return f.Proprietary
}
