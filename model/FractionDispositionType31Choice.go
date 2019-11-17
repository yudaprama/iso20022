package model

// Choice between a standard code or a proprietary code to specify the type of fraction disposition.
type FractionDispositionType31Choice struct {

	// Standard code to specify how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	Code *FractionDispositionType8Code `xml:"Cd"`

	// Proprietary identification of the type of fraction disposition.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *FractionDispositionType31Choice) SetCode(value string) {
	f.Code = (*FractionDispositionType8Code)(&value)
}

func (f *FractionDispositionType31Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
