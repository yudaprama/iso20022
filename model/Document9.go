package model

// Information about a document.
type Document9 struct {

	// Type of document or template.
	Type *UndertakingDocumentType1Choice `xml:"Tp"`

	// Identification of the document or template.
	Identification *Max35Text `xml:"Id"`

	// Format of the document  or template, such as PDF, XML, XSLT.
	Format *DocumentFormat1Choice `xml:"Frmt,omitempty"`

	// Binary file representing the enclosed document or template, such as a PDF file, image file, XML file, MT message.
	Enclosure *Max2MBBinary `xml:"Nclsr"`

	// Digital signature of the enclosed binary file.
	DigitalSignature *PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (d *Document9) AddType() *UndertakingDocumentType1Choice {
	d.Type = new(UndertakingDocumentType1Choice)
	return d.Type
}

func (d *Document9) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *Document9) AddFormat() *DocumentFormat1Choice {
	d.Format = new(DocumentFormat1Choice)
	return d.Format
}

func (d *Document9) SetEnclosure(value string) {
	d.Enclosure = (*Max2MBBinary)(&value)
}

func (d *Document9) AddDigitalSignature() *PartyAndSignature2 {
	d.DigitalSignature = new(PartyAndSignature2)
	return d.DigitalSignature
}
