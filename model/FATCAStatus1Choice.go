package model

// Choice of formats for the Foreign Account Tax Compliance Act (FATCA) status.
type FATCAStatus1Choice struct {

	// Foreign Account Tax Compliance Act (FATCA) status expressed as a code.
	Code *FATCAStatus1Code `xml:"Cd"`

	// Foreign Account Tax Compliance Act (FATCA) status expressed as a proprietary code.
	Proprietary *GenericIdentification29 `xml:"Prtry"`
}

func (f *FATCAStatus1Choice) SetCode(value string) {
	f.Code = (*FATCAStatus1Code)(&value)
}

func (f *FATCAStatus1Choice) AddProprietary() *GenericIdentification29 {
	f.Proprietary = new(GenericIdentification29)
	return f.Proprietary
}
