package model

// Choice between a standard code or proprietary code to specify the type of the additional business process, that is, a tax refund.
type AdditionalBusinessProcessFormat11Choice struct {

	// Standard code to specify the additional business process "tax refund" linked to a corporate action event.
	Code *AdditionalBusinessProcess7Code `xml:"Cd"`

	// Proprietary identification of the additional business process "tax refund" linked to a corporate action event.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (a *AdditionalBusinessProcessFormat11Choice) SetCode(value string) {
	a.Code = (*AdditionalBusinessProcess7Code)(&value)
}

func (a *AdditionalBusinessProcessFormat11Choice) AddProprietary() *GenericIdentification30 {
	a.Proprietary = new(GenericIdentification30)
	return a.Proprietary
}
