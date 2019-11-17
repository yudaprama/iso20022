package model

// Choice of format for benchmark curve name.
type BenchmarkCurveName4Choice struct {

	// International Securities Identification Number (ISIN), when it exists for the reference rate.
	ISIN *ISINOct2015Identifier `xml:"ISIN"`

	// Identifier of the index/benchmark of a floating rate bond, when an identifier exists.
	Index *BenchmarkCurveName2Code `xml:"Indx"`

	// Where no identifier exists, standardized name of the index, including its term (such as ‘EURIBOR6M’, ‘LIBOR3M’).
	Name *Max25Text `xml:"Nm"`
}

func (b *BenchmarkCurveName4Choice) SetISIN(value string) {
	b.ISIN = (*ISINOct2015Identifier)(&value)
}

func (b *BenchmarkCurveName4Choice) SetIndex(value string) {
	b.Index = (*BenchmarkCurveName2Code)(&value)
}

func (b *BenchmarkCurveName4Choice) SetName(value string) {
	b.Name = (*Max25Text)(&value)
}
