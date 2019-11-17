package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03900104 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.04 Document"`
	Message *CorporateActionCancellationAdviceV04 `xml:"CorpActnCxlAdvc"`
}

func (d *Document03900104) AddMessage() *CorporateActionCancellationAdviceV04 {
	d.Message = new(CorporateActionCancellationAdviceV04)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionCancellationAdvice message to an account owner or its designated agent to cancel a previously announced corporate action event in case of error from the account servicer or in case of withdrawal by the issuer.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionCancellationAdviceV04 struct {

	// General information about the event cancellation status and cancellation reason.
	CancellationAdviceGeneralInformation *model.CorporateActionCancellation1 `xml:"CxlAdvcGnlInf"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation56 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account and the account owner.
	AccountsDetails *model.AccountIdentification13Choice `xml:"AcctsDtls"`

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

func (c *CorporateActionCancellationAdviceV04) AddCancellationAdviceGeneralInformation() *model.CorporateActionCancellation1 {
	c.CancellationAdviceGeneralInformation = new(model.CorporateActionCancellation1)
	return c.CancellationAdviceGeneralInformation
}

func (c *CorporateActionCancellationAdviceV04) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation56 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation56)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionCancellationAdviceV04) AddAccountsDetails() *model.AccountIdentification13Choice {
	c.AccountsDetails = new(model.AccountIdentification13Choice)
	return c.AccountsDetails
}

func (c *CorporateActionCancellationAdviceV04) AddIssuerAgent() *model.PartyIdentification46Choice {
	newValue := new(model.PartyIdentification46Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV04) AddPayingAgent() *model.PartyIdentification46Choice {
	newValue := new(model.PartyIdentification46Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV04) AddSubPayingAgent() *model.PartyIdentification46Choice {
	newValue := new(model.PartyIdentification46Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV04) AddRegistrar() *model.PartyIdentification46Choice {
	c.Registrar = new(model.PartyIdentification46Choice)
	return c.Registrar
}

func (c *CorporateActionCancellationAdviceV04) AddResellingAgent() *model.PartyIdentification46Choice {
	newValue := new(model.PartyIdentification46Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV04) AddPhysicalSecuritiesAgent() *model.PartyIdentification46Choice {
	c.PhysicalSecuritiesAgent = new(model.PartyIdentification46Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionCancellationAdviceV04) AddDropAgent() *model.PartyIdentification46Choice {
	c.DropAgent = new(model.PartyIdentification46Choice)
	return c.DropAgent
}

func (c *CorporateActionCancellationAdviceV04) AddSolicitationAgent() *model.PartyIdentification46Choice {
	newValue := new(model.PartyIdentification46Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV04) AddInformationAgent() *model.PartyIdentification46Choice {
	c.InformationAgent = new(model.PartyIdentification46Choice)
	return c.InformationAgent
}

func (c *CorporateActionCancellationAdviceV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
