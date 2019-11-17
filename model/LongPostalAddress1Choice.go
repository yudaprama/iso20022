package model

// Information that locates and identifies a specific address, as defined by postal services.
type LongPostalAddress1Choice struct {

	// Information that locates and identifies a specific address, as defined by postal services, that is presented in free format text.
	Unstructured *Max140Text `xml:"Ustrd"`

	// Information that locates and identifies a specific address, as defined by postal services, that is presented in a formal structure.
	Structured *StructuredLongPostalAddress1 `xml:"Strd"`
}

func (l *LongPostalAddress1Choice) SetUnstructured(value string) {
	l.Unstructured = (*Max140Text)(&value)
}

func (l *LongPostalAddress1Choice) AddStructured() *StructuredLongPostalAddress1 {
	l.Structured = new(StructuredLongPostalAddress1)
	return l.Structured
}
