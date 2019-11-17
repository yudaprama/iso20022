package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300102 struct {
	XMLName xml.Name        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.013.001.02 Document"`
	Message *SwitchOrderV02 `xml:"setr.013.001.02"`
}

func (d *Document01300102) AddMessage() *SwitchOrderV02 {
	d.Message = new(SwitchOrderV02)
	return d.Message
}

// Scope
// The SwitchOrder message is sent by an instructing party, eg, an investment manager or its authorised representative, to an executing party, eg, a transfer agent. There may be one or more intermediary parties between the instructing party and the executing party. The intermediary party is, for example, an intermediary or a concentrator.
// This message is used to instruct the executing party to switch from a specified amount/quantity of specified financial instruments to a specified amount/quantity of different financial instruments.
// Usage
// The SwitchOrder message is used when the instructing party, ie, an investor, wants to change its investments within the same fund family, according to the terms of the prospectus.
type SwitchOrderV02 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *model.AdditionalReference3 `xml:"MstrRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Information related to the switch order
	SwitchOrderDetails *model.SwitchOrder2 `xml:"SwtchOrdrDtls"`

	// The information related to an intermediary.
	IntermediaryDetails []*model.Intermediary4 `xml:"IntrmyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation1 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SwitchOrderV02) AddMasterReference() *model.AdditionalReference3 {
	s.MasterReference = new(model.AdditionalReference3)
	return s.MasterReference
}

func (s *SwitchOrderV02) AddPoolReference() *model.AdditionalReference3 {
	s.PoolReference = new(model.AdditionalReference3)
	return s.PoolReference
}

func (s *SwitchOrderV02) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SwitchOrderV02) AddSwitchOrderDetails() *model.SwitchOrder2 {
	s.SwitchOrderDetails = new(model.SwitchOrder2)
	return s.SwitchOrderDetails
}

func (s *SwitchOrderV02) AddIntermediaryDetails() *model.Intermediary4 {
	newValue := new(model.Intermediary4)
	s.IntermediaryDetails = append(s.IntermediaryDetails, newValue)
	return newValue
}

func (s *SwitchOrderV02) AddCopyDetails() *model.CopyInformation1 {
	s.CopyDetails = new(model.CopyInformation1)
	return s.CopyDetails
}

func (s *SwitchOrderV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
