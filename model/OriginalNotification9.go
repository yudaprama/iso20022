package model

// Identifies the original notification and to provide the status.
type OriginalNotification9 struct {

	// Point to point reference, as assigned by the original sender, to unambiguously identify the original notification to receive message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Identification of the original notification.
	OriginalNotificationIdentification *Max35Text `xml:"OrgnlNtfctnId"`

	// Specifies the status of the notification in a coded form.
	NotificationStatus *NotificationStatus3Code `xml:"NtfctnSts,omitempty"`

	// Further details of the notification status.
	AdditionalStatusInformation *Max140Text `xml:"AddtlStsInf,omitempty"`

	// Identifies the original notification item and provides the status.
	OriginalNotificationReference []*OriginalNotificationReference7 `xml:"OrgnlNtfctnRef,omitempty"`
}

func (o *OriginalNotification9) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification9) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalNotification9) SetOriginalNotificationIdentification(value string) {
	o.OriginalNotificationIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification9) SetNotificationStatus(value string) {
	o.NotificationStatus = (*NotificationStatus3Code)(&value)
}

func (o *OriginalNotification9) SetAdditionalStatusInformation(value string) {
	o.AdditionalStatusInformation = (*Max140Text)(&value)
}

func (o *OriginalNotification9) AddOriginalNotificationReference() *OriginalNotificationReference7 {
	newValue := new(OriginalNotificationReference7)
	o.OriginalNotificationReference = append(o.OriginalNotificationReference, newValue)
	return newValue
}
