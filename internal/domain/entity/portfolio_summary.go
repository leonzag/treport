package entity

import (
	"slices"
	"time"
)

type PortfolioSummary struct {
	createdAt time.Time

	Account                   *Account
	Instruments               []*Instrument
	Portfolio                 *Portfolio
	PortfolioPositions        []*PortfolioPosition
	PortfolioPositionsVirtual []*VirtualPortfolioPosition
}

func NewPortfolioSummary(
	account *Account,
	portfolio *Portfolio,
	instruments []*Instrument,
) *PortfolioSummary {
	if account == nil || portfolio == nil {
		return nil
	}
	return &PortfolioSummary{
		createdAt: time.Now(),

		Account:                   account,
		Instruments:               instruments,
		Portfolio:                 portfolio,
		PortfolioPositions:        portfolio.Positions,
		PortfolioPositionsVirtual: portfolio.VirtualPositions,
	}
}

func (p *PortfolioSummary) CreatedAt() time.Time {
	return p.createdAt
}

func (p *PortfolioSummary) InstrumentByUid(uid string) (*Instrument, error) {
	var found *Instrument
	slices.ContainsFunc(p.Instruments, func(i *Instrument) bool {
		if i != nil && i.Uid == uid {
			found = i
			return true
		}
		return false
	})
	if found != nil {
		return found, nil
	}
	return nil, ErrInstrumentNotFound
}
