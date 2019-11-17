package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01500104 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:setr.015.001.04 Document"`
	Message *SwitchOrderConfirmationV04 `xml:"SwtchOrdrConf"`
}

func (d *Document01500104) AddMessage() *SwitchOrderConfirmationV04 {
	d.Message = new(SwitchOrderConfirmationV04)
	return d.Message
}

// Scope
// The SwitchOrderConfirmation message is sent by an executing party, for example, a transfer agent, to the instructing party, for example, an investment manager or its authorised representative, to confirm the details of the execution of a previously received SwitchOrder instruction.
// Usage
// The SwitchOrderConfirmation message is used to confirm that all the legs of the previously instructed switch order have been executed. The reference of the switch order confirmation is identified in DealReference.
// The reference of the original switch order is specified in OrderReference. The message identification of the SwitchOrder message in which the switch order was conveyed may also be quoted in RelatedReference but this is not recommended.
// When the message is used to convey a confirmation amendment/s, the AmendmentIndicator must be present with the value ‘true’ or ‘1’. When this is the case, the message must only contain a confirmation amendment/s and not contain both a confirmation amendment/s and a ‘new’ confirmation/s.
type SwitchOrderConfirmationV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference9 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference8 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference8 `xml:"RltdRef,omitempty"`

	// Information related to a switch execution.
	SwitchExecutionDetails []*model.SwitchExecution7 `xml:"SwtchExctnDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation4 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SwitchOrderConfirmationV04) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SwitchOrderConfirmationV04) AddPoolReference() *model.AdditionalReference9 {
	s.PoolReference = new(model.AdditionalReference9)
	return s.PoolReference
}

func (s *SwitchOrderConfirmationV04) AddPreviousReference() *model.AdditionalReference8 {
	newValue := new(model.AdditionalReference8)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SwitchOrderConfirmationV04) AddRelatedReference() *model.AdditionalReference8 {
	s.RelatedReference = new(model.AdditionalReference8)
	return s.RelatedReference
}

func (s *SwitchOrderConfirmationV04) AddSwitchExecutionDetails() *model.SwitchExecution7 {
	newValue := new(model.SwitchExecution7)
	s.SwitchExecutionDetails = append(s.SwitchExecutionDetails, newValue)
	return newValue
}

func (s *SwitchOrderConfirmationV04) AddCopyDetails() *model.CopyInformation4 {
	s.CopyDetails = new(model.CopyInformation4)
	return s.CopyDetails
}

func (s *SwitchOrderConfirmationV04) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
