package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700103 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.03 Document"`
	Message *ForwardDataSetSubmissionReportV03 `xml:"FwdDataSetSubmissnRpt"`
}

func (d *Document01700103) AddMessage() *ForwardDataSetSubmissionReportV03 {
	d.Message = new(ForwardDataSetSubmissionReportV03)
	return d.Message
}

// Scope
// The ForwardDataSetSubmissionReport message is sent by the matching application to the counterparty(ies) of the submitter of data sets.
// This message is used to pass on information related to the purchasing agreement(s) covered by the transaction(s) referred to in the message.
// Usage
// The ForwardDataSetSubmission message can be sent by the matching application to forward the details of a DataSetSubmission message that it has obtained.
type ForwardDataSetSubmissionReportV03 struct {

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
	CommercialDataSet *model.CommercialDataSet3 `xml:"ComrclDataSet,omitempty"`

	// Transport information that has been submitted to the matching application by the other party.
	TransportDataSet *model.TransportDataSet3 `xml:"TrnsprtDataSet,omitempty"`

	// Insurance information that has been submitted to the matching application by the other party.
	InsuranceDataSet *model.InsuranceDataSet1 `xml:"InsrncDataSet,omitempty"`

	// Certificate information that has been submitted to the matching application by the other party.
	CertificateDataSet []*model.CertificateDataSet1 `xml:"CertDataSet,omitempty"`

	// Other certificate information that has been submitted to the matching application by the other party.
	OtherCertificateDataSet []*model.OtherCertificateDataSet1 `xml:"OthrCertDataSet,omitempty"`

	// Next processing step required.
	RequestForAction *model.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (f *ForwardDataSetSubmissionReportV03) AddReportIdentification() *model.MessageIdentification1 {
	f.ReportIdentification = new(model.MessageIdentification1)
	return f.ReportIdentification
}

func (f *ForwardDataSetSubmissionReportV03) AddRelatedTransactionReferences() *model.DataSetSubmissionReferences4 {
	newValue := new(model.DataSetSubmissionReferences4)
	f.RelatedTransactionReferences = append(f.RelatedTransactionReferences, newValue)
	return newValue
}

func (f *ForwardDataSetSubmissionReportV03) AddCommonSubmissionReference() *model.SimpleIdentificationInformation {
	f.CommonSubmissionReference = new(model.SimpleIdentificationInformation)
	return f.CommonSubmissionReference
}

func (f *ForwardDataSetSubmissionReportV03) AddSubmitter() *model.BICIdentification1 {
	f.Submitter = new(model.BICIdentification1)
	return f.Submitter
}

func (f *ForwardDataSetSubmissionReportV03) AddBuyerBank() *model.BICIdentification1 {
	f.BuyerBank = new(model.BICIdentification1)
	return f.BuyerBank
}

func (f *ForwardDataSetSubmissionReportV03) AddSellerBank() *model.BICIdentification1 {
	f.SellerBank = new(model.BICIdentification1)
	return f.SellerBank
}

func (f *ForwardDataSetSubmissionReportV03) AddCommercialDataSet() *model.CommercialDataSet3 {
	f.CommercialDataSet = new(model.CommercialDataSet3)
	return f.CommercialDataSet
}

func (f *ForwardDataSetSubmissionReportV03) AddTransportDataSet() *model.TransportDataSet3 {
	f.TransportDataSet = new(model.TransportDataSet3)
	return f.TransportDataSet
}

func (f *ForwardDataSetSubmissionReportV03) AddInsuranceDataSet() *model.InsuranceDataSet1 {
	f.InsuranceDataSet = new(model.InsuranceDataSet1)
	return f.InsuranceDataSet
}

func (f *ForwardDataSetSubmissionReportV03) AddCertificateDataSet() *model.CertificateDataSet1 {
	newValue := new(model.CertificateDataSet1)
	f.CertificateDataSet = append(f.CertificateDataSet, newValue)
	return newValue
}

func (f *ForwardDataSetSubmissionReportV03) AddOtherCertificateDataSet() *model.OtherCertificateDataSet1 {
	newValue := new(model.OtherCertificateDataSet1)
	f.OtherCertificateDataSet = append(f.OtherCertificateDataSet, newValue)
	return newValue
}

func (f *ForwardDataSetSubmissionReportV03) AddRequestForAction() *model.PendingActivity2 {
	f.RequestForAction = new(model.PendingActivity2)
	return f.RequestForAction
}
