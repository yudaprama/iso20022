package model

// Choice between a standard code or a proprietary code to specify the type of fraction disposition.
type FractionDispositionType2Choice struct {

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	Code *FractionDispositionType3Code `xml:"Cd"`

	// Proprietary identification of the type of fraction disposition.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (f *FractionDispositionType2Choice) SetCode(value string) {
	f.Code = (*FractionDispositionType3Code)(&value)
}

func (f *FractionDispositionType2Choice) AddProprietary() *GenericIdentification20 {
	f.Proprietary = new(GenericIdentification20)
	return f.Proprietary
}
