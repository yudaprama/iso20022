package model

// Choice of formats for the waiving instruction.
type WaivingInstruction1Choice struct {

	// Type of waiving instruction expressed as a code.
	Code *WaivingInstruction1Code `xml:"Cd"`

	// Type of waiving instruction expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (w *WaivingInstruction1Choice) SetCode(value string) {
	w.Code = (*WaivingInstruction1Code)(&value)
}

func (w *WaivingInstruction1Choice) AddProprietary() *GenericIdentification47 {
	w.Proprietary = new(GenericIdentification47)
	return w.Proprietary
}
