package caaa

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200102 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.012.001.02 Document"`
	Message *AcceptorBatchTransferResponseV02 `xml:"AccptrBtchTrfRspn"`
}

func (d *Document01200102) AddMessage() *AcceptorBatchTransferResponseV02 {
	d.Message = new(AcceptorBatchTransferResponseV02)
	return d.Message
}

// The AcceptorBatchTransferResponse is sent by the acquirer (or its agent) to inform the acceptor (or its agent) of the transfer in a previous AcceptorBatchTransfer of a collection of transactions.
type AcceptorBatchTransferResponseV02 struct {

	// Capture advice response message management information.
	Header *model.Header3 `xml:"Hdr"`

	// Information related to the previously sent set of transaction.
	BatchTransferResponse *model.CardPaymentBatchTransferResponse1 `xml:"BtchTrfRspn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType4 `xml:"SctyTrlr"`
}

func (a *AcceptorBatchTransferResponseV02) AddHeader() *model.Header3 {
	a.Header = new(model.Header3)
	return a.Header
}

func (a *AcceptorBatchTransferResponseV02) AddBatchTransferResponse() *model.CardPaymentBatchTransferResponse1 {
	a.BatchTransferResponse = new(model.CardPaymentBatchTransferResponse1)
	return a.BatchTransferResponse
}

func (a *AcceptorBatchTransferResponseV02) AddSecurityTrailer() *model.ContentInformationType4 {
	a.SecurityTrailer = new(model.ContentInformationType4)
	return a.SecurityTrailer
}
