package model

// Information related to the request of certificate management.
type CertificateManagementRequest1 struct {

	// Identification of the terminal or system using the certificate management service.
	POIIdentification *GenericIdentification72 `xml:"POIId"`

	// Identification of the TM or the MTM providing the Certificate Authority service.
	TMIdentification *GenericIdentification72 `xml:"TMId,omitempty"`

	// Requested certificate management service.
	CertificateService *CardPaymentServiceType10Code `xml:"CertSvc"`

	// Identification of the client and server public key infrastructures containing the certificate. In addition, it may identify specific requirements of the customer.
	SecurityDomain *Max70Text `xml:"SctyDomn,omitempty"`

	// PKCS#10 (Public Key Certificate Standard 10) certification request coded in base64 ASN.1/DER (Abstract Syntax Notation 1, Distinguished Encoding Rules) or PEM (Privacy Enhanced Message) format.
	BinaryCertificationRequest *Max20000Text `xml:"BinryCertfctnReq,omitempty"`

	// Certification request PKCS#10 (Public Key Certificate Standard 10) for creation or renewal of an X.509 certificate.
	CertificationRequest *CertificationRequest1 `xml:"CertfctnReq,omitempty"`

	// Created certificate. The certificate is ASN.1/DER encoded, for renewal or revocation of certificate.
	ClientCertificate *Max10KBinary `xml:"ClntCert,omitempty"`

	// Identification of the white list element, for white list addition or removal.
	WhiteListIdentification *PointOfInteraction6 `xml:"WhtListId,omitempty"`
}

func (c *CertificateManagementRequest1) AddPOIIdentification() *GenericIdentification72 {
	c.POIIdentification = new(GenericIdentification72)
	return c.POIIdentification
}

func (c *CertificateManagementRequest1) AddTMIdentification() *GenericIdentification72 {
	c.TMIdentification = new(GenericIdentification72)
	return c.TMIdentification
}

func (c *CertificateManagementRequest1) SetCertificateService(value string) {
	c.CertificateService = (*CardPaymentServiceType10Code)(&value)
}

func (c *CertificateManagementRequest1) SetSecurityDomain(value string) {
	c.SecurityDomain = (*Max70Text)(&value)
}

func (c *CertificateManagementRequest1) SetBinaryCertificationRequest(value string) {
	c.BinaryCertificationRequest = (*Max20000Text)(&value)
}

func (c *CertificateManagementRequest1) AddCertificationRequest() *CertificationRequest1 {
	c.CertificationRequest = new(CertificationRequest1)
	return c.CertificationRequest
}

func (c *CertificateManagementRequest1) SetClientCertificate(value string) {
	c.ClientCertificate = (*Max10KBinary)(&value)
}

func (c *CertificateManagementRequest1) AddWhiteListIdentification() *PointOfInteraction6 {
	c.WhiteListIdentification = new(PointOfInteraction6)
	return c.WhiteListIdentification
}
