package model

// Choice between a standard code or proprietary code to specify the type of the additional business process, that is, a tax refund.
type AdditionalBusinessProcessFormat14Choice struct {

	// Standard code to specify the additional business process "tax refund" linked to a corporate action event.
	Code *AdditionalBusinessProcess7Code `xml:"Cd"`

	// Proprietary identification of the additional business process "tax refund" linked to a corporate action event.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AdditionalBusinessProcessFormat14Choice) SetCode(value string) {
	a.Code = (*AdditionalBusinessProcess7Code)(&value)
}

func (a *AdditionalBusinessProcessFormat14Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
