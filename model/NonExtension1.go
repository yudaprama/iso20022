package model

// Non-extension information.
type NonExtension1 struct {

	// Minimum number of days prior to the then current expiry date by which notice of non-extension must be sent.
	NotificationPeriod *Number `xml:"NtfctnPrd,omitempty"`

	// Method by which the notice of non-extension is intended to be delivered.
	NotificationMethod *CommunicationMethod1Choice `xml:"NtfctnMtd,omitempty"`

	// Type of party to whom the notice of non-extension is intended to be delivered.
	NotificationRecipientType *PartyType1Choice `xml:"NtfctnRcptTp,omitempty"`

	// Name of party to whom the notice of non-extension is intended to be delivered.
	NotificationRecipientName *Max140Text `xml:"NtfctnRcptNm,omitempty"`

	// Address of party to whom the notice of non-extension is intended to be delivered.
	NotificationRecipientAddress *PostalAddress6 `xml:"NtfctnRcptAdr,omitempty"`
}

func (n *NonExtension1) SetNotificationPeriod(value string) {
	n.NotificationPeriod = (*Number)(&value)
}

func (n *NonExtension1) AddNotificationMethod() *CommunicationMethod1Choice {
	n.NotificationMethod = new(CommunicationMethod1Choice)
	return n.NotificationMethod
}

func (n *NonExtension1) AddNotificationRecipientType() *PartyType1Choice {
	n.NotificationRecipientType = new(PartyType1Choice)
	return n.NotificationRecipientType
}

func (n *NonExtension1) SetNotificationRecipientName(value string) {
	n.NotificationRecipientName = (*Max140Text)(&value)
}

func (n *NonExtension1) AddNotificationRecipientAddress() *PostalAddress6 {
	n.NotificationRecipientAddress = new(PostalAddress6)
	return n.NotificationRecipientAddress
}
