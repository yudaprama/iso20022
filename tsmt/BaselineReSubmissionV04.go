package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200104 struct {
	XMLName xml.Name                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.012.001.04 Document"`
	Message *BaselineReSubmissionV04 `xml:"BaselnReSubmissn"`
}

func (d *Document01200104) AddMessage() *BaselineReSubmissionV04 {
	d.Message = new(BaselineReSubmissionV04)
	return d.Message
}

// Scope
// The BaselineReSubmission message is sent by either the counterparty or the initiator of a transaction (baseline) to the matching application.
// This message is used by the counterparty to respond on the registration of a push-through transaction in the matching application or by the initiator or counterparty to re-send earlier mis-matched baseline information.
// Usage
// The BaselineReSubmission message can be sent by the counterparty of a transaction to the matching application in response to a FullPushThroughReport message received from the matching application conveying the details of an InitialBaselineSubmission message. The objective of the BaselineReSubmission message sent in the outlined scenario is to achieve a successful match of two baseline initiation messages in order to establish a transaction in the matching application.
// or
// The BaselineReSubmission message can be sent by the initiator of a transaction to the matching application in response to a BaselineMatchReport message indicating mis-matches. The objective of the BaselineReSubmission message sent in the outlined scenario is to correct an InitialBaselineSubmission or BaselineReSubmission message submitted earlier in order to achieve the establishment of a transaction in the matching application.
// or
// The BaselineReSubmission message can be sent by the counterparty of a transaction to the matching application in response to a BaselineMatchReport message indicating mis-matches. The objective of the BaselineReSubmission message sent in the outlined scenario is to correct a BaselineReSubmission message submitted earlier in order to achieve the establishment of a transaction in the matching application.
type BaselineReSubmissionV04 struct {

	// Identifies the submitted information
	SubmissionIdentification *model.MessageIdentification1 `xml:"SubmissnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *model.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Specifies the commercial details of the underlying transaction.
	Baseline *model.Baseline4 `xml:"Baseln"`

	// Person to be contacted in the organisation of the buyer.
	BuyerContactPerson []*model.ContactIdentification1 `xml:"BuyrCtctPrsn,omitempty"`

	// Person to be contacted in the organisation of the seller.
	SellerContactPerson []*model.ContactIdentification1 `xml:"SellrCtctPrsn,omitempty"`

	// Person to be contacted in the seller's bank or buyer's bank.
	BankContactPerson *model.BankContactPerson1Choice `xml:"BkCtctPrsn"`

	// Person to be contacted in another bank than the seller or buyer's bank.
	OtherBankContactPerson []*model.ContactIdentification3 `xml:"OthrBkCtctPrsn,omitempty"`
}

func (b *BaselineReSubmissionV04) AddSubmissionIdentification() *model.MessageIdentification1 {
	b.SubmissionIdentification = new(model.MessageIdentification1)
	return b.SubmissionIdentification
}

func (b *BaselineReSubmissionV04) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	b.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return b.TransactionIdentification
}

func (b *BaselineReSubmissionV04) AddSubmitterTransactionReference() *model.SimpleIdentificationInformation {
	b.SubmitterTransactionReference = new(model.SimpleIdentificationInformation)
	return b.SubmitterTransactionReference
}

func (b *BaselineReSubmissionV04) AddBaseline() *model.Baseline4 {
	b.Baseline = new(model.Baseline4)
	return b.Baseline
}

func (b *BaselineReSubmissionV04) AddBuyerContactPerson() *model.ContactIdentification1 {
	newValue := new(model.ContactIdentification1)
	b.BuyerContactPerson = append(b.BuyerContactPerson, newValue)
	return newValue
}

func (b *BaselineReSubmissionV04) AddSellerContactPerson() *model.ContactIdentification1 {
	newValue := new(model.ContactIdentification1)
	b.SellerContactPerson = append(b.SellerContactPerson, newValue)
	return newValue
}

func (b *BaselineReSubmissionV04) AddBankContactPerson() *model.BankContactPerson1Choice {
	b.BankContactPerson = new(model.BankContactPerson1Choice)
	return b.BankContactPerson
}

func (b *BaselineReSubmissionV04) AddOtherBankContactPerson() *model.ContactIdentification3 {
	newValue := new(model.ContactIdentification3)
	b.OtherBankContactPerson = append(b.OtherBankContactPerson, newValue)
	return newValue
}
