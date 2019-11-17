package model

// Identifies the original notification, to which the cancellation advice refers.
type OriginalNotification6 struct {

	// Point to point reference, as assigned by the original sender, to unambiguously identify the original notification to receive message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Identification of the original notification.
	OriginalNotificationIdentification *Max35Text `xml:"OrgnlNtfctnId"`

	// Indicates whether the cancellation applies to the complete original notification or to individual items within the original notification.
	NotificationCancellation *GroupCancellationIndicator `xml:"NtfctnCxl,omitempty"`

	// Identifies the original notification item, to which the cancellation advice refers.
	OriginalNotificationReference []*OriginalNotificationReference4 `xml:"OrgnlNtfctnRef,omitempty"`
}

func (o *OriginalNotification6) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification6) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalNotification6) SetOriginalNotificationIdentification(value string) {
	o.OriginalNotificationIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification6) SetNotificationCancellation(value string) {
	o.NotificationCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalNotification6) AddOriginalNotificationReference() *OriginalNotificationReference4 {
	newValue := new(OriginalNotificationReference4)
	o.OriginalNotificationReference = append(o.OriginalNotificationReference, newValue)
	return newValue
}
