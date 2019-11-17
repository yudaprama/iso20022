package model

// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
type Organisation23 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Name of the organisation in short form.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress []*PostalAddress21 `xml:"PstlAdr"`
}

func (o *Organisation23) SetName(value string) {
	o.Name = (*Max350Text)(&value)
}

func (o *Organisation23) SetShortName(value string) {
	o.ShortName = (*Max35Text)(&value)
}

func (o *Organisation23) AddPostalAddress() *PostalAddress21 {
	newValue := new(PostalAddress21)
	o.PostalAddress = append(o.PostalAddress, newValue)
	return newValue
}
