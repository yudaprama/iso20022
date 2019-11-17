package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03500103 struct {
	XMLName xml.Name                                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.035.001.03 Document"`
	Message *CorporateActionMovementPreliminaryAdviceV03 `xml:"CorpActnMvmntPrlimryAdvc"`
}

func (d *Document03500103) AddMessage() *CorporateActionMovementPreliminaryAdviceV03 {
	d.Message = new(CorporateActionMovementPreliminaryAdviceV03)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionMovementPreliminaryAdvice message to an account owner or its designated agent to pre-advise upcoming posting or reversal of securities and/or cash postings.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionMovementPreliminaryAdviceV03 struct {

	// Page number of the message and  continuation indicator to indicate that the multi-parts preliminary advice is to continue or that the message is the last page of the multi-parts preliminary advice.
	Pagination *model.Pagination `xml:"Pgntn,omitempty"`

	// General information about the movement preliminary advice document.
	MovementPreliminaryAdviceGeneralInformation *model.CorporateActionPreliminaryAdviceType1 `xml:"MvmntPrlimryAdvcGnlInf"`

	// Identification of a previously sent movement preliminary advice document.
	PreviousMovementPreliminaryAdviceIdentification *model.DocumentIdentification15 `xml:"PrvsMvmntPrlimryAdvcId,omitempty"`

	// Identification of a previously sent notification document.
	NotificationIdentification *model.DocumentIdentification15 `xml:"NtfctnId,omitempty"`

	// Identification of a previously sent movement confirmation document.
	MovementConfirmationIdentification *model.DocumentIdentification15 `xml:"MvmntConfId,omitempty"`

	// Identification of a related instruction document.
	InstructionIdentification *model.DocumentIdentification9 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification13 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*model.CorporateActionEventReference1 `xml:"EvtsLkg,omitempty"`

	// Reason for the reversal.
	ReversalReason *model.CorporateActionReversalReason1 `xml:"RvslRsn,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation38 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *model.AccountIdentification18Choice `xml:"AcctDtls"`

	// Information about the corporate action option.
	CorporateActionMovementDetails []*model.CorporateActionOption35 `xml:"CorpActnMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative6 `xml:"AddtlInf,omitempty"`

	// Party appointed to administer the event on behalf of the issuer company/offeror. The party may be contacted for more information about the event.
	IssuerAgent []*model.PartyIdentification46Choice `xml:"IssrAgt,omitempty"`

	// Agent (principal or fiscal paying agent) appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	PayingAgent []*model.PartyIdentification46Choice `xml:"PngAgt,omitempty"`

	// Sub-agent appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	SubPayingAgent []*model.PartyIdentification46Choice `xml:"SubPngAgt,omitempty"`

	// Party/agent responsible for maintaining the register of a security.
	Registrar *model.PartyIdentification46Choice `xml:"Regar,omitempty"`

	// A broker-dealer responsible for reselling to new investors securities (usually bonds) that have been tendered for purchase by their owner.
	ResellingAgent []*model.PartyIdentification46Choice `xml:"RsellngAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to accept presentations of instruments, usually bonds, for transfer and or exchange.
	PhysicalSecuritiesAgent *model.PartyIdentification46Choice `xml:"PhysSctiesAgt,omitempty"`

	// A trust company, bank or similar financial institution who acts on behalf of an out of town agent or event agent where securities can be delivered in person.
	DropAgent *model.PartyIdentification46Choice `xml:"DrpAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to maintain records of investors and account balances and transactions for the consent of a material change.
	SolicitationAgent []*model.PartyIdentification46Choice `xml:"SlctnAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an Issuer to provide information and copies of the offering documentation.
	InformationAgent *model.PartyIdentification46Choice `xml:"InfAgt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddPagination() *model.Pagination {
	c.Pagination = new(model.Pagination)
	return c.Pagination
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddMovementPreliminaryAdviceGeneralInformation() *model.CorporateActionPreliminaryAdviceType1 {
	c.MovementPreliminaryAdviceGeneralInformation = new(model.CorporateActionPreliminaryAdviceType1)
	return c.MovementPreliminaryAdviceGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddPreviousMovementPreliminaryAdviceIdentification() *model.DocumentIdentification15 {
	c.PreviousMovementPreliminaryAdviceIdentification = new(model.DocumentIdentification15)
	return c.PreviousMovementPreliminaryAdviceIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddNotificationIdentification() *model.DocumentIdentification15 {
	c.NotificationIdentification = new(model.DocumentIdentification15)
	return c.NotificationIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddMovementConfirmationIdentification() *model.DocumentIdentification15 {
	c.MovementConfirmationIdentification = new(model.DocumentIdentification15)
	return c.MovementConfirmationIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddInstructionIdentification() *model.DocumentIdentification9 {
	c.InstructionIdentification = new(model.DocumentIdentification9)
	return c.InstructionIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddOtherDocumentIdentification() *model.DocumentIdentification13 {
	newValue := new(model.DocumentIdentification13)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddEventsLinkage() *model.CorporateActionEventReference1 {
	newValue := new(model.CorporateActionEventReference1)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddReversalReason() *model.CorporateActionReversalReason1 {
	c.ReversalReason = new(model.CorporateActionReversalReason1)
	return c.ReversalReason
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation38 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation38)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddAccountDetails() *model.AccountIdentification18Choice {
	c.AccountDetails = new(model.AccountIdentification18Choice)
	return c.AccountDetails
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddCorporateActionMovementDetails() *model.CorporateActionOption35 {
	newValue := new(model.CorporateActionOption35)
	c.CorporateActionMovementDetails = append(c.CorporateActionMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddAdditionalInformation() *model.CorporateActionNarrative6 {
	c.AdditionalInformation = new(model.CorporateActionNarrative6)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddIssuerAgent() *model.PartyIdentification46Choice {
	newValue := new(model.PartyIdentification46Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddPayingAgent() *model.PartyIdentification46Choice {
	newValue := new(model.PartyIdentification46Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddSubPayingAgent() *model.PartyIdentification46Choice {
	newValue := new(model.PartyIdentification46Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddRegistrar() *model.PartyIdentification46Choice {
	c.Registrar = new(model.PartyIdentification46Choice)
	return c.Registrar
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddResellingAgent() *model.PartyIdentification46Choice {
	newValue := new(model.PartyIdentification46Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddPhysicalSecuritiesAgent() *model.PartyIdentification46Choice {
	c.PhysicalSecuritiesAgent = new(model.PartyIdentification46Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddDropAgent() *model.PartyIdentification46Choice {
	c.DropAgent = new(model.PartyIdentification46Choice)
	return c.DropAgent
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddSolicitationAgent() *model.PartyIdentification46Choice {
	newValue := new(model.PartyIdentification46Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddInformationAgent() *model.PartyIdentification46Choice {
	c.InformationAgent = new(model.PartyIdentification46Choice)
	return c.InformationAgent
}

func (c *CorporateActionMovementPreliminaryAdviceV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
