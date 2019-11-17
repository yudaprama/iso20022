package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300103 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.013.001.03 Document"`
	Message *DataSetMatchReportV03 `xml:"DataSetMtchRpt"`
}

func (d *Document01300103) AddMessage() *DataSetMatchReportV03 {
	d.Message = new(DataSetMatchReportV03)
	return d.Message
}

// Scope
// The DataSetMatchReport message is sent by the matching application to the parties involved in a data set match.
// This message is used to either
// - inform about the successful match of data sets submitted with the instruction match or pre-match (DataSetSubmission message) and the related baseline,or
// - inform about mis-matches found between data sets submitted with the instruction match or pre-match (DataSetSubmission message) and the related baseline.
// Usage
// The DataSetMatchReport message can be sent by the matching application to the party requesting a data set pre-match for a transaction established in the push-through mode. In the outlined scenario, the DataSetMatchReport message will either inform about the successful pre-match or list the mis-matches between the data set(s) conveyed with the DataSetSubmission message and the related baseline.
// or
// The DataSetMatchReport message can be sent by the matching application to the parties involved in a data set match for a transaction established in the push-through mode. In the outlined scenario, the DataSetMatchReport message will either inform about the successful match or list the mis-matches between the data set(s) conveyed with the DataSetSubmission message and the related baseline.
// or
// The DataSetMatchReport message can be sent by the matching application to the party requesting a data set match or pre-match for a transaction established in the lodge mode. In the outlined scenario, the DataSetMatchReport will either inform about the successful match or list the mis-matches between the data set(s) conveyed with the DataSetSubmission message and the related baseline.
type DataSetMatchReportV03 struct {

	// Identifies the report.
	ReportIdentification *model.MessageIdentification1 `xml:"RptId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *model.DocumentIdentification3 `xml:"EstblishdBaselnId"`

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

	// Identifies the documents compared in this report.
	ComparedDocumentReference []*model.DocumentIdentification10 `xml:"CmpardDocRef"`

	// Specifies whether the data set was submitted for match or pre-match.
	SubmissionType *model.ReportType3 `xml:"SubmissnTp"`

	// Description of the differences between the data set(s) and the baseline.
	Report *model.MisMatchReport3 `xml:"Rpt"`

	// Information on the next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (d *DataSetMatchReportV03) AddReportIdentification() *model.MessageIdentification1 {
	d.ReportIdentification = new(model.MessageIdentification1)
	return d.ReportIdentification
}

func (d *DataSetMatchReportV03) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	d.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return d.TransactionIdentification
}

func (d *DataSetMatchReportV03) AddEstablishedBaselineIdentification() *model.DocumentIdentification3 {
	d.EstablishedBaselineIdentification = new(model.DocumentIdentification3)
	return d.EstablishedBaselineIdentification
}

func (d *DataSetMatchReportV03) AddTransactionStatus() *model.TransactionStatus4 {
	d.TransactionStatus = new(model.TransactionStatus4)
	return d.TransactionStatus
}

func (d *DataSetMatchReportV03) AddUserTransactionReference() *model.DocumentIdentification5 {
	newValue := new(model.DocumentIdentification5)
	d.UserTransactionReference = append(d.UserTransactionReference, newValue)
	return newValue
}

func (d *DataSetMatchReportV03) AddBuyer() *model.PartyIdentification26 {
	d.Buyer = new(model.PartyIdentification26)
	return d.Buyer
}

func (d *DataSetMatchReportV03) AddSeller() *model.PartyIdentification26 {
	d.Seller = new(model.PartyIdentification26)
	return d.Seller
}

func (d *DataSetMatchReportV03) AddBuyerBank() *model.BICIdentification1 {
	d.BuyerBank = new(model.BICIdentification1)
	return d.BuyerBank
}

func (d *DataSetMatchReportV03) AddSellerBank() *model.BICIdentification1 {
	d.SellerBank = new(model.BICIdentification1)
	return d.SellerBank
}

func (d *DataSetMatchReportV03) AddComparedDocumentReference() *model.DocumentIdentification10 {
	newValue := new(model.DocumentIdentification10)
	d.ComparedDocumentReference = append(d.ComparedDocumentReference, newValue)
	return newValue
}

func (d *DataSetMatchReportV03) AddSubmissionType() *model.ReportType3 {
	d.SubmissionType = new(model.ReportType3)
	return d.SubmissionType
}

func (d *DataSetMatchReportV03) AddReport() *model.MisMatchReport3 {
	d.Report = new(model.MisMatchReport3)
	return d.Report
}

func (d *DataSetMatchReportV03) AddRequestForAction() *model.PendingActivity2 {
	d.RequestForAction = new(model.PendingActivity2)
	return d.RequestForAction
}
