package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400104 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.04 Document"`
	Message *DataSetSubmissionV04 `xml:"DataSetSubmissn"`
}

func (d *Document01400104) AddMessage() *DataSetSubmissionV04 {
	d.Message = new(DataSetSubmissionV04)
	return d.Message
}

// Scope
// The DataSetSubmission message is sent by a party involved in the transaction to the matching application.
// This message is used to trigger either a match or a pre-match of the information submitted with the message.
// Usage
// The DataSetSubmission message can be sent by either party with the instruction pre-match. In the outlined scenario, the matching application will compare the data set(s) conveyed by the DataSetSubmission message with the established baseline and report the matching result to the requester of the data set pre-match by sending a DataSetMatchReport message.
// or
// The DataSetSubmission message can be sent by the party specified in the baseline as data set submitter with the instruction match. In the outlined scenario, the matching application will compare the data set(s) conveyed by the DataSetSubmission message with the established baseline and report the matching result to
// - the parties involved in a transaction established in the push-through mode, or
// - the initiator of a transaction established in the lodge mode.
// The DataSetSubmission message can be used to submit multiple data sets for multiple transactions (baselines) at the same time. However, all transactions (baselines) covered by the message must be for the same parties, that is transaction initiator and counterparty, as well as for the same buyer and seller.
// The DataSetSubmission message consists of data reflecting trade information related to the purchasing agreement covered by the transaction(s), for example shipment date, invoice amount.
type DataSetSubmissionV04 struct {

	// Identifies the submitted information.
	SubmissionIdentification *model.MessageIdentification1 `xml:"SubmissnId"`

	// Identifies the transactions that this submission relates to and provides associated information.
	RelatedTransactionReferences []*model.DataSetSubmissionReferences3 `xml:"RltdTxRefs"`

	// This reference must be used for all data sets belonging to the same submission group.
	CommonSubmissionReference *model.SimpleIdentificationInformation `xml:"CmonSubmissnRef"`

	// Specifies the instruction given by the submitter.
	Instruction *model.InstructionType3 `xml:"Instr"`

	// The financial institution of the buyer, uniquely identified by its BIC.
	BuyerBank *model.BICIdentification1 `xml:"BuyrBk"`

	// The financial institution of the seller, uniquely identified by its BIC.
	SellerBank *model.BICIdentification1 `xml:"SellrBk"`

	// Commercial information that is submitted to the matching application for processing.
	CommercialDataSet *model.CommercialDataSet4 `xml:"ComrclDataSet,omitempty"`

	// Transport information that is submitted to the matching application for processing.
	TransportDataSet *model.TransportDataSet4 `xml:"TrnsprtDataSet,omitempty"`

	// Insurance information that is submitted to the matching application for processing.
	InsuranceDataSet *model.InsuranceDataSet1 `xml:"InsrncDataSet,omitempty"`

	// Certificate information that is submitted to the matching application for processing.
	CertificateDataSet []*model.CertificateDataSet2 `xml:"CertDataSet,omitempty"`

	// Other certificate information that is submitted to the matching application for processing.
	OtherCertificateDataSet []*model.OtherCertificateDataSet1 `xml:"OthrCertDataSet,omitempty"`
}

func (d *DataSetSubmissionV04) AddSubmissionIdentification() *model.MessageIdentification1 {
	d.SubmissionIdentification = new(model.MessageIdentification1)
	return d.SubmissionIdentification
}

func (d *DataSetSubmissionV04) AddRelatedTransactionReferences() *model.DataSetSubmissionReferences3 {
	newValue := new(model.DataSetSubmissionReferences3)
	d.RelatedTransactionReferences = append(d.RelatedTransactionReferences, newValue)
	return newValue
}

func (d *DataSetSubmissionV04) AddCommonSubmissionReference() *model.SimpleIdentificationInformation {
	d.CommonSubmissionReference = new(model.SimpleIdentificationInformation)
	return d.CommonSubmissionReference
}

func (d *DataSetSubmissionV04) AddInstruction() *model.InstructionType3 {
	d.Instruction = new(model.InstructionType3)
	return d.Instruction
}

func (d *DataSetSubmissionV04) AddBuyerBank() *model.BICIdentification1 {
	d.BuyerBank = new(model.BICIdentification1)
	return d.BuyerBank
}

func (d *DataSetSubmissionV04) AddSellerBank() *model.BICIdentification1 {
	d.SellerBank = new(model.BICIdentification1)
	return d.SellerBank
}

func (d *DataSetSubmissionV04) AddCommercialDataSet() *model.CommercialDataSet4 {
	d.CommercialDataSet = new(model.CommercialDataSet4)
	return d.CommercialDataSet
}

func (d *DataSetSubmissionV04) AddTransportDataSet() *model.TransportDataSet4 {
	d.TransportDataSet = new(model.TransportDataSet4)
	return d.TransportDataSet
}

func (d *DataSetSubmissionV04) AddInsuranceDataSet() *model.InsuranceDataSet1 {
	d.InsuranceDataSet = new(model.InsuranceDataSet1)
	return d.InsuranceDataSet
}

func (d *DataSetSubmissionV04) AddCertificateDataSet() *model.CertificateDataSet2 {
	newValue := new(model.CertificateDataSet2)
	d.CertificateDataSet = append(d.CertificateDataSet, newValue)
	return newValue
}

func (d *DataSetSubmissionV04) AddOtherCertificateDataSet() *model.OtherCertificateDataSet1 {
	newValue := new(model.OtherCertificateDataSet1)
	d.OtherCertificateDataSet = append(d.OtherCertificateDataSet, newValue)
	return newValue
}
