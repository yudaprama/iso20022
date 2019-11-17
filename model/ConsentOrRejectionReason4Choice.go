package model

// Choice of format for the rejection reason.
type ConsentOrRejectionReason4Choice struct {

	// Specifies the reason why the counterparty response has a rejection status.
	Code *CounterpartyResponseStatusReason1Code `xml:"Cd"`

	// Specifies the reason why the counterparty response has a  rejection status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *ConsentOrRejectionReason4Choice) SetCode(value string) {
	c.Code = (*CounterpartyResponseStatusReason1Code)(&value)
}

func (c *ConsentOrRejectionReason4Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
