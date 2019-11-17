package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03500108 struct {
	XMLName xml.Name                                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.035.001.08 Document"`
	Message *CorporateActionMovementPreliminaryAdviceV08 `xml:"CorpActnMvmntPrlimryAdvc"`
}

func (d *Document03500108) AddMessage() *CorporateActionMovementPreliminaryAdviceV08 {
	d.Message = new(CorporateActionMovementPreliminaryAdviceV08)
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
type CorporateActionMovementPreliminaryAdviceV08 struct {

	// Page number of the message and  continuation indicator to indicate that the multi-parts preliminary advice is to continue or that the message is the last page of the multi-parts preliminary advice.
	Pagination *model.Pagination `xml:"Pgntn,omitempty"`

	// General information about the movement preliminary advice document.
	MovementPreliminaryAdviceGeneralInformation *model.CorporateActionPreliminaryAdviceType2 `xml:"MvmntPrlimryAdvcGnlInf"`

	// Identification of a previously sent movement preliminary advice document.
	PreviousMovementPreliminaryAdviceIdentification *model.DocumentIdentification31 `xml:"PrvsMvmntPrlimryAdvcId,omitempty"`

	// Identification of a previously sent notification document.
	NotificationIdentification *model.DocumentIdentification31 `xml:"NtfctnId,omitempty"`

	// Identification of a previously sent movement confirmation document.
	MovementConfirmationIdentification *model.DocumentIdentification31 `xml:"MvmntConfId,omitempty"`

	// Identification of a related instruction document.
	InstructionIdentification *model.DocumentIdentification9 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification32 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*model.CorporateActionEventReference3 `xml:"EvtsLkg,omitempty"`

	// Reason for the reversal.
	ReversalReason *model.CorporateActionReversalReason3 `xml:"RvslRsn,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation106 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *model.AccountIdentification32Choice `xml:"AcctDtls"`

	// Information about the corporate action event.
	CorporateActionDetails *model.CorporateAction32 `xml:"CorpActnDtls,omitempty"`

	// Information about the corporate action option.
	CorporateActionMovementDetails []*model.CorporateActionOption129 `xml:"CorpActnMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative28 `xml:"AddtlInf,omitempty"`

	// Party appointed to administer the event on behalf of the issuer company/offeror. The party may be contacted for more information about the event.
	IssuerAgent []*model.PartyIdentification71Choice `xml:"IssrAgt,omitempty"`

	// Agent (principal or fiscal paying agent) appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	PayingAgent []*model.PartyIdentification71Choice `xml:"PngAgt,omitempty"`

	// Sub-agent appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	SubPayingAgent []*model.PartyIdentification71Choice `xml:"SubPngAgt,omitempty"`

	// Party/agent responsible for maintaining the register of a security.
	Registrar *model.PartyIdentification71Choice `xml:"Regar,omitempty"`

	// A broker-dealer responsible for reselling to new investors securities (usually bonds) that have been tendered for purchase by their owner.
	ResellingAgent []*model.PartyIdentification71Choice `xml:"RsellngAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to accept presentations of instruments, usually bonds, for transfer and or exchange.
	PhysicalSecuritiesAgent *model.PartyIdentification71Choice `xml:"PhysSctiesAgt,omitempty"`

	// A trust company, bank or similar financial institution who acts on behalf of an out of town agent or event agent where securities can be delivered in person.
	DropAgent *model.PartyIdentification71Choice `xml:"DrpAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to maintain records of investors and account balances and transactions for the consent of a material change.
	SolicitationAgent []*model.PartyIdentification71Choice `xml:"SlctnAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an Issuer to provide information and copies of the offering documentation.
	InformationAgent *model.PartyIdentification71Choice `xml:"InfAgt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddPagination() *model.Pagination {
	c.Pagination = new(model.Pagination)
	return c.Pagination
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddMovementPreliminaryAdviceGeneralInformation() *model.CorporateActionPreliminaryAdviceType2 {
	c.MovementPreliminaryAdviceGeneralInformation = new(model.CorporateActionPreliminaryAdviceType2)
	return c.MovementPreliminaryAdviceGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddPreviousMovementPreliminaryAdviceIdentification() *model.DocumentIdentification31 {
	c.PreviousMovementPreliminaryAdviceIdentification = new(model.DocumentIdentification31)
	return c.PreviousMovementPreliminaryAdviceIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddNotificationIdentification() *model.DocumentIdentification31 {
	c.NotificationIdentification = new(model.DocumentIdentification31)
	return c.NotificationIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddMovementConfirmationIdentification() *model.DocumentIdentification31 {
	c.MovementConfirmationIdentification = new(model.DocumentIdentification31)
	return c.MovementConfirmationIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddInstructionIdentification() *model.DocumentIdentification9 {
	c.InstructionIdentification = new(model.DocumentIdentification9)
	return c.InstructionIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddOtherDocumentIdentification() *model.DocumentIdentification32 {
	newValue := new(model.DocumentIdentification32)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddEventsLinkage() *model.CorporateActionEventReference3 {
	newValue := new(model.CorporateActionEventReference3)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddReversalReason() *model.CorporateActionReversalReason3 {
	c.ReversalReason = new(model.CorporateActionReversalReason3)
	return c.ReversalReason
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation106 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation106)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddAccountDetails() *model.AccountIdentification32Choice {
	c.AccountDetails = new(model.AccountIdentification32Choice)
	return c.AccountDetails
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddCorporateActionDetails() *model.CorporateAction32 {
	c.CorporateActionDetails = new(model.CorporateAction32)
	return c.CorporateActionDetails
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddCorporateActionMovementDetails() *model.CorporateActionOption129 {
	newValue := new(model.CorporateActionOption129)
	c.CorporateActionMovementDetails = append(c.CorporateActionMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddAdditionalInformation() *model.CorporateActionNarrative28 {
	c.AdditionalInformation = new(model.CorporateActionNarrative28)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddIssuerAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddPayingAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddSubPayingAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddRegistrar() *model.PartyIdentification71Choice {
	c.Registrar = new(model.PartyIdentification71Choice)
	return c.Registrar
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddResellingAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddPhysicalSecuritiesAgent() *model.PartyIdentification71Choice {
	c.PhysicalSecuritiesAgent = new(model.PartyIdentification71Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddDropAgent() *model.PartyIdentification71Choice {
	c.DropAgent = new(model.PartyIdentification71Choice)
	return c.DropAgent
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddSolicitationAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddInformationAgent() *model.PartyIdentification71Choice {
	c.InformationAgent = new(model.PartyIdentification71Choice)
	return c.InformationAgent
}

func (c *CorporateActionMovementPreliminaryAdviceV08) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
