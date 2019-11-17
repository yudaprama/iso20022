package model

// Information that locates and identifies a specific branch of a financial institution.
type BranchData2 struct {

	// Unique and unambiguous identification of a branch of a financial institution.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Name by which an agent is known and which is usually used to identify that agent.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress6 `xml:"PstlAdr,omitempty"`
}

func (b *BranchData2) SetIdentification(value string) {
	b.Identification = (*Max35Text)(&value)
}

func (b *BranchData2) SetName(value string) {
	b.Name = (*Max140Text)(&value)
}

func (b *BranchData2) AddPostalAddress() *PostalAddress6 {
	b.PostalAddress = new(PostalAddress6)
	return b.PostalAddress
}
