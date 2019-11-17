package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300104 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.003.001.04 Document"`
	Message *MeetingEntitlementNotificationV04 `xml:"MtgEntitlmntNtfctn"`
}

func (d *Document00300104) AddMessage() *MeetingEntitlementNotificationV04 {
	d.Message = new(MeetingEntitlementNotificationV04)
	return d.Message
}

// Scope
// An account servicer sends the MeetingEntitlementNotification to an issuer, its agent, an intermediary or an account owner to advise the entitlement in relation to a shareholders meeting.
// Usage
// This message is sent to advise the quantity of securities held by an account owner. The balance is specified for the securities for which the meeting is taking place.
// This entitlement message is sent by the account servicer or the registrar to an intermediary, the issuer's agent or the issuer. It is also sent between the account servicer and the account owner or the party holding the right to vote.
// The message is also used to amend a previously sent MeetingEntitlementNotification. To notify an update, the RelatedReference must be included in the message.
type MeetingEntitlementNotificationV04 struct {

	// Identifies the notification of entitlement instruction.
	Identification *model.MessageIdentification1 `xml:"Id"`

	// Identifies the meeting entitlement message to be modified.
	RelatedReference *model.MessageIdentification `xml:"RltdRef,omitempty"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference4 `xml:"MtgRef"`

	// Party notifying the entitlement.
	NotifyingParty *model.PartyIdentification9Choice `xml:"NtifngPty"`

	// Identifies the security for which the meeting is organised, the account and the positions of the security holder.
	Security []*model.SecurityPosition7 `xml:"Scty"`

	// Defines the dates determining eligibility.
	Eligibility *model.EligibilityDates1 `xml:"Elgblty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (m *MeetingEntitlementNotificationV04) AddIdentification() *model.MessageIdentification1 {
	m.Identification = new(model.MessageIdentification1)
	return m.Identification
}

func (m *MeetingEntitlementNotificationV04) AddRelatedReference() *model.MessageIdentification {
	m.RelatedReference = new(model.MessageIdentification)
	return m.RelatedReference
}

func (m *MeetingEntitlementNotificationV04) AddMeetingReference() *model.MeetingReference4 {
	m.MeetingReference = new(model.MeetingReference4)
	return m.MeetingReference
}

func (m *MeetingEntitlementNotificationV04) AddNotifyingParty() *model.PartyIdentification9Choice {
	m.NotifyingParty = new(model.PartyIdentification9Choice)
	return m.NotifyingParty
}

func (m *MeetingEntitlementNotificationV04) AddSecurity() *model.SecurityPosition7 {
	newValue := new(model.SecurityPosition7)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingEntitlementNotificationV04) AddEligibility() *model.EligibilityDates1 {
	m.Eligibility = new(model.EligibilityDates1)
	return m.Eligibility
}

func (m *MeetingEntitlementNotificationV04) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	m.Extension = append(m.Extension, newValue)
	return newValue
}
