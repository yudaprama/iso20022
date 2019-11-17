package caam

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300102 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.003.001.02 Document"`
	Message *ATMKeyDownloadRequestV02 `xml:"ATMKeyDwnldReq"`
}

func (d *Document00300102) AddMessage() *ATMKeyDownloadRequestV02 {
	d.Message = new(ATMKeyDownloadRequestV02)
	return d.Message
}

// The ATMKeyDownloadRequest message is sent by an ATM to an ATM manager to initiate the download of one or several cryptographic keys.
type ATMKeyDownloadRequestV02 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header31 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMKeyDownloadRequest *model.ContentInformationType10 `xml:"PrtctdATMKeyDwnldReq,omitempty"`

	// Information related to the request of a key download from an ATM.
	ATMKeyDownloadRequest *model.ATMKeyDownloadRequest2 `xml:"ATMKeyDwnldReq,omitempty"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType13 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMKeyDownloadRequestV02) AddHeader() *model.Header31 {
	a.Header = new(model.Header31)
	return a.Header
}

func (a *ATMKeyDownloadRequestV02) AddProtectedATMKeyDownloadRequest() *model.ContentInformationType10 {
	a.ProtectedATMKeyDownloadRequest = new(model.ContentInformationType10)
	return a.ProtectedATMKeyDownloadRequest
}

func (a *ATMKeyDownloadRequestV02) AddATMKeyDownloadRequest() *model.ATMKeyDownloadRequest2 {
	a.ATMKeyDownloadRequest = new(model.ATMKeyDownloadRequest2)
	return a.ATMKeyDownloadRequest
}

func (a *ATMKeyDownloadRequestV02) AddSecurityTrailer() *model.ContentInformationType13 {
	a.SecurityTrailer = new(model.ContentInformationType13)
	return a.SecurityTrailer
}
