package model

// Choice between a standard code or proprietary code to specify the type of distribution.
type DistributionTypeFormat3Choice struct {

	// Standard code to specify whether the proceeds of the event will be distributed on a rolling basis rather than on a specific date.
	Code *DistributionType2Code `xml:"Cd"`

	// Proprietary identification of the type of distribution.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (d *DistributionTypeFormat3Choice) SetCode(value string) {
	d.Code = (*DistributionType2Code)(&value)
}

func (d *DistributionTypeFormat3Choice) AddProprietary() *GenericIdentification20 {
	d.Proprietary = new(GenericIdentification20)
	return d.Proprietary
}
