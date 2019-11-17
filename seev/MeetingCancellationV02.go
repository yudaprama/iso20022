package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200102 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.002.001.02 Document"`
	Message *MeetingCancellationV02 `xml:"MtgCxl"`
}

func (d *Document00200102) AddMessage() *MeetingCancellationV02 {
	d.Message = new(MeetingCancellationV02)
	return d.Message
}

// Scope
// The MeetingCancellation message is sent by the party that sent the MeetingNotification message to the original receiver. It is sent to cancel the previous MeetingNotification message or to advise the cancellation of a meeting.
// Usage
// The MeetingCancellation message is used in two different situations.
// First, it is used to cancel a previously sent MeetingNotification message. In this case, the MessageCancellation, the MeetingReference and the Reason building blocks need to be present.
// Second, it is used to advise that the meeting is cancelled. In this case, only the MeetingReference and Reason building blocks need to be present.
type MeetingCancellationV02 struct {

	// Identifies the cancellation message.
	CancellationIdentification *model.MessageIdentification1 `xml:"CxlId"`

	// Information indicating that the cancellation of a message previously sent is requested (and not the cancellation of the meeting).
	MessageCancellation *model.AmendInformation1 `xml:"MsgCxl,omitempty"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference2 `xml:"MtgRef"`

	// Party notifying the cancellation of the meeting.
	NotifyingParty *model.PartyIdentification9Choice `xml:"NtifngPty,omitempty"`

	// Identifies the security for which the meeting was organised.
	Security []*model.SecurityPosition5 `xml:"Scty,omitempty"`

	// Defines the justification for the cancellation.
	Reason *model.MeetingCancellationReason1 `xml:"Rsn"`
}

func (m *MeetingCancellationV02) AddCancellationIdentification() *model.MessageIdentification1 {
	m.CancellationIdentification = new(model.MessageIdentification1)
	return m.CancellationIdentification
}

func (m *MeetingCancellationV02) AddMessageCancellation() *model.AmendInformation1 {
	m.MessageCancellation = new(model.AmendInformation1)
	return m.MessageCancellation
}

func (m *MeetingCancellationV02) AddMeetingReference() *model.MeetingReference2 {
	m.MeetingReference = new(model.MeetingReference2)
	return m.MeetingReference
}

func (m *MeetingCancellationV02) AddNotifyingParty() *model.PartyIdentification9Choice {
	m.NotifyingParty = new(model.PartyIdentification9Choice)
	return m.NotifyingParty
}

func (m *MeetingCancellationV02) AddSecurity() *model.SecurityPosition5 {
	newValue := new(model.SecurityPosition5)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingCancellationV02) AddReason() *model.MeetingCancellationReason1 {
	m.Reason = new(model.MeetingCancellationReason1)
	return m.Reason
}
