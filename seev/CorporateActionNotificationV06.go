package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03100106 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.031.001.06 Document"`
	Message *CorporateActionNotificationV06 `xml:"CorpActnNtfctn"`
}

func (d *Document03100106) AddMessage() *CorporateActionNotificationV06 {
	d.Message = new(CorporateActionNotificationV06)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionNotification message to an account owner or its designated agent to notify details of a corporate action event and optionally account information, eligible balance and entitlements.
// It may also include possible elections or choices available to the account owner. The account servicer can initially send the CorporateActionNotification message as a preliminary advice, subsequently replaced by another CorporateActionNotification message with complete or confirmed information.
// It may also be sent to an account owner or its designated agent, to remind of event details and/or of missing or incomplete instructions for a corporate action event.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionNotificationV06 struct {

	// Page number of the message and  continuation indicator to indicate that the multi-parts notification is to continue or that the message is the last page of the multi-parts notification.
	Pagination *model.Pagination `xml:"Pgntn,omitempty"`

	// General information about the event notification type, status and contents.
	NotificationGeneralInformation *model.CorporateActionNotification5 `xml:"NtfctnGnlInf"`

	// Identification of a previously sent notification document.
	PreviousNotificationIdentification *model.DocumentIdentification31 `xml:"PrvsNtfctnId,omitempty"`

	// Identification of a related instruction document.
	InstructionIdentification *model.DocumentIdentification9 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification32 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*model.CorporateActionEventReference3 `xml:"EvtsLkg,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation85 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *model.AccountIdentification32Choice `xml:"AcctDtls"`

	// Provides details on rights credited to the account as for instance trading period, expiry date, renounceability.
	IntermediateSecurity *model.FinancialInstrumentAttributes68 `xml:"IntrmdtScty,omitempty"`

	// Information about the corporate action event.
	CorporateActionDetails *model.CorporateAction31 `xml:"CorpActnDtls,omitempty"`

	// Information about the corporate action option.
	CorporateActionOptionDetails []*model.CorporateActionOption114 `xml:"CorpActnOptnDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative27 `xml:"AddtlInf,omitempty"`

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

func (c *CorporateActionNotificationV06) AddPagination() *model.Pagination {
	c.Pagination = new(model.Pagination)
	return c.Pagination
}

func (c *CorporateActionNotificationV06) AddNotificationGeneralInformation() *model.CorporateActionNotification5 {
	c.NotificationGeneralInformation = new(model.CorporateActionNotification5)
	return c.NotificationGeneralInformation
}

func (c *CorporateActionNotificationV06) AddPreviousNotificationIdentification() *model.DocumentIdentification31 {
	c.PreviousNotificationIdentification = new(model.DocumentIdentification31)
	return c.PreviousNotificationIdentification
}

func (c *CorporateActionNotificationV06) AddInstructionIdentification() *model.DocumentIdentification9 {
	c.InstructionIdentification = new(model.DocumentIdentification9)
	return c.InstructionIdentification
}

func (c *CorporateActionNotificationV06) AddOtherDocumentIdentification() *model.DocumentIdentification32 {
	newValue := new(model.DocumentIdentification32)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionNotificationV06) AddEventsLinkage() *model.CorporateActionEventReference3 {
	newValue := new(model.CorporateActionEventReference3)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionNotificationV06) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation85 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation85)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionNotificationV06) AddAccountDetails() *model.AccountIdentification32Choice {
	c.AccountDetails = new(model.AccountIdentification32Choice)
	return c.AccountDetails
}

func (c *CorporateActionNotificationV06) AddIntermediateSecurity() *model.FinancialInstrumentAttributes68 {
	c.IntermediateSecurity = new(model.FinancialInstrumentAttributes68)
	return c.IntermediateSecurity
}

func (c *CorporateActionNotificationV06) AddCorporateActionDetails() *model.CorporateAction31 {
	c.CorporateActionDetails = new(model.CorporateAction31)
	return c.CorporateActionDetails
}

func (c *CorporateActionNotificationV06) AddCorporateActionOptionDetails() *model.CorporateActionOption114 {
	newValue := new(model.CorporateActionOption114)
	c.CorporateActionOptionDetails = append(c.CorporateActionOptionDetails, newValue)
	return newValue
}

func (c *CorporateActionNotificationV06) AddAdditionalInformation() *model.CorporateActionNarrative27 {
	c.AdditionalInformation = new(model.CorporateActionNarrative27)
	return c.AdditionalInformation
}

func (c *CorporateActionNotificationV06) AddIssuerAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionNotificationV06) AddPayingAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionNotificationV06) AddSubPayingAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionNotificationV06) AddRegistrar() *model.PartyIdentification71Choice {
	c.Registrar = new(model.PartyIdentification71Choice)
	return c.Registrar
}

func (c *CorporateActionNotificationV06) AddResellingAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionNotificationV06) AddPhysicalSecuritiesAgent() *model.PartyIdentification71Choice {
	c.PhysicalSecuritiesAgent = new(model.PartyIdentification71Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionNotificationV06) AddDropAgent() *model.PartyIdentification71Choice {
	c.DropAgent = new(model.PartyIdentification71Choice)
	return c.DropAgent
}

func (c *CorporateActionNotificationV06) AddSolicitationAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionNotificationV06) AddInformationAgent() *model.PartyIdentification71Choice {
	c.InformationAgent = new(model.PartyIdentification71Choice)
	return c.InformationAgent
}

func (c *CorporateActionNotificationV06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
