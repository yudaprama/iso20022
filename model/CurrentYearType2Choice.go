package model

// Choice between the selections of individual saving accounts issued during the current fiscal year.
type CurrentYearType2Choice struct {

	// Current year ISA is an ISA that was issued during the current fiscal year.
	CurrentYearType *ISAType2Code `xml:"CurYrTp"`

	// Current year ISA is an ISA that was issued during the current fiscal year.
	ExtendedCurrentYearType *Extended350Code `xml:"XtndedCurYrTp"`
}

func (c *CurrentYearType2Choice) SetCurrentYearType(value string) {
	c.CurrentYearType = (*ISAType2Code)(&value)
}

func (c *CurrentYearType2Choice) SetExtendedCurrentYearType(value string) {
	c.ExtendedCurrentYearType = (*Extended350Code)(&value)
}
