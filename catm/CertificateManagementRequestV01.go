package catm

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700101 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.007.001.01 Document"`
	Message *CertificateManagementRequestV01 `xml:"CertMgmtReq"`
}

func (d *Document00700101) AddMessage() *CertificateManagementRequestV01 {
	d.Message = new(CertificateManagementRequestV01)
	return d.Message
}

// The certificate management request message is sent by a POI terminal or any intermediary entity either to a terminal manager acting as a certificate authority for  managing X.509 certificate of a public key owned by the initiating party, or for requesting the inclusion or the removal of the POI to a white list of the terminal manager.
//
type CertificateManagementRequestV01 struct {

	// Information related to the protocol management.
	Header *model.Header29 `xml:"Hdr"`

	// Information related to the request of certificate management.
	CertificateManagementRequest *model.CertificateManagementRequest1 `xml:"CertMgmtReq"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType13 `xml:"SctyTrlr,omitempty"`
}

func (c *CertificateManagementRequestV01) AddHeader() *model.Header29 {
	c.Header = new(model.Header29)
	return c.Header
}

func (c *CertificateManagementRequestV01) AddCertificateManagementRequest() *model.CertificateManagementRequest1 {
	c.CertificateManagementRequest = new(model.CertificateManagementRequest1)
	return c.CertificateManagementRequest
}

func (c *CertificateManagementRequestV01) AddSecurityTrailer() *model.ContentInformationType13 {
	c.SecurityTrailer = new(model.ContentInformationType13)
	return c.SecurityTrailer
}
