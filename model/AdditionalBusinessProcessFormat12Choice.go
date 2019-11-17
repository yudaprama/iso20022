package model

// Choice between a standard code or proprietary code to specify the type of the additional business process.
type AdditionalBusinessProcessFormat12Choice struct {

	// Standard code to specify the additional business process linked to a corporate action event.
	Code *AdditionalBusinessProcess5Code `xml:"Cd"`

	// Proprietary identification of the additional business process linked to a corporate action event.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AdditionalBusinessProcessFormat12Choice) SetCode(value string) {
	a.Code = (*AdditionalBusinessProcess5Code)(&value)
}

func (a *AdditionalBusinessProcessFormat12Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
