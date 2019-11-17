package model

// Choice between a standard code or proprietary code to specify the type of the additional business process, that is, a tax refund.
type AdditionalBusinessProcessFormat7Choice struct {

	// Standard code to specify the additional business process "tax refund" linked to a corporate action event.
	Code *AdditionalBusinessProcess4Code `xml:"Cd"`

	// Proprietary identification of the additional business process "tax refund" linked to a corporate action event.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (a *AdditionalBusinessProcessFormat7Choice) SetCode(value string) {
	a.Code = (*AdditionalBusinessProcess4Code)(&value)
}

func (a *AdditionalBusinessProcessFormat7Choice) AddProprietary() *GenericIdentification20 {
	a.Proprietary = new(GenericIdentification20)
	return a.Proprietary
}
