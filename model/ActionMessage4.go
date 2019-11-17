package model

// Information to display, print or log.
type ActionMessage4 struct {

	// Information format.
	Format *OutputFormat2Code `xml:"Frmt,omitempty"`

	// Content of the message.
	Message *Max20000Text `xml:"Msg,omitempty"`

	// Message content if this is a message reference or screen reference.
	Reference *Max35Text `xml:"Ref,omitempty"`

	// Device to be used.
	Device *ATMDevice1Code `xml:"Dvc,omitempty"`

	// Electronic signature of the message to display or print.
	MessageContentSignature *Max35Binary `xml:"MsgCnttSgntr,omitempty"`
}

func (a *ActionMessage4) SetFormat(value string) {
	a.Format = (*OutputFormat2Code)(&value)
}

func (a *ActionMessage4) SetMessage(value string) {
	a.Message = (*Max20000Text)(&value)
}

func (a *ActionMessage4) SetReference(value string) {
	a.Reference = (*Max35Text)(&value)
}

func (a *ActionMessage4) SetDevice(value string) {
	a.Device = (*ATMDevice1Code)(&value)
}

func (a *ActionMessage4) SetMessageContentSignature(value string) {
	a.MessageContentSignature = (*Max35Binary)(&value)
}
