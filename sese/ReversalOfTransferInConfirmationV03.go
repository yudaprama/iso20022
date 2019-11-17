package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00800103 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.008.001.03 Document"`
	Message *ReversalOfTransferInConfirmationV03 `xml:"RvslOfTrfInConf"`
}

func (d *Document00800103) AddMessage() *ReversalOfTransferInConfirmationV03 {
	d.Message = new(ReversalOfTransferInConfirmationV03)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the ReversalOfTransferInConfirmation message to the instructing party, for example, an investment manager or its authorised representative, to cancel a previously sent TransferInConfirmation message.
// Usage
// The ReversalOfTransferInConfirmation message is used to reverse a previously sent TransferInConfirmation.
// There are two ways to specify the reversal of the transfer in confirmation. Either:
// - the business references, for example, TransferReference, TransferConfirmationIdentification, of the transfer confirmation are quoted, or,
// - all the details of the transfer confirmation (this includes TransferReference and TransferConfirmationIdentification) are quoted but this is not recommended.
// The message identification of the TransferInConfirmation message in which the transfer confirmation was conveyed may also be quoted in PreviousReference.
// The message reference (MessageIdentification) of the TransferInInstruction message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
// It is also possible to request a reversal of a TransferInConfirmation by quoting its message reference (MessageIdentification) in PreviousReference.
type ReversalOfTransferInConfirmationV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to the transaction identifier issued by the counterparty. Building block may also be used to reference a previous transaction, or tie a set of messages together.
	References []*model.References11 `xml:"Refs"`

	// Reference of the transfer in confirmation to be reversed.
	ReversalByReference *model.TransferReference2 `xml:"RvslByRef,omitempty"`

	// Copy of the transfer in confirmation to reverse.
	ReversalByTransferInConfirmationDetails *model.TransferIn6 `xml:"RvslByTrfInConfDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`
}

func (r *ReversalOfTransferInConfirmationV03) AddMessageIdentification() *model.MessageIdentification1 {
	r.MessageIdentification = new(model.MessageIdentification1)
	return r.MessageIdentification
}

func (r *ReversalOfTransferInConfirmationV03) AddReferences() *model.References11 {
	newValue := new(model.References11)
	r.References = append(r.References, newValue)
	return newValue
}

func (r *ReversalOfTransferInConfirmationV03) AddReversalByReference() *model.TransferReference2 {
	r.ReversalByReference = new(model.TransferReference2)
	return r.ReversalByReference
}

func (r *ReversalOfTransferInConfirmationV03) AddReversalByTransferInConfirmationDetails() *model.TransferIn6 {
	r.ReversalByTransferInConfirmationDetails = new(model.TransferIn6)
	return r.ReversalByTransferInConfirmationDetails
}

func (r *ReversalOfTransferInConfirmationV03) AddCopyDetails() *model.CopyInformation2 {
	r.CopyDetails = new(model.CopyInformation2)
	return r.CopyDetails
}
