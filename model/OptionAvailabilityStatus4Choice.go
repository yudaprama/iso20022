package model

// Choice between a standard code or proprietary code to specify the option availability status.
type OptionAvailabilityStatus4Choice struct {

	// Standard code to specify the status of the option availability.
	Code *OptionAvailabilityStatus1Code `xml:"Cd"`

	// Proprietary identification of the status of the option availability.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (o *OptionAvailabilityStatus4Choice) SetCode(value string) {
	o.Code = (*OptionAvailabilityStatus1Code)(&value)
}

func (o *OptionAvailabilityStatus4Choice) AddProprietary() *GenericIdentification47 {
	o.Proprietary = new(GenericIdentification47)
	return o.Proprietary
}
