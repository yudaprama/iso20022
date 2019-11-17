package model

// Identifies a party that notifies a financial document, the party to be notified, and whether notified party must send an acknowledgement and to whom.
type FinancingNotificationParties1 struct {

	// Party that notifies a third party.
	NotifyingParty *QualifiedPartyIdentification1 `xml:"NtifngPty"`

	// Party (to be) notified.
	NotificationReceiver *QualifiedPartyIdentification1 `xml:"NtfctnRcvr"`

	// Party to whom a notification acknowledgement has to be sent by the notification receiver.
	AcknowledgementReceiver []*QualifiedPartyIdentification1 `xml:"AckRcvr,omitempty"`
}

func (f *FinancingNotificationParties1) AddNotifyingParty() *QualifiedPartyIdentification1 {
	f.NotifyingParty = new(QualifiedPartyIdentification1)
	return f.NotifyingParty
}

func (f *FinancingNotificationParties1) AddNotificationReceiver() *QualifiedPartyIdentification1 {
	f.NotificationReceiver = new(QualifiedPartyIdentification1)
	return f.NotificationReceiver
}

func (f *FinancingNotificationParties1) AddAcknowledgementReceiver() *QualifiedPartyIdentification1 {
	newValue := new(QualifiedPartyIdentification1)
	f.AcknowledgementReceiver = append(f.AcknowledgementReceiver, newValue)
	return newValue
}
