package model

// Provide information on the status reason of the record.
type ValidationStatusReason1 struct {

	// Party that issues the status.
	Originator *PartyIdentification77 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the status.
	Reason *StatusReason6Choice `xml:"Rsn,omitempty"`

	// Provides details about the rule which could not be validated.
	ValidationRule []*GenericValidationRuleIdentification1 `xml:"VldtnRule,omitempty"`

	// Further details on the status reason.
	//
	// Usage: Additional information can be used for several purposes such as the reporting of repaired information.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (v *ValidationStatusReason1) AddOriginator() *PartyIdentification77 {
	v.Originator = new(PartyIdentification77)
	return v.Originator
}

func (v *ValidationStatusReason1) AddReason() *StatusReason6Choice {
	v.Reason = new(StatusReason6Choice)
	return v.Reason
}

func (v *ValidationStatusReason1) AddValidationRule() *GenericValidationRuleIdentification1 {
	newValue := new(GenericValidationRuleIdentification1)
	v.ValidationRule = append(v.ValidationRule, newValue)
	return newValue
}

func (v *ValidationStatusReason1) AddAdditionalInformation(value string) {
	v.AdditionalInformation = append(v.AdditionalInformation, (*Max105Text)(&value))
}
