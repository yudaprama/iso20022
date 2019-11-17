package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700107 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.007.001.07 Document"`
	Message *TransferInConfirmationV07 `xml:"TrfInConf"`
}

func (d *Document00700107) AddMessage() *TransferInConfirmationV07 {
	d.Message = new(TransferInConfirmationV07)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the TransferInConfirmation message to the instructing party, for example, an investment manager or its authorised representative, to confirm the receipt of a financial instrument, free of payment, on a given date, from a specified party.
// This message may also be used to confirm the receipt of a financial instrument, free of payment, from another of the instructing parties own accounts or from a third party.
// This message may also be used as an advice, that is, the message is used to provide account information.
// Usage
// The TransferInConfirmation message is used to confirm receipt of a financial instrument, either from another account owned by the instructing party or from a third party. The reference of the transfer confirmation is identified in TransferConfirmationReference.
// The reference of the original transfer instruction is specified in TransferReference. The message identification of the TransferInInstruction message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
type TransferInConfirmationV07 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference6 `xml:"PoolRef,omitempty"`

	// Reference of the linked message that was previously sent.
	PreviousReference *model.AdditionalReference6 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference6 `xml:"RltdRef,omitempty"`

	// Function of the transfer in, that is, whether the transfer in message is used as a confirmation or as or an advice. The absence of Function indicates the message is a confirmation.
	Function *model.TransferInFunction2Code `xml:"Fctn,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *model.Max35Text `xml:"MstrRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*model.Transfer33 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *model.InvestmentAccount56 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *model.DeliverInformation17 `xml:"SttlmDtls,omitempty"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation4 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferInConfirmationV07) AddMessageIdentification() *model.MessageIdentification1 {
	t.MessageIdentification = new(model.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferInConfirmationV07) AddPoolReference() *model.AdditionalReference6 {
	t.PoolReference = new(model.AdditionalReference6)
	return t.PoolReference
}

func (t *TransferInConfirmationV07) AddPreviousReference() *model.AdditionalReference6 {
	t.PreviousReference = new(model.AdditionalReference6)
	return t.PreviousReference
}

func (t *TransferInConfirmationV07) AddRelatedReference() *model.AdditionalReference6 {
	t.RelatedReference = new(model.AdditionalReference6)
	return t.RelatedReference
}

func (t *TransferInConfirmationV07) SetFunction(value string) {
	t.Function = (*model.TransferInFunction2Code)(&value)
}

func (t *TransferInConfirmationV07) SetMasterReference(value string) {
	t.MasterReference = (*model.Max35Text)(&value)
}

func (t *TransferInConfirmationV07) AddTransferDetails() *model.Transfer33 {
	newValue := new(model.Transfer33)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferInConfirmationV07) AddAccountDetails() *model.InvestmentAccount56 {
	t.AccountDetails = new(model.InvestmentAccount56)
	return t.AccountDetails
}

func (t *TransferInConfirmationV07) AddSettlementDetails() *model.DeliverInformation17 {
	t.SettlementDetails = new(model.DeliverInformation17)
	return t.SettlementDetails
}

func (t *TransferInConfirmationV07) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	t.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return t.MarketPracticeVersion
}

func (t *TransferInConfirmationV07) AddCopyDetails() *model.CopyInformation4 {
	t.CopyDetails = new(model.CopyInformation4)
	return t.CopyDetails
}

func (t *TransferInConfirmationV07) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
