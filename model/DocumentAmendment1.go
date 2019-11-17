package model

// Indicates whether the supporting document is amending an original document or not, and the reference of the original supporting document, when applicable.
type DocumentAmendment1 struct {

	// Provides the correction sequence number used to identify the amendment.
	CorrectionIdentification *Number `xml:"CrrctnId"`

	// Identification of the original document being amended.
	OriginalDocumentIdentification *Max35Text `xml:"OrgnlDocId,omitempty"`
}

func (d *DocumentAmendment1) SetCorrectionIdentification(value string) {
	d.CorrectionIdentification = (*Number)(&value)
}

func (d *DocumentAmendment1) SetOriginalDocumentIdentification(value string) {
	d.OriginalDocumentIdentification = (*Max35Text)(&value)
}
