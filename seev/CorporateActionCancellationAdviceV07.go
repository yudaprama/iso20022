package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03900107 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.07 Document"`
	Message *CorporateActionCancellationAdviceV07 `xml:"CorpActnCxlAdvc"`
}

func (d *Document03900107) AddMessage() *CorporateActionCancellationAdviceV07 {
	d.Message = new(CorporateActionCancellationAdviceV07)
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
type CorporateActionCancellationAdviceV07 struct {

	// General information about the event cancellation status and cancellation reason.
	CancellationAdviceGeneralInformation *model.CorporateActionCancellation3 `xml:"CxlAdvcGnlInf"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation108 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account and the account owner.
	AccountsDetails *model.AccountIdentification29Choice `xml:"AcctsDtls"`

	// Information about the corporate action event.
	CorporateActionDetails *model.CorporateAction34 `xml:"CorpActnDtls,omitempty"`

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

func (c *CorporateActionCancellationAdviceV07) AddCancellationAdviceGeneralInformation() *model.CorporateActionCancellation3 {
	c.CancellationAdviceGeneralInformation = new(model.CorporateActionCancellation3)
	return c.CancellationAdviceGeneralInformation
}

func (c *CorporateActionCancellationAdviceV07) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation108 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation108)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionCancellationAdviceV07) AddAccountsDetails() *model.AccountIdentification29Choice {
	c.AccountsDetails = new(model.AccountIdentification29Choice)
	return c.AccountsDetails
}

func (c *CorporateActionCancellationAdviceV07) AddCorporateActionDetails() *model.CorporateAction34 {
	c.CorporateActionDetails = new(model.CorporateAction34)
	return c.CorporateActionDetails
}

func (c *CorporateActionCancellationAdviceV07) AddIssuerAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV07) AddPayingAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV07) AddSubPayingAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV07) AddRegistrar() *model.PartyIdentification71Choice {
	c.Registrar = new(model.PartyIdentification71Choice)
	return c.Registrar
}

func (c *CorporateActionCancellationAdviceV07) AddResellingAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV07) AddPhysicalSecuritiesAgent() *model.PartyIdentification71Choice {
	c.PhysicalSecuritiesAgent = new(model.PartyIdentification71Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionCancellationAdviceV07) AddDropAgent() *model.PartyIdentification71Choice {
	c.DropAgent = new(model.PartyIdentification71Choice)
	return c.DropAgent
}

func (c *CorporateActionCancellationAdviceV07) AddSolicitationAgent() *model.PartyIdentification71Choice {
	newValue := new(model.PartyIdentification71Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdviceV07) AddInformationAgent() *model.PartyIdentification71Choice {
	c.InformationAgent = new(model.PartyIdentification71Choice)
	return c.InformationAgent
}

func (c *CorporateActionCancellationAdviceV07) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
