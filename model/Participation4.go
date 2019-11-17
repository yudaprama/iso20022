package model

// Specifies the level of participation to a shareholders meeting.
type Participation4 struct {

	// Number of rights admitted to the vote.
	TotalNumberOfVotingRights *Number `xml:"TtlNbOfVtngRghts,omitempty"`

	// Percentage of rights participating to the vote versus total voting rights.
	PercentageOfVotingRights *PercentageRate `xml:"PctgOfVtngRghts,omitempty"`

	// Number of securities admitted to the vote, expressed as an amount and a currency.
	TotalNumberOfSecuritiesOutstanding *UnitOrFaceAmount1Choice `xml:"TtlNbOfSctiesOutsdng,omitempty"`

	// Date of calculation of the total number of outstanding securities.
	CalculationDate *ISODate `xml:"ClctnDt,omitempty"`
}

func (p *Participation4) SetTotalNumberOfVotingRights(value string) {
	p.TotalNumberOfVotingRights = (*Number)(&value)
}

func (p *Participation4) SetPercentageOfVotingRights(value string) {
	p.PercentageOfVotingRights = (*PercentageRate)(&value)
}

func (p *Participation4) AddTotalNumberOfSecuritiesOutstanding() *UnitOrFaceAmount1Choice {
	p.TotalNumberOfSecuritiesOutstanding = new(UnitOrFaceAmount1Choice)
	return p.TotalNumberOfSecuritiesOutstanding
}

func (p *Participation4) SetCalculationDate(value string) {
	p.CalculationDate = (*ISODate)(&value)
}
