package model

// Choice of formats for the type of Foreign Account Tax Compliance Act (FATCA) form.
type FATCAForm1Choice struct {

	// Type of Foreign Account Tax Compliance Act (FATCA) form expressed as a code.
	Code *FATCAFormType1Code `xml:"Cd"`

	// Type of Foreign Account Tax Compliance Act (FATCA) form expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *FATCAForm1Choice) SetCode(value string) {
	f.Code = (*FATCAFormType1Code)(&value)
}

func (f *FATCAForm1Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
