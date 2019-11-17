package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03300107 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.001.07 Document"`
	Message *CorporateActionInstructionV07 `xml:"CorpActnInstr"`
}

func (d *Document03300107) AddMessage() *CorporateActionInstructionV07 {
	d.Message = new(CorporateActionInstructionV07)
	return d.Message
}

// Scope
// An account owner sends the CorporateActionInstruction message to an account servicer to instruct election on a corporate action event.
// This message is used to provide the custodian with instructions on how the account owner wishes to proceed with a corporate action event. Instructions include investment decisions regarding the exercise of rights issues, the election of stock or cash when the option is available, and decisions on the conversion or tendering of securities.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionInstructionV07 struct {

	// When used in a corporate action instruction, indicates that the current instruction is replacing a previous one that was cancelled earlier. When used in a corporate action instruction cancellation request, indicates that cancelled instruction will be replaced by a new corporate action instruction to be sent later.
	ChangeInstructionIndicator *model.YesNoIndicator `xml:"ChngInstrInd,omitempty"`

	// Identification of a previously sent cancelled instruction document.
	CancelledInstructionIdentification *model.DocumentIdentification31 `xml:"CancInstrId,omitempty"`

	// Identification of a previously sent instruction cancellation request document.
	InstructionCancellationRequestIdentification *model.DocumentIdentification31 `xml:"InstrCxlReqId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification32 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*model.CorporateActionEventReference3 `xml:"EvtsLkg,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation107 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *model.AccountAndBalance35 `xml:"AcctDtls"`

	// Provides information about the beneficial owner of the securities.
	BeneficialOwnerDetails []*model.PartyIdentification93 `xml:"BnfclOwnrDtls,omitempty"`

	// Information about the corporate action instruction.
	CorporateActionInstruction *model.CorporateActionOption131 `xml:"CorpActnInstr"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative30 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionInstructionV07) SetChangeInstructionIndicator(value string) {
	c.ChangeInstructionIndicator = (*model.YesNoIndicator)(&value)
}

func (c *CorporateActionInstructionV07) AddCancelledInstructionIdentification() *model.DocumentIdentification31 {
	c.CancelledInstructionIdentification = new(model.DocumentIdentification31)
	return c.CancelledInstructionIdentification
}

func (c *CorporateActionInstructionV07) AddInstructionCancellationRequestIdentification() *model.DocumentIdentification31 {
	c.InstructionCancellationRequestIdentification = new(model.DocumentIdentification31)
	return c.InstructionCancellationRequestIdentification
}

func (c *CorporateActionInstructionV07) AddOtherDocumentIdentification() *model.DocumentIdentification32 {
	newValue := new(model.DocumentIdentification32)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionInstructionV07) AddEventsLinkage() *model.CorporateActionEventReference3 {
	newValue := new(model.CorporateActionEventReference3)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionInstructionV07) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation107 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation107)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionInstructionV07) AddAccountDetails() *model.AccountAndBalance35 {
	c.AccountDetails = new(model.AccountAndBalance35)
	return c.AccountDetails
}

func (c *CorporateActionInstructionV07) AddBeneficialOwnerDetails() *model.PartyIdentification93 {
	newValue := new(model.PartyIdentification93)
	c.BeneficialOwnerDetails = append(c.BeneficialOwnerDetails, newValue)
	return newValue
}

func (c *CorporateActionInstructionV07) AddCorporateActionInstruction() *model.CorporateActionOption131 {
	c.CorporateActionInstruction = new(model.CorporateActionOption131)
	return c.CorporateActionInstruction
}

func (c *CorporateActionInstructionV07) AddAdditionalInformation() *model.CorporateActionNarrative30 {
	c.AdditionalInformation = new(model.CorporateActionNarrative30)
	return c.AdditionalInformation
}

func (c *CorporateActionInstructionV07) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
