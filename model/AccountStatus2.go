package model

// Status of an account.
type AccountStatus2 struct {

	// Account can be used for its intended purpose.
	Enabled *EnabledStatusReason1Choice `xml:"Nbld,omitempty"`

	// Account cannot temporarily be used for its intended purpose.
	Disabled *DisabledStatusReason1Choice `xml:"Dsbld,omitempty"`

	// Account change is pending approval.
	Pending *PendingStatusReason1Choice `xml:"Pdg,omitempty"`

	// Account opening is pending.
	PendingOpening *PendingOpeningStatusReason1Choice `xml:"PdgOpng,omitempty"`

	// Account is temporary and can be partially used for its intended purpose. The account will be fully available for use when the account servicer has received all relevant documents.
	Proforma *ProformaStatusReason1Choice `xml:"Profrm,omitempty"`

	// Account is closed.
	Closed *ClosedStatusReason1Choice `xml:"Clsd,omitempty"`

	// Account closure is pending.
	ClosurePending *ClosurePendingStatusReason1Choice `xml:"ClsrPdg,omitempty"`

	// Status is a bilaterally agreed status.
	Other []*OtherAccountStatus1 `xml:"Othr,omitempty"`
}

func (a *AccountStatus2) AddEnabled() *EnabledStatusReason1Choice {
	a.Enabled = new(EnabledStatusReason1Choice)
	return a.Enabled
}

func (a *AccountStatus2) AddDisabled() *DisabledStatusReason1Choice {
	a.Disabled = new(DisabledStatusReason1Choice)
	return a.Disabled
}

func (a *AccountStatus2) AddPending() *PendingStatusReason1Choice {
	a.Pending = new(PendingStatusReason1Choice)
	return a.Pending
}

func (a *AccountStatus2) AddPendingOpening() *PendingOpeningStatusReason1Choice {
	a.PendingOpening = new(PendingOpeningStatusReason1Choice)
	return a.PendingOpening
}

func (a *AccountStatus2) AddProforma() *ProformaStatusReason1Choice {
	a.Proforma = new(ProformaStatusReason1Choice)
	return a.Proforma
}

func (a *AccountStatus2) AddClosed() *ClosedStatusReason1Choice {
	a.Closed = new(ClosedStatusReason1Choice)
	return a.Closed
}

func (a *AccountStatus2) AddClosurePending() *ClosurePendingStatusReason1Choice {
	a.ClosurePending = new(ClosurePendingStatusReason1Choice)
	return a.ClosurePending
}

func (a *AccountStatus2) AddOther() *OtherAccountStatus1 {
	newValue := new(OtherAccountStatus1)
	a.Other = append(a.Other, newValue)
	return newValue
}
