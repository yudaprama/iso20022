package model

// Securitised right for entitlement, for example, equity or bond.
type UnderlyingSecurity1 struct {

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification11 `xml:"SctyId"`
}

func (u *UnderlyingSecurity1) AddSecurityIdentification() *SecurityIdentification11 {
	u.SecurityIdentification = new(SecurityIdentification11)
	return u.SecurityIdentification
}
