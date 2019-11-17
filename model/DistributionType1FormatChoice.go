package model

// Choice of formats to  express the distribution type.
type DistributionType1FormatChoice struct {

	// Standard code to  specify the distribution type.
	Code *DistributionType1Code `xml:"Cd"`

	// Proprietary code to  express the distribution type.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (d *DistributionType1FormatChoice) SetCode(value string) {
	d.Code = (*DistributionType1Code)(&value)
}

func (d *DistributionType1FormatChoice) AddProprietary() *GenericIdentification13 {
	d.Proprietary = new(GenericIdentification13)
	return d.Proprietary
}
