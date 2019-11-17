package model

// Choice between a standard code or proprietary code to specify the type of dividend.
type DividendTypeFormat3Choice struct {

	// Standard code to specify the frequency of the corporate action event.
	Code *CorporateActionFrequencyType2Code `xml:"Cd"`

	// Proprietary identification of the frequency of the corporate action event.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (d *DividendTypeFormat3Choice) SetCode(value string) {
	d.Code = (*CorporateActionFrequencyType2Code)(&value)
}

func (d *DividendTypeFormat3Choice) AddProprietary() *GenericIdentification20 {
	d.Proprietary = new(GenericIdentification20)
	return d.Proprietary
}
