package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03200101 struct {
	XMLName xml.Name                                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.032.001.01 Document"`
	Message *CorporateActionEventProcessingStatusAdviceV01 `xml:"CorpActnEvtPrcgStsAdvc"`
}

func (d *Document03200101) AddMessage() *CorporateActionEventProcessingStatusAdviceV01 {
	d.Message = new(CorporateActionEventProcessingStatusAdviceV01)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionEventProcessingStatusAdvice message to an account owner or its designated agent to report processing status of a corporate action event.
// The account servicer uses this message to provide a reason as to why a corporate action event has not been completed by the announced payment dates.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionEventProcessingStatusAdviceV01 struct {

	// Information that unambiguously identifies a CorporateActionEventProcessingStatusAdvice message as know by the account servicer.
	Identification *model.DocumentIdentification11 `xml:"Id"`

	// Identification of a previously sent notification document.
	NotificationIdentification *model.DocumentIdentification9 `xml:"NtfctnId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification14 `xml:"OthrDocId,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation9 `xml:"CorpActnGnlInf"`

	// Information about the status of a corporate action.
	EventProcessingStatus []*model.EventProcessingStatus1Choice `xml:"EvtPrcgSts"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative10 `xml:"AddtlInf,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (c *CorporateActionEventProcessingStatusAdviceV01) AddIdentification() *model.DocumentIdentification11 {
	c.Identification = new(model.DocumentIdentification11)
	return c.Identification
}

func (c *CorporateActionEventProcessingStatusAdviceV01) AddNotificationIdentification() *model.DocumentIdentification9 {
	c.NotificationIdentification = new(model.DocumentIdentification9)
	return c.NotificationIdentification
}

func (c *CorporateActionEventProcessingStatusAdviceV01) AddOtherDocumentIdentification() *model.DocumentIdentification14 {
	newValue := new(model.DocumentIdentification14)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionEventProcessingStatusAdviceV01) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation9 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation9)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionEventProcessingStatusAdviceV01) AddEventProcessingStatus() *model.EventProcessingStatus1Choice {
	newValue := new(model.EventProcessingStatus1Choice)
	c.EventProcessingStatus = append(c.EventProcessingStatus, newValue)
	return newValue
}

func (c *CorporateActionEventProcessingStatusAdviceV01) AddAdditionalInformation() *model.CorporateActionNarrative10 {
	c.AdditionalInformation = new(model.CorporateActionNarrative10)
	return c.AdditionalInformation
}

func (c *CorporateActionEventProcessingStatusAdviceV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	c.MessageOriginator = new(model.PartyIdentification10Choice)
	return c.MessageOriginator
}

func (c *CorporateActionEventProcessingStatusAdviceV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	c.MessageRecipient = new(model.PartyIdentification10Choice)
	return c.MessageRecipient
}

func (c *CorporateActionEventProcessingStatusAdviceV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	c.Extension = append(c.Extension, newValue)
	return newValue
}
