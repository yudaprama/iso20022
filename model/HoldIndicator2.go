package model

// Specifies whether the transaction is on hold/blocked/frozen.
type HoldIndicator2 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Specifies the reason of the registration staus.
	Reason []*RegistrationReason1 `xml:"Rsn,omitempty"`
}

func (h *HoldIndicator2) SetIndicator(value string) {
	h.Indicator = (*YesNoIndicator)(&value)
}

func (h *HoldIndicator2) AddReason() *RegistrationReason1 {
	newValue := new(RegistrationReason1)
	h.Reason = append(h.Reason, newValue)
	return newValue
}
