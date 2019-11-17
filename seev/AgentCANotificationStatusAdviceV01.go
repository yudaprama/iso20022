package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100101 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.011.001.01 Document"`
	Message *AgentCANotificationStatusAdviceV01 `xml:"AgtCANtfctnStsAdvc"`
}

func (d *Document01100101) AddMessage() *AgentCANotificationStatusAdviceV01 {
	d.Message = new(AgentCANotificationStatusAdviceV01)
	return d.Message
}

// Scope
// This message is sent by a CSD to an issuer (or its agent) to report the status, or change in status, of a notification advice or notification cancellation request.
// Usage
// When this message is used to report the status of a notification advice then the building block Agent Corporate Action Notification Advice Identification must be present.
// When this message is used to provide the status of a notification cancellation request then the building block Notification Cancellation Request Identification must be present.
type AgentCANotificationStatusAdviceV01 struct {

	// Identification assigned by the Sender to unambiguously identify the status advice.
	Identification *model.DocumentIdentification8 `xml:"Id"`

	// Identification of the linked Agent CA Notification Advice for which a status is given.
	AgentCANotificationAdviceIdentification *model.DocumentIdentification8 `xml:"AgtCANtfctnAdvcId"`

	// Identification of the linked Agent CA Notification Cancellation Request for which a status is given.
	AgentCANotificationCancellationRequestIdentification *model.DocumentIdentification8 `xml:"AgtCANtfctnCxlReqId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionInformation2 `xml:"CorpActnGnlInf"`

	// Status of the Notification Cancellation Request sent by the issuer (agent).
	NotificationCancellationRequestStatus *model.NotificationCancellationRequestStatus1Choice `xml:"NtfctnCxlReqSts"`

	// Status of the notification advice sent by the issuer (agent).
	NotificationAdviceStatus *model.NotificationAdviceStatus1Choice `xml:"NtfctnAdvcSts"`
}

func (a *AgentCANotificationStatusAdviceV01) AddIdentification() *model.DocumentIdentification8 {
	a.Identification = new(model.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCANotificationStatusAdviceV01) AddAgentCANotificationAdviceIdentification() *model.DocumentIdentification8 {
	a.AgentCANotificationAdviceIdentification = new(model.DocumentIdentification8)
	return a.AgentCANotificationAdviceIdentification
}

func (a *AgentCANotificationStatusAdviceV01) AddAgentCANotificationCancellationRequestIdentification() *model.DocumentIdentification8 {
	a.AgentCANotificationCancellationRequestIdentification = new(model.DocumentIdentification8)
	return a.AgentCANotificationCancellationRequestIdentification
}

func (a *AgentCANotificationStatusAdviceV01) AddCorporateActionGeneralInformation() *model.CorporateActionInformation2 {
	a.CorporateActionGeneralInformation = new(model.CorporateActionInformation2)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCANotificationStatusAdviceV01) AddNotificationCancellationRequestStatus() *model.NotificationCancellationRequestStatus1Choice {
	a.NotificationCancellationRequestStatus = new(model.NotificationCancellationRequestStatus1Choice)
	return a.NotificationCancellationRequestStatus
}

func (a *AgentCANotificationStatusAdviceV01) AddNotificationAdviceStatus() *model.NotificationAdviceStatus1Choice {
	a.NotificationAdviceStatus = new(model.NotificationAdviceStatus1Choice)
	return a.NotificationAdviceStatus
}
