package model

// Specifies whether the transaction is on hold/blocked/frozen.
type HoldIndicator6 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Specifies the reason of the registration staus.
	Reason []*RegistrationReason5 `xml:"Rsn,omitempty"`
}

func (h *HoldIndicator6) SetIndicator(value string) {
	h.Indicator = (*YesNoIndicator)(&value)
}

func (h *HoldIndicator6) AddReason() *RegistrationReason5 {
	newValue := new(RegistrationReason5)
	h.Reason = append(h.Reason, newValue)
	return newValue
}
