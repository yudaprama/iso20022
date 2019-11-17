package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100105 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.001.001.05 Document"`
	Message *MeetingNotificationV05 `xml:"MtgNtfctn"`
}

func (d *Document00100105) AddMessage() *MeetingNotificationV05 {
	d.Message = new(MeetingNotificationV05)
	return d.Message
}

// Scope
// A notifying party, for example, an issuer, its agent or an intermediary, sends the MeetingNotification message to a party holding the right to vote, to announce a shareholders meeting.
// Usage
// The MeetingNotification message is used to announce a shareholders meeting, for example, it provides information on the participation details and requirements for the meeting, the vote parameters and the resolutions. The MeetingNotification message may also be used to announce an update.
// To notify an update, the Amendment building block must be filled in. Any building block that is modified must be included in the amendment message. The information previously notified and not repeated in the amendment message remains valid.
// To update the resolutions of the agenda, the complete list of resolutions must be repeated in the amendment message. The resolutions that are deleted should be assigned the status Withdrawn.
// This message definition is intended for use with the Business Application Header (head.001.001.01).
type MeetingNotificationV05 struct {

	// Information specific to an amendment.
	Amendment *model.AmendInformation1 `xml:"Amdmnt,omitempty"`

	// Defines the global status of the event contained in the notification.
	NotificationStatus *model.NotificationStatus2 `xml:"NtfctnSts"`

	// Specifies information about the meeting. This component contains meeting identifications, various deadlines, contact persons, electronic and postal locations for accessing information and proxy assignment parameters.
	Meeting *model.MeetingNotice4 `xml:"Mtg"`

	// Dates and details of the shareholders meeting.
	MeetingDetails []*model.Meeting4 `xml:"MtgDtls"`

	// Specifies the institution that is the issuer of the security to which the meeting applies.
	Issuer *model.IssuerInformation2 `xml:"Issr"`

	// Agents of the issuer.
	IssuerAgent []*model.IssuerAgent2 `xml:"IssrAgt,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	Security []*model.SecurityPosition8 `xml:"Scty"`

	// Detailed information of a resolution proposed to the vote.
	Resolution []*model.Resolution3 `xml:"Rsltn,omitempty"`

	// Specifies the conditions to be allowed to vote, the different voting methods and options, the voting deadlines and the parameters of the incentive premium.
	Vote *model.VoteParameters4 `xml:"Vote,omitempty"`

	// Specifies the entitlement ratio and the different deadlines for calculating the entitlement.
	EntitlementSpecification *model.EntitlementAssessment3 `xml:"EntitlmntSpcfctn,omitempty"`

	// Specifies requirements relative to the use of Power of Attorney.
	PowerOfAttorneyRequirements *model.PowerOfAttorneyRequirements3 `xml:"PwrOfAttnyRqrmnts,omitempty"`

	// Provides additional narrative information about the corporate event.
	AdditionalInformation *model.CorporateEventNarrative2 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MeetingNotificationV05) AddAmendment() *model.AmendInformation1 {
	m.Amendment = new(model.AmendInformation1)
	return m.Amendment
}

func (m *MeetingNotificationV05) AddNotificationStatus() *model.NotificationStatus2 {
	m.NotificationStatus = new(model.NotificationStatus2)
	return m.NotificationStatus
}

func (m *MeetingNotificationV05) AddMeeting() *model.MeetingNotice4 {
	m.Meeting = new(model.MeetingNotice4)
	return m.Meeting
}

func (m *MeetingNotificationV05) AddMeetingDetails() *model.Meeting4 {
	newValue := new(model.Meeting4)
	m.MeetingDetails = append(m.MeetingDetails, newValue)
	return newValue
}

func (m *MeetingNotificationV05) AddIssuer() *model.IssuerInformation2 {
	m.Issuer = new(model.IssuerInformation2)
	return m.Issuer
}

func (m *MeetingNotificationV05) AddIssuerAgent() *model.IssuerAgent2 {
	newValue := new(model.IssuerAgent2)
	m.IssuerAgent = append(m.IssuerAgent, newValue)
	return newValue
}

func (m *MeetingNotificationV05) AddSecurity() *model.SecurityPosition8 {
	newValue := new(model.SecurityPosition8)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingNotificationV05) AddResolution() *model.Resolution3 {
	newValue := new(model.Resolution3)
	m.Resolution = append(m.Resolution, newValue)
	return newValue
}

func (m *MeetingNotificationV05) AddVote() *model.VoteParameters4 {
	m.Vote = new(model.VoteParameters4)
	return m.Vote
}

func (m *MeetingNotificationV05) AddEntitlementSpecification() *model.EntitlementAssessment3 {
	m.EntitlementSpecification = new(model.EntitlementAssessment3)
	return m.EntitlementSpecification
}

func (m *MeetingNotificationV05) AddPowerOfAttorneyRequirements() *model.PowerOfAttorneyRequirements3 {
	m.PowerOfAttorneyRequirements = new(model.PowerOfAttorneyRequirements3)
	return m.PowerOfAttorneyRequirements
}

func (m *MeetingNotificationV05) AddAdditionalInformation() *model.CorporateEventNarrative2 {
	m.AdditionalInformation = new(model.CorporateEventNarrative2)
	return m.AdditionalInformation
}

func (m *MeetingNotificationV05) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
