package model

// Information that locates and identifies a specific address, as defined by postal services either in a structured or unstructured way.
type LongPostalAddress2Choice struct {

	// Information that locates and identifies a specific address, as defined by postal services, that is presented in free format text.
	Unstructured *Max350Text `xml:"Ustrd"`

	// Information that locates and identifies a specific address, as defined by postal services, that is presented in a formal structure.
	Structured *PostalAddress1 `xml:"Strd"`
}

func (l *LongPostalAddress2Choice) SetUnstructured(value string) {
	l.Unstructured = (*Max350Text)(&value)
}

func (l *LongPostalAddress2Choice) AddStructured() *PostalAddress1 {
	l.Structured = new(PostalAddress1)
	return l.Structured
}
