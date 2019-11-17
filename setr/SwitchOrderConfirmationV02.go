package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01500102 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:setr.015.001.02 Document"`
	Message *SwitchOrderConfirmationV02 `xml:"setr.015.001.02"`
}

func (d *Document01500102) AddMessage() *SwitchOrderConfirmationV02 {
	d.Message = new(SwitchOrderConfirmationV02)
	return d.Message
}

// Scope
// The SwitchOrderConfirmation message is sent by an executing party, eg, a transfer agent, to an instruction party, eg, an investment manager or its authorised representative. There may be one or more intermediary parties between the instructing party and the executing party. The intermediary party is, for example, an intermediary or a concentrator.
// This message is used to confirm the details of the execution of a SwitchOrder message.
// Usage
// The SwitchOrderConfirmation message is sent to confirm that all the legs of the switch have been executed.
type SwitchOrderConfirmationV02 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *model.AdditionalReference3 `xml:"MstrRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef"`

	// Information related to a switch execution.
	SwitchExecutionDetails *model.SwitchExecution3 `xml:"SwtchExctnDtls"`

	// Confirmation of the information related to an intermediary.
	IntermediaryDetails []*model.Intermediary4 `xml:"IntrmyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation1 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SwitchOrderConfirmationV02) AddMasterReference() *model.AdditionalReference3 {
	s.MasterReference = new(model.AdditionalReference3)
	return s.MasterReference
}

func (s *SwitchOrderConfirmationV02) AddPoolReference() *model.AdditionalReference3 {
	s.PoolReference = new(model.AdditionalReference3)
	return s.PoolReference
}

func (s *SwitchOrderConfirmationV02) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SwitchOrderConfirmationV02) AddRelatedReference() *model.AdditionalReference3 {
	s.RelatedReference = new(model.AdditionalReference3)
	return s.RelatedReference
}

func (s *SwitchOrderConfirmationV02) AddSwitchExecutionDetails() *model.SwitchExecution3 {
	s.SwitchExecutionDetails = new(model.SwitchExecution3)
	return s.SwitchExecutionDetails
}

func (s *SwitchOrderConfirmationV02) AddIntermediaryDetails() *model.Intermediary4 {
	newValue := new(model.Intermediary4)
	s.IntermediaryDetails = append(s.IntermediaryDetails, newValue)
	return newValue
}

func (s *SwitchOrderConfirmationV02) AddCopyDetails() *model.CopyInformation1 {
	s.CopyDetails = new(model.CopyInformation1)
	return s.CopyDetails
}

func (s *SwitchOrderConfirmationV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
