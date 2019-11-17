package model

// Specifies if the occurrence of the event contained in the notification is confirmed or unconfirmed. Details of the event can be complete or incomplete.
type NotificationStatus1 struct {

	// Status to define if the occurrence of the event contained in the notification is confirmed or unconfirmed.
	Status *NotificationStatus2Code `xml:"Sts"`
}

func (n *NotificationStatus1) SetStatus(value string) {
	n.Status = (*NotificationStatus2Code)(&value)
}
