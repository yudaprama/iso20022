package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200103 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.002.001.03 Document"`
	Message *MeetingCancellationV03 `xml:"MtgCxl"`
}

func (d *Document00200103) AddMessage() *MeetingCancellationV03 {
	d.Message = new(MeetingCancellationV03)
	return d.Message
}

// Scope
// The MeetingCancellation message is sent by the party that sent the MeetingNotification message to the original receiver. It is sent to cancel the previous MeetingNotification message or to advise the cancellation of a meeting.
// Usage
// The MeetingCancellation message is used in two different situations.
// First, it is used to cancel a previously sent MeetingNotification message. In this case, the MessageCancellation, the MeetingReference and the Reason building blocks need to be present.
// Second, it is used to advise that the meeting is cancelled. In this case, only the MeetingReference and Reason building blocks need to be present.
type MeetingCancellationV03 struct {

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
}

func (m *MeetingCancellationV03) AddIdentification() *model.MessageIdentification1 {
	m.Identification = new(model.MessageIdentification1)
	return m.Identification
}

func (m *MeetingCancellationV03) AddMessageCancellation() *model.AmendInformation1 {
	m.MessageCancellation = new(model.AmendInformation1)
	return m.MessageCancellation
}

func (m *MeetingCancellationV03) AddMeetingReference() *model.MeetingReference5 {
	m.MeetingReference = new(model.MeetingReference5)
	return m.MeetingReference
}

func (m *MeetingCancellationV03) AddNotifyingParty() *model.PartyIdentification9Choice {
	m.NotifyingParty = new(model.PartyIdentification9Choice)
	return m.NotifyingParty
}

func (m *MeetingCancellationV03) AddSecurity() *model.SecurityPosition6 {
	newValue := new(model.SecurityPosition6)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingCancellationV03) AddReason() *model.MeetingCancellationReason2 {
	m.Reason = new(model.MeetingCancellationReason2)
	return m.Reason
}
