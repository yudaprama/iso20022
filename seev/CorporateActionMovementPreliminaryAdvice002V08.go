package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03500208 struct {
	XMLName xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.035.002.08 Document"`
	Message *CorporateActionMovementPreliminaryAdvice002V08 `xml:"CorpActnMvmntPrlimryAdvc"`
}

func (d *Document03500208) AddMessage() *CorporateActionMovementPreliminaryAdvice002V08 {
	d.Message = new(CorporateActionMovementPreliminaryAdvice002V08)
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
type CorporateActionMovementPreliminaryAdvice002V08 struct {

	// Page number of the message and  continuation indicator to indicate that the multi-parts preliminary advice is to continue or that the message is the last page of the multi-parts preliminary advice.
	Pagination *model.Pagination `xml:"Pgntn,omitempty"`

	// General information about the movement preliminary advice document.
	MovementPreliminaryAdviceGeneralInformation *model.CorporateActionPreliminaryAdviceType2 `xml:"MvmntPrlimryAdvcGnlInf"`

	// Identification of a previously sent movement preliminary advice document.
	PreviousMovementPreliminaryAdviceIdentification *model.DocumentIdentification37 `xml:"PrvsMvmntPrlimryAdvcId,omitempty"`

	// Identification of a previously sent notification document.
	NotificationIdentification *model.DocumentIdentification37 `xml:"NtfctnId,omitempty"`

	// Identification of a previously sent movement confirmation document.
	MovementConfirmationIdentification *model.DocumentIdentification37 `xml:"MvmntConfId,omitempty"`

	// Identification of a related instruction document.
	InstructionIdentification *model.DocumentIdentification17 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification38 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*model.CorporateActionEventReference4 `xml:"EvtsLkg,omitempty"`

	// Reason for the reversal.
	ReversalReason *model.CorporateActionReversalReason4 `xml:"RvslRsn,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation117 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *model.AccountIdentification36Choice `xml:"AcctDtls"`

	// Information about the corporate action event.
	CorporateActionDetails *model.CorporateAction38 `xml:"CorpActnDtls,omitempty"`

	// Information about the corporate action option.
	CorporateActionMovementDetails []*model.CorporateActionOption135 `xml:"CorpActnMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative37 `xml:"AddtlInf,omitempty"`

	// Party appointed to administer the event on behalf of the issuer company/offeror. The party may be contacted for more information about the event.
	IssuerAgent []*model.PartyIdentification104Choice `xml:"IssrAgt,omitempty"`

	// Agent (principal or fiscal paying agent) appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	PayingAgent []*model.PartyIdentification104Choice `xml:"PngAgt,omitempty"`

	// Sub-agent appointed to execute the payment for the corporate action event on behalf of the issuer company/offeror.
	SubPayingAgent []*model.PartyIdentification104Choice `xml:"SubPngAgt,omitempty"`

	// Party/agent responsible for maintaining the register of a security.
	Registrar *model.PartyIdentification104Choice `xml:"Regar,omitempty"`

	// A broker-dealer responsible for reselling to new investors securities (usually bonds) that have been tendered for purchase by their owner.
	ResellingAgent []*model.PartyIdentification104Choice `xml:"RsellngAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to accept presentations of instruments, usually bonds, for transfer and or exchange.
	PhysicalSecuritiesAgent *model.PartyIdentification104Choice `xml:"PhysSctiesAgt,omitempty"`

	// A trust company, bank or similar financial institution who acts on behalf of an out of town agent or event agent where securities can be delivered in person.
	DropAgent *model.PartyIdentification104Choice `xml:"DrpAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to maintain records of investors and account balances and transactions for the consent of a material change.
	SolicitationAgent []*model.PartyIdentification104Choice `xml:"SlctnAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an Issuer to provide information and copies of the offering documentation.
	InformationAgent *model.PartyIdentification104Choice `xml:"InfAgt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddPagination() *model.Pagination {
	c.Pagination = new(model.Pagination)
	return c.Pagination
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddMovementPreliminaryAdviceGeneralInformation() *model.CorporateActionPreliminaryAdviceType2 {
	c.MovementPreliminaryAdviceGeneralInformation = new(model.CorporateActionPreliminaryAdviceType2)
	return c.MovementPreliminaryAdviceGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddPreviousMovementPreliminaryAdviceIdentification() *model.DocumentIdentification37 {
	c.PreviousMovementPreliminaryAdviceIdentification = new(model.DocumentIdentification37)
	return c.PreviousMovementPreliminaryAdviceIdentification
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddNotificationIdentification() *model.DocumentIdentification37 {
	c.NotificationIdentification = new(model.DocumentIdentification37)
	return c.NotificationIdentification
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddMovementConfirmationIdentification() *model.DocumentIdentification37 {
	c.MovementConfirmationIdentification = new(model.DocumentIdentification37)
	return c.MovementConfirmationIdentification
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddInstructionIdentification() *model.DocumentIdentification17 {
	c.InstructionIdentification = new(model.DocumentIdentification17)
	return c.InstructionIdentification
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddOtherDocumentIdentification() *model.DocumentIdentification38 {
	newValue := new(model.DocumentIdentification38)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddEventsLinkage() *model.CorporateActionEventReference4 {
	newValue := new(model.CorporateActionEventReference4)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddReversalReason() *model.CorporateActionReversalReason4 {
	c.ReversalReason = new(model.CorporateActionReversalReason4)
	return c.ReversalReason
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation117 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation117)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddAccountDetails() *model.AccountIdentification36Choice {
	c.AccountDetails = new(model.AccountIdentification36Choice)
	return c.AccountDetails
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddCorporateActionDetails() *model.CorporateAction38 {
	c.CorporateActionDetails = new(model.CorporateAction38)
	return c.CorporateActionDetails
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddCorporateActionMovementDetails() *model.CorporateActionOption135 {
	newValue := new(model.CorporateActionOption135)
	c.CorporateActionMovementDetails = append(c.CorporateActionMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddAdditionalInformation() *model.CorporateActionNarrative37 {
	c.AdditionalInformation = new(model.CorporateActionNarrative37)
	return c.AdditionalInformation
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddIssuerAgent() *model.PartyIdentification104Choice {
	newValue := new(model.PartyIdentification104Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddPayingAgent() *model.PartyIdentification104Choice {
	newValue := new(model.PartyIdentification104Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddSubPayingAgent() *model.PartyIdentification104Choice {
	newValue := new(model.PartyIdentification104Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddRegistrar() *model.PartyIdentification104Choice {
	c.Registrar = new(model.PartyIdentification104Choice)
	return c.Registrar
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddResellingAgent() *model.PartyIdentification104Choice {
	newValue := new(model.PartyIdentification104Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddPhysicalSecuritiesAgent() *model.PartyIdentification104Choice {
	c.PhysicalSecuritiesAgent = new(model.PartyIdentification104Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddDropAgent() *model.PartyIdentification104Choice {
	c.DropAgent = new(model.PartyIdentification104Choice)
	return c.DropAgent
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddSolicitationAgent() *model.PartyIdentification104Choice {
	newValue := new(model.PartyIdentification104Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddInformationAgent() *model.PartyIdentification104Choice {
	c.InformationAgent = new(model.PartyIdentification104Choice)
	return c.InformationAgent
}

func (c *CorporateActionMovementPreliminaryAdvice002V08) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
