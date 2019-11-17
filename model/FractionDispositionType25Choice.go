package model

// Choice between a standard code or a proprietary code to specify the type of fraction disposition.
type FractionDispositionType25Choice struct {

	// Standard code to specify how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	Code *FractionDispositionType9Code `xml:"Cd"`

	// Proprietary identification of the type of fraction disposition.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (f *FractionDispositionType25Choice) SetCode(value string) {
	f.Code = (*FractionDispositionType9Code)(&value)
}

func (f *FractionDispositionType25Choice) AddProprietary() *GenericIdentification30 {
	f.Proprietary = new(GenericIdentification30)
	return f.Proprietary
}
