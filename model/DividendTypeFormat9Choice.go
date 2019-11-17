package model

// Choice between a standard code or proprietary code to specify the type of dividend.
type DividendTypeFormat9Choice struct {

	// Standard code to specify the frequency of the corporate action event.
	Code *CorporateActionFrequencyType5Code `xml:"Cd"`

	// Proprietary identification of the frequency of the corporate action event.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (d *DividendTypeFormat9Choice) SetCode(value string) {
	d.Code = (*CorporateActionFrequencyType5Code)(&value)
}

func (d *DividendTypeFormat9Choice) AddProprietary() *GenericIdentification30 {
	d.Proprietary = new(GenericIdentification30)
	return d.Proprietary
}
