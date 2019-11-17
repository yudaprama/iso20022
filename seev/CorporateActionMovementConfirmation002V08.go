package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03600208 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.036.002.08 Document"`
	Message *CorporateActionMovementConfirmation002V08 `xml:"CorpActnMvmntConf"`
}

func (d *Document03600208) AddMessage() *CorporateActionMovementConfirmation002V08 {
	d.Message = new(CorporateActionMovementConfirmation002V08)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionMovementConfirmation message to an account owner or its designated agent to confirm posting of securities or cash as a result of a corporate action event.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionMovementConfirmation002V08 struct {

	// Identification of a previously sent notification document.
	NotificationIdentification *model.DocumentIdentification37 `xml:"NtfctnId,omitempty"`

	// Identification of a previously sent movement preliminary advice document.
	MovementPreliminaryAdviceIdentification *model.DocumentIdentification37 `xml:"MvmntPrlimryAdvcId,omitempty"`

	// Identification of a related instruction document.
	InstructionIdentification *model.DocumentIdentification17 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification38 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely  linked to the processing of the event notified in this document.
	EventsLinkage []*model.CorporateActionEventReference4 `xml:"EvtsLkg,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation118 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *model.AccountAndBalance38 `xml:"AcctDtls"`

	// Information about the corporate action event.
	CorporateActionDetails *model.CorporateAction36 `xml:"CorpActnDtls,omitempty"`

	// Information about the corporate action option.
	CorporateActionConfirmationDetails *model.CorporateActionOption136 `xml:"CorpActnConfDtls"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative35 `xml:"AddtlInf,omitempty"`

	// Party appointed to administer the event on behalf of the issuer company/offeror. The party may be contacted for more information about the event.
	IssuerAgent []*model.PartyIdentification104Choice `xml:"IssrAgt,omitempty"`

	// Agent (principal or fiscal paying agent) appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	PayingAgent []*model.PartyIdentification104Choice `xml:"PngAgt,omitempty"`

	// Sub-agent appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	SubPayingAgent []*model.PartyIdentification104Choice `xml:"SubPngAgt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionMovementConfirmation002V08) AddNotificationIdentification() *model.DocumentIdentification37 {
	c.NotificationIdentification = new(model.DocumentIdentification37)
	return c.NotificationIdentification
}

func (c *CorporateActionMovementConfirmation002V08) AddMovementPreliminaryAdviceIdentification() *model.DocumentIdentification37 {
	c.MovementPreliminaryAdviceIdentification = new(model.DocumentIdentification37)
	return c.MovementPreliminaryAdviceIdentification
}

func (c *CorporateActionMovementConfirmation002V08) AddInstructionIdentification() *model.DocumentIdentification17 {
	c.InstructionIdentification = new(model.DocumentIdentification17)
	return c.InstructionIdentification
}

func (c *CorporateActionMovementConfirmation002V08) AddOtherDocumentIdentification() *model.DocumentIdentification38 {
	newValue := new(model.DocumentIdentification38)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmation002V08) AddEventsLinkage() *model.CorporateActionEventReference4 {
	newValue := new(model.CorporateActionEventReference4)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmation002V08) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation118 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation118)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementConfirmation002V08) AddAccountDetails() *model.AccountAndBalance38 {
	c.AccountDetails = new(model.AccountAndBalance38)
	return c.AccountDetails
}

func (c *CorporateActionMovementConfirmation002V08) AddCorporateActionDetails() *model.CorporateAction36 {
	c.CorporateActionDetails = new(model.CorporateAction36)
	return c.CorporateActionDetails
}

func (c *CorporateActionMovementConfirmation002V08) AddCorporateActionConfirmationDetails() *model.CorporateActionOption136 {
	c.CorporateActionConfirmationDetails = new(model.CorporateActionOption136)
	return c.CorporateActionConfirmationDetails
}

func (c *CorporateActionMovementConfirmation002V08) AddAdditionalInformation() *model.CorporateActionNarrative35 {
	c.AdditionalInformation = new(model.CorporateActionNarrative35)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementConfirmation002V08) AddIssuerAgent() *model.PartyIdentification104Choice {
	newValue := new(model.PartyIdentification104Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmation002V08) AddPayingAgent() *model.PartyIdentification104Choice {
	newValue := new(model.PartyIdentification104Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmation002V08) AddSubPayingAgent() *model.PartyIdentification104Choice {
	newValue := new(model.PartyIdentification104Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmation002V08) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
