package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300102 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.003.001.02 Document"`
	Message *MeetingEntitlementNotificationV02 `xml:"MtgEntitlmntNtfctn"`
}

func (d *Document00300102) AddMessage() *MeetingEntitlementNotificationV02 {
	d.Message = new(MeetingEntitlementNotificationV02)
	return d.Message
}

// Scope
// An account servicer sends the MeetingEntitlementNotification to an issuer, its agent, an intermediary or an account owner to advise the entitlement in relation to a shareholders meeting.
// Usage
// This message is sent to advise the quantity of securities held by an account owner. The balance is specified for the securities for which the meeting is taking place.
// This entitlement message is sent by the account servicer or the registrar to an intermediary, the issuer's agent or the issuer. It is also sent between the account servicer and the account owner or the party holding the right to vote.
// The message is also used to amend a previously sent MeetingEntitlementNotification. To notify an update, the RelatedReference must be included in the message.
type MeetingEntitlementNotificationV02 struct {

	// Identifies the notification of entitlement instruction.
	EntitlementNotificationIdentification *model.MessageIdentification1 `xml:"EntitlmntNtfctnId"`

	// Identifies the meeting entitlement message to be modified.
	RelatedReference *model.MessageIdentification `xml:"RltdRef,omitempty"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference3 `xml:"MtgRef"`

	// Party notifying the entitlement.
	NotifyingParty *model.PartyIdentification9Choice `xml:"NtifngPty"`

	// Identifies the security for which the meeting is organised, the account and the positions of the security holder.
	Security []*model.SecurityPosition5 `xml:"Scty"`

	// Defines the dates determining eligibility.
	Eligibility *model.EligibilityDates1 `xml:"Elgblty"`
}

func (m *MeetingEntitlementNotificationV02) AddEntitlementNotificationIdentification() *model.MessageIdentification1 {
	m.EntitlementNotificationIdentification = new(model.MessageIdentification1)
	return m.EntitlementNotificationIdentification
}

func (m *MeetingEntitlementNotificationV02) AddRelatedReference() *model.MessageIdentification {
	m.RelatedReference = new(model.MessageIdentification)
	return m.RelatedReference
}

func (m *MeetingEntitlementNotificationV02) AddMeetingReference() *model.MeetingReference3 {
	m.MeetingReference = new(model.MeetingReference3)
	return m.MeetingReference
}

func (m *MeetingEntitlementNotificationV02) AddNotifyingParty() *model.PartyIdentification9Choice {
	m.NotifyingParty = new(model.PartyIdentification9Choice)
	return m.NotifyingParty
}

func (m *MeetingEntitlementNotificationV02) AddSecurity() *model.SecurityPosition5 {
	newValue := new(model.SecurityPosition5)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingEntitlementNotificationV02) AddEligibility() *model.EligibilityDates1 {
	m.Eligibility = new(model.EligibilityDates1)
	return m.Eligibility
}
