package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200105 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.002.001.05 Document"`
	Message *MeetingCancellationV05 `xml:"MtgCxl"`
}

func (d *Document00200105) AddMessage() *MeetingCancellationV05 {
	d.Message = new(MeetingCancellationV05)
	return d.Message
}

// Scope
// The MeetingCancellation message is sent by the party that sent the MeetingNotification message to the original receiver. It is sent to cancel the previous MeetingNotification message or to advise the cancellation of a meeting.
// Usage
// The MeetingCancellation message is used in two different situations.
// First, it is used to cancel a previously sent MeetingNotification message. In this case, the MessageCancellation, the MeetingReference and the Reason building blocks need to be present.
// Second, it is used to advise that the meeting is cancelled. In this case, only the MeetingReference and Reason building blocks need to be present.
// This message definition is intended for use with the Business Application Header (head.001.001.01).
type MeetingCancellationV05 struct {

	// Series of elements which allow to identify a meeting.
	MeetingReference *model.MeetingReference6 `xml:"MtgRef"`

	// Identifies the security for which the meeting was organised.
	Security []*model.SecurityPosition8 `xml:"Scty,omitempty"`

	// Defines the justification for the cancellation.
	Reason *model.MeetingCancellationReason2 `xml:"Rsn"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MeetingCancellationV05) AddMeetingReference() *model.MeetingReference6 {
	m.MeetingReference = new(model.MeetingReference6)
	return m.MeetingReference
}

func (m *MeetingCancellationV05) AddSecurity() *model.SecurityPosition8 {
	newValue := new(model.SecurityPosition8)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingCancellationV05) AddReason() *model.MeetingCancellationReason2 {
	m.Reason = new(model.MeetingCancellationReason2)
	return m.Reason
}

func (m *MeetingCancellationV05) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
