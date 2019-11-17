package model

// Choice between a standard code or proprietary code to specify the type of lottery.
type LotteryTypeFormat1Choice struct {

	// Standard code to specify the type of lottery announced.
	Code *LotteryType1Code `xml:"Cd"`

	// Proprietary identification of the type of lottery announced.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (l *LotteryTypeFormat1Choice) SetCode(value string) {
	l.Code = (*LotteryType1Code)(&value)
}

func (l *LotteryTypeFormat1Choice) AddProprietary() *GenericIdentification20 {
	l.Proprietary = new(GenericIdentification20)
	return l.Proprietary
}
