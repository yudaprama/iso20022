package model

// Choice between a standard code or proprietary code to specify the type of the additional business process, that is, a tax refund.
type AdditionalBusinessProcessFormat13Choice struct {

	// Standard code to specify the additional business process "tax refund" linked to a corporate action event.
	Code *AdditionalBusinessProcess6Code `xml:"Cd"`

	// Proprietary identification of the additional business process "tax refund" linked to a corporate action event.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AdditionalBusinessProcessFormat13Choice) SetCode(value string) {
	a.Code = (*AdditionalBusinessProcess6Code)(&value)
}

func (a *AdditionalBusinessProcessFormat13Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
