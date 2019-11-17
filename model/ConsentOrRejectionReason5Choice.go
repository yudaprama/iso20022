package model

// Choice of format for the rejection reason.
type ConsentOrRejectionReason5Choice struct {

	// Specifies the reason why the counterparty response has a rejection status.
	Code *CounterpartyResponseStatusReason1Code `xml:"Cd"`

	// Specifies the reason why the counterparty response has a  rejection status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *ConsentOrRejectionReason5Choice) SetCode(value string) {
	c.Code = (*CounterpartyResponseStatusReason1Code)(&value)
}

func (c *ConsentOrRejectionReason5Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
