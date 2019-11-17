package model

// Information that locates and identifies a specific branch of a financial institution.
type BranchData struct {

	// Unique and unambiguous identification of a branch of a financial institution.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress1 `xml:"PstlAdr,omitempty"`
}

func (b *BranchData) SetIdentification(value string) {
	b.Identification = (*Max35Text)(&value)
}

func (b *BranchData) SetName(value string) {
	b.Name = (*Max35Text)(&value)
}

func (b *BranchData) AddPostalAddress() *PostalAddress1 {
	b.PostalAddress = new(PostalAddress1)
	return b.PostalAddress
}
