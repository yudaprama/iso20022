package model

// Number used to sequence pages when it is not possible for data to be conveyed in a single message and the data has to be split across several pages (messages).
type Pagination struct {

	// Page number.
	PageNumber *Max5NumericText `xml:"PgNb"`

	// Indicates the last page.
	LastPageIndicator *YesNoIndicator `xml:"LastPgInd"`
}

func (p *Pagination) SetPageNumber(value string) {
	p.PageNumber = (*Max5NumericText)(&value)
}

func (p *Pagination) SetLastPageIndicator(value string) {
	p.LastPageIndicator = (*YesNoIndicator)(&value)
}
