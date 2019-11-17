package model

// Information related to registration of securities.
type RegistrationParameters5 struct {

	// Identification assigned to a deposit.
	CertificationIdentification *RestrictedFINXMax16Text `xml:"CertfctnId,omitempty"`

	// Date/time at which the certificates in the deposit were validated by the agent.
	CertificationDateTime *DateAndDateTimeChoice `xml:"CertfctnDtTm,omitempty"`

	// Account at the registrar where financial instruments are registered.
	RegistrarAccount *RestrictedFINXMax35Text `xml:"RegarAcct,omitempty"`

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	CertificateNumber []*SecuritiesCertificate5 `xml:"CertNb,omitempty"`
}

func (r *RegistrationParameters5) SetCertificationIdentification(value string) {
	r.CertificationIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *RegistrationParameters5) AddCertificationDateTime() *DateAndDateTimeChoice {
	r.CertificationDateTime = new(DateAndDateTimeChoice)
	return r.CertificationDateTime
}

func (r *RegistrationParameters5) SetRegistrarAccount(value string) {
	r.RegistrarAccount = (*RestrictedFINXMax35Text)(&value)
}

func (r *RegistrationParameters5) AddCertificateNumber() *SecuritiesCertificate5 {
	newValue := new(SecuritiesCertificate5)
	r.CertificateNumber = append(r.CertificateNumber, newValue)
	return newValue
}
