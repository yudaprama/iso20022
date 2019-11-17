package model

// Information related to registration of securities.
type RegistrationParameters4 struct {

	// Identification assigned to a deposit.
	CertificationIdentification *Max35Text `xml:"CertfctnId,omitempty"`

	// Date/time at which the certificates in the deposit were validated by the agent.
	CertificationDateTime *DateAndDateTimeChoice `xml:"CertfctnDtTm,omitempty"`

	// Account at the registrar where financial instruments are registered.
	RegistrarAccount *Max35Text `xml:"RegarAcct,omitempty"`

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	CertificateNumber []*SecuritiesCertificate4 `xml:"CertNb,omitempty"`
}

func (r *RegistrationParameters4) SetCertificationIdentification(value string) {
	r.CertificationIdentification = (*Max35Text)(&value)
}

func (r *RegistrationParameters4) AddCertificationDateTime() *DateAndDateTimeChoice {
	r.CertificationDateTime = new(DateAndDateTimeChoice)
	return r.CertificationDateTime
}

func (r *RegistrationParameters4) SetRegistrarAccount(value string) {
	r.RegistrarAccount = (*Max35Text)(&value)
}

func (r *RegistrationParameters4) AddCertificateNumber() *SecuritiesCertificate4 {
	newValue := new(SecuritiesCertificate4)
	r.CertificateNumber = append(r.CertificateNumber, newValue)
	return newValue
}
