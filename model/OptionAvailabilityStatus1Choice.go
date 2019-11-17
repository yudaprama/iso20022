package model

// Choice between a standard code or proprietary code to specify the option availability status.
type OptionAvailabilityStatus1Choice struct {

	// Standard code to specify the status of the option availability.
	Code *OptionAvailabilityStatus1Code `xml:"Cd"`

	// Proprietary identification of the status of the option availability.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (o *OptionAvailabilityStatus1Choice) SetCode(value string) {
	o.Code = (*OptionAvailabilityStatus1Code)(&value)
}

func (o *OptionAvailabilityStatus1Choice) AddProprietary() *GenericIdentification20 {
	o.Proprietary = new(GenericIdentification20)
	return o.Proprietary
}
