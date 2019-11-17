package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03600106 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.036.001.06 Document"`
	Message *CorporateActionMovementConfirmationV06 `xml:"CorpActnMvmntConf"`
}

func (d *Document03600106) AddMessage() *CorporateActionMovementConfirmationV06 {
	d.Message = new(CorporateActionMovementConfirmationV06)
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
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionMovementConfirmationV06 struct {

	// Identification of a previously sent notification document.
	NotificationIdentification *model.DocumentIdentification15 `xml:"NtfctnId,omitempty"`

	// Identification of a previously sent movement preliminary advice document.
	MovementPreliminaryAdviceIdentification *model.DocumentIdentification15 `xml:"MvmntPrlimryAdvcId,omitempty"`

	// Identification of a related instruction document.
	InstructionIdentification *model.DocumentIdentification9 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification13 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely  linked to the processing of the event notified in this document.
	EventsLinkage []*model.CorporateActionEventReference1 `xml:"EvtsLkg,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation79 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *model.AccountAndBalance26 `xml:"AcctDtls"`

	// Information about the corporate action event.
	CorporateActionDetails *model.CorporateAction13 `xml:"CorpActnDtls,omitempty"`

	// Information about the corporate action option.
	CorporateActionConfirmationDetails *model.CorporateActionOption111 `xml:"CorpActnConfDtls"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative4 `xml:"AddtlInf,omitempty"`

	// Party appointed to administer the event on behalf of the issuer company/offeror. The party may be contacted for more information about the event.
	IssuerAgent []*model.PartyIdentification40Choice `xml:"IssrAgt,omitempty"`

	// Agent (principal or fiscal paying agent) appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	PayingAgent []*model.PartyIdentification40Choice `xml:"PngAgt,omitempty"`

	// Sub-agent appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	SubPayingAgent []*model.PartyIdentification40Choice `xml:"SubPngAgt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionMovementConfirmationV06) AddNotificationIdentification() *model.DocumentIdentification15 {
	c.NotificationIdentification = new(model.DocumentIdentification15)
	return c.NotificationIdentification
}

func (c *CorporateActionMovementConfirmationV06) AddMovementPreliminaryAdviceIdentification() *model.DocumentIdentification15 {
	c.MovementPreliminaryAdviceIdentification = new(model.DocumentIdentification15)
	return c.MovementPreliminaryAdviceIdentification
}

func (c *CorporateActionMovementConfirmationV06) AddInstructionIdentification() *model.DocumentIdentification9 {
	c.InstructionIdentification = new(model.DocumentIdentification9)
	return c.InstructionIdentification
}

func (c *CorporateActionMovementConfirmationV06) AddOtherDocumentIdentification() *model.DocumentIdentification13 {
	newValue := new(model.DocumentIdentification13)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmationV06) AddEventsLinkage() *model.CorporateActionEventReference1 {
	newValue := new(model.CorporateActionEventReference1)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmationV06) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation79 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation79)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementConfirmationV06) AddAccountDetails() *model.AccountAndBalance26 {
	c.AccountDetails = new(model.AccountAndBalance26)
	return c.AccountDetails
}

func (c *CorporateActionMovementConfirmationV06) AddCorporateActionDetails() *model.CorporateAction13 {
	c.CorporateActionDetails = new(model.CorporateAction13)
	return c.CorporateActionDetails
}

func (c *CorporateActionMovementConfirmationV06) AddCorporateActionConfirmationDetails() *model.CorporateActionOption111 {
	c.CorporateActionConfirmationDetails = new(model.CorporateActionOption111)
	return c.CorporateActionConfirmationDetails
}

func (c *CorporateActionMovementConfirmationV06) AddAdditionalInformation() *model.CorporateActionNarrative4 {
	c.AdditionalInformation = new(model.CorporateActionNarrative4)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementConfirmationV06) AddIssuerAgent() *model.PartyIdentification40Choice {
	newValue := new(model.PartyIdentification40Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmationV06) AddPayingAgent() *model.PartyIdentification40Choice {
	newValue := new(model.PartyIdentification40Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmationV06) AddSubPayingAgent() *model.PartyIdentification40Choice {
	newValue := new(model.PartyIdentification40Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementConfirmationV06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
