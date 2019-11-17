package model

// Choice of formats for the specification of the risk level.
type RiskLevel1Choice struct {

	// Risk level expressed as code.
	Code *RiskLevel1Code `xml:"Cd"`

	// Risk level expressed as a proprietary code.
	Proprietary *GenericIdentification41 `xml:"Prtry"`
}

func (r *RiskLevel1Choice) SetCode(value string) {
	r.Code = (*RiskLevel1Code)(&value)
}

func (r *RiskLevel1Choice) AddProprietary() *GenericIdentification41 {
	r.Proprietary = new(GenericIdentification41)
	return r.Proprietary
}
