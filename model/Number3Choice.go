package model

// Choice of 3 and 5 exact numeric number.
type Number3Choice struct {

	// Number of maximum 3 numeric text.
	Short *Exact3NumericText `xml:"Shrt"`

	// Number of maximum 5 numeric text. Is only to be used in a delta statement.
	Long *Exact5NumericText `xml:"Lng"`
}

func (n *Number3Choice) SetShort(value string) {
	n.Short = (*Exact3NumericText)(&value)
}

func (n *Number3Choice) SetLong(value string) {
	n.Long = (*Exact5NumericText)(&value)
}
