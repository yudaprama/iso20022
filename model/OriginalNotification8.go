package model

// Identifies the original notification, to which the cancellation advice refers.
type OriginalNotification8 struct {

	// Point to point reference, as assigned by the original sender, to unambiguously identify the original notification to receive message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Identification of the original notification.
	OriginalNotificationIdentification *Max35Text `xml:"OrgnlNtfctnId"`

	// Indicates whether the cancellation applies to the complete original notification or to individual items within the original notification.
	NotificationCancellation *GroupCancellationIndicator `xml:"NtfctnCxl,omitempty"`

	// Identifies the original notification item, to which the cancellation advice refers.
	OriginalNotificationReference []*OriginalNotificationReference6 `xml:"OrgnlNtfctnRef,omitempty"`
}

func (o *OriginalNotification8) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification8) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalNotification8) SetOriginalNotificationIdentification(value string) {
	o.OriginalNotificationIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification8) SetNotificationCancellation(value string) {
	o.NotificationCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalNotification8) AddOriginalNotificationReference() *OriginalNotificationReference6 {
	newValue := new(OriginalNotificationReference6)
	o.OriginalNotificationReference = append(o.OriginalNotificationReference, newValue)
	return newValue
}
