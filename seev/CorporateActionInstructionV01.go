package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03300101 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.01 Document"`
	Message *CorporateActionInstructionV01 `xml:"CorpActnInstr"`
}

func (d *Document03300101) AddMessage() *CorporateActionInstructionV01 {
	d.Message = new(CorporateActionInstructionV01)
	return d.Message
}

// Scope
// An account owner sends the CorporateActionInstruction message to an account servicer to instruct election on a corporate action event.
// This message is used to provide the custodian with instructions on how the account owner wishes to proceed with a corporate action event. Instructions include investment decisions regarding the exercise of rights issues, the election of stock or cash when the option is available, and decisions on the conversion or tendering of securities.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionInstructionV01 struct {

	// Information that unambiguously identifies a CorporateActionInstruction message as know by the account owner (or the instructing party acting on its behalf).
	Identification *model.DocumentIdentification12 `xml:"Id"`

	// Identification of a previously sent cancelled instruction document.
	CancelledInstructionIdentification *model.DocumentIdentification15 `xml:"CancInstrId,omitempty"`

	// Identification of a previously sent instruction cancellation request document.
	InstructionCancellationRequestIdentification *model.DocumentIdentification15 `xml:"InstrCxlReqId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification13 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*model.CorporateActionEventReference1 `xml:"EvtsLkg,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation6 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *model.AccountAndBalance2 `xml:"AcctDtls"`

	// Provides information about the beneficial owner of the securities.
	BeneficialOwnerDetails []*model.PartyIdentification33 `xml:"BnfclOwnrDtls,omitempty"`

	// Information about the corporate action instruction.
	CorporateActionInstruction *model.CorporateActionOption5 `xml:"CorpActnInstr"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative7 `xml:"AddtlInf,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (c *CorporateActionInstructionV01) AddIdentification() *model.DocumentIdentification12 {
	c.Identification = new(model.DocumentIdentification12)
	return c.Identification
}

func (c *CorporateActionInstructionV01) AddCancelledInstructionIdentification() *model.DocumentIdentification15 {
	c.CancelledInstructionIdentification = new(model.DocumentIdentification15)
	return c.CancelledInstructionIdentification
}

func (c *CorporateActionInstructionV01) AddInstructionCancellationRequestIdentification() *model.DocumentIdentification15 {
	c.InstructionCancellationRequestIdentification = new(model.DocumentIdentification15)
	return c.InstructionCancellationRequestIdentification
}

func (c *CorporateActionInstructionV01) AddOtherDocumentIdentification() *model.DocumentIdentification13 {
	newValue := new(model.DocumentIdentification13)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionInstructionV01) AddEventsLinkage() *model.CorporateActionEventReference1 {
	newValue := new(model.CorporateActionEventReference1)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionInstructionV01) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation6 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation6)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionInstructionV01) AddAccountDetails() *model.AccountAndBalance2 {
	c.AccountDetails = new(model.AccountAndBalance2)
	return c.AccountDetails
}

func (c *CorporateActionInstructionV01) AddBeneficialOwnerDetails() *model.PartyIdentification33 {
	newValue := new(model.PartyIdentification33)
	c.BeneficialOwnerDetails = append(c.BeneficialOwnerDetails, newValue)
	return newValue
}

func (c *CorporateActionInstructionV01) AddCorporateActionInstruction() *model.CorporateActionOption5 {
	c.CorporateActionInstruction = new(model.CorporateActionOption5)
	return c.CorporateActionInstruction
}

func (c *CorporateActionInstructionV01) AddAdditionalInformation() *model.CorporateActionNarrative7 {
	c.AdditionalInformation = new(model.CorporateActionNarrative7)
	return c.AdditionalInformation
}

func (c *CorporateActionInstructionV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	c.MessageOriginator = new(model.PartyIdentification10Choice)
	return c.MessageOriginator
}

func (c *CorporateActionInstructionV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	c.MessageRecipient = new(model.PartyIdentification10Choice)
	return c.MessageRecipient
}

func (c *CorporateActionInstructionV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	c.Extension = append(c.Extension, newValue)
	return newValue
}
