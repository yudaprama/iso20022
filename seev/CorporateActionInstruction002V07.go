package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03300207 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.033.002.07 Document"`
	Message *CorporateActionInstruction002V07 `xml:"CorpActnInstr"`
}

func (d *Document03300207) AddMessage() *CorporateActionInstruction002V07 {
	d.Message = new(CorporateActionInstruction002V07)
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
type CorporateActionInstruction002V07 struct {

	// When used in a corporate action instruction, indicates that the current instruction is replacing a previous one that was cancelled earlier. When used in a corporate action instruction cancellation request, indicates that cancelled instruction will be replaced by a new corporate action instruction to be sent later.
	ChangeInstructionIndicator *model.YesNoIndicator `xml:"ChngInstrInd,omitempty"`

	// Identification of a previously sent cancelled instruction document.
	CancelledInstructionIdentification *model.DocumentIdentification37 `xml:"CancInstrId,omitempty"`

	// Identification of a previously sent instruction cancellation request document.
	InstructionCancellationRequestIdentification *model.DocumentIdentification37 `xml:"InstrCxlReqId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification38 `xml:"OthrDocId,omitempty"`

	// Identification of an other corporate action event that needs to be closely linked to the processing of the event notified in this document.
	EventsLinkage []*model.CorporateActionEventReference4 `xml:"EvtsLkg,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation115 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account, owner and account balance.
	AccountDetails *model.AccountAndBalance37 `xml:"AcctDtls"`

	// Provides information about the beneficial owner of the securities.
	BeneficialOwnerDetails []*model.PartyIdentification101 `xml:"BnfclOwnrDtls,omitempty"`

	// Information about the corporate action instruction.
	CorporateActionInstruction *model.CorporateActionOption134 `xml:"CorpActnInstr"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative34 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionInstruction002V07) SetChangeInstructionIndicator(value string) {
	c.ChangeInstructionIndicator = (*model.YesNoIndicator)(&value)
}

func (c *CorporateActionInstruction002V07) AddCancelledInstructionIdentification() *model.DocumentIdentification37 {
	c.CancelledInstructionIdentification = new(model.DocumentIdentification37)
	return c.CancelledInstructionIdentification
}

func (c *CorporateActionInstruction002V07) AddInstructionCancellationRequestIdentification() *model.DocumentIdentification37 {
	c.InstructionCancellationRequestIdentification = new(model.DocumentIdentification37)
	return c.InstructionCancellationRequestIdentification
}

func (c *CorporateActionInstruction002V07) AddOtherDocumentIdentification() *model.DocumentIdentification38 {
	newValue := new(model.DocumentIdentification38)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionInstruction002V07) AddEventsLinkage() *model.CorporateActionEventReference4 {
	newValue := new(model.CorporateActionEventReference4)
	c.EventsLinkage = append(c.EventsLinkage, newValue)
	return newValue
}

func (c *CorporateActionInstruction002V07) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation115 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation115)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionInstruction002V07) AddAccountDetails() *model.AccountAndBalance37 {
	c.AccountDetails = new(model.AccountAndBalance37)
	return c.AccountDetails
}

func (c *CorporateActionInstruction002V07) AddBeneficialOwnerDetails() *model.PartyIdentification101 {
	newValue := new(model.PartyIdentification101)
	c.BeneficialOwnerDetails = append(c.BeneficialOwnerDetails, newValue)
	return newValue
}

func (c *CorporateActionInstruction002V07) AddCorporateActionInstruction() *model.CorporateActionOption134 {
	c.CorporateActionInstruction = new(model.CorporateActionOption134)
	return c.CorporateActionInstruction
}

func (c *CorporateActionInstruction002V07) AddAdditionalInformation() *model.CorporateActionNarrative34 {
	c.AdditionalInformation = new(model.CorporateActionNarrative34)
	return c.AdditionalInformation
}

func (c *CorporateActionInstruction002V07) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
