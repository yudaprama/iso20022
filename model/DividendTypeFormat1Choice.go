package model

// Choice between a standard code or proprietary code to specify the type of dividend.
type DividendTypeFormat1Choice struct {

	// Standard code to specify the frequency of the corporate action event.
	Code *CorporateActionFrequencyType1Code `xml:"Cd"`

	// Proprietary identification of the frequency of the corporate action event.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (d *DividendTypeFormat1Choice) SetCode(value string) {
	d.Code = (*CorporateActionFrequencyType1Code)(&value)
}

func (d *DividendTypeFormat1Choice) AddProprietary() *GenericIdentification20 {
	d.Proprietary = new(GenericIdentification20)
	return d.Proprietary
}
