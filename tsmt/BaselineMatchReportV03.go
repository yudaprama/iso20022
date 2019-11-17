package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01000103 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.010.001.03 Document"`
	Message *BaselineMatchReportV03 `xml:"BaselnMtchRpt"`
}

func (d *Document01000103) AddMessage() *BaselineMatchReportV03 {
	d.Message = new(BaselineMatchReportV03)
	return d.Message
}

// Scope
// The BaselineMatchReport message is sent by the matching application to the parties involved in the establishment of a transaction.
// The message is used to inform about either the successful establishment of a transaction (baseline) or the mis-matches found between two baseline initiation messages.
// Usage
// The BaselineMatchReport message can be sent by the matching application to
// - the parties involved in the establishment of a transaction in the push-through mode, that is the senders of InitialBaselineSubmission and BaselineReSubmission messages including the instruction push-through. In the outlined scenario the message is used to inform either about the successful establishment of a transaction in the matching application or about mis-matches found between two baseline initiation messages,or
// - the party establishing a transaction in the lodge mode, that is the sender of an InitialBaselineSubmission message including the instruction lodge. In the outlined scenario the message is used to inform about the successful establishment of a transaction in the matching application.
type BaselineMatchReportV03 struct {

	// Identifies the report.
	ReportIdentification *model.MessageIdentification1 `xml:"RptId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *model.DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

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

	// Specifies the number of matching trials for a baseline.
	BaselineEstablishmentTrials *model.Limit1 `xml:"BaselnEstblishmtTrils"`

	// Identifies the two baselines compared in this report.
	ComparedDocumentReference []*model.DocumentIdentification4 `xml:"CmpardDocRef"`

	// Description of the differences between the two proposed baselines
	Report *model.MisMatchReport3 `xml:"Rpt"`

	// Information on the next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (b *BaselineMatchReportV03) AddReportIdentification() *model.MessageIdentification1 {
	b.ReportIdentification = new(model.MessageIdentification1)
	return b.ReportIdentification
}

func (b *BaselineMatchReportV03) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	b.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return b.TransactionIdentification
}

func (b *BaselineMatchReportV03) AddEstablishedBaselineIdentification() *model.DocumentIdentification3 {
	b.EstablishedBaselineIdentification = new(model.DocumentIdentification3)
	return b.EstablishedBaselineIdentification
}

func (b *BaselineMatchReportV03) AddTransactionStatus() *model.TransactionStatus4 {
	b.TransactionStatus = new(model.TransactionStatus4)
	return b.TransactionStatus
}

func (b *BaselineMatchReportV03) AddUserTransactionReference() *model.DocumentIdentification5 {
	newValue := new(model.DocumentIdentification5)
	b.UserTransactionReference = append(b.UserTransactionReference, newValue)
	return newValue
}

func (b *BaselineMatchReportV03) AddBuyer() *model.PartyIdentification26 {
	b.Buyer = new(model.PartyIdentification26)
	return b.Buyer
}

func (b *BaselineMatchReportV03) AddSeller() *model.PartyIdentification26 {
	b.Seller = new(model.PartyIdentification26)
	return b.Seller
}

func (b *BaselineMatchReportV03) AddBuyerBank() *model.BICIdentification1 {
	b.BuyerBank = new(model.BICIdentification1)
	return b.BuyerBank
}

func (b *BaselineMatchReportV03) AddSellerBank() *model.BICIdentification1 {
	b.SellerBank = new(model.BICIdentification1)
	return b.SellerBank
}

func (b *BaselineMatchReportV03) AddBaselineEstablishmentTrials() *model.Limit1 {
	b.BaselineEstablishmentTrials = new(model.Limit1)
	return b.BaselineEstablishmentTrials
}

func (b *BaselineMatchReportV03) AddComparedDocumentReference() *model.DocumentIdentification4 {
	newValue := new(model.DocumentIdentification4)
	b.ComparedDocumentReference = append(b.ComparedDocumentReference, newValue)
	return newValue
}

func (b *BaselineMatchReportV03) AddReport() *model.MisMatchReport3 {
	b.Report = new(model.MisMatchReport3)
	return b.Report
}

func (b *BaselineMatchReportV03) AddRequestForAction() *model.PendingActivity2 {
	b.RequestForAction = new(model.PendingActivity2)
	return b.RequestForAction
}
