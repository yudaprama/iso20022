package model

// Conditions applicable when the investor is not covered by the "de minimis" exemption.
type DeMinimusNotApplicable1 struct {

	// Reason for the restricted person.
	RestrictedPersonReason *Max350Text `xml:"RstrctdPrsnRsn"`
}

func (d *DeMinimusNotApplicable1) SetRestrictedPersonReason(value string) {
	d.RestrictedPersonReason = (*Max350Text)(&value)
}
