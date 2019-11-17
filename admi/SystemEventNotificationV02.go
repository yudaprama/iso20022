package admi

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400102 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 Document"`
	Message *SystemEventNotificationV02 `xml:"SysEvtNtfctn"`
}

func (d *Document00400102) AddMessage() *SystemEventNotificationV02 {
	d.Message = new(SystemEventNotificationV02)
	return d.Message
}

// Scope
// The SystemEventNotification message is sent by a central system to notify the occurrence of an event in a central system.
// Usage
// The message can be used by a central settlement system to inform its participants of an event that is going to occur in the system, for instance that the system will be down at a certain time, etc.
type SystemEventNotificationV02 struct {

	// Detailed information about a system event.
	EventInformation *model.Event2 `xml:"EvtInf"`
}

func (s *SystemEventNotificationV02) AddEventInformation() *model.Event2 {
	s.EventInformation = new(model.Event2)
	return s.EventInformation
}
