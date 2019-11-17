package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00800105 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.008.001.05 Document"`
	Message *ReversalOfTransferInConfirmationV05 `xml:"RvslOfTrfInConf"`
}

func (d *Document00800105) AddMessage() *ReversalOfTransferInConfirmationV05 {
	d.Message = new(ReversalOfTransferInConfirmationV05)
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
//
type ReversalOfTransferInConfirmationV05 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to the transaction identifier issued by the counterparty. Building block may also be used to reference a previous transaction, or tie a set of messages together.
	References []*model.References15 `xml:"Refs,omitempty"`

	// Choice between reversal by reference or by reversal details.
	Reversal *model.Reversal4Choice `xml:"Rvsl"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`
}

func (r *ReversalOfTransferInConfirmationV05) AddMessageIdentification() *model.MessageIdentification1 {
	r.MessageIdentification = new(model.MessageIdentification1)
	return r.MessageIdentification
}

func (r *ReversalOfTransferInConfirmationV05) AddReferences() *model.References15 {
	newValue := new(model.References15)
	r.References = append(r.References, newValue)
	return newValue
}

func (r *ReversalOfTransferInConfirmationV05) AddReversal() *model.Reversal4Choice {
	r.Reversal = new(model.Reversal4Choice)
	return r.Reversal
}

func (r *ReversalOfTransferInConfirmationV05) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	r.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return r.MarketPracticeVersion
}

func (r *ReversalOfTransferInConfirmationV05) AddCopyDetails() *model.CopyInformation2 {
	r.CopyDetails = new(model.CopyInformation2)
	return r.CopyDetails
}
