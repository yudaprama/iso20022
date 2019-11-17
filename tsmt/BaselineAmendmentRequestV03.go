package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00900103 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.009.001.03 Document"`
	Message *BaselineAmendmentRequestV03 `xml:"BaselnAmdmntReq"`
}

func (d *Document00900103) AddMessage() *BaselineAmendmentRequestV03 {
	d.Message = new(BaselineAmendmentRequestV03)
	return d.Message
}

// Scope
// The BaselineAmendmentRequest message is sent by a primary party involved in a transaction to the matching application.
// The message is used to request the amendment of an established baseline.
// Usage
// The BaselineAmendmentRequest message may only be sent if the transaction is in the state Established or Active.
// The BaselineAmendmentRequest message can be sent to the matching application by one of the primary parties involved in a transaction established in the push-through mode to request the amendment of an established baseline.
// The matching application acknowledges the receipt of the amendment request by sending a DeltaReport message to the submitter of the BaselineAmendmentRequest message. It passes on the newly proposed baseline to the counterparty by sending a FullPushThroughReport message, a DeltaReport message and a pre-calculated BaselineReport message.
// The counterparty is expected to either accept or reject the amendment request by submitting an AmendmentAcceptance or AmendmentRejection message.
// or
// The BaselineAmendmentRequest message can be sent by the party involved in a transaction established in the lodge mode to the matching application to amend an established baseline.
// The matching application amends the baseline according to the BaselineAmendmentRequest message and confirms the execution of the request by sending a DeltaReport and calculated BaselineReport message to the requester of the amendment.
type BaselineAmendmentRequestV03 struct {

	// Identifies the request message.
	RequestIdentification *model.MessageIdentification1 `xml:"ReqId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *model.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Specifies the commercial details of the underlying transaction.
	Baseline *model.Baseline3 `xml:"Baseln"`

	// Person to be contacted in the organisation of the buyer.
	BuyerContactPerson []*model.ContactIdentification1 `xml:"BuyrCtctPrsn,omitempty"`

	// Person to be contacted in the organisation of the seller.
	SellerContactPerson []*model.ContactIdentification1 `xml:"SellrCtctPrsn,omitempty"`

	// Person to be contacted in the buyer's bank.
	BuyerBankContactPerson []*model.ContactIdentification1 `xml:"BuyrBkCtctPrsn,omitempty"`

	// Person to be contacted in the seller's bank.
	SellerBankContactPerson []*model.ContactIdentification1 `xml:"SellrBkCtctPrsn,omitempty"`

	// Person to be contacted in another bank than the seller or buyer's bank.
	OtherBankContactPerson []*model.ContactIdentification3 `xml:"OthrBkCtctPrsn,omitempty"`
}

func (b *BaselineAmendmentRequestV03) AddRequestIdentification() *model.MessageIdentification1 {
	b.RequestIdentification = new(model.MessageIdentification1)
	return b.RequestIdentification
}

func (b *BaselineAmendmentRequestV03) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	b.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return b.TransactionIdentification
}

func (b *BaselineAmendmentRequestV03) AddSubmitterTransactionReference() *model.SimpleIdentificationInformation {
	b.SubmitterTransactionReference = new(model.SimpleIdentificationInformation)
	return b.SubmitterTransactionReference
}

func (b *BaselineAmendmentRequestV03) AddBaseline() *model.Baseline3 {
	b.Baseline = new(model.Baseline3)
	return b.Baseline
}

func (b *BaselineAmendmentRequestV03) AddBuyerContactPerson() *model.ContactIdentification1 {
	newValue := new(model.ContactIdentification1)
	b.BuyerContactPerson = append(b.BuyerContactPerson, newValue)
	return newValue
}

func (b *BaselineAmendmentRequestV03) AddSellerContactPerson() *model.ContactIdentification1 {
	newValue := new(model.ContactIdentification1)
	b.SellerContactPerson = append(b.SellerContactPerson, newValue)
	return newValue
}

func (b *BaselineAmendmentRequestV03) AddBuyerBankContactPerson() *model.ContactIdentification1 {
	newValue := new(model.ContactIdentification1)
	b.BuyerBankContactPerson = append(b.BuyerBankContactPerson, newValue)
	return newValue
}

func (b *BaselineAmendmentRequestV03) AddSellerBankContactPerson() *model.ContactIdentification1 {
	newValue := new(model.ContactIdentification1)
	b.SellerBankContactPerson = append(b.SellerBankContactPerson, newValue)
	return newValue
}

func (b *BaselineAmendmentRequestV03) AddOtherBankContactPerson() *model.ContactIdentification3 {
	newValue := new(model.ContactIdentification3)
	b.OtherBankContactPerson = append(b.OtherBankContactPerson, newValue)
	return newValue
}
