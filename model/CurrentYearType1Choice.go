package model

// Choice between the selections of individual saving accounts issued during the current fiscal year.
type CurrentYearType1Choice struct {

	// Current year ISA is an ISA that was issued during the current fiscal year.
	CurrentYearType *ISAType1Code `xml:"CurYrTp"`

	// Current year ISA is an ISA that was issued during the current fiscal year.
	ExtendedCurrentYearType *Extended350Code `xml:"XtndedCurYrTp"`
}

func (c *CurrentYearType1Choice) SetCurrentYearType(value string) {
	c.CurrentYearType = (*ISAType1Code)(&value)
}

func (c *CurrentYearType1Choice) SetExtendedCurrentYearType(value string) {
	c.ExtendedCurrentYearType = (*Extended350Code)(&value)
}
