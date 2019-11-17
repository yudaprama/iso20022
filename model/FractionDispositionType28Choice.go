package model

// Choice between a standard code or a proprietary code to specify the type of fraction disposition.
type FractionDispositionType28Choice struct {

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	Code *FractionDispositionType10Code `xml:"Cd"`

	// Proprietary identification of the type of fraction disposition.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (f *FractionDispositionType28Choice) SetCode(value string) {
	f.Code = (*FractionDispositionType10Code)(&value)
}

func (f *FractionDispositionType28Choice) AddProprietary() *GenericIdentification30 {
	f.Proprietary = new(GenericIdentification30)
	return f.Proprietary
}
