package model

// Information about the status of a specific message.
type ValidationStatusInformation1 struct {

	// The result of the technical validation (e.g. Accepted, Reception error) executed on the  request message.
	Status *TechnicalValidationStatus1Code `xml:"Sts"`

	// The reason for the validation status.
	StatusReason *StatusReason4Choice `xml:"StsRsn,omitempty"`

	// Further details on the validation status reason.
	AdditionalStatusReasonInformation []*Max105Text `xml:"AddtlStsRsnInf,omitempty"`
}

func (v *ValidationStatusInformation1) SetStatus(value string) {
	v.Status = (*TechnicalValidationStatus1Code)(&value)
}

func (v *ValidationStatusInformation1) AddStatusReason() *StatusReason4Choice {
	v.StatusReason = new(StatusReason4Choice)
	return v.StatusReason
}

func (v *ValidationStatusInformation1) AddAdditionalStatusReasonInformation(value string) {
	v.AdditionalStatusReasonInformation = append(v.AdditionalStatusReasonInformation, (*Max105Text)(&value))
}
