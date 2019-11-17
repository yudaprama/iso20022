package model

// Conditions applicable when the investor is covered by the "de minimis" exemption.
type DeMinimusApplicable1 struct {

	// Indicates whether the investor permits its beneficial owners that are restricted persons, if any, to participate in profits and losses allocated to the investor that are attribute to new issue securities.
	NewIssuePermission *YesNoIndicator `xml:"NewIssePrmssn"`

	// Percentage of the new issue profits and losses that it receives to beneficial owners that are restricted persons.
	Percentage *PercentageRate `xml:"Pctg,omitempty"`
}

func (d *DeMinimusApplicable1) SetNewIssuePermission(value string) {
	d.NewIssuePermission = (*YesNoIndicator)(&value)
}

func (d *DeMinimusApplicable1) SetPercentage(value string) {
	d.Percentage = (*PercentageRate)(&value)
}
