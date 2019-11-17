package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03900206 struct {
	XMLName xml.Name                                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.002.06 Document"`
	Message *CorporateActionCancellationAdvice002V06 `xml:"CorpActnCxlAdvc"`
}

func (d *Document03900206) AddMessage() *CorporateActionCancellationAdvice002V06 {
	d.Message = new(CorporateActionCancellationAdvice002V06)
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
type CorporateActionCancellationAdvice002V06 struct {

	// General information about the event cancellation status and cancellation reason.
	CancellationAdviceGeneralInformation *model.CorporateActionCancellation4 `xml:"CxlAdvcGnlInf"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation93 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account and the account owner.
	AccountsDetails *model.AccountIdentification34Choice `xml:"AcctsDtls"`

	// Information about the corporate action event.
	CorporateActionDetails *model.CorporateAction35 `xml:"CorpActnDtls,omitempty"`

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

func (c *CorporateActionCancellationAdvice002V06) AddCancellationAdviceGeneralInformation() *model.CorporateActionCancellation4 {
	c.CancellationAdviceGeneralInformation = new(model.CorporateActionCancellation4)
	return c.CancellationAdviceGeneralInformation
}

func (c *CorporateActionCancellationAdvice002V06) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation93 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation93)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionCancellationAdvice002V06) AddAccountsDetails() *model.AccountIdentification34Choice {
	c.AccountsDetails = new(model.AccountIdentification34Choice)
	return c.AccountsDetails
}

func (c *CorporateActionCancellationAdvice002V06) AddCorporateActionDetails() *model.CorporateAction35 {
	c.CorporateActionDetails = new(model.CorporateAction35)
	return c.CorporateActionDetails
}

func (c *CorporateActionCancellationAdvice002V06) AddIssuerAgent() *model.PartyIdentification104Choice {
	newValue := new(model.PartyIdentification104Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdvice002V06) AddPayingAgent() *model.PartyIdentification104Choice {
	newValue := new(model.PartyIdentification104Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdvice002V06) AddSubPayingAgent() *model.PartyIdentification104Choice {
	newValue := new(model.PartyIdentification104Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdvice002V06) AddRegistrar() *model.PartyIdentification104Choice {
	c.Registrar = new(model.PartyIdentification104Choice)
	return c.Registrar
}

func (c *CorporateActionCancellationAdvice002V06) AddResellingAgent() *model.PartyIdentification104Choice {
	newValue := new(model.PartyIdentification104Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdvice002V06) AddPhysicalSecuritiesAgent() *model.PartyIdentification104Choice {
	c.PhysicalSecuritiesAgent = new(model.PartyIdentification104Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionCancellationAdvice002V06) AddDropAgent() *model.PartyIdentification104Choice {
	c.DropAgent = new(model.PartyIdentification104Choice)
	return c.DropAgent
}

func (c *CorporateActionCancellationAdvice002V06) AddSolicitationAgent() *model.PartyIdentification104Choice {
	newValue := new(model.PartyIdentification104Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionCancellationAdvice002V06) AddInformationAgent() *model.PartyIdentification104Choice {
	c.InformationAgent = new(model.PartyIdentification104Choice)
	return c.InformationAgent
}

func (c *CorporateActionCancellationAdvice002V06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
