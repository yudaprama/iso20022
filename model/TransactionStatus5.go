package model

// Identifies the future status of the transaction by means of a code and a period.
type TransactionStatus5 struct {

	// Identifies the status of the transaction by means of a code.
	Status *BaselineStatus3Code `xml:"Sts"`

	// Date and time at which the current status will change.
	ChangeDateTime *ISODateTime `xml:"ChngDtTm"`

	// Additional information on the reason for the time-out.
	Description *Max140Text `xml:"Desc,omitempty"`
}

func (t *TransactionStatus5) SetStatus(value string) {
	t.Status = (*BaselineStatus3Code)(&value)
}

func (t *TransactionStatus5) SetChangeDateTime(value string) {
	t.ChangeDateTime = (*ISODateTime)(&value)
}

func (t *TransactionStatus5) SetDescription(value string) {
	t.Description = (*Max140Text)(&value)
}
