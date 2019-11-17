package model

// Specifies whether the transaction is on hold/blocked/frozen.
type HoldIndicator4 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Specifies the reason of the registration staus.
	Reason []*RegistrationReason3 `xml:"Rsn,omitempty"`
}

func (h *HoldIndicator4) SetIndicator(value string) {
	h.Indicator = (*YesNoIndicator)(&value)
}

func (h *HoldIndicator4) AddReason() *RegistrationReason3 {
	newValue := new(RegistrationReason3)
	h.Reason = append(h.Reason, newValue)
	return newValue
}
