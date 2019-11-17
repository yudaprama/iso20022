package model

// Choice between the specification of the data in a structured or unstructured form.
type DataFormat2Choice struct {

	// Specification of data in a structured form.
	Structured *GenericIdentification1 `xml:"Strd"`

	// Specification of data for which there isn't a structured form.
	Unstructured *Max140Text `xml:"Ustrd"`
}

func (d *DataFormat2Choice) AddStructured() *GenericIdentification1 {
	d.Structured = new(GenericIdentification1)
	return d.Structured
}

func (d *DataFormat2Choice) SetUnstructured(value string) {
	d.Unstructured = (*Max140Text)(&value)
}
