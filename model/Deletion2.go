package model

// Deletion of the current element.
type Deletion2 struct {

	// Content of the deleted element.
	DeletedValue *Max350Text `xml:"DeltdVal,omitempty"`
}

func (d *Deletion2) SetDeletedValue(value string) {
	d.DeletedValue = (*Max350Text)(&value)
}
