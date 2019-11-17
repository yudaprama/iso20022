package model

// Choice between a standard code or proprietary code to specify the type of the additional business process.
type AdditionalBusinessProcessFormat9Choice struct {

	// Standard code to specify the additional business process linked to a corporate action event.
	Code *AdditionalBusinessProcess5Code `xml:"Cd"`

	// Proprietary identification of the additional business process linked to a corporate action event.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (a *AdditionalBusinessProcessFormat9Choice) SetCode(value string) {
	a.Code = (*AdditionalBusinessProcess5Code)(&value)
}

func (a *AdditionalBusinessProcessFormat9Choice) AddProprietary() *GenericIdentification30 {
	a.Proprietary = new(GenericIdentification30)
	return a.Proprietary
}
