package model

// Choice of format for the rejection reason.
type ConsentOrRejectionReason2Choice struct {

	// Specifies the reason why the counterparty response has a rejection status.
	Code *CounterpartyResponseStatusReason1Code `xml:"Cd"`

	// Specifies the reason why the counterparty response has a  rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *ConsentOrRejectionReason2Choice) SetCode(value string) {
	c.Code = (*CounterpartyResponseStatusReason1Code)(&value)
}

func (c *ConsentOrRejectionReason2Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
