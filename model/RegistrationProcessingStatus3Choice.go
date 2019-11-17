package model

// Choice of format for the registration processing status
type RegistrationProcessingStatus3Choice struct {

	// Provides the status of the registration processing.
	Code *RegistrationProcessingStatus1Code `xml:"Cd"`

	// Provides the status of the registration processing.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RegistrationProcessingStatus3Choice) SetCode(value string) {
	r.Code = (*RegistrationProcessingStatus1Code)(&value)
}

func (r *RegistrationProcessingStatus3Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
