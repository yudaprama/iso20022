package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01500103 struct {
	XMLName xml.Name        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.015.001.03 Document"`
	Message *DeltaReportV03 `xml:"DltaRpt"`
}

func (d *Document01500103) AddMessage() *DeltaReportV03 {
	d.Message = new(DeltaReportV03)
	return d.Message
}

// Scope
// The DeltaReport message is sent by the matching application to the parties involved in the request of a baseline amendment.
// The message is used to list the differences between the established and the newly proposed baseline.
// Usage
// The DeltaReport message can be sent by the matching application to
// - the parties involved in the amendment of a baseline that has been established in the push-through mode. In the outlined scenario the message is sent to the requester of the amendment to acknowledge the receipt of the request and to list the differences between the established and the newly proposed baseline and to the counterparty to list the differences between the established and the newly proposed baseline and to request the acceptance or rejection of the amendment request,
// or
// - the party that has requested the amendment of a baseline established in the lodge mode. In the outlined scenario the message is used to confirm the changes to the baseline and to list the differences between the amended baseline and the baseline established earlier.
type DeltaReportV03 struct {

	// Identifies the report.
	ReportIdentification *model.MessageIdentification1 `xml:"RptId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *model.DocumentIdentification3 `xml:"EstblishdBaselnId"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *model.TransactionStatus4 `xml:"TxSts"`

	// Sequence number of the proposed baseline amendment.
	AmendmentNumber *model.Count1 `xml:"AmdmntNb"`

	// Reference to the transaction for each financial institution which is a party to the transaction.
	UserTransactionReference []*model.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Party that buys goods or services, or a financial instrument.
	Buyer *model.PartyIdentification26 `xml:"Buyr"`

	// Party that sells goods or services, or a financial instrument.
	Seller *model.PartyIdentification26 `xml:"Sellr"`

	// The financial institution of the buyer, uniquely identified by its BIC.
	BuyerBank *model.BICIdentification1 `xml:"BuyrBk"`

	// The financial institution of the seller, uniquely identified by its BIC.
	SellerBank *model.BICIdentification1 `xml:"SellrBk"`

	// Reference to the identification of the baseline included in the amendment request.
	SubmitterProposedBaselineReference *model.DocumentIdentification1 `xml:"SubmitrPropsdBaselnRef"`

	// Detailed comparison between the currently established baseline elements and the proposed ones.
	UpdatedElement []*model.ComparisonResult2 `xml:"UpdtdElmt"`

	// Information on the next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (d *DeltaReportV03) AddReportIdentification() *model.MessageIdentification1 {
	d.ReportIdentification = new(model.MessageIdentification1)
	return d.ReportIdentification
}

func (d *DeltaReportV03) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	d.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return d.TransactionIdentification
}

func (d *DeltaReportV03) AddEstablishedBaselineIdentification() *model.DocumentIdentification3 {
	d.EstablishedBaselineIdentification = new(model.DocumentIdentification3)
	return d.EstablishedBaselineIdentification
}

func (d *DeltaReportV03) AddTransactionStatus() *model.TransactionStatus4 {
	d.TransactionStatus = new(model.TransactionStatus4)
	return d.TransactionStatus
}

func (d *DeltaReportV03) AddAmendmentNumber() *model.Count1 {
	d.AmendmentNumber = new(model.Count1)
	return d.AmendmentNumber
}

func (d *DeltaReportV03) AddUserTransactionReference() *model.DocumentIdentification5 {
	newValue := new(model.DocumentIdentification5)
	d.UserTransactionReference = append(d.UserTransactionReference, newValue)
	return newValue
}

func (d *DeltaReportV03) AddBuyer() *model.PartyIdentification26 {
	d.Buyer = new(model.PartyIdentification26)
	return d.Buyer
}

func (d *DeltaReportV03) AddSeller() *model.PartyIdentification26 {
	d.Seller = new(model.PartyIdentification26)
	return d.Seller
}

func (d *DeltaReportV03) AddBuyerBank() *model.BICIdentification1 {
	d.BuyerBank = new(model.BICIdentification1)
	return d.BuyerBank
}

func (d *DeltaReportV03) AddSellerBank() *model.BICIdentification1 {
	d.SellerBank = new(model.BICIdentification1)
	return d.SellerBank
}

func (d *DeltaReportV03) AddSubmitterProposedBaselineReference() *model.DocumentIdentification1 {
	d.SubmitterProposedBaselineReference = new(model.DocumentIdentification1)
	return d.SubmitterProposedBaselineReference
}

func (d *DeltaReportV03) AddUpdatedElement() *model.ComparisonResult2 {
	newValue := new(model.ComparisonResult2)
	d.UpdatedElement = append(d.UpdatedElement, newValue)
	return newValue
}

func (d *DeltaReportV03) AddRequestForAction() *model.PendingActivity2 {
	d.RequestForAction = new(model.PendingActivity2)
	return d.RequestForAction
}
