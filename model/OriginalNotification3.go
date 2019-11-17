package model

// Identify the original notification and to provide the status.
type OriginalNotification3 struct {

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
	OriginalNotificationReference []*OriginalNotificationReference2 `xml:"OrgnlNtfctnRef,omitempty"`
}

func (o *OriginalNotification3) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification3) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalNotification3) SetOriginalNotificationIdentification(value string) {
	o.OriginalNotificationIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification3) SetNotificationStatus(value string) {
	o.NotificationStatus = (*NotificationStatus3Code)(&value)
}

func (o *OriginalNotification3) SetAdditionalStatusInformation(value string) {
	o.AdditionalStatusInformation = (*Max140Text)(&value)
}

func (o *OriginalNotification3) AddOriginalNotificationReference() *OriginalNotificationReference2 {
	newValue := new(OriginalNotificationReference2)
	o.OriginalNotificationReference = append(o.OriginalNotificationReference, newValue)
	return newValue
}
