package model

// Choice between a matching system unique identification or the related reference.
type MatchingSystemReference1Choice struct {

	// Reference to the unique system identification assigned to the trade  by the central matching system.
	MatchingSystemUniqueReference *Max35Text `xml:"MtchgSysUnqRef"`

	// Refers to the identification of a previous event in the life of a trade which is amended or cancelled.
	RelatedReference *Max35Text `xml:"RltdRef"`
}

func (m *MatchingSystemReference1Choice) SetMatchingSystemUniqueReference(value string) {
	m.MatchingSystemUniqueReference = (*Max35Text)(&value)
}

func (m *MatchingSystemReference1Choice) SetRelatedReference(value string) {
	m.RelatedReference = (*Max35Text)(&value)
}
