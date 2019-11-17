package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200101 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.012.001.01 Document"`
	Message *AcceptorBatchTransferResponseV01 `xml:"AccptrBtchTrfRspn"`
}

func (d *Document01200101) AddMessage() *AcceptorBatchTransferResponseV01 {
	d.Message = new(AcceptorBatchTransferResponseV01)
	return d.Message
}

// Scope
// The AcceptorBatchTransferResponse message is sent by the acquirer to the card acceptor to acknowledge the proper reception of the AcceptorBatchTransfer.
// Usage
// The AcceptorBatchTransferResponse message is used by an acquirer to inform the card acceptor of the card payment transactions that could not be captured in the AcceptorBatchTransfer.
type AcceptorBatchTransferResponseV01 struct {

	// Capture advice response message management information.
	Header *model.Header3 `xml:"Hdr"`

	// Information related to the previously sent set of transaction.
	DataSet []*model.CardPaymentDataSet2 `xml:"DataSet"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType1 `xml:"SctyTrlr"`
}

func (a *AcceptorBatchTransferResponseV01) AddHeader() *model.Header3 {
	a.Header = new(model.Header3)
	return a.Header
}

func (a *AcceptorBatchTransferResponseV01) AddDataSet() *model.CardPaymentDataSet2 {
	newValue := new(model.CardPaymentDataSet2)
	a.DataSet = append(a.DataSet, newValue)
	return newValue
}

func (a *AcceptorBatchTransferResponseV01) AddSecurityTrailer() *model.ContentInformationType1 {
	a.SecurityTrailer = new(model.ContentInformationType1)
	return a.SecurityTrailer
}
