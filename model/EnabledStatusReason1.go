package model

// Reason for an enabled status.
type EnabledStatusReason1 struct {

	// Reason for the enabled account status.
	Code *EnabledStatusReason2Choice `xml:"Cd"`

	// Additional information about the reason for the enabled account status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (e *EnabledStatusReason1) AddCode() *EnabledStatusReason2Choice {
	e.Code = new(EnabledStatusReason2Choice)
	return e.Code
}

func (e *EnabledStatusReason1) SetAdditionalInformation(value string) {
	e.AdditionalInformation = (*Max350Text)(&value)
}
