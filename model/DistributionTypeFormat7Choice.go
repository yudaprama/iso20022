package model

// Choice between a standard code or proprietary code to specify the type of distribution.
type DistributionTypeFormat7Choice struct {

	// Standard code to specify whether the proceeds of the event will be distributed on a rolling basis rather than on a specific date.
	Code *DistributionType3Code `xml:"Cd"`

	// Proprietary identification of the type of distribution.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (d *DistributionTypeFormat7Choice) SetCode(value string) {
	d.Code = (*DistributionType3Code)(&value)
}

func (d *DistributionTypeFormat7Choice) AddProprietary() *GenericIdentification30 {
	d.Proprietary = new(GenericIdentification30)
	return d.Proprietary
}
