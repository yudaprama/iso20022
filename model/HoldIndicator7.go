package model

// Specifies whether the transaction is on hold/blocked/frozen.
type HoldIndicator7 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Specifies the reason of the registration staus.
	Reason []*RegistrationReason6 `xml:"Rsn,omitempty"`
}

func (h *HoldIndicator7) SetIndicator(value string) {
	h.Indicator = (*YesNoIndicator)(&value)
}

func (h *HoldIndicator7) AddReason() *RegistrationReason6 {
	newValue := new(RegistrationReason6)
	h.Reason = append(h.Reason, newValue)
	return newValue
}
