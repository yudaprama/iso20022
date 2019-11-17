package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04000106 struct {
	XMLName xml.Name                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.040.001.06 Document"`
	Message *CorporateActionInstructionCancellationRequestV06 `xml:"CorpActnInstrCxlReq"`
}

func (d *Document04000106) AddMessage() *CorporateActionInstructionCancellationRequestV06 {
	d.Message = new(CorporateActionInstructionCancellationRequestV06)
	return d.Message
}

// Scope
// An account owner sends the CorporateActionInstructionCancellationRequest message to an account servicer to request cancellation of a previously sent corporate action election instruction.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionInstructionCancellationRequestV06 struct {

	// When used in a corporate action instruction, indicates that the current instruction is replacing a previous one that was cancelled earlier. When used in a corporate action instruction cancellation request, indicates that cancelled instruction will be replaced by a new corporate action instruction to be sent later.
	ChangeInstructionIndicator *model.YesNoIndicator `xml:"ChngInstrInd,omitempty"`

	// Identification of a previously sent instruction document.
	InstructionIdentification *model.DocumentIdentification31 `xml:"InstrId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation90 `xml:"CorpActnGnlInf"`

	// General information about the safekeeping account and the account owner.
	AccountDetails *model.AccountIdentification31 `xml:"AcctDtls"`

	// Information about the corporate action option.
	CorporateActionInstruction *model.CorporateActionOption120 `xml:"CorpActnInstr"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionInstructionCancellationRequestV06) SetChangeInstructionIndicator(value string) {
	c.ChangeInstructionIndicator = (*model.YesNoIndicator)(&value)
}

func (c *CorporateActionInstructionCancellationRequestV06) AddInstructionIdentification() *model.DocumentIdentification31 {
	c.InstructionIdentification = new(model.DocumentIdentification31)
	return c.InstructionIdentification
}

func (c *CorporateActionInstructionCancellationRequestV06) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation90 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation90)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionInstructionCancellationRequestV06) AddAccountDetails() *model.AccountIdentification31 {
	c.AccountDetails = new(model.AccountIdentification31)
	return c.AccountDetails
}

func (c *CorporateActionInstructionCancellationRequestV06) AddCorporateActionInstruction() *model.CorporateActionOption120 {
	c.CorporateActionInstruction = new(model.CorporateActionOption120)
	return c.CorporateActionInstruction
}

func (c *CorporateActionInstructionCancellationRequestV06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
