package mapper

import (
	"github.com/leonzag/treport/internal/domain/entity"
	"github.com/leonzag/treport/internal/domain/enum"
	investapi "github.com/russianinvestments/invest-api-go-sdk/proto"
)

type AccountMapper struct{}

func (m *AccountMapper) TypeToRequest(t enum.AccountType) investapi.AccountType {
	switch t {
	case enum.AccountTypeTINKOFF:
		return investapi.AccountType_ACCOUNT_TYPE_TINKOFF

	case enum.AccountTypeTINKOFF_IIS:
		return investapi.AccountType_ACCOUNT_TYPE_TINKOFF_IIS

	case enum.AccountTypeINVEST_BOX:
		return investapi.AccountType_ACCOUNT_TYPE_INVEST_BOX

	case enum.AccountTypeINVEST_FUND:
		return investapi.AccountType_ACCOUNT_TYPE_INVEST_FUND

	default:
		return investapi.AccountType_ACCOUNT_TYPE_UNSPECIFIED
	}
}

func (m *AccountMapper) TypeToDomain(t investapi.AccountType) enum.AccountType {
	switch t {
	case investapi.AccountType_ACCOUNT_TYPE_TINKOFF:
		return enum.AccountTypeTINKOFF

	case investapi.AccountType_ACCOUNT_TYPE_TINKOFF_IIS:
		return enum.AccountTypeTINKOFF_IIS

	case investapi.AccountType_ACCOUNT_TYPE_INVEST_BOX:
		return enum.AccountTypeINVEST_BOX

	case investapi.AccountType_ACCOUNT_TYPE_INVEST_FUND:
		return enum.AccountTypeINVEST_FUND

	default:
		return enum.AccountTypeUNSPECIFIED
	}
}

func (m *AccountMapper) StatusToRequest(s enum.AccountStatus) investapi.AccountStatus {
	switch s {
	case enum.AccountStatusUNSPECIFIED:
		return investapi.AccountStatus_ACCOUNT_STATUS_UNSPECIFIED

	case enum.AccountStatusNEW:
		return investapi.AccountStatus_ACCOUNT_STATUS_NEW

	case enum.AccountStatusOPEN:
		return investapi.AccountStatus_ACCOUNT_STATUS_OPEN

	case enum.AccountStatusCLOSED:
		return investapi.AccountStatus_ACCOUNT_STATUS_CLOSED

	default:
		return investapi.AccountStatus_ACCOUNT_STATUS_ALL
	}
}

func (m *AccountMapper) StatusToDomain(s investapi.AccountStatus) enum.AccountStatus {
	switch s {
	case investapi.AccountStatus_ACCOUNT_STATUS_UNSPECIFIED:
		return enum.AccountStatusUNSPECIFIED

	case investapi.AccountStatus_ACCOUNT_STATUS_NEW:
		return enum.AccountStatusNEW

	case investapi.AccountStatus_ACCOUNT_STATUS_OPEN:
		return enum.AccountStatusOPEN

	case investapi.AccountStatus_ACCOUNT_STATUS_CLOSED:
		return enum.AccountStatusCLOSED

	default:
		return enum.AccountStatusALL
	}
}

func (m *AccountMapper) AccessLevelToRequest(l enum.AccessLevel) investapi.AccessLevel {
	switch l {
	case enum.AccessLevelFULLACCESS:
		return investapi.AccessLevel_ACCOUNT_ACCESS_LEVEL_FULL_ACCESS

	case enum.AccessLevelREADONLY:
		return investapi.AccessLevel_ACCOUNT_ACCESS_LEVEL_READ_ONLY

	case enum.AccessLevelNOACCESS:
		return investapi.AccessLevel_ACCOUNT_ACCESS_LEVEL_NO_ACCESS

	default:
		return investapi.AccessLevel_ACCOUNT_ACCESS_LEVEL_UNSPECIFIED
	}
}

func (m *AccountMapper) AccessLevelToDomain(l investapi.AccessLevel) enum.AccessLevel {
	switch l {
	case investapi.AccessLevel_ACCOUNT_ACCESS_LEVEL_FULL_ACCESS:
		return enum.AccessLevelFULLACCESS

	case investapi.AccessLevel_ACCOUNT_ACCESS_LEVEL_READ_ONLY:
		return enum.AccessLevelREADONLY

	case investapi.AccessLevel_ACCOUNT_ACCESS_LEVEL_NO_ACCESS:
		return enum.AccessLevelNOACCESS

	default:
		return enum.AccessLevelUNSPECIFIED
	}
}

func (m *AccountMapper) AccountToDomain(a *investapi.Account) *entity.Account {
	return &entity.Account{
		OpenedDate:  a.GetOpenedDate().AsTime(),
		ClosedDate:  a.GetClosedDate().AsTime(),
		Id:          a.GetId(),
		Name:        a.GetName(),
		Type:        m.TypeToDomain(a.GetType()),
		Status:      m.StatusToDomain(a.GetStatus()),
		AccessLevel: m.AccessLevelToDomain(a.GetAccessLevel()),
	}
}
