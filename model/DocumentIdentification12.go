package model

// Provides information about identification of the document.
type DocumentIdentification12 struct {

	// Unique identifier of the document (message) assigned by the sender of the document.
	Identification *Max35Text `xml:"Id"`

	// Date and time at which the document (message) was created by the sender.
	CreationDateTime *DateAndDateTimeChoice `xml:"CreDtTm,omitempty"`

	// Specifies if this document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicate *CopyDuplicate1Code `xml:"CpyDplct,omitempty"`

	// When used in a corporate action instruction, indicates that the current instruction is replacing a previous one that was cancelled earlier. When used in a corporate action instruction cancellation request, indicates that cancelled instruction will be replaced by a new corporate action instruction to be sent later.
	ChangeInstructionIndicator *YesNoIndicator `xml:"ChngInstrInd,omitempty"`
}

func (d *DocumentIdentification12) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification12) AddCreationDateTime() *DateAndDateTimeChoice {
	d.CreationDateTime = new(DateAndDateTimeChoice)
	return d.CreationDateTime
}

func (d *DocumentIdentification12) SetCopyDuplicate(value string) {
	d.CopyDuplicate = (*CopyDuplicate1Code)(&value)
}

func (d *DocumentIdentification12) SetChangeInstructionIndicator(value string) {
	d.ChangeInstructionIndicator = (*YesNoIndicator)(&value)
}
