package model

// Choice between specification of the balance type in structured or free text format.
type SecuritiesBalanceType1Choice struct {

	// Balance reason in structured format.
	Structured *SecuritiesBalanceType1Code `xml:"Strd"`

	// Balance reason in free text form.
	Unstructured *Max35Text `xml:"Ustrd"`
}

func (s *SecuritiesBalanceType1Choice) SetStructured(value string) {
	s.Structured = (*SecuritiesBalanceType1Code)(&value)
}

func (s *SecuritiesBalanceType1Choice) SetUnstructured(value string) {
	s.Unstructured = (*Max35Text)(&value)
}
