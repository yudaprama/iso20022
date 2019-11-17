package model

// Provides information about the corporate action event.
type CorporateAction14 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate30 `xml:"DtDtls,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat1Choice `xml:"LtryTp,omitempty"`
}

func (c *CorporateAction14) AddDateDetails() *CorporateActionDate30 {
	c.DateDetails = new(CorporateActionDate30)
	return c.DateDetails
}

func (c *CorporateAction14) AddLotteryType() *LotteryTypeFormat1Choice {
	c.LotteryType = new(LotteryTypeFormat1Choice)
	return c.LotteryType
}
