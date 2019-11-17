package model

// Message to be displayed to the cardholder or the cashier.
type ActionMessage1 struct {

	// Destination of the message to be displayed or printed.
	MessageDestination *UserInterface1Code `xml:"MsgDstn"`

	// Text or graphic data to be display or printed to the cardholder or the cashier.
	MessageContent *Max256Text `xml:"MsgCntt"`

	// Electronic signature of the message to display or print.
	MessageContentSignature *Max70Text `xml:"MsgCnttSgntr,omitempty"`
}

func (a *ActionMessage1) SetMessageDestination(value string) {
	a.MessageDestination = (*UserInterface1Code)(&value)
}

func (a *ActionMessage1) SetMessageContent(value string) {
	a.MessageContent = (*Max256Text)(&value)
}

func (a *ActionMessage1) SetMessageContentSignature(value string) {
	a.MessageContentSignature = (*Max70Text)(&value)
}
