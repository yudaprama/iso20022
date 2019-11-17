package caam

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300101 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.003.001.01 Document"`
	Message *ATMKeyDownloadRequestV01 `xml:"ATMKeyDwnldReq"`
}

func (d *Document00300101) AddMessage() *ATMKeyDownloadRequestV01 {
	d.Message = new(ATMKeyDownloadRequestV01)
	return d.Message
}

// The ATMKeyDownloadRequest message is sent by an ATM to an ATM manager to initiate the download of one or several cryptographic keys.
type ATMKeyDownloadRequestV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *model.Header20 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMKeyDownloadRequest *model.ContentInformationType10 `xml:"PrtctdATMKeyDwnldReq,omitempty"`

	// Information related to the request of a key download from an ATM.
	ATMKeyDownloadRequest *model.ATMKeyDownloadRequest1 `xml:"ATMKeyDwnldReq,omitempty"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType13 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMKeyDownloadRequestV01) AddHeader() *model.Header20 {
	a.Header = new(model.Header20)
	return a.Header
}

func (a *ATMKeyDownloadRequestV01) AddProtectedATMKeyDownloadRequest() *model.ContentInformationType10 {
	a.ProtectedATMKeyDownloadRequest = new(model.ContentInformationType10)
	return a.ProtectedATMKeyDownloadRequest
}

func (a *ATMKeyDownloadRequestV01) AddATMKeyDownloadRequest() *model.ATMKeyDownloadRequest1 {
	a.ATMKeyDownloadRequest = new(model.ATMKeyDownloadRequest1)
	return a.ATMKeyDownloadRequest
}

func (a *ATMKeyDownloadRequestV01) AddSecurityTrailer() *model.ContentInformationType13 {
	a.SecurityTrailer = new(model.ContentInformationType13)
	return a.SecurityTrailer
}
