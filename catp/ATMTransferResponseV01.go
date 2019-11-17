package catp

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700101 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:catp.017.001.01 Document"`
	Message *ATMTransferResponseV01 `xml:"ATMTrfRspn"`
}

func (d *Document01700101) AddMessage() *ATMTransferResponseV01 {
	d.Message = new(ATMTransferResponseV01)
	return d.Message
}

// The ATMTransferResponse message is sent by an acquirer or its agent to inform the ATM of the approval or decline of the transfer transaction.
type ATMTransferResponseV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMTransferResponse *model.ContentInformationType10 `xml:"PrtctdATMTrfRspn,omitempty"`

	// Information related to the response of an ATM transfer from an ATM manager.
	ATMTransferResponse *model.ATMTransferResponse1 `xml:"ATMTrfRspn,omitempty"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMTransferResponseV01) AddHeader() *model.Header31 {
	a.Header = new(model.Header31)
	return a.Header
}

func (a *ATMTransferResponseV01) AddProtectedATMTransferResponse() *model.ContentInformationType10 {
	a.ProtectedATMTransferResponse = new(model.ContentInformationType10)
	return a.ProtectedATMTransferResponse
}

func (a *ATMTransferResponseV01) AddATMTransferResponse() *model.ATMTransferResponse1 {
	a.ATMTransferResponse = new(model.ATMTransferResponse1)
	return a.ATMTransferResponse
}

func (a *ATMTransferResponseV01) AddSecurityTrailer() *model.ContentInformationType15 {
	a.SecurityTrailer = new(model.ContentInformationType15)
	return a.SecurityTrailer
}
