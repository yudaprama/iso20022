package model

// Choice between pre-defined and user-defined sorting criteria.
type SortCriteria1Choice struct {

	// Pre-defined way of indicating how the information is broken down.
	Predefined *ReportSortedType1Code `xml:"Prdfnd"`

	// User-defined way of indicating how the information is broken down.
	UserDefined *DataFormat1Choice `xml:"UsrDfnd"`
}

func (s *SortCriteria1Choice) SetPredefined(value string) {
	s.Predefined = (*ReportSortedType1Code)(&value)
}

func (s *SortCriteria1Choice) AddUserDefined() *DataFormat1Choice {
	s.UserDefined = new(DataFormat1Choice)
	return s.UserDefined
}
