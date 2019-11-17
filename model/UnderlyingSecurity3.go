package model

// Securitised right for entitlement, for example, equity or bond.
type UnderlyingSecurity3 struct {

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId"`
}

func (u *UnderlyingSecurity3) AddSecurityIdentification() *SecurityIdentification14 {
	u.SecurityIdentification = new(SecurityIdentification14)
	return u.SecurityIdentification
}
