package model

// Gives the name and the reference of the query.
type QueryReference struct {

	// Unique and unambiguous identification of the query.
	QueryReference *Max35Text `xml:"QryRef"`

	// Name of the query.
	QueryName *Max35Text `xml:"QryNm,omitempty"`
}

func (q *QueryReference) SetQueryReference(value string) {
	q.QueryReference = (*Max35Text)(&value)
}

func (q *QueryReference) SetQueryName(value string) {
	q.QueryName = (*Max35Text)(&value)
}
