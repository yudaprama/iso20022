package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200103 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.012.001.03 Document"`
	Message *AcceptorBatchTransferResponseV03 `xml:"AccptrBtchTrfRspn"`
}

func (d *Document01200103) AddMessage() *AcceptorBatchTransferResponseV03 {
	d.Message = new(AcceptorBatchTransferResponseV03)
	return d.Message
}

// The AcceptorBatchTransferResponse is sent by the acquirer (or its agent) to inform the acceptor (or its agent) of the transfer in a previous AcceptorBatchTransfer of a collection of transactions.
type AcceptorBatchTransferResponseV03 struct {

	// Capture advice response message management information.
	Header *model.Header3 `xml:"Hdr"`

	// Information related to the previously sent set of transaction.
	BatchTransferResponse *model.CardPaymentBatchTransferResponse2 `xml:"BtchTrfRspn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType9 `xml:"SctyTrlr"`
}

func (a *AcceptorBatchTransferResponseV03) AddHeader() *model.Header3 {
	a.Header = new(model.Header3)
	return a.Header
}

func (a *AcceptorBatchTransferResponseV03) AddBatchTransferResponse() *model.CardPaymentBatchTransferResponse2 {
	a.BatchTransferResponse = new(model.CardPaymentBatchTransferResponse2)
	return a.BatchTransferResponse
}

func (a *AcceptorBatchTransferResponseV03) AddSecurityTrailer() *model.ContentInformationType9 {
	a.SecurityTrailer = new(model.ContentInformationType9)
	return a.SecurityTrailer
}
