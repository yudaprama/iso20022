package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01800103 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.018.001.03 Document"`
	Message *FullPushThroughReportV03 `xml:"FullPushThrghRpt"`
}

func (d *Document01800103) AddMessage() *FullPushThroughReportV03 {
	d.Message = new(FullPushThroughReportV03)
	return d.Message
}

// Scope
// The FullPushThroughReport message is sent by the matching application to a party involved in a transaction.
// This message is used to pass on information that the matching application has received from the submitter. The forwarded information can originate from an InitialBaselineSubmission or BaselineReSubmission or BaselineAmendmentRequest message.
// Usage
// The FullPushThroughReport message can be sent by the matching application to a party to convey
// - the details of an InitialBaselineSubmission message that it has obtained,or
// - the details of a BaselineResubmission message that it has obtained,or
// - the details of a BaselineAmendmentRequest message that it has obtained.
type FullPushThroughReportV03 struct {

	// Identifies the report.
	ReportIdentification *model.MessageIdentification1 `xml:"RptId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *model.DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *model.TransactionStatus4 `xml:"TxSts"`

	// Reference to the transaction for the financial institution which submitted the baseline.
	UserTransactionReference []*model.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Specifies the type of report.
	ReportPurpose *model.ReportType1 `xml:"RptPurp"`

	// Specifies the commercial details of the underlying transaction.
	PushedThroughBaseline *model.Baseline3 `xml:"PushdThrghBaseln"`

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

	// Information on the next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (f *FullPushThroughReportV03) AddReportIdentification() *model.MessageIdentification1 {
	f.ReportIdentification = new(model.MessageIdentification1)
	return f.ReportIdentification
}

func (f *FullPushThroughReportV03) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	f.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return f.TransactionIdentification
}

func (f *FullPushThroughReportV03) AddEstablishedBaselineIdentification() *model.DocumentIdentification3 {
	f.EstablishedBaselineIdentification = new(model.DocumentIdentification3)
	return f.EstablishedBaselineIdentification
}

func (f *FullPushThroughReportV03) AddTransactionStatus() *model.TransactionStatus4 {
	f.TransactionStatus = new(model.TransactionStatus4)
	return f.TransactionStatus
}

func (f *FullPushThroughReportV03) AddUserTransactionReference() *model.DocumentIdentification5 {
	newValue := new(model.DocumentIdentification5)
	f.UserTransactionReference = append(f.UserTransactionReference, newValue)
	return newValue
}

func (f *FullPushThroughReportV03) AddReportPurpose() *model.ReportType1 {
	f.ReportPurpose = new(model.ReportType1)
	return f.ReportPurpose
}

func (f *FullPushThroughReportV03) AddPushedThroughBaseline() *model.Baseline3 {
	f.PushedThroughBaseline = new(model.Baseline3)
	return f.PushedThroughBaseline
}

func (f *FullPushThroughReportV03) AddBuyerContactPerson() *model.ContactIdentification1 {
	newValue := new(model.ContactIdentification1)
	f.BuyerContactPerson = append(f.BuyerContactPerson, newValue)
	return newValue
}

func (f *FullPushThroughReportV03) AddSellerContactPerson() *model.ContactIdentification1 {
	newValue := new(model.ContactIdentification1)
	f.SellerContactPerson = append(f.SellerContactPerson, newValue)
	return newValue
}

func (f *FullPushThroughReportV03) AddBuyerBankContactPerson() *model.ContactIdentification1 {
	newValue := new(model.ContactIdentification1)
	f.BuyerBankContactPerson = append(f.BuyerBankContactPerson, newValue)
	return newValue
}

func (f *FullPushThroughReportV03) AddSellerBankContactPerson() *model.ContactIdentification1 {
	newValue := new(model.ContactIdentification1)
	f.SellerBankContactPerson = append(f.SellerBankContactPerson, newValue)
	return newValue
}

func (f *FullPushThroughReportV03) AddOtherBankContactPerson() *model.ContactIdentification3 {
	newValue := new(model.ContactIdentification3)
	f.OtherBankContactPerson = append(f.OtherBankContactPerson, newValue)
	return newValue
}

func (f *FullPushThroughReportV03) AddRequestForAction() *model.PendingActivity2 {
	f.RequestForAction = new(model.PendingActivity2)
	return f.RequestForAction
}
