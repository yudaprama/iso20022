package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04000101 struct {
	XMLName xml.Name                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.01 Document"`
	Message *CorporateActionInstructionCancellationRequestV01 `xml:"CorpActnInstrCxlReq"`
}

func (d *Document04000101) AddMessage() *CorporateActionInstructionCancellationRequestV01 {
	d.Message = new(CorporateActionInstructionCancellationRequestV01)
	return d.Message
}

// Scope
// An account owner sends the CorporateActionInstructionCancellationRequest message to an account servicer to request cancellation of a previously sent corporate action election instruction.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionInstructionCancellationRequestV01 struct {

	// Information that unambiguously identifies a CorporateActionInstructionCancellationRequest message as know by the account owner (or the instructing party acting on its behalf).
	Identification *model.DocumentIdentification12 `xml:"Id"`

	// Identification of a previously sent instruction document.
	InstructionIdentification *model.DocumentIdentification15 `xml:"InstrId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation7 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account and the account owner.
	AccountDetails *model.AccountIdentification8 `xml:"AcctDtls"`

	// Information about the corporate action option.
	CorporateActionInstruction *model.CorporateActionOption6 `xml:"CorpActnInstr"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (c *CorporateActionInstructionCancellationRequestV01) AddIdentification() *model.DocumentIdentification12 {
	c.Identification = new(model.DocumentIdentification12)
	return c.Identification
}

func (c *CorporateActionInstructionCancellationRequestV01) AddInstructionIdentification() *model.DocumentIdentification15 {
	c.InstructionIdentification = new(model.DocumentIdentification15)
	return c.InstructionIdentification
}

func (c *CorporateActionInstructionCancellationRequestV01) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation7 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation7)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionInstructionCancellationRequestV01) AddAccountDetails() *model.AccountIdentification8 {
	c.AccountDetails = new(model.AccountIdentification8)
	return c.AccountDetails
}

func (c *CorporateActionInstructionCancellationRequestV01) AddCorporateActionInstruction() *model.CorporateActionOption6 {
	c.CorporateActionInstruction = new(model.CorporateActionOption6)
	return c.CorporateActionInstruction
}

func (c *CorporateActionInstructionCancellationRequestV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	c.MessageOriginator = new(model.PartyIdentification10Choice)
	return c.MessageOriginator
}

func (c *CorporateActionInstructionCancellationRequestV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	c.MessageRecipient = new(model.PartyIdentification10Choice)
	return c.MessageRecipient
}

func (c *CorporateActionInstructionCancellationRequestV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	c.Extension = append(c.Extension, newValue)
	return newValue
}
