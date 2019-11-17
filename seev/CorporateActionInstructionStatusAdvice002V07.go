package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03400207 struct {
	XMLName xml.Name                                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.002.07 Document"`
	Message *CorporateActionInstructionStatusAdvice002V07 `xml:"CorpActnInstrStsAdvc"`
}

func (d *Document03400207) AddMessage() *CorporateActionInstructionStatusAdvice002V07 {
	d.Message = new(CorporateActionInstructionStatusAdvice002V07)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionInstructionStatusAdvice message to an account owner or its designated agent, to report status of a received corporate action election instruction.
// This message is used to advise the status, or a change in status, of a corporate action-related transaction previously instructed by, or executed on behalf of, the account owner. This will include the acknowledgement/rejection of a corporate action instruction.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionInstructionStatusAdvice002V07 struct {

	// Identification of a related instruction document.
	InstructionIdentification *model.DocumentIdentification17 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification34 `xml:"OthrDocId,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation94 `xml:"CorpActnGnlInf"`

	// Provides information about the processing status of the instruction.
	InstructionProcessingStatus []*model.InstructionProcessingStatus25Choice `xml:"InstrPrcgSts"`

	// Information about the corporate action instruction.
	CorporateActionInstruction *model.CorporateActionOption121 `xml:"CorpActnInstr,omitempty"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative19 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionInstructionStatusAdvice002V07) AddInstructionIdentification() *model.DocumentIdentification17 {
	c.InstructionIdentification = new(model.DocumentIdentification17)
	return c.InstructionIdentification
}

func (c *CorporateActionInstructionStatusAdvice002V07) AddOtherDocumentIdentification() *model.DocumentIdentification34 {
	newValue := new(model.DocumentIdentification34)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionInstructionStatusAdvice002V07) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation94 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation94)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionInstructionStatusAdvice002V07) AddInstructionProcessingStatus() *model.InstructionProcessingStatus25Choice {
	newValue := new(model.InstructionProcessingStatus25Choice)
	c.InstructionProcessingStatus = append(c.InstructionProcessingStatus, newValue)
	return newValue
}

func (c *CorporateActionInstructionStatusAdvice002V07) AddCorporateActionInstruction() *model.CorporateActionOption121 {
	c.CorporateActionInstruction = new(model.CorporateActionOption121)
	return c.CorporateActionInstruction
}

func (c *CorporateActionInstructionStatusAdvice002V07) AddAdditionalInformation() *model.CorporateActionNarrative19 {
	c.AdditionalInformation = new(model.CorporateActionNarrative19)
	return c.AdditionalInformation
}

func (c *CorporateActionInstructionStatusAdvice002V07) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
