package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100101 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.011.001.01 Document"`
	Message *AcceptorBatchTransferV01 `xml:"AccptrBtchTrf"`
}

func (d *Document01100101) AddMessage() *AcceptorBatchTransferV01 {
	d.Message = new(AcceptorBatchTransferV01)
	return d.Message
}

// Scope
// The AcceptorBatchTransfer message is sent by the card acceptor to the acquirer to capture a collection of previously completed card payment transactions.
// Usage
// The AcceptorBatchTransfer message embeds the information required for transferring to the acquirer the data needed to perform the financial settlement of these transactions (capture).
type AcceptorBatchTransferV01 struct {

	// Batch capture message management information.
	Header *model.Header3 `xml:"Hdr"`

	// Information related to the set of transaction.
	DataSet []*model.CardPaymentDataSet1 `xml:"DataSet"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType1 `xml:"SctyTrlr"`
}

func (a *AcceptorBatchTransferV01) AddHeader() *model.Header3 {
	a.Header = new(model.Header3)
	return a.Header
}

func (a *AcceptorBatchTransferV01) AddDataSet() *model.CardPaymentDataSet1 {
	newValue := new(model.CardPaymentDataSet1)
	a.DataSet = append(a.DataSet, newValue)
	return newValue
}

func (a *AcceptorBatchTransferV01) AddSecurityTrailer() *model.ContentInformationType1 {
	a.SecurityTrailer = new(model.ContentInformationType1)
	return a.SecurityTrailer
}
