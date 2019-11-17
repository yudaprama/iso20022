package model

// Choice between a standard code or a proprietary code to specify the type of fraction disposition.
type FractionDispositionType3Choice struct {

	// Standard code to specify how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	Code *FractionDispositionType2Code `xml:"Cd"`

	// Proprietary identification of the type of fraction disposition.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (f *FractionDispositionType3Choice) SetCode(value string) {
	f.Code = (*FractionDispositionType2Code)(&value)
}

func (f *FractionDispositionType3Choice) AddProprietary() *GenericIdentification20 {
	f.Proprietary = new(GenericIdentification20)
	return f.Proprietary
}
