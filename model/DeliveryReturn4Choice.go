package model

// Choice of format for the delivery return information.
type DeliveryReturn4Choice struct {

	// Delivery return expressed as an ISO 20022 code.
	Code *DeliveryReturn1Code `xml:"Cd"`

	// Delivery return expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (d *DeliveryReturn4Choice) SetCode(value string) {
	d.Code = (*DeliveryReturn1Code)(&value)
}

func (d *DeliveryReturn4Choice) AddProprietary() *GenericIdentification47 {
	d.Proprietary = new(GenericIdentification47)
	return d.Proprietary
}
