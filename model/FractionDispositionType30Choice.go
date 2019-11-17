package model

// Choice between a standard code or a proprietary code to specify the type of fraction disposition.
type FractionDispositionType30Choice struct {

	// Standard code to specify how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	Code *FractionDispositionType11Code `xml:"Cd"`

	// Proprietary identification of the type of fraction disposition.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *FractionDispositionType30Choice) SetCode(value string) {
	f.Code = (*FractionDispositionType11Code)(&value)
}

func (f *FractionDispositionType30Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
