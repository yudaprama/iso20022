package model

// Information related to registration of securities.
type RegistrationParameters3 struct {

	// Identification assigned to a deposit.
	CertificationIdentification *Max35Text `xml:"CertfctnId,omitempty"`

	// Date/time at which the certificates in the deposit were validated by the agent.
	CertificationDateTime *DateAndDateTime1Choice `xml:"CertfctnDtTm,omitempty"`

	// Account at the registrar where financial instruments are registered.
	RegistrarAccount *Max35Text `xml:"RegarAcct,omitempty"`

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	CertificateNumber []*SecuritiesCertificate3 `xml:"CertNb,omitempty"`
}

func (r *RegistrationParameters3) SetCertificationIdentification(value string) {
	r.CertificationIdentification = (*Max35Text)(&value)
}

func (r *RegistrationParameters3) AddCertificationDateTime() *DateAndDateTime1Choice {
	r.CertificationDateTime = new(DateAndDateTime1Choice)
	return r.CertificationDateTime
}

func (r *RegistrationParameters3) SetRegistrarAccount(value string) {
	r.RegistrarAccount = (*Max35Text)(&value)
}

func (r *RegistrationParameters3) AddCertificateNumber() *SecuritiesCertificate3 {
	newValue := new(SecuritiesCertificate3)
	r.CertificateNumber = append(r.CertificateNumber, newValue)
	return newValue
}
