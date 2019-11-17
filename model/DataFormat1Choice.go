package model

// Choice between specification of the data in structured or free text format.
type DataFormat1Choice struct {

	// Specification of data in structured form.
	Structured *GenericIdentification2 `xml:"Strd"`

	// Specification of data in free text form.
	Unstructured *Max140Text `xml:"Ustrd"`
}

func (d *DataFormat1Choice) AddStructured() *GenericIdentification2 {
	d.Structured = new(GenericIdentification2)
	return d.Structured
}

func (d *DataFormat1Choice) SetUnstructured(value string) {
	d.Unstructured = (*Max140Text)(&value)
}
