package model

// Choice between a standard code or proprietary code to specify the type of distribution.
type DistributionTypeFormat8Choice struct {

	// Standard code to specify whether the proceeds of the event will be distributed on a rolling basis rather than on a specific date.
	Code *DistributionType3Code `xml:"Cd"`

	// Proprietary identification of the type of distribution.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (d *DistributionTypeFormat8Choice) SetCode(value string) {
	d.Code = (*DistributionType3Code)(&value)
}

func (d *DistributionTypeFormat8Choice) AddProprietary() *GenericIdentification47 {
	d.Proprietary = new(GenericIdentification47)
	return d.Proprietary
}
