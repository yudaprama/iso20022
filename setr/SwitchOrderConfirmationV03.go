package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01500103 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:setr.015.001.03 Document"`
	Message *SwitchOrderConfirmationV03 `xml:"SwtchOrdrConfV03"`
}

func (d *Document01500103) AddMessage() *SwitchOrderConfirmationV03 {
	d.Message = new(SwitchOrderConfirmationV03)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the SwitchOrderConfirmation message to the instructing party, for example, an investment manager or its authorised representative to confirm the details of the execution of a previously received SwitchOrder instruction.
// Usage
// The SwitchOrderConfirmation message is used to confirm that all the legs of the previously instructed switch order have been executed. The reference of the switch order confirmation is identified in DealReference.
// The reference of the original switch order is specified in OrderReference. The message identification of the SwitchOrder message in which the switch order was conveyed may also be quoted in RelatedReference.
// If the SwitchOrderConfirmation Details sequence is present more than once, then the redemption leg details and subscription leg details sequences may only be present once.
// If Switch Confirmation Order Details\Investment Account is used, then the Investment Account Details sequences in Subscription Leg Details and Redemption Leg Details are not allowed.
// If the SwitchOrderConfirmation Details sequence is present more than once, then the redemption leg details and subscription leg details sequences may only be present once.
// If SwitchOrderConfirmationDetails\InvestmentAccount is used, then the InvestmentAccountDetails sequences in SubscriptionLegDetails and RedemptionLegDetails are not allowed. This functionality is to be used by institutions that set up two accounts per investor, rather than one investment account.
//
type SwitchOrderConfirmationV03 struct {

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

func (s *SwitchOrderConfirmationV03) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SwitchOrderConfirmationV03) AddPoolReference() *model.AdditionalReference3 {
	s.PoolReference = new(model.AdditionalReference3)
	return s.PoolReference
}

func (s *SwitchOrderConfirmationV03) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SwitchOrderConfirmationV03) AddRelatedReference() *model.AdditionalReference3 {
	s.RelatedReference = new(model.AdditionalReference3)
	return s.RelatedReference
}

func (s *SwitchOrderConfirmationV03) AddSwitchExecutionDetails() *model.SwitchExecution4 {
	newValue := new(model.SwitchExecution4)
	s.SwitchExecutionDetails = append(s.SwitchExecutionDetails, newValue)
	return newValue
}

func (s *SwitchOrderConfirmationV03) AddCopyDetails() *model.CopyInformation2 {
	s.CopyDetails = new(model.CopyInformation2)
	return s.CopyDetails
}

func (s *SwitchOrderConfirmationV03) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
