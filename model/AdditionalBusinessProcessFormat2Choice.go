package model

// Choice between a standard code or proprietary code to specify the type of the additional business process, that is, a tax refund.
type AdditionalBusinessProcessFormat2Choice struct {

	// Standard code to specify the additional business process "tax refund" linked to a corporate action event.
	Code *AdditionalBusinessProcess2Code `xml:"Cd"`

	// Proprietary identification of the additional business process "tax refund" linked to a corporate action event.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (a *AdditionalBusinessProcessFormat2Choice) SetCode(value string) {
	a.Code = (*AdditionalBusinessProcess2Code)(&value)
}

func (a *AdditionalBusinessProcessFormat2Choice) AddProprietary() *GenericIdentification20 {
	a.Proprietary = new(GenericIdentification20)
	return a.Proprietary
}
