package model

// Choice of formats for the waiving instruction.
type WaivingInstruction2Choice struct {

	// Type of discount or waiving expressed as a code.
	Code *WaivingInstruction1Code `xml:"Cd"`

	// Type of discount or waiving expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (w *WaivingInstruction2Choice) SetCode(value string) {
	w.Code = (*WaivingInstruction1Code)(&value)
}

func (w *WaivingInstruction2Choice) AddProprietary() *GenericIdentification47 {
	w.Proprietary = new(GenericIdentification47)
	return w.Proprietary
}
