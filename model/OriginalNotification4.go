package model

// Identify the original notification, to which the cancellation advice refers.
type OriginalNotification4 struct {

	// Point to point reference, as assigned by the original sender, to unambiguously identify the original notification to receive message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Identification of the original notification.
	OriginalNotificationIdentification *Max35Text `xml:"OrgnlNtfctnId"`

	// Indicates whether the cancellation applies to the complete original notification or to individual items within the original notification.
	NotificationCancellation *GroupCancellationIndicator `xml:"NtfctnCxl,omitempty"`

	// Identifies the original notification item, to which the cancellation advice refers.
	OriginalNotificationReference []*OriginalNotificationReference1 `xml:"OrgnlNtfctnRef,omitempty"`
}

func (o *OriginalNotification4) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification4) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalNotification4) SetOriginalNotificationIdentification(value string) {
	o.OriginalNotificationIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification4) SetNotificationCancellation(value string) {
	o.NotificationCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalNotification4) AddOriginalNotificationReference() *OriginalNotificationReference1 {
	newValue := new(OriginalNotificationReference1)
	o.OriginalNotificationReference = append(o.OriginalNotificationReference, newValue)
	return newValue
}
