package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400102 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.004.001.02 Document"`
	Message *ReversalOfTransferOutConfirmationV02 `xml:"RvslOfTrfOutConfV02"`
}

func (d *Document00400102) AddMessage() *ReversalOfTransferOutConfirmationV02 {
	d.Message = new(ReversalOfTransferOutConfirmationV02)
	return d.Message
}

// Scope
// An executing party, eg, a transfer agent, sends the ReversalOfTransferOutConfirmation message to the instructing party, eg, an investment manager or its authorised representative, to cancel a previously sent TransferOutConfirmation message.
// Usage
// The ReversalOfTransferOutConfirmation message is used to reverse a previously sent TransferOutConfirmation.
// There are two ways to specify the reversal of the transfer out confirmation. Either:
// - the business references, eg, TransferReference, TransferConfirmationIdentification, of the transfer confirmation are quoted, or,
// - all the details of the transfer confirmation (this includes TransferReference and TransferConfirmationIdentification) are quoted but this is not recommended.
// The message identification of the TransferOutConfirmation message in which the transfer out confirmation was conveyed may also be quoted in PreviousReference. The message identification of the TransferOutInstruction message in which the transfer out instruction was conveyed may also be quoted in RelatedReference.
type ReversalOfTransferOutConfirmationV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Reference of the transfer out confirmation to be reversed.
	ReversalByReference *model.TransferReference2 `xml:"RvslByRef,omitempty"`

	// Copy of the transfer out confirmation to reverse.
	ReversalByTransferOutConfirmationDetails *model.TransferOut6 `xml:"RvslByTrfOutConfDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`
}

func (r *ReversalOfTransferOutConfirmationV02) AddMessageIdentification() *model.MessageIdentification1 {
	r.MessageIdentification = new(model.MessageIdentification1)
	return r.MessageIdentification
}

func (r *ReversalOfTransferOutConfirmationV02) AddPreviousReference() *model.AdditionalReference2 {
	r.PreviousReference = new(model.AdditionalReference2)
	return r.PreviousReference
}

func (r *ReversalOfTransferOutConfirmationV02) AddPoolReference() *model.AdditionalReference2 {
	r.PoolReference = new(model.AdditionalReference2)
	return r.PoolReference
}

func (r *ReversalOfTransferOutConfirmationV02) AddRelatedReference() *model.AdditionalReference2 {
	r.RelatedReference = new(model.AdditionalReference2)
	return r.RelatedReference
}

func (r *ReversalOfTransferOutConfirmationV02) AddReversalByReference() *model.TransferReference2 {
	r.ReversalByReference = new(model.TransferReference2)
	return r.ReversalByReference
}

func (r *ReversalOfTransferOutConfirmationV02) AddReversalByTransferOutConfirmationDetails() *model.TransferOut6 {
	r.ReversalByTransferOutConfirmationDetails = new(model.TransferOut6)
	return r.ReversalByTransferOutConfirmationDetails
}

func (r *ReversalOfTransferOutConfirmationV02) AddCopyDetails() *model.CopyInformation2 {
	r.CopyDetails = new(model.CopyInformation2)
	return r.CopyDetails
}
