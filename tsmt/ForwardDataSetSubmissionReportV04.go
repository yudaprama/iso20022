package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700104 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.04 Document"`
	Message *ForwardDataSetSubmissionReportV04 `xml:"FwdDataSetSubmissnRpt"`
}

func (d *Document01700104) AddMessage() *ForwardDataSetSubmissionReportV04 {
	d.Message = new(ForwardDataSetSubmissionReportV04)
	return d.Message
}

// Scope
// The ForwardDataSetSubmissionReport message is sent by the matching application to the counterparty(ies) of the submitter of data sets.
// This message is used to pass on information related to the purchasing agreement(s) covered by the transaction(s) referred to in the message.
// Usage
// The ForwardDataSetSubmission message can be sent by the matching application to forward the details of a DataSetSubmission message that it has obtained.
type ForwardDataSetSubmissionReportV04 struct {

	// Identifies the report.
	ReportIdentification *model.MessageIdentification1 `xml:"RptId"`

	// Identifies the transactions that this submission relates to and provides associated information.
	RelatedTransactionReferences []*model.DataSetSubmissionReferences4 `xml:"RltdTxRefs"`

	// This reference must be used for all data sets belonging to the same submission group.
	CommonSubmissionReference *model.SimpleIdentificationInformation `xml:"CmonSubmissnRef"`

	// The financial institution that has submitted the data sets to the matching engine.
	Submitter *model.BICIdentification1 `xml:"Submitr"`

	// The financial institution of the buyer, uniquely identified by its BIC.
	BuyerBank *model.BICIdentification1 `xml:"BuyrBk"`

	// The financial institution of the seller, uniquely identified by its BIC.
	SellerBank *model.BICIdentification1 `xml:"SellrBk"`

	// Commercial information that has been submitted to the matching application by the other party.
	CommercialDataSet *model.CommercialDataSet4 `xml:"ComrclDataSet,omitempty"`

	// Transport information that has been submitted to the matching application by the other party.
	TransportDataSet *model.TransportDataSet4 `xml:"TrnsprtDataSet,omitempty"`

	// Insurance information that has been submitted to the matching application by the other party.
	InsuranceDataSet *model.InsuranceDataSet1 `xml:"InsrncDataSet,omitempty"`

	// Certificate information that has been submitted to the matching application by the other party.
	CertificateDataSet []*model.CertificateDataSet2 `xml:"CertDataSet,omitempty"`

	// Other certificate information that has been submitted to the matching application by the other party.
	OtherCertificateDataSet []*model.OtherCertificateDataSet1 `xml:"OthrCertDataSet,omitempty"`

	// Next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (f *ForwardDataSetSubmissionReportV04) AddReportIdentification() *model.MessageIdentification1 {
	f.ReportIdentification = new(model.MessageIdentification1)
	return f.ReportIdentification
}

func (f *ForwardDataSetSubmissionReportV04) AddRelatedTransactionReferences() *model.DataSetSubmissionReferences4 {
	newValue := new(model.DataSetSubmissionReferences4)
	f.RelatedTransactionReferences = append(f.RelatedTransactionReferences, newValue)
	return newValue
}

func (f *ForwardDataSetSubmissionReportV04) AddCommonSubmissionReference() *model.SimpleIdentificationInformation {
	f.CommonSubmissionReference = new(model.SimpleIdentificationInformation)
	return f.CommonSubmissionReference
}

func (f *ForwardDataSetSubmissionReportV04) AddSubmitter() *model.BICIdentification1 {
	f.Submitter = new(model.BICIdentification1)
	return f.Submitter
}

func (f *ForwardDataSetSubmissionReportV04) AddBuyerBank() *model.BICIdentification1 {
	f.BuyerBank = new(model.BICIdentification1)
	return f.BuyerBank
}

func (f *ForwardDataSetSubmissionReportV04) AddSellerBank() *model.BICIdentification1 {
	f.SellerBank = new(model.BICIdentification1)
	return f.SellerBank
}

func (f *ForwardDataSetSubmissionReportV04) AddCommercialDataSet() *model.CommercialDataSet4 {
	f.CommercialDataSet = new(model.CommercialDataSet4)
	return f.CommercialDataSet
}

func (f *ForwardDataSetSubmissionReportV04) AddTransportDataSet() *model.TransportDataSet4 {
	f.TransportDataSet = new(model.TransportDataSet4)
	return f.TransportDataSet
}

func (f *ForwardDataSetSubmissionReportV04) AddInsuranceDataSet() *model.InsuranceDataSet1 {
	f.InsuranceDataSet = new(model.InsuranceDataSet1)
	return f.InsuranceDataSet
}

func (f *ForwardDataSetSubmissionReportV04) AddCertificateDataSet() *model.CertificateDataSet2 {
	newValue := new(model.CertificateDataSet2)
	f.CertificateDataSet = append(f.CertificateDataSet, newValue)
	return newValue
}

func (f *ForwardDataSetSubmissionReportV04) AddOtherCertificateDataSet() *model.OtherCertificateDataSet1 {
	newValue := new(model.OtherCertificateDataSet1)
	f.OtherCertificateDataSet = append(f.OtherCertificateDataSet, newValue)
	return newValue
}

func (f *ForwardDataSetSubmissionReportV04) AddRequestForAction() *model.PendingActivity2 {
	f.RequestForAction = new(model.PendingActivity2)
	return f.RequestForAction
}
