package model

// Specification of the waiving type.
type WaivingType1 struct {

	// Structured format.
	Structured *WaivingInstruction2Code `xml:"Strd"`

	// Additional information about the type of waiving.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (w *WaivingType1) SetStructured(value string) {
	w.Structured = (*WaivingInstruction2Code)(&value)
}

func (w *WaivingType1) SetAdditionalInformation(value string) {
	w.AdditionalInformation = (*Max350Text)(&value)
}
