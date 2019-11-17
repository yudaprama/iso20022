package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200104 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.002.001.04 Document"`
	Message *MeetingCancellationV04 `xml:"MtgCxl"`
}

func (d *Document00200104) AddMessage() *MeetingCancellationV04 {
	d.Message = new(MeetingCancellationV04)
	return d.Message
}

// Scope
// The MeetingCancellation message is sent by the party that sent the MeetingNotification message to the original receiver. It is sent to cancel the previous MeetingNotification message or to advise the cancellation of a meeting.
// Usage
// The MeetingCancellation message is used in two different situations.
// First, it is used to cancel a previously sent MeetingNotification message. In this case, the MessageCancellation, the MeetingReference and the Reason building blocks need to be present.
// Second, it is used to advise that the meeting is cancelled. In this case, only the MeetingReference and Reason building blocks need to be present.
type MeetingCancellationV04 struct {

	// Identifies the cancellation message.
	Identification *model.MessageIdentification1 `xml:"Id"`

	// Information indicating that the cancellation of a message previously sent is requested (and not the cancellation of the meeting).
	MessageCancellation *model.AmendInformation1 `xml:"MsgCxl,omitempty"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference5 `xml:"MtgRef"`

	// Party notifying the cancellation of the meeting.
	NotifyingParty *model.PartyIdentification9Choice `xml:"NtifngPty,omitempty"`

	// Identifies the security for which the meeting was organised.
	Security []*model.SecurityPosition6 `xml:"Scty,omitempty"`

	// Defines the justification for the cancellation.
	Reason *model.MeetingCancellationReason2 `xml:"Rsn"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (m *MeetingCancellationV04) AddIdentification() *model.MessageIdentification1 {
	m.Identification = new(model.MessageIdentification1)
	return m.Identification
}

func (m *MeetingCancellationV04) AddMessageCancellation() *model.AmendInformation1 {
	m.MessageCancellation = new(model.AmendInformation1)
	return m.MessageCancellation
}

func (m *MeetingCancellationV04) AddMeetingReference() *model.MeetingReference5 {
	m.MeetingReference = new(model.MeetingReference5)
	return m.MeetingReference
}

func (m *MeetingCancellationV04) AddNotifyingParty() *model.PartyIdentification9Choice {
	m.NotifyingParty = new(model.PartyIdentification9Choice)
	return m.NotifyingParty
}

func (m *MeetingCancellationV04) AddSecurity() *model.SecurityPosition6 {
	newValue := new(model.SecurityPosition6)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingCancellationV04) AddReason() *model.MeetingCancellationReason2 {
	m.Reason = new(model.MeetingCancellationReason2)
	return m.Reason
}

func (m *MeetingCancellationV04) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	m.Extension = append(m.Extension, newValue)
	return newValue
}
