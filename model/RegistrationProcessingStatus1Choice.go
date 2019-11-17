package model

// Choice of format for the registration processing status
type RegistrationProcessingStatus1Choice struct {

	// Provides the status of the registration processing.
	Code *RegistrationProcessingStatus1Code `xml:"Cd"`

	// Provides the status of the registration processing.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RegistrationProcessingStatus1Choice) SetCode(value string) {
	r.Code = (*RegistrationProcessingStatus1Code)(&value)
}

func (r *RegistrationProcessingStatus1Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
