package model

// Choice of format for the registration processing status
type RegistrationProcessingStatus4Choice struct {

	// Provides the status of the registration processing.
	Code *RegistrationProcessingStatus1Code `xml:"Cd"`

	// Provides the status of the registration processing.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RegistrationProcessingStatus4Choice) SetCode(value string) {
	r.Code = (*RegistrationProcessingStatus1Code)(&value)
}

func (r *RegistrationProcessingStatus4Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
