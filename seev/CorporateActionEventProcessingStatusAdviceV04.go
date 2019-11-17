package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03200104 struct {
	XMLName xml.Name                                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.032.001.04 Document"`
	Message *CorporateActionEventProcessingStatusAdviceV04 `xml:"CorpActnEvtPrcgStsAdvc"`
}

func (d *Document03200104) AddMessage() *CorporateActionEventProcessingStatusAdviceV04 {
	d.Message = new(CorporateActionEventProcessingStatusAdviceV04)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionEventProcessingStatusAdvice message to an account owner or its designated agent to report processing status of a corporate action event.
// The account servicer uses this message to provide a reason as to why a corporate action event has not been completed by the announced payment dates.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionEventProcessingStatusAdviceV04 struct {

	// Identification of a previously sent notification document.
	NotificationIdentification *model.DocumentIdentification9 `xml:"NtfctnId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification14 `xml:"OthrDocId,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation52 `xml:"CorpActnGnlInf"`

	// Information about the status of a corporate action.
	EventProcessingStatus []*model.EventProcessingStatus1Choice `xml:"EvtPrcgSts"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative10 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionEventProcessingStatusAdviceV04) AddNotificationIdentification() *model.DocumentIdentification9 {
	c.NotificationIdentification = new(model.DocumentIdentification9)
	return c.NotificationIdentification
}

func (c *CorporateActionEventProcessingStatusAdviceV04) AddOtherDocumentIdentification() *model.DocumentIdentification14 {
	newValue := new(model.DocumentIdentification14)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionEventProcessingStatusAdviceV04) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation52 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation52)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionEventProcessingStatusAdviceV04) AddEventProcessingStatus() *model.EventProcessingStatus1Choice {
	newValue := new(model.EventProcessingStatus1Choice)
	c.EventProcessingStatus = append(c.EventProcessingStatus, newValue)
	return newValue
}

func (c *CorporateActionEventProcessingStatusAdviceV04) AddAdditionalInformation() *model.CorporateActionNarrative10 {
	c.AdditionalInformation = new(model.CorporateActionNarrative10)
	return c.AdditionalInformation
}

func (c *CorporateActionEventProcessingStatusAdviceV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
