package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05500101 struct {
	XMLName xml.Name                                           `xml:"urn:iso:std:iso:20022:tech:xsd:setr.055.001.01 Document"`
	Message *SwitchOrderConfirmationCancellationInstructionV01 `xml:"SwtchOrdrConfCxlInstrV01"`
}

func (d *Document05500101) AddMessage() *SwitchOrderConfirmationCancellationInstructionV01 {
	d.Message = new(SwitchOrderConfirmationCancellationInstructionV01)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the SwitchOrderConfirmationCancellationInstruction message to the instructing party, for example, an investment manager or its authorised representative to amend a previously sent SwitchOrderConfirmation message.
// Usage
// The SwitchOrderConfirmationCancellationInstruction message is used to cancel a previously sent SwitchOrderConfirmation.
// The amendment indicator element is used to specify whether the switch order confirmation cancellation is to be followed by a switch order confirmation amendment.
// There are two ways to specify the switch order confirmation cancellation. Either:
// - the business references, for example, OrderReference, DealReference, of the switch order confirmation are quoted, or,
// - all the details of the switch order confirmation (this includes the OrderReference and DealReference) are quoted, but this is not recommended.
// The message identification of the SwitchOrderConfirmation message may also be quoted in PreviousReference.
// It is also possible to instruct the cancellation of the confirmation message by quoting its message identification in PreviousReference, but this is not recommended.
type SwitchOrderConfirmationCancellationInstructionV01 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// References of the switch orders to be cancelled.
	CancellationByReference *model.InvestmentFundOrderExecution1 `xml:"CxlByRef,omitempty"`

	// Common information related to all the switch orders confirmations to be cancelled.
	CancellationByOrderConfirmationDetails *model.SwitchOrderConfirmation1 `xml:"CxlByOrdrConfDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`
}

func (s *SwitchOrderConfirmationCancellationInstructionV01) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SwitchOrderConfirmationCancellationInstructionV01) AddPoolReference() *model.AdditionalReference3 {
	s.PoolReference = new(model.AdditionalReference3)
	return s.PoolReference
}

func (s *SwitchOrderConfirmationCancellationInstructionV01) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SwitchOrderConfirmationCancellationInstructionV01) AddRelatedReference() *model.AdditionalReference3 {
	s.RelatedReference = new(model.AdditionalReference3)
	return s.RelatedReference
}

func (s *SwitchOrderConfirmationCancellationInstructionV01) AddCancellationByReference() *model.InvestmentFundOrderExecution1 {
	s.CancellationByReference = new(model.InvestmentFundOrderExecution1)
	return s.CancellationByReference
}

func (s *SwitchOrderConfirmationCancellationInstructionV01) AddCancellationByOrderConfirmationDetails() *model.SwitchOrderConfirmation1 {
	s.CancellationByOrderConfirmationDetails = new(model.SwitchOrderConfirmation1)
	return s.CancellationByOrderConfirmationDetails
}

func (s *SwitchOrderConfirmationCancellationInstructionV01) AddCopyDetails() *model.CopyInformation2 {
	s.CopyDetails = new(model.CopyInformation2)
	return s.CopyDetails
}
