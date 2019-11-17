package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100104 struct {
	XMLName xml.Name           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Document"`
	Message *BaselineReportV04 `xml:"BaselnRpt"`
}

func (d *Document01100104) AddMessage() *BaselineReportV04 {
	d.Message = new(BaselineReportV04)
	return d.Message
}

// Scope
// The BaselineReport message is sent by the matching application to the parties involved in an amendment request or to the parties involved in a data set match.
// The message is used to report either a pre-calculation or final calculation of the dynamic part of an established baseline.
// Usage
// The BaselineReport message can be sent by the matching application to the parties involved in an amendment request for a transaction established in the push-through mode. In the outlined scenario, the message is sent
// - to the party requested to accept or reject an amendment request after the matching application has received a BaselineAmendmentRequest message. The message informs about the provisional status of the dynamic part of the baseline.
// - to the requester and the accepter of an amendment request after the matching application has received an AmendmentAcceptance message conveying the acceptance of the amendment request. The message informs about the actual status of the dynamic part of the baseline.
// or
// The BaselineReport message can be sent by the matching application to the party which has sent an amendment request for a transaction established in the lodge mode. In the outlined scenario the message is used to inform about the actual status of the dynamic part of the baseline.
// or
// The BaselineReport message can be sent by the matching application to the parties involved in a data set match. In the outlined scenario, the message is sent
// - to the submitter of the data set(s) in the case of a data set match for a transaction established in the lodge mode.
// - to the submitter of the data set(s) and to the counterparty in case of a data set match for a transaction established in the push-through mode.The message can be sent after a successful data-set match or after the acceptance of mis-matched data sets to inform about the actual status of the dynamic part of the baseline.
type BaselineReportV04 struct {

	// Identifies the report.
	ReportIdentification *model.MessageIdentification1 `xml:"RptId"`

	// Reference to the related message at the origin of the report or sent at the same time than the report.
	RelatedMessageReference *model.MessageIdentification1 `xml:"RltdMsgRef,omitempty"`

	// Type of baseline report.
	ReportType *model.ReportType2 `xml:"RptTp"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	//
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *model.DocumentIdentification6 `xml:"EstblishdBaselnId"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *model.TransactionStatus4 `xml:"TxSts"`

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

	// Information on the goods
	ReportedLineItem *model.LineItem14 `xml:"RptdLineItm"`

	// Information on the next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (b *BaselineReportV04) AddReportIdentification() *model.MessageIdentification1 {
	b.ReportIdentification = new(model.MessageIdentification1)
	return b.ReportIdentification
}

func (b *BaselineReportV04) AddRelatedMessageReference() *model.MessageIdentification1 {
	b.RelatedMessageReference = new(model.MessageIdentification1)
	return b.RelatedMessageReference
}

func (b *BaselineReportV04) AddReportType() *model.ReportType2 {
	b.ReportType = new(model.ReportType2)
	return b.ReportType
}

func (b *BaselineReportV04) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	b.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return b.TransactionIdentification
}

func (b *BaselineReportV04) AddEstablishedBaselineIdentification() *model.DocumentIdentification6 {
	b.EstablishedBaselineIdentification = new(model.DocumentIdentification6)
	return b.EstablishedBaselineIdentification
}

func (b *BaselineReportV04) AddTransactionStatus() *model.TransactionStatus4 {
	b.TransactionStatus = new(model.TransactionStatus4)
	return b.TransactionStatus
}

func (b *BaselineReportV04) AddUserTransactionReference() *model.DocumentIdentification5 {
	newValue := new(model.DocumentIdentification5)
	b.UserTransactionReference = append(b.UserTransactionReference, newValue)
	return newValue
}

func (b *BaselineReportV04) AddBuyer() *model.PartyIdentification26 {
	b.Buyer = new(model.PartyIdentification26)
	return b.Buyer
}

func (b *BaselineReportV04) AddSeller() *model.PartyIdentification26 {
	b.Seller = new(model.PartyIdentification26)
	return b.Seller
}

func (b *BaselineReportV04) AddBuyerBank() *model.BICIdentification1 {
	b.BuyerBank = new(model.BICIdentification1)
	return b.BuyerBank
}

func (b *BaselineReportV04) AddSellerBank() *model.BICIdentification1 {
	b.SellerBank = new(model.BICIdentification1)
	return b.SellerBank
}

func (b *BaselineReportV04) AddReportedLineItem() *model.LineItem14 {
	b.ReportedLineItem = new(model.LineItem14)
	return b.ReportedLineItem
}

func (b *BaselineReportV04) AddRequestForAction() *model.PendingActivity2 {
	b.RequestForAction = new(model.PendingActivity2)
	return b.RequestForAction
}
