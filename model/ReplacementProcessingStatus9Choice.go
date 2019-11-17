package model

// Choice of format for the replacement processing status
type ReplacementProcessingStatus9Choice struct {

	// Provides the processing status of the replacement request.
	Code *ReplacementProcessingStatus1Code `xml:"Cd"`

	// Provides the processing status of the replacement request.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *ReplacementProcessingStatus9Choice) SetCode(value string) {
	r.Code = (*ReplacementProcessingStatus1Code)(&value)
}

func (r *ReplacementProcessingStatus9Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
