package model

// Choice between a standard code or proprietary code to specify the type of dividend.
type DividendTypeFormat5Choice struct {

	// Standard code to specify the frequency of the corporate action event.
	Code *CorporateActionFrequencyType3Code `xml:"Cd"`

	// Proprietary identification of the frequency of the corporate action event.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (d *DividendTypeFormat5Choice) SetCode(value string) {
	d.Code = (*CorporateActionFrequencyType3Code)(&value)
}

func (d *DividendTypeFormat5Choice) AddProprietary() *GenericIdentification20 {
	d.Proprietary = new(GenericIdentification20)
	return d.Proprietary
}
