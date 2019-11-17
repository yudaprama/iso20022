package model

// Choice of format for the replacement processing status
type ReplacementProcessingStatus8Choice struct {

	// Provides the processing status of the replacement request.
	Code *ReplacementProcessingStatus1Code `xml:"Cd"`

	// Provides the processing status of the replacement request.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *ReplacementProcessingStatus8Choice) SetCode(value string) {
	r.Code = (*ReplacementProcessingStatus1Code)(&value)
}

func (r *ReplacementProcessingStatus8Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
