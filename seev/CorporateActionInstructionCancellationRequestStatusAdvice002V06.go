package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04100206 struct {
	XMLName xml.Name                                                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.002.06 Document"`
	Message *CorporateActionInstructionCancellationRequestStatusAdvice002V06 `xml:"CorpActnInstrCxlReqStsAdvc"`
}

func (d *Document04100206) AddMessage() *CorporateActionInstructionCancellationRequestStatusAdvice002V06 {
	d.Message = new(CorporateActionInstructionCancellationRequestStatusAdvice002V06)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionInstructionCancellationRequestStatusAdvice message to an account owner or its designated agent to report status of a previously received CorporateActionInstructionCancellationRequest message sent by the account owner. This will include the acknowledgement/rejection of a request to cancel an outstanding instruction.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionInstructionCancellationRequestStatusAdvice002V06 struct {

	// Identification of a related instruction cancellation request document.
	InstructionCancellationRequestIdentification *model.DocumentIdentification17 `xml:"InstrCxlReqId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification34 `xml:"OthrDocId,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation94 `xml:"CorpActnGnlInf"`

	// Provides information about the processing status of the instruction cancellation request.
	InstructionCancellationRequestStatus []*model.InstructionCancellationRequestStatus10Choice `xml:"InstrCxlReqSts"`

	// Information about the corporate action option.
	CorporateActionInstruction *model.CorporateActionOption121 `xml:"CorpActnInstr,omitempty"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative19 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionInstructionCancellationRequestStatusAdvice002V06) AddInstructionCancellationRequestIdentification() *model.DocumentIdentification17 {
	c.InstructionCancellationRequestIdentification = new(model.DocumentIdentification17)
	return c.InstructionCancellationRequestIdentification
}

func (c *CorporateActionInstructionCancellationRequestStatusAdvice002V06) AddOtherDocumentIdentification() *model.DocumentIdentification34 {
	newValue := new(model.DocumentIdentification34)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionInstructionCancellationRequestStatusAdvice002V06) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation94 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation94)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionInstructionCancellationRequestStatusAdvice002V06) AddInstructionCancellationRequestStatus() *model.InstructionCancellationRequestStatus10Choice {
	newValue := new(model.InstructionCancellationRequestStatus10Choice)
	c.InstructionCancellationRequestStatus = append(c.InstructionCancellationRequestStatus, newValue)
	return newValue
}

func (c *CorporateActionInstructionCancellationRequestStatusAdvice002V06) AddCorporateActionInstruction() *model.CorporateActionOption121 {
	c.CorporateActionInstruction = new(model.CorporateActionOption121)
	return c.CorporateActionInstruction
}

func (c *CorporateActionInstructionCancellationRequestStatusAdvice002V06) AddAdditionalInformation() *model.CorporateActionNarrative19 {
	c.AdditionalInformation = new(model.CorporateActionNarrative19)
	return c.AdditionalInformation
}

func (c *CorporateActionInstructionCancellationRequestStatusAdvice002V06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
