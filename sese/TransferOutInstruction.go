package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100101 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.001.001.01 Document"`
	Message *TransferOutInstruction `xml:"sese.001.001.01"`
}

func (d *Document00100101) AddMessage() *TransferOutInstruction {
	d.Message = new(TransferOutInstruction)
	return d.Message
}

// Scope
// The TransferOutInstruction message is sent by an instructing party, or an instructing party's designated agent, to the executing party.
// This message is used to instruct the delivery of a financial instrument, free of payment, at a given date, to a specified party. This message can be used to instruct the transfer of a financial instrument to an own account or to a third party.
// Usage
// The TransferOutInstruction message is used by an instructing party to instruct the executing party to withdraw a financial instrument from one account and deliver it to either another account or to a third party.
type TransferOutInstruction struct {

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference of the linked message that was previously sent.
	PreviousReference *model.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails *model.Transfer1 `xml:"TrfDtls"`

	// Information related to the financial instrument to be withdrawn.
	FinancialInstrumentDetails *model.FinancialInstrument3 `xml:"FinInstrmDtls"`

	// Information related to the account from which the financial instrument is to be withdrawn.
	AccountDetails *model.InvestmentAccount10 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *model.ReceiveInformation1 `xml:"SttlmDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOutInstruction) AddPoolReference() *model.AdditionalReference2 {
	t.PoolReference = new(model.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferOutInstruction) AddPreviousReference() *model.AdditionalReference2 {
	t.PreviousReference = new(model.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferOutInstruction) AddRelatedReference() *model.AdditionalReference2 {
	t.RelatedReference = new(model.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferOutInstruction) AddTransferDetails() *model.Transfer1 {
	t.TransferDetails = new(model.Transfer1)
	return t.TransferDetails
}

func (t *TransferOutInstruction) AddFinancialInstrumentDetails() *model.FinancialInstrument3 {
	t.FinancialInstrumentDetails = new(model.FinancialInstrument3)
	return t.FinancialInstrumentDetails
}

func (t *TransferOutInstruction) AddAccountDetails() *model.InvestmentAccount10 {
	t.AccountDetails = new(model.InvestmentAccount10)
	return t.AccountDetails
}

func (t *TransferOutInstruction) AddSettlementDetails() *model.ReceiveInformation1 {
	t.SettlementDetails = new(model.ReceiveInformation1)
	return t.SettlementDetails
}

func (t *TransferOutInstruction) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
