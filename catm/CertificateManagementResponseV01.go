package catm

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00800101 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.008.001.01 Document"`
	Message *CertificateManagementResponseV01 `xml:"CertMgmtRspn"`
}

func (d *Document00800101) AddMessage() *CertificateManagementResponseV01 {
	d.Message = new(CertificateManagementResponseV01)
	return d.Message
}

// The CertificateManagementResponse is sent by a terminal manager in response to a CertificateManagementRequest to provide the outcome of the requested service.
type CertificateManagementResponseV01 struct {

	// Information related to the protocol management.
	Header *model.Header29 `xml:"Hdr"`

	// Information related to the result of the certificate management request.
	CertificateManagementResponse *model.CertificateManagementResponse1 `xml:"CertMgmtRspn"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *model.ContentInformationType13 `xml:"SctyTrlr,omitempty"`
}

func (c *CertificateManagementResponseV01) AddHeader() *model.Header29 {
	c.Header = new(model.Header29)
	return c.Header
}

func (c *CertificateManagementResponseV01) AddCertificateManagementResponse() *model.CertificateManagementResponse1 {
	c.CertificateManagementResponse = new(model.CertificateManagementResponse1)
	return c.CertificateManagementResponse
}

func (c *CertificateManagementResponseV01) AddSecurityTrailer() *model.ContentInformationType13 {
	c.SecurityTrailer = new(model.ContentInformationType13)
	return c.SecurityTrailer
}
