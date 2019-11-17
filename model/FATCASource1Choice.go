package model

// Choice of formats for the source of the Foreign Account Tax Compliance Act (FATCA)
type FATCASource1Choice struct {

	// Source of the Foreign Account Tax Compliance Act (FATCA) status expressed as a code.
	Code *FATCASourceStatus1Code `xml:"Cd"`

	// Source of Foreign Account Tax Compliance Act (FATCA) status expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *FATCASource1Choice) SetCode(value string) {
	f.Code = (*FATCASourceStatus1Code)(&value)
}

func (f *FATCASource1Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
