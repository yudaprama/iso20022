package model

// Information about the cancellation of a notification advice or the withdrawal of a CA event.
type NotificationCancellation1 struct {

	// The function of the notification e.g. new notification.
	NotificationCancellationType *CorporateActionNotificationType2Code `xml:"NtfctnCxlTp"`

	// The identification of the linked notification advice.
	LinkedAgentCANotificationAdviceIdentification *DocumentIdentification8 `xml:"LkdAgtCANtfctnAdvcId"`
}

func (n *NotificationCancellation1) SetNotificationCancellationType(value string) {
	n.NotificationCancellationType = (*CorporateActionNotificationType2Code)(&value)
}

func (n *NotificationCancellation1) AddLinkedAgentCANotificationAdviceIdentification() *DocumentIdentification8 {
	n.LinkedAgentCANotificationAdviceIdentification = new(DocumentIdentification8)
	return n.LinkedAgentCANotificationAdviceIdentification
}
