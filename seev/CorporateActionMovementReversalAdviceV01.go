package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03700101 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.037.001.01 Document"`
	Message *CorporateActionMovementReversalAdviceV01 `xml:"CorpActnMvmntRvslAdvc"`
}

func (d *Document03700101) AddMessage() *CorporateActionMovementReversalAdviceV01 {
	d.Message = new(CorporateActionMovementReversalAdviceV01)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionMovementReversalAdvice message to an account owner or its designated agent to reverse previously confirmed posting of securities or cash.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionMovementReversalAdviceV01 struct {

	// Information that unambiguously identifies a CorporateActionMovementReversalAdvice message as know by the account servicer.
	Identification *model.DocumentIdentification11 `xml:"Id"`

	// Identification of a previously sent movement confirmation document.
	MovementConfirmationIdentification *model.DocumentIdentification15 `xml:"MvmntConfId"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification13 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*model.CorporateActionEventReference1 `xml:"EvtsLkg,omitempty"`

	// Reason for the reversal.
	ReversalReason *model.CorporateActionReversalReason1 `xml:"RvslRsn,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation4 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *model.AccountAndBalance4 `xml:"AcctDtls"`

	// Information about the corporate action option.
	CorporateActionConfirmationDetails *model.CorporateActionOption12 `xml:"CorpActnConfDtls"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative4 `xml:"AddtlInf,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Party appointed to administer the event on behalf of the issuer company/offeror. The party may be contacted for more information about the event.
	IssuerAgent []*model.PartyIdentification10Choice `xml:"IssrAgt,omitempty"`

	// Agent (principal or fiscal paying agent) appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	PayingAgent []*model.PartyIdentification10Choice `xml:"PngAgt,omitempty"`

	// Sub-agent appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	SubPayingAgent []*model.PartyIdentification10Choice `xml:"SubPngAgt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (c *CorporateActionMovementReversalAdviceV01) AddIdentification() *model.DocumentIdentification11 {
	c.Identification = new(model.DocumentIdentification11)
	return c.Identification
}

func (c *CorporateActionMovementReversalAdviceV01) AddMovementConfirmationIdentification() *model.DocumentIdentification15 {
	c.MovementConfirmationIdentification = new(model.DocumentIdentification15)
	return c.MovementConfirmationIdentification
}

func (c *CorporateActionMovementReversalAdviceV01) AddOtherDocumentIdentification() *model.DocumentIdentification13 {
	newValue := new(model.DocumentIdentification13)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV01) AddEventsLinkage() *model.CorporateActionEventReference1 {
	newValue := new(model.CorporateActionEventReference1)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV01) AddReversalReason() *model.CorporateActionReversalReason1 {
	c.ReversalReason = new(model.CorporateActionReversalReason1)
	return c.ReversalReason
}

func (c *CorporateActionMovementReversalAdviceV01) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation4 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation4)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementReversalAdviceV01) AddAccountDetails() *model.AccountAndBalance4 {
	c.AccountDetails = new(model.AccountAndBalance4)
	return c.AccountDetails
}

func (c *CorporateActionMovementReversalAdviceV01) AddCorporateActionConfirmationDetails() *model.CorporateActionOption12 {
	c.CorporateActionConfirmationDetails = new(model.CorporateActionOption12)
	return c.CorporateActionConfirmationDetails
}

func (c *CorporateActionMovementReversalAdviceV01) AddAdditionalInformation() *model.CorporateActionNarrative4 {
	c.AdditionalInformation = new(model.CorporateActionNarrative4)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementReversalAdviceV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	c.MessageOriginator = new(model.PartyIdentification10Choice)
	return c.MessageOriginator
}

func (c *CorporateActionMovementReversalAdviceV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	c.MessageRecipient = new(model.PartyIdentification10Choice)
	return c.MessageRecipient
}

func (c *CorporateActionMovementReversalAdviceV01) AddIssuerAgent() *model.PartyIdentification10Choice {
	newValue := new(model.PartyIdentification10Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV01) AddPayingAgent() *model.PartyIdentification10Choice {
	newValue := new(model.PartyIdentification10Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV01) AddSubPayingAgent() *model.PartyIdentification10Choice {
	newValue := new(model.PartyIdentification10Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementReversalAdviceV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	c.Extension = append(c.Extension, newValue)
	return newValue
}
