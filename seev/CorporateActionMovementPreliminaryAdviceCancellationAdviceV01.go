package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04400101 struct {
	XMLName xml.Name                                                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.01 Document"`
	Message *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01 `xml:"CorpActnMvmntPrlimryAdvcCxlAdvc"`
}

func (d *Document04400101) AddMessage() *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01 {
	d.Message = new(CorporateActionMovementPreliminaryAdviceCancellationAdviceV01)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionMovementPreliminaryAdviceCancellationAdvice message to an account owner or its designated agent to cancel a previously announced CorporateActionMovementPreliminaryAdvice.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionMovementPreliminaryAdviceCancellationAdviceV01 struct {

	// Information that unambiguously identifies a CorporateActionMovementPreliminaryAdviceCancellationAdvice message as know by the account servicer.
	Identification *model.DocumentIdentification11 `xml:"Id"`

	// Specifies the status of the details of the event.
	CancellationAdviceGeneralInformation *model.CorporateActionProcessingStatus1Choice `xml:"CxlAdvcGnlInf"`

	// Identification of a previously sent movement preliminary advice document.
	MovementPreliminaryAdviceIdentification *model.DocumentIdentification15 `xml:"MvmntPrlimryAdvcId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation8 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account and the account owner.
	AccountDetails *model.AccountIdentification6Choice `xml:"AcctDtls"`

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

	// Party/agent responsible for maintaining the register of a security.
	Registrar *model.PartyIdentification10Choice `xml:"Regar,omitempty"`

	// A broker-dealer responsible for reselling to new investors securities (usually bonds) that have been tendered for purchase by their owner.
	ResellingAgent []*model.PartyIdentification10Choice `xml:"RsellngAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to accept presentations of instruments, usually bonds, for transfer and or exchange.
	PhysicalSecuritiesAgent *model.PartyIdentification10Choice `xml:"PhysSctiesAgt,omitempty"`

	// A trust company, bank or similar financial institution who acts on behalf of an out of town agent or event agent where securities can be delivered in person.
	DropAgent *model.PartyIdentification10Choice `xml:"DrpAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an issuer to maintain records of investors and account balances and transactions for the consent of a material change.
	SolicitationAgent []*model.PartyIdentification10Choice `xml:"SlctnAgt,omitempty"`

	// A trust company, bank or similar financial institution assigned by an Issuer to provide information and copies of the offering documentation.
	InformationAgent *model.PartyIdentification10Choice `xml:"InfAgt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01) AddIdentification() *model.DocumentIdentification11 {
	c.Identification = new(model.DocumentIdentification11)
	return c.Identification
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01) AddCancellationAdviceGeneralInformation() *model.CorporateActionProcessingStatus1Choice {
	c.CancellationAdviceGeneralInformation = new(model.CorporateActionProcessingStatus1Choice)
	return c.CancellationAdviceGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01) AddMovementPreliminaryAdviceIdentification() *model.DocumentIdentification15 {
	c.MovementPreliminaryAdviceIdentification = new(model.DocumentIdentification15)
	return c.MovementPreliminaryAdviceIdentification
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation8 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation8)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01) AddAccountDetails() *model.AccountIdentification6Choice {
	c.AccountDetails = new(model.AccountIdentification6Choice)
	return c.AccountDetails
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	c.MessageOriginator = new(model.PartyIdentification10Choice)
	return c.MessageOriginator
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	c.MessageRecipient = new(model.PartyIdentification10Choice)
	return c.MessageRecipient
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01) AddIssuerAgent() *model.PartyIdentification10Choice {
	newValue := new(model.PartyIdentification10Choice)
	c.IssuerAgent = append(c.IssuerAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01) AddPayingAgent() *model.PartyIdentification10Choice {
	newValue := new(model.PartyIdentification10Choice)
	c.PayingAgent = append(c.PayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01) AddSubPayingAgent() *model.PartyIdentification10Choice {
	newValue := new(model.PartyIdentification10Choice)
	c.SubPayingAgent = append(c.SubPayingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01) AddRegistrar() *model.PartyIdentification10Choice {
	c.Registrar = new(model.PartyIdentification10Choice)
	return c.Registrar
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01) AddResellingAgent() *model.PartyIdentification10Choice {
	newValue := new(model.PartyIdentification10Choice)
	c.ResellingAgent = append(c.ResellingAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01) AddPhysicalSecuritiesAgent() *model.PartyIdentification10Choice {
	c.PhysicalSecuritiesAgent = new(model.PartyIdentification10Choice)
	return c.PhysicalSecuritiesAgent
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01) AddDropAgent() *model.PartyIdentification10Choice {
	c.DropAgent = new(model.PartyIdentification10Choice)
	return c.DropAgent
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01) AddSolicitationAgent() *model.PartyIdentification10Choice {
	newValue := new(model.PartyIdentification10Choice)
	c.SolicitationAgent = append(c.SolicitationAgent, newValue)
	return newValue
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01) AddInformationAgent() *model.PartyIdentification10Choice {
	c.InformationAgent = new(model.PartyIdentification10Choice)
	return c.InformationAgent
}

func (c *CorporateActionMovementPreliminaryAdviceCancellationAdviceV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	c.Extension = append(c.Extension, newValue)
	return newValue
}
