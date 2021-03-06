package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03700108 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.08 Document"`
	Message *CorporateActionMovementReversalAdviceV08 `xml:"CorpActnMvmntRvslAdvc"`
}

func (d *Document03700108) AddMessage() *CorporateActionMovementReversalAdviceV08 {
	d.Message = new(CorporateActionMovementReversalAdviceV08)
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
type CorporateActionMovementReversalAdviceV08 struct {

	// Identification of a previously sent movement confirmation document.
	MovementConfirmationIdentification *model.DocumentIdentification31 `xml:"MvmntConfId"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification32 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*model.CorporateActionEventReference3 `xml:"EvtsLkg,omitempty"`

	// Reason for the reversal.
	ReversalReason *model.CorporateActionReversalReason3 `xml:"RvslRsn,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation111 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *model.AccountAndBalance36 `xml:"AcctDtls"`

	// Information about the corporate action event.
	CorporateActionDetails *model.CorporateAction33 `xml:"CorpActnDtls,omitempty"`

	// Information about the corporate action option.
	CorporateActionConfirmationDetails *model.CorporateActionOption119 `xml:"CorpActnConfDtls"`

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

func (c *CorporateActionMovementReversalAdviceV08) AddMovementConfirmationIdentification() *model.DocumentIdentification31 {
	c.MovementConfirmationIdentification = new(model.DocumentIdentification31)
	return c.MovementConfirmationIdentification
}

func (c *CorporateActionMovementReversalAdviceV08) AddOtherDocumentIdentification() *model.DocumentIdentification32 {
	newValue := new(model.DocumentIdentification32)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV08) AddEventsLinkage() *model.CorporateActionEventReference3 {
	newValue := new(model.CorporateActionEventReference3)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV08) AddReversalReason() *model.CorporateActionReversalReason3 {
	c.ReversalReason = new(model.CorporateActionReversalReason3)
	return c.ReversalReason
}

func (c *CorporateActionMovementReversalAdviceV08) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation111 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation111)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementReversalAdviceV08) AddAccountDetails() *model.AccountAndBalance36 {
	c.AccountDetails = new(model.AccountAndBalance36)
	return c.AccountDetails
}

func (c *CorporateActionMovementReversalAdviceV08) AddCorporateActionDetails() *model.CorporateAction33 {
	c.CorporateActionDetails = new(model.CorporateAction33)
	return c.CorporateActionDetails
}

func (c *CorporateActionMovementReversalAdviceV08) AddCorporateActionConfirmationDetails() *model.CorporateActionOption119 {
	c.CorporateActionConfirmationDetails = new(model.CorporateActionOption119)
	return c.CorporateActionConfirmationDetails
}

func (c *CorporateActionMovementReversalAdviceV08) AddAdditionalInformation() *model.CorporateActionNarrative31 {
	c.AdditionalInformation = new(model.CorporateActionNarrative31)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementReversalAdviceV08) AddIssuerAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV08) AddPayingAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV08) AddSubPayingAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV08) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
