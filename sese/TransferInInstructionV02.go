package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500102 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.005.001.02 Document"`
	Message *TransferInInstructionV02 `xml:"TrfInInstrV02"`
}

func (d *Document00500102) AddMessage() *TransferInInstructionV02 {
	d.Message = new(TransferInInstructionV02)
	return d.Message
}

// Scope
// An instructing party, eg, an investment manager or its authorised representative, sends the TransferInInstruction message to the executing party, eg, a transfer agent, to instruct the receipt of a financial instrument, free of payment, on a given date from a specified party.
// This message may also be used to instruct the receipt of a financial instrument, free of payment, from another of the instructing parties own accounts or from a third party.
// Usage
// The TransferInInstruction message is used to instruct the receipt of a financial instrument from another account, either owned by the instructing party or by a third party.
type TransferInInstructionV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference of the linked message which was previously sent.
	PreviousReference *model.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails *model.Transfer5 `xml:"TrfDtls"`

	// Information related to the financial instrument to be received.
	FinancialInstrumentDetails *model.FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *model.InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *model.DeliverInformation3 `xml:"SttlmDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferInInstructionV02) AddMessageIdentification() *model.MessageIdentification1 {
	t.MessageIdentification = new(model.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferInInstructionV02) AddPoolReference() *model.AdditionalReference2 {
	t.PoolReference = new(model.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferInInstructionV02) AddPreviousReference() *model.AdditionalReference2 {
	t.PreviousReference = new(model.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferInInstructionV02) AddRelatedReference() *model.AdditionalReference2 {
	t.RelatedReference = new(model.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferInInstructionV02) AddTransferDetails() *model.Transfer5 {
	t.TransferDetails = new(model.Transfer5)
	return t.TransferDetails
}

func (t *TransferInInstructionV02) AddFinancialInstrumentDetails() *model.FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(model.FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *TransferInInstructionV02) AddAccountDetails() *model.InvestmentAccount22 {
	t.AccountDetails = new(model.InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferInInstructionV02) AddSettlementDetails() *model.DeliverInformation3 {
	t.SettlementDetails = new(model.DeliverInformation3)
	return t.SettlementDetails
}

func (t *TransferInInstructionV02) AddCopyDetails() *model.CopyInformation2 {
	t.CopyDetails = new(model.CopyInformation2)
	return t.CopyDetails
}

func (t *TransferInInstructionV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
