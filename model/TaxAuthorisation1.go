package model

// Details of the authorised tax paying party.
type TaxAuthorisation1 struct {

	// Title or position of debtor or the debtor's authorised representative.
	Title *Max35Text `xml:"Titl,omitempty"`

	// Name of the debtor or the debtor's authorised representative.
	Name *Max140Text `xml:"Nm,omitempty"`
}

func (t *TaxAuthorisation1) SetTitle(value string) {
	t.Title = (*Max35Text)(&value)
}

func (t *TaxAuthorisation1) SetName(value string) {
	t.Name = (*Max140Text)(&value)
}
