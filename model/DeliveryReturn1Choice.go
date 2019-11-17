package model

// Choice of format for the delivery return information.
type DeliveryReturn1Choice struct {

	// Delivery return expressed as an ISO 20022 code.
	Code *DeliveryReturn1Code `xml:"Cd"`

	// Delivery return expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (d *DeliveryReturn1Choice) SetCode(value string) {
	d.Code = (*DeliveryReturn1Code)(&value)
}

func (d *DeliveryReturn1Choice) AddProprietary() *GenericIdentification20 {
	d.Proprietary = new(GenericIdentification20)
	return d.Proprietary
}
