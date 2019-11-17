package model

// Structured information to be communicated to other parties in the transaction.
type Notification1 struct {

	// Type of the notification.
	Type *NotificationType1Code `xml:"Tp"`

	// Additional and important information to qualify and describe the notification.
	AdditionalInformation *Max140Text `xml:"AddtlInf"`
}

func (n *Notification1) SetType(value string) {
	n.Type = (*NotificationType1Code)(&value)
}

func (n *Notification1) SetAdditionalInformation(value string) {
	n.AdditionalInformation = (*Max140Text)(&value)
}
