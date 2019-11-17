package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00900101 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.009.001.01 Document"`
	Message *AgentCANotificationAdviceV01 `xml:"AgtCANtfctnAdvc"`
}

func (d *Document00900101) AddMessage() *AgentCANotificationAdviceV01 {
	d.Message = new(AgentCANotificationAdviceV01)
	return d.Message
}

// Scope
// This message is sent by an issuer (or its agent) to a CSD to:
// - Provide a CSD with the details of a corporate action along with the possible options available to the clients of that CSD; and
// - to update a corporate action notification. A notification advice can be initially sent as a preliminary advice and subsequently replaced by another notification advice with updated information.
// Usage
// This message is used:
// - to provide a CSD with the details of a corporate action along with the possible options available to the clients of that CSD. The information can be complete or incomplete.
// - to update a corporate action notification advice. A notification advice can be initially sent as a preliminary advice and subsequently replaced by another notification advice with updated information. As per SMPG recommendation, all the information should be provided in the update, not only updated information.
// An Agent Corporate Action Notification Status Advice is sent in reply to the Agent Corporate Action Notification Advice.
// Note: The amendment of a corporate action notification is done through a replacement mechanism in line with both the ISO 15022 messages used in the flow between the CSD and its clients, and the ISO 20022 proxy voting messages.
type AgentCANotificationAdviceV01 struct {

	// Identification assigned by the Sender to unambiguously identify the advice.
	Identification *model.DocumentIdentification8 `xml:"Id"`

	// Provides information about the type of  notification advice and linked message.
	NotificationTypeAndLinkage *model.LinkedCorporateAction1 `xml:"NtfctnTpAndLkg"`

	// Provides general information about the notification advice.
	NotificationGeneralInformation *model.CorporateActionNotification1 `xml:"NtfctnGnlInf"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionInformation2 `xml:"CorpActnGnlInf"`

	// Provides details information about the CA event.
	CorporateActionDetails *model.CorporateAction2 `xml:"CorpActnDtls"`

	// Provides detailed information about the option of the CA event.
	CorporateActionOptionDetails []*model.CorporateActionOption1 `xml:"CorpActnOptnDtls,omitempty"`

	// Provides information about the contact responsible for the transaction identified in the message.
	ContactDetails []*model.ContactPerson1 `xml:"CtctDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative2 `xml:"AddtlInf,omitempty"`
}

func (a *AgentCANotificationAdviceV01) AddIdentification() *model.DocumentIdentification8 {
	a.Identification = new(model.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCANotificationAdviceV01) AddNotificationTypeAndLinkage() *model.LinkedCorporateAction1 {
	a.NotificationTypeAndLinkage = new(model.LinkedCorporateAction1)
	return a.NotificationTypeAndLinkage
}

func (a *AgentCANotificationAdviceV01) AddNotificationGeneralInformation() *model.CorporateActionNotification1 {
	a.NotificationGeneralInformation = new(model.CorporateActionNotification1)
	return a.NotificationGeneralInformation
}

func (a *AgentCANotificationAdviceV01) AddCorporateActionGeneralInformation() *model.CorporateActionInformation2 {
	a.CorporateActionGeneralInformation = new(model.CorporateActionInformation2)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCANotificationAdviceV01) AddCorporateActionDetails() *model.CorporateAction2 {
	a.CorporateActionDetails = new(model.CorporateAction2)
	return a.CorporateActionDetails
}

func (a *AgentCANotificationAdviceV01) AddCorporateActionOptionDetails() *model.CorporateActionOption1 {
	newValue := new(model.CorporateActionOption1)
	a.CorporateActionOptionDetails = append(a.CorporateActionOptionDetails, newValue)
	return newValue
}

func (a *AgentCANotificationAdviceV01) AddContactDetails() *model.ContactPerson1 {
	newValue := new(model.ContactPerson1)
	a.ContactDetails = append(a.ContactDetails, newValue)
	return newValue
}

func (a *AgentCANotificationAdviceV01) AddAdditionalInformation() *model.CorporateActionNarrative2 {
	a.AdditionalInformation = new(model.CorporateActionNarrative2)
	return a.AdditionalInformation
}
