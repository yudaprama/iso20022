package model

// Choice of format for the delivery return information.
type DeliveryReturn3Choice struct {

	// Delivery return expressed as an ISO 20022 code.
	Code *DeliveryReturn1Code `xml:"Cd"`

	// Delivery return expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (d *DeliveryReturn3Choice) SetCode(value string) {
	d.Code = (*DeliveryReturn1Code)(&value)
}

func (d *DeliveryReturn3Choice) AddProprietary() *GenericIdentification30 {
	d.Proprietary = new(GenericIdentification30)
	return d.Proprietary
}
