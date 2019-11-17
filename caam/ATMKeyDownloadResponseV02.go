package caam

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400102 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.02 Document"`
	Message *ATMKeyDownloadResponseV02 `xml:"ATMKeyDwnldRspn"`
}

func (d *Document00400102) AddMessage() *ATMKeyDownloadResponseV02 {
	d.Message = new(ATMKeyDownloadResponseV02)
	return d.Message
}

// The ATMKeyDownloadResponse message is sent from an acquirer to an ATM in response to an ATMKeyDownloadRequest message, to download of one or several cryptographic keys.
type ATMKeyDownloadResponseV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMKeyDownloadResponse *model.ContentInformationType10 `xml:"PrtctdATMKeyDwnldRspn,omitempty"`

	// Information related to the response of an ATM key download from an ATM manager.
	ATMKeyDownloadResponse *model.ATMKeyDownloadResponse2 `xml:"ATMKeyDwnldRspn,omitempty"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType13 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMKeyDownloadResponseV02) AddHeader() *model.Header31 {
	a.Header = new(model.Header31)
	return a.Header
}

func (a *ATMKeyDownloadResponseV02) AddProtectedATMKeyDownloadResponse() *model.ContentInformationType10 {
	a.ProtectedATMKeyDownloadResponse = new(model.ContentInformationType10)
	return a.ProtectedATMKeyDownloadResponse
}

func (a *ATMKeyDownloadResponseV02) AddATMKeyDownloadResponse() *model.ATMKeyDownloadResponse2 {
	a.ATMKeyDownloadResponse = new(model.ATMKeyDownloadResponse2)
	return a.ATMKeyDownloadResponse
}

func (a *ATMKeyDownloadResponseV02) AddSecurityTrailer() *model.ContentInformationType13 {
	a.SecurityTrailer = new(model.ContentInformationType13)
	return a.SecurityTrailer
}
