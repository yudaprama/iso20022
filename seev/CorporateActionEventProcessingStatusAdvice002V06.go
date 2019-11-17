package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03200206 struct {
	XMLName xml.Name                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.032.002.06 Document"`
	Message *CorporateActionEventProcessingStatusAdvice002V06 `xml:"CorpActnEvtPrcgStsAdvc"`
}

func (d *Document03200206) AddMessage() *CorporateActionEventProcessingStatusAdvice002V06 {
	d.Message = new(CorporateActionEventProcessingStatusAdvice002V06)
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
type CorporateActionEventProcessingStatusAdvice002V06 struct {

	// Identification of a previously sent notification document.
	NotificationIdentification *model.DocumentIdentification17 `xml:"NtfctnId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification34 `xml:"OthrDocId,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation114 `xml:"CorpActnGnlInf"`

	// Information about the status of a corporate action.
	EventProcessingStatus []*model.EventProcessingStatus4Choice `xml:"EvtPrcgSts"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative19 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionEventProcessingStatusAdvice002V06) AddNotificationIdentification() *model.DocumentIdentification17 {
	c.NotificationIdentification = new(model.DocumentIdentification17)
	return c.NotificationIdentification
}

func (c *CorporateActionEventProcessingStatusAdvice002V06) AddOtherDocumentIdentification() *model.DocumentIdentification34 {
	newValue := new(model.DocumentIdentification34)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionEventProcessingStatusAdvice002V06) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation114 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation114)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionEventProcessingStatusAdvice002V06) AddEventProcessingStatus() *model.EventProcessingStatus4Choice {
	newValue := new(model.EventProcessingStatus4Choice)
	c.EventProcessingStatus = append(c.EventProcessingStatus, newValue)
	return newValue
}

func (c *CorporateActionEventProcessingStatusAdvice002V06) AddAdditionalInformation() *model.CorporateActionNarrative19 {
	c.AdditionalInformation = new(model.CorporateActionNarrative19)
	return c.AdditionalInformation
}

func (c *CorporateActionEventProcessingStatusAdvice002V06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
