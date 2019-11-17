package model

// Information related to the result of the certificate management request.
type CertificateManagementResponse1 struct {

	// Identification of the terminal or system using the certificate management service.
	POIIdentification *GenericIdentification72 `xml:"POIId"`

	// Identification of the TM or the MTM providing the Certificate Authority service.
	TMIdentification *GenericIdentification72 `xml:"TMId,omitempty"`

	// Requested certificate management service.
	CertificateService *CardPaymentServiceType10Code `xml:"CertSvc"`

	// Outcome of the certificate service processing.
	Result *ResponseType6 `xml:"Rslt"`

	// Identification of the security profile, for creation, renewal or revocation of certificate.
	SecurityProfile *Max35Text `xml:"SctyPrfl,omitempty"`

	// Created or renewed certificate. The certificate is ASN.1/DER encoded.
	ClientCertificate *Max3000Binary `xml:"ClntCert,omitempty"`

	// Certificate of the client certificate path, from the CA (Certificate Authority) certificate, to the root certificate, for renewal or revocation of certificate.
	ClientCertificatePath []*Max10KBinary `xml:"ClntCertPth,omitempty"`

	// Certificate of the server certificate path, from the CA (Certificate Authority) certificate, to the root certificate, for renewal or revocation of certificate.
	ServerCertificatePath []*Max10KBinary `xml:"SvrCertPth,omitempty"`
}

func (c *CertificateManagementResponse1) AddPOIIdentification() *GenericIdentification72 {
	c.POIIdentification = new(GenericIdentification72)
	return c.POIIdentification
}

func (c *CertificateManagementResponse1) AddTMIdentification() *GenericIdentification72 {
	c.TMIdentification = new(GenericIdentification72)
	return c.TMIdentification
}

func (c *CertificateManagementResponse1) SetCertificateService(value string) {
	c.CertificateService = (*CardPaymentServiceType10Code)(&value)
}

func (c *CertificateManagementResponse1) AddResult() *ResponseType6 {
	c.Result = new(ResponseType6)
	return c.Result
}

func (c *CertificateManagementResponse1) SetSecurityProfile(value string) {
	c.SecurityProfile = (*Max35Text)(&value)
}

func (c *CertificateManagementResponse1) SetClientCertificate(value string) {
	c.ClientCertificate = (*Max3000Binary)(&value)
}

func (c *CertificateManagementResponse1) AddClientCertificatePath(value string) {
	c.ClientCertificatePath = append(c.ClientCertificatePath, (*Max10KBinary)(&value))
}

func (c *CertificateManagementResponse1) AddServerCertificatePath(value string) {
	c.ServerCertificatePath = append(c.ServerCertificatePath, (*Max10KBinary)(&value))
}
