package model

// Describes the comparison between the currently established baseline elements and the proposed ones.
type ComparisonResult2 struct {

	// Sequence number assigned to the element.
	ElementSequenceNumber *Number `xml:"ElmtSeqNb"`

	// Specifies from the root of the message the complete path of the element.
	ElementPath *Max350Text `xml:"ElmtPth"`

	// Name of the element.
	ElementName *Max35Text `xml:"ElmtNm"`

	// Replacement of an existing content by a different one
	Replacement *Replacement2 `xml:"Rplcmnt"`

	// Deletion of the current element.
	Deletion *Deletion2 `xml:"Deltn"`

	// Addition of a new element.
	Addition *Addition2 `xml:"Addtn"`
}

func (c *ComparisonResult2) SetElementSequenceNumber(value string) {
	c.ElementSequenceNumber = (*Number)(&value)
}

func (c *ComparisonResult2) SetElementPath(value string) {
	c.ElementPath = (*Max350Text)(&value)
}

func (c *ComparisonResult2) SetElementName(value string) {
	c.ElementName = (*Max35Text)(&value)
}

func (c *ComparisonResult2) AddReplacement() *Replacement2 {
	c.Replacement = new(Replacement2)
	return c.Replacement
}

func (c *ComparisonResult2) AddDeletion() *Deletion2 {
	c.Deletion = new(Deletion2)
	return c.Deletion
}

func (c *ComparisonResult2) AddAddition() *Addition2 {
	c.Addition = new(Addition2)
	return c.Addition
}
