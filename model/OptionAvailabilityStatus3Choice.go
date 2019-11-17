package model

// Choice between a standard code or proprietary code to specify the option availability status.
type OptionAvailabilityStatus3Choice struct {

	// Standard code to specify the status of the option availability.
	Code *OptionAvailabilityStatus1Code `xml:"Cd"`

	// Proprietary identification of the status of the option availability.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (o *OptionAvailabilityStatus3Choice) SetCode(value string) {
	o.Code = (*OptionAvailabilityStatus1Code)(&value)
}

func (o *OptionAvailabilityStatus3Choice) AddProprietary() *GenericIdentification30 {
	o.Proprietary = new(GenericIdentification30)
	return o.Proprietary
}
