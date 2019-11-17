package model

// Provides details on the legal basis of the request.
type LegalMandate1 struct {

	// Identifies the legal mandate paragraph in law which gives power to the authority's request.
	Paragraph *Max35Text `xml:"Prgrph"`

	// Specifies any additional information describing how or why the paragraph of law should be applied.
	Disclaimer *Max350Text `xml:"Dsclmr,omitempty"`
}

func (l *LegalMandate1) SetParagraph(value string) {
	l.Paragraph = (*Max35Text)(&value)
}

func (l *LegalMandate1) SetDisclaimer(value string) {
	l.Disclaimer = (*Max350Text)(&value)
}
