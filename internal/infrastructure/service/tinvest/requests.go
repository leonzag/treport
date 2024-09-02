package tinvest

import (
	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/internal/domain/enum"
)

func (c *tinvestService) accounts(status enum.AccountStatus) ([]*entity.Account, error) {
	srv := c.conn.NewUsersServiceClient()
	st := c.mapper.Account.StatusToRequest(status)
	resp, err := srv.GetAccounts(&st)
	if err != nil {
		return nil, err
	}

	accs := make([]*entity.Account, 0, len(resp.Accounts))
	for _, a := range resp.Accounts {
		accs = append(accs, c.mapper.Account.AccountToDomain(a))
	}

	return accs, nil
}

func (c *tinvestService) instrument(uid string) (*entity.Instrument, error) {
	if c.UseCache() {
		if i, ok := c.instrumentFromCache(uid); ok {
			return i, nil
		}
	}

	instr := c.conn.NewInstrumentsServiceClient()
	resp, err := instr.InstrumentByUid(uid)
	if err != nil {
		return nil, err
	}

	i := c.mapper.Instrument.InstrumentToDomain(resp.Instrument)
	if c.UseCache() {
		c.instrumentToCache(uid, i)
	}

	return i, nil
}

func (c *tinvestService) instrumentFromCache(uid string) (*entity.Instrument, bool) {
	c.initCache()
	if i, ok := c.instrumentCache.Get(uid); ok {
		c.logger.Debugf("tinvest: fetch instrument uid=%s from cache", uid)
		return i, true
	}

	return nil, false
}

func (c *tinvestService) instrumentToCache(uid string, i *entity.Instrument) {
	c.logger.Debugf("tinvest: add instrument uid=%s to cache", uid)
	c.initCache()

	c.instrumentCache.Add(uid, i)
}

func (c *tinvestService) portfolio(accId string, crc enum.Currency) (*entity.Portfolio, error) {
	ops := c.conn.NewOperationsServiceClient()
	crcReq := c.mapper.Portfolio.CurrencyToRequest(crc)

	resp, err := ops.GetPortfolio(accId, crcReq)
	if err != nil {
		return nil, err
	}

	return c.mapper.Portfolio.ResponseToDomain(resp), nil
}
