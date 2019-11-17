package model

// Acquirer of the withdrawal transaction, in charge of the funds settlement with the issuer.
type Acquirer7 struct {

	// Identification of the acquirer.
	AcquiringInstitution *Max35Text `xml:"AcqrgInstn,omitempty"`

	// Identification of the acquirer branch.
	Branch *Max35Text `xml:"Brnch,omitempty"`
}

func (a *Acquirer7) SetAcquiringInstitution(value string) {
	a.AcquiringInstitution = (*Max35Text)(&value)
}

func (a *Acquirer7) SetBranch(value string) {
	a.Branch = (*Max35Text)(&value)
}
