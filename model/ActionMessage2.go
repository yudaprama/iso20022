package model

// Information to display, print or store.
type ActionMessage2 struct {

	// Destination of the message.
	MessageDestination *UserInterface4Code `xml:"MsgDstn"`

	// Message format.
	Format *OutputFormat1Code `xml:"Frmt,omitempty"`

	// Content or reference of the message.
	MessageContent *Max20000Text `xml:"MsgCntt"`

	// Digital signature of the message.
	MessageContentSignature *Max140Binary `xml:"MsgCnttSgntr,omitempty"`
}

func (a *ActionMessage2) SetMessageDestination(value string) {
	a.MessageDestination = (*UserInterface4Code)(&value)
}

func (a *ActionMessage2) SetFormat(value string) {
	a.Format = (*OutputFormat1Code)(&value)
}

func (a *ActionMessage2) SetMessageContent(value string) {
	a.MessageContent = (*Max20000Text)(&value)
}

func (a *ActionMessage2) SetMessageContentSignature(value string) {
	a.MessageContentSignature = (*Max140Binary)(&value)
}
