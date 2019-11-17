package model

// Choice of formats for the Foreign Account Tax Compliance Act (FATCA) status.
type FATCAStatus2Choice struct {

	// Foreign Account Tax Compliance Act (FATCA) status expressed as a code.
	Code *FATCAStatus1Code `xml:"Cd"`

	// Foreign Account Tax Compliance Act (FATCA) status expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *FATCAStatus2Choice) SetCode(value string) {
	f.Code = (*FATCAStatus1Code)(&value)
}

func (f *FATCAStatus2Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
