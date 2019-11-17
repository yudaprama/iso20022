package model

// Parameters of a physical delivery.
type DeliveryParameters3 struct {

	// Address for physical delivery.
	Address *NameAndAddress4 `xml:"Adr"`

	// Certificate representing a security that is delivered.
	IssuedCertificateNumber *Max35Text `xml:"IssdCertNb,omitempty"`
}

func (d *DeliveryParameters3) AddAddress() *NameAndAddress4 {
	d.Address = new(NameAndAddress4)
	return d.Address
}

func (d *DeliveryParameters3) SetIssuedCertificateNumber(value string) {
	d.IssuedCertificateNumber = (*Max35Text)(&value)
}
