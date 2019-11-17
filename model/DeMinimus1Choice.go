package model

// De minimus applicability conditions.
type DeMinimus1Choice struct {

	// Conditions applicable when the investor is covered by the "de minimis" exemption.
	DeMinimusApplicable *DeMinimusApplicable1 `xml:"DeMnmsAplbl"`

	// Conditions applicable when the investor is not covered by the "de minimis" exemption.
	DeMinimusNotApplicable *DeMinimusNotApplicable1 `xml:"DeMnmsNotAplbl"`
}

func (d *DeMinimus1Choice) AddDeMinimusApplicable() *DeMinimusApplicable1 {
	d.DeMinimusApplicable = new(DeMinimusApplicable1)
	return d.DeMinimusApplicable
}

func (d *DeMinimus1Choice) AddDeMinimusNotApplicable() *DeMinimusNotApplicable1 {
	d.DeMinimusNotApplicable = new(DeMinimusNotApplicable1)
	return d.DeMinimusNotApplicable
}
