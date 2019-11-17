package model

// Information about a document.
type Document10 struct {

	// Type of document.
	DocumentType *UndertakingDocumentType2Choice `xml:"DocTp"`

	// Channel through which the document should be presented.
	PresentationChannel *Channel1Choice `xml:"PresntnChanl,omitempty"`

	// Format of the document.
	DocumentFormat *DocumentFormat1Choice `xml:"DocFrmt,omitempty"`

	// Indication whether the document may be a copy of the original document.
	CopyIndicator *YesNoIndicator `xml:"CpyInd,omitempty"`

	// Indication whether the document must be signed.
	SignedIndicator *YesNoIndicator `xml:"SgndInd,omitempty"`
}

func (d *Document10) AddDocumentType() *UndertakingDocumentType2Choice {
	d.DocumentType = new(UndertakingDocumentType2Choice)
	return d.DocumentType
}

func (d *Document10) AddPresentationChannel() *Channel1Choice {
	d.PresentationChannel = new(Channel1Choice)
	return d.PresentationChannel
}

func (d *Document10) AddDocumentFormat() *DocumentFormat1Choice {
	d.DocumentFormat = new(DocumentFormat1Choice)
	return d.DocumentFormat
}

func (d *Document10) SetCopyIndicator(value string) {
	d.CopyIndicator = (*YesNoIndicator)(&value)
}

func (d *Document10) SetSignedIndicator(value string) {
	d.SignedIndicator = (*YesNoIndicator)(&value)
}
