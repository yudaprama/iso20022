package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03600108 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.036.001.08 Document"`
	Message *CorporateActionMovementConfirmationV08 `xml:"CorpActnMvmntConf"`
}

func (d *Document03600108) AddMessage() *CorporateActionMovementConfirmationV08 {
	d.Message = new(CorporateActionMovementConfirmationV08)
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
type CorporateActionMovementConfirmationV08 struct {

	// Identification of a previously sent notification document.
	NotificationIdentification *model.DocumentIdentification31 `xml:"NtfctnId,omitempty"`

	// Identification of a previously sent movement preliminary advice document.
	MovementPreliminaryAdviceIdentification *model.DocumentIdentification31 `xml:"MvmntPrlimryAdvcId,omitempty"`

	// Identification of a related instruction document.
	InstructionIdentification *model.DocumentIdentification9 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification32 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely  linked to the processing of the event notified in this document.
	EventsLinkage []*model.CorporateActionEventReference3 `xml:"EvtsLkg,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation111 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *model.AccountAndBalance34 `xml:"AcctDtls"`

	// Information about the corporate action event.
	CorporateActionDetails *model.CorporateAction33 `xml:"CorpActnDtls,omitempty"`

	// Information about the corporate action option.
	CorporateActionConfirmationDetails *model.CorporateActionOption132 `xml:"CorpActnConfDtls"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative31 `xml:"AddtlInf,omitempty"`

	// Party appointed to administer the event on behalf of the issuer company/offeror. The party may be contacted for more information about the event.
	IssuerAgent []*model.PartyIdentification71Choice `xml:"IssrAgt,omitempty"`

	// Agent (principal or fiscal paying agent) appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	PayingAgent []*model.PartyIdentification71Choice `xml:"PngAgt,omitempty"`

	// Sub-agent appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	SubPayingAgent []*model.PartyIdentification71Choice `xml:"SubPngAgt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionMovementConfirmationV08) AddNotificationIdentification() *model.DocumentIdentification31 {
	c.NotificationIdentification = new(model.DocumentIdentification31)
	return c.NotificationIdentification
}

func (c *CorporateActionMovementConfirmationV08) AddMovementPreliminaryAdviceIdentification() *model.DocumentIdentification31 {
	c.MovementPreliminaryAdviceIdentification = new(model.DocumentIdentification31)
	return c.MovementPreliminaryAdviceIdentification
}

func (c *CorporateActionMovementConfirmationV08) AddInstructionIdentification() *model.DocumentIdentification9 {
	c.InstructionIdentification = new(model.DocumentIdentification9)
	return c.InstructionIdentification
}

func (c *CorporateActionMovementConfirmationV08) AddOtherDocumentIdentification() *model.DocumentIdentification32 {
	newValue := new(model.DocumentIdentification32)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmationV08) AddEventsLinkage() *model.CorporateActionEventReference3 {
	newValue := new(model.CorporateActionEventReference3)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmationV08) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation111 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation111)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementConfirmationV08) AddAccountDetails() *model.AccountAndBalance34 {
	c.AccountDetails = new(model.AccountAndBalance34)
	return c.AccountDetails
}

func (c *CorporateActionMovementConfirmationV08) AddCorporateActionDetails() *model.CorporateAction33 {
	c.CorporateActionDetails = new(model.CorporateAction33)
	return c.CorporateActionDetails
}

func (c *CorporateActionMovementConfirmationV08) AddCorporateActionConfirmationDetails() *model.CorporateActionOption132 {
	c.CorporateActionConfirmationDetails = new(model.CorporateActionOption132)
	return c.CorporateActionConfirmationDetails
}

func (c *CorporateActionMovementConfirmationV08) AddAdditionalInformation() *model.CorporateActionNarrative31 {
	c.AdditionalInformation = new(model.CorporateActionNarrative31)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementConfirmationV08) AddIssuerAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmationV08) AddPayingAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmationV08) AddSubPayingAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmationV08) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
