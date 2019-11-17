package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500103 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.03 Document"`
	Message *TransferInInstructionV03 `xml:"TrfInInstr"`
}

func (d *Document00500103) AddMessage() *TransferInInstructionV03 {
	d.Message = new(TransferInInstructionV03)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative, sends the TransferInInstruction message to the executing party, for example, a transfer agent, to instruct the receipt of a financial instrument, free of payment, on a given date from a specified party.
// This message may also be used to instruct the receipt of a financial instrument, free of payment, from another of the instructing parties own accounts or from a third party.
// Usage
// The TransferInInstruction message is used to instruct the receipt of a financial instrument from another account, either owned by the instructing party or by a third party.
type TransferInInstructionV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference of the linked message that was previously sent.
	PreviousReference *model.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *model.Max35Text `xml:"MstrRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*model.Transfer15 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *model.InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *model.DeliverInformation5 `xml:"SttlmDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferInInstructionV03) AddMessageIdentification() *model.MessageIdentification1 {
	t.MessageIdentification = new(model.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferInInstructionV03) AddPoolReference() *model.AdditionalReference2 {
	t.PoolReference = new(model.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferInInstructionV03) AddPreviousReference() *model.AdditionalReference2 {
	t.PreviousReference = new(model.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferInInstructionV03) AddRelatedReference() *model.AdditionalReference2 {
	t.RelatedReference = new(model.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferInInstructionV03) SetMasterReference(value string) {
	t.MasterReference = (*model.Max35Text)(&value)
}

func (t *TransferInInstructionV03) AddTransferDetails() *model.Transfer15 {
	newValue := new(model.Transfer15)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferInInstructionV03) AddAccountDetails() *model.InvestmentAccount22 {
	t.AccountDetails = new(model.InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferInInstructionV03) AddSettlementDetails() *model.DeliverInformation5 {
	t.SettlementDetails = new(model.DeliverInformation5)
	return t.SettlementDetails
}

func (t *TransferInInstructionV03) AddCopyDetails() *model.CopyInformation2 {
	t.CopyDetails = new(model.CopyInformation2)
	return t.CopyDetails
}

func (t *TransferInInstructionV03) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
