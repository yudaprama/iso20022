package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03800101 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.01 Document"`
	Message *CorporateActionNarrativeV01 `xml:"CorpActnNrrtv"`
}

func (d *Document03800101) AddMessage() *CorporateActionNarrativeV01 {
	d.Message = new(CorporateActionNarrativeV01)
	return d.Message
}

// Scope
// The CorporateActionNarrative message is sent between an account servicer and an account owner or its designated agent to cater for tax reclaims, restrictions, documentation requirements. This message is bi-directional.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
//
type CorporateActionNarrativeV01 struct {

	// Information that unambiguously identifies a CorporateActionNarrative message as know by the account servicer.
	Identification *model.DocumentIdentification11 `xml:"Id"`

	// General information about the safekeeping account and the account owner.
	AccountDetails *model.AccountIdentification9Choice `xml:"AcctDtls,omitempty"`

	// Provides information about the securitised right for entitlement.
	UnderlyingSecurity *model.UnderlyingSecurity1 `xml:"UndrlygScty,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation10 `xml:"CorpActnGnlInf"`

	// Provides additional information.
	AdditionalInformation *model.UpdatedAdditionalInformation2 `xml:"AddtlInf"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (c *CorporateActionNarrativeV01) AddIdentification() *model.DocumentIdentification11 {
	c.Identification = new(model.DocumentIdentification11)
	return c.Identification
}

func (c *CorporateActionNarrativeV01) AddAccountDetails() *model.AccountIdentification9Choice {
	c.AccountDetails = new(model.AccountIdentification9Choice)
	return c.AccountDetails
}

func (c *CorporateActionNarrativeV01) AddUnderlyingSecurity() *model.UnderlyingSecurity1 {
	c.UnderlyingSecurity = new(model.UnderlyingSecurity1)
	return c.UnderlyingSecurity
}

func (c *CorporateActionNarrativeV01) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation10 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation10)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionNarrativeV01) AddAdditionalInformation() *model.UpdatedAdditionalInformation2 {
	c.AdditionalInformation = new(model.UpdatedAdditionalInformation2)
	return c.AdditionalInformation
}

func (c *CorporateActionNarrativeV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	c.MessageOriginator = new(model.PartyIdentification10Choice)
	return c.MessageOriginator
}

func (c *CorporateActionNarrativeV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	c.MessageRecipient = new(model.PartyIdentification10Choice)
	return c.MessageRecipient
}

func (c *CorporateActionNarrativeV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	c.Extension = append(c.Extension, newValue)
	return newValue
}
