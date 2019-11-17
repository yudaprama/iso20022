package model

// Computer file stored in a binary format.
type BinaryFile1 struct {

	// Code specifying the Multipurpose Internet Mail Extensions (MIME) type for this attached binary file. Reference IANA (Internet Assigned Numbers Authority) - MIME Media Types (www.iana.org/assignments/media-types).
	MIMEType *Max35Text `xml:"MIMETp,omitempty"`

	// Specifies the encoding algorithm used for this attached binary file. Reference IANA (Internet Assigned Numbers Authority) - Transfer Encodings (www.iana.org/assignments/transfer-encodings).
	EncodingType *Max35Text `xml:"NcodgTp,omitempty"`

	// Specifies a code signifying the particular character set used for this attached binary file. Reference IANA (Internet Assigned Numbers Authority) - Character Sets (www.iana.org/assignments/character-sets).
	CharacterSet *Max35Text `xml:"CharSet,omitempty"`

	// Binary object included in this attached binary file.
	IncludedBinaryObject *Max100KBinary `xml:"InclBinryObjct,omitempty"`
}

func (b *BinaryFile1) SetMIMEType(value string) {
	b.MIMEType = (*Max35Text)(&value)
}

func (b *BinaryFile1) SetEncodingType(value string) {
	b.EncodingType = (*Max35Text)(&value)
}

func (b *BinaryFile1) SetCharacterSet(value string) {
	b.CharacterSet = (*Max35Text)(&value)
}

func (b *BinaryFile1) SetIncludedBinaryObject(value string) {
	b.IncludedBinaryObject = (*Max100KBinary)(&value)
}
