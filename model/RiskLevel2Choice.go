package model

// Choice of formats for the specification of the risk level.
type RiskLevel2Choice struct {

	// Risk level expressed as code.
	Code *RiskLevel1Code `xml:"Cd"`

	// Risk level expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RiskLevel2Choice) SetCode(value string) {
	r.Code = (*RiskLevel1Code)(&value)
}

func (r *RiskLevel2Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
