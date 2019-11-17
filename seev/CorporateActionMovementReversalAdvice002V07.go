package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03700207 struct {
	XMLName xml.Name                                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.002.07 Document"`
	Message *CorporateActionMovementReversalAdvice002V07 `xml:"CorpActnMvmntRvslAdvc"`
}

func (d *Document03700207) AddMessage() *CorporateActionMovementReversalAdvice002V07 {
	d.Message = new(CorporateActionMovementReversalAdvice002V07)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionMovementReversalAdvice message to an account owner or its designated agent to reverse previously confirmed posting of securities or cash.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionMovementReversalAdvice002V07 struct {

	// Identification of a previously sent movement confirmation document.
	MovementConfirmationIdentification *model.DocumentIdentification37 `xml:"MvmntConfId"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification38 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*model.CorporateActionEventReference4 `xml:"EvtsLkg,omitempty"`

	// Reason for the reversal.
	ReversalReason *model.CorporateActionReversalReason4 `xml:"RvslRsn,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation98 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *model.AccountAndBalance40 `xml:"AcctDtls"`

	// Information about the corporate action event.
	CorporateActionDetails *model.CorporateAction36 `xml:"CorpActnDtls,omitempty"`

	// Information about the corporate action option.
	CorporateActionConfirmationDetails *model.CorporateActionOption126 `xml:"CorpActnConfDtls"`

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

func (c *CorporateActionMovementReversalAdvice002V07) AddMovementConfirmationIdentification() *model.DocumentIdentification37 {
	c.MovementConfirmationIdentification = new(model.DocumentIdentification37)
	return c.MovementConfirmationIdentification
}

func (c *CorporateActionMovementReversalAdvice002V07) AddOtherDocumentIdentification() *model.DocumentIdentification38 {
	newValue := new(model.DocumentIdentification38)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdvice002V07) AddEventsLinkage() *model.CorporateActionEventReference4 {
	newValue := new(model.CorporateActionEventReference4)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdvice002V07) AddReversalReason() *model.CorporateActionReversalReason4 {
	c.ReversalReason = new(model.CorporateActionReversalReason4)
	return c.ReversalReason
}

func (c *CorporateActionMovementReversalAdvice002V07) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation98 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation98)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementReversalAdvice002V07) AddAccountDetails() *model.AccountAndBalance40 {
	c.AccountDetails = new(model.AccountAndBalance40)
	return c.AccountDetails
}

func (c *CorporateActionMovementReversalAdvice002V07) AddCorporateActionDetails() *model.CorporateAction36 {
	c.CorporateActionDetails = new(model.CorporateAction36)
	return c.CorporateActionDetails
}

func (c *CorporateActionMovementReversalAdvice002V07) AddCorporateActionConfirmationDetails() *model.CorporateActionOption126 {
	c.CorporateActionConfirmationDetails = new(model.CorporateActionOption126)
	return c.CorporateActionConfirmationDetails
}

func (c *CorporateActionMovementReversalAdvice002V07) AddAdditionalInformation() *model.CorporateActionNarrative35 {
	c.AdditionalInformation = new(model.CorporateActionNarrative35)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementReversalAdvice002V07) AddIssuerAgent() *model.PartyIdentification104Choice {
	newValue := new(model.PartyIdentification104Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdvice002V07) AddPayingAgent() *model.PartyIdentification104Choice {
	newValue := new(model.PartyIdentification104Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdvice002V07) AddSubPayingAgent() *model.PartyIdentification104Choice {
	newValue := new(model.PartyIdentification104Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdvice002V07) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
