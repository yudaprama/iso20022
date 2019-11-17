package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05600101 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:setr.056.001.01 Document"`
	Message *SwitchOrderConfirmationAmendmentV01 `xml:"SwtchOrdrConfAmdmntV01"`
}

func (d *Document05600101) AddMessage() *SwitchOrderConfirmationAmendmentV01 {
	d.Message = new(SwitchOrderConfirmationAmendmentV01)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the SwitchOrderConfirmationAmendment message to the instructing party, for example, an investment manager or its authorised representative to amend a previously sent SwitchOrderConfirmation message.
// Usage
// The SwitchOrderConfirmationAmendment message is used to amend a previously sent switch order confirmation.
// Each order confirmation amendment specified is identified in DealReference. The reference of the original order is specified in OrderReference.
// The message identification of the SwitchOrder message in which the switch order was conveyed may also be quoted in RelatedReference. The message identification of the SwitchOrderConfirmation message in which the original switch order confirmation was conveyed may also be quoted in PreviousReference.
type SwitchOrderConfirmationAmendmentV01 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Information related to a switch execution.
	SwitchExecutionDetails []*model.SwitchExecution4 `xml:"SwtchExctnDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SwitchOrderConfirmationAmendmentV01) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SwitchOrderConfirmationAmendmentV01) AddPoolReference() *model.AdditionalReference3 {
	s.PoolReference = new(model.AdditionalReference3)
	return s.PoolReference
}

func (s *SwitchOrderConfirmationAmendmentV01) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SwitchOrderConfirmationAmendmentV01) AddRelatedReference() *model.AdditionalReference3 {
	s.RelatedReference = new(model.AdditionalReference3)
	return s.RelatedReference
}

func (s *SwitchOrderConfirmationAmendmentV01) AddSwitchExecutionDetails() *model.SwitchExecution4 {
	newValue := new(model.SwitchExecution4)
	s.SwitchExecutionDetails = append(s.SwitchExecutionDetails, newValue)
	return newValue
}

func (s *SwitchOrderConfirmationAmendmentV01) AddCopyDetails() *model.CopyInformation2 {
	s.CopyDetails = new(model.CopyInformation2)
	return s.CopyDetails
}

func (s *SwitchOrderConfirmationAmendmentV01) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
