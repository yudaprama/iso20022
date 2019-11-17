package model

// Provides information about the identification and the creation date of a notification.
type NotificationIdentification1 struct {

	// Unique identifier of the last notification document (message) assigned by the sender of the document.
	Identification *Max35Text `xml:"Id"`

	// Date and time at which the last notification document (message) was created by the sender.
	CreationDateTime *DateAndDateTimeChoice `xml:"CreDtTm,omitempty"`
}

func (n *NotificationIdentification1) SetIdentification(value string) {
	n.Identification = (*Max35Text)(&value)
}

func (n *NotificationIdentification1) AddCreationDateTime() *DateAndDateTimeChoice {
	n.CreationDateTime = new(DateAndDateTimeChoice)
	return n.CreationDateTime
}
