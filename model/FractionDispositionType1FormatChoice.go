package model

// Choice of formats to  express how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
type FractionDispositionType1FormatChoice struct {

	// Standard code to  specify how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	Code *FractionDispositionType1Code `xml:"Cd"`

	// Proprietary code to  express how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (f *FractionDispositionType1FormatChoice) SetCode(value string) {
	f.Code = (*FractionDispositionType1Code)(&value)
}

func (f *FractionDispositionType1FormatChoice) AddProprietary() *GenericIdentification13 {
	f.Proprietary = new(GenericIdentification13)
	return f.Proprietary
}
