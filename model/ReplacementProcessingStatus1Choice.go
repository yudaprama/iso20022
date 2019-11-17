package model

// Choice of format for the replacement processing status
type ReplacementProcessingStatus1Choice struct {

	// Provides the processing status of the replacement request.
	Code *ReplacementProcessingStatus1Code `xml:"Cd"`

	// Provides the processing status of the replacement request.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *ReplacementProcessingStatus1Choice) SetCode(value string) {
	r.Code = (*ReplacementProcessingStatus1Code)(&value)
}

func (r *ReplacementProcessingStatus1Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
