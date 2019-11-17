package model

// Choice between a standard code or proprietary code to specify the type of the additional business process, that is, a tax refund.
type AdditionalBusinessProcessFormat10Choice struct {

	// Standard code to specify the additional business process "tax refund" linked to a corporate action event.
	Code *AdditionalBusinessProcess6Code `xml:"Cd"`

	// Proprietary identification of the additional business process "tax refund" linked to a corporate action event.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (a *AdditionalBusinessProcessFormat10Choice) SetCode(value string) {
	a.Code = (*AdditionalBusinessProcess6Code)(&value)
}

func (a *AdditionalBusinessProcessFormat10Choice) AddProprietary() *GenericIdentification30 {
	a.Proprietary = new(GenericIdentification30)
	return a.Proprietary
}
