package model

// Information to log.
type ActionMessage3 struct {

	// Destination of the information.
	Destination *UserInterface3Code `xml:"Dstn"`

	// Format of the content.
	Format *OutputFormat1Code `xml:"Frmt,omitempty"`

	// Content of the information.
	Content *Max20000Text `xml:"Cntt"`
}

func (a *ActionMessage3) SetDestination(value string) {
	a.Destination = (*UserInterface3Code)(&value)
}

func (a *ActionMessage3) SetFormat(value string) {
	a.Format = (*OutputFormat1Code)(&value)
}

func (a *ActionMessage3) SetContent(value string) {
	a.Content = (*Max20000Text)(&value)
}
