package model

// Choice between a standard code or proprietary code to specify the type of dividend.
type DividendTypeFormat10Choice struct {

	// Standard code to specify the frequency of the corporate action event.
	Code *CorporateActionFrequencyType5Code `xml:"Cd"`

	// Proprietary identification of the frequency of the corporate action event.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (d *DividendTypeFormat10Choice) SetCode(value string) {
	d.Code = (*CorporateActionFrequencyType5Code)(&value)
}

func (d *DividendTypeFormat10Choice) AddProprietary() *GenericIdentification47 {
	d.Proprietary = new(GenericIdentification47)
	return d.Proprietary
}
