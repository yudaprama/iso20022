package model

// Message to be displayed to the cardholder or the cashier.
type ActionMessage5 struct {

	// Format of the content.
	Format *OutputFormat1Code `xml:"Frmt,omitempty"`

	// Text or graphic data to be display or printed to the cardholder or the cashier.
	MessageContent *Max20000Text `xml:"MsgCntt"`
}

func (a *ActionMessage5) SetFormat(value string) {
	a.Format = (*OutputFormat1Code)(&value)
}

func (a *ActionMessage5) SetMessageContent(value string) {
	a.MessageContent = (*Max20000Text)(&value)
}
