package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03400106 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.034.001.06 Document"`
	Message *CorporateActionInstructionStatusAdviceV06 `xml:"CorpActnInstrStsAdvc"`
}

func (d *Document03400106) AddMessage() *CorporateActionInstructionStatusAdviceV06 {
	d.Message = new(CorporateActionInstructionStatusAdviceV06)
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
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionInstructionStatusAdviceV06 struct {

	// Identification of a related instruction document.
	InstructionIdentification *model.DocumentIdentification9 `xml:"InstrId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification14 `xml:"OthrDocId,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation52 `xml:"CorpActnGnlInf"`

	// Provides information about the processing status of the instruction.
	InstructionProcessingStatus []*model.InstructionProcessingStatus20Choice `xml:"InstrPrcgSts"`

	// Information about the corporate action instruction.
	CorporateActionInstruction *model.CorporateActionOption41 `xml:"CorpActnInstr,omitempty"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative10 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionInstructionStatusAdviceV06) AddInstructionIdentification() *model.DocumentIdentification9 {
	c.InstructionIdentification = new(model.DocumentIdentification9)
	return c.InstructionIdentification
}

func (c *CorporateActionInstructionStatusAdviceV06) AddOtherDocumentIdentification() *model.DocumentIdentification14 {
	newValue := new(model.DocumentIdentification14)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionInstructionStatusAdviceV06) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation52 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation52)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionInstructionStatusAdviceV06) AddInstructionProcessingStatus() *model.InstructionProcessingStatus20Choice {
	newValue := new(model.InstructionProcessingStatus20Choice)
	c.InstructionProcessingStatus = append(c.InstructionProcessingStatus, newValue)
	return newValue
}

func (c *CorporateActionInstructionStatusAdviceV06) AddCorporateActionInstruction() *model.CorporateActionOption41 {
	c.CorporateActionInstruction = new(model.CorporateActionOption41)
	return c.CorporateActionInstruction
}

func (c *CorporateActionInstructionStatusAdviceV06) AddAdditionalInformation() *model.CorporateActionNarrative10 {
	c.AdditionalInformation = new(model.CorporateActionNarrative10)
	return c.AdditionalInformation
}

func (c *CorporateActionInstructionStatusAdviceV06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
