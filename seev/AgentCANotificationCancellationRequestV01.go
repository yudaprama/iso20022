package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01000101 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.010.001.01 Document"`
	Message *AgentCANotificationCancellationRequestV01 `xml:"AgtCANtfctnCxlReq"`
}

func (d *Document01000101) AddMessage() *AgentCANotificationCancellationRequestV01 {
	d.Message = new(AgentCANotificationCancellationRequestV01)
	return d.Message
}

// Scope
// This message is sent by an issuer (or its agent) to a CSD to request the cancellation of a notification advice message.
// Usage
// When this message is used to request the cancellation of a notification advice message, the function of the message must be cancellation.
// When this message is used to request the withdrawal of a Corporate Action event or option, then the function of the message must be withdrawal.
// In both cases, the building block notification advice identification must be present to link this cancellation request to the notification advice that was previously sent.
type AgentCANotificationCancellationRequestV01 struct {

	// Identification assigned by the Sender to unambiguously identify the cancellation request.
	Identification *model.DocumentIdentification8 `xml:"Id"`

	// Provides information about the type of  notification cancellation request and linked message.
	NotificationCancellationTypeAndLinkage *model.NotificationCancellation1 `xml:"NtfctnCxlTpAndLkg"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionInformation2 `xml:"CorpActnGnlInf"`

	// Detailed information of the notification advice to be cancelled.
	CorporateActionNotificationDetails *model.CorporateActionNotificationAdvice1 `xml:"CorpActnNtfctnDtls,omitempty"`
}

func (a *AgentCANotificationCancellationRequestV01) AddIdentification() *model.DocumentIdentification8 {
	a.Identification = new(model.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCANotificationCancellationRequestV01) AddNotificationCancellationTypeAndLinkage() *model.NotificationCancellation1 {
	a.NotificationCancellationTypeAndLinkage = new(model.NotificationCancellation1)
	return a.NotificationCancellationTypeAndLinkage
}

func (a *AgentCANotificationCancellationRequestV01) AddCorporateActionGeneralInformation() *model.CorporateActionInformation2 {
	a.CorporateActionGeneralInformation = new(model.CorporateActionInformation2)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCANotificationCancellationRequestV01) AddCorporateActionNotificationDetails() *model.CorporateActionNotificationAdvice1 {
	a.CorporateActionNotificationDetails = new(model.CorporateActionNotificationAdvice1)
	return a.CorporateActionNotificationDetails
}
