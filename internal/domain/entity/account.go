package entity

import (
	"fmt"
	"time"

	"github.com/leonzag/treport/internal/domain/enum"
)

type Account struct {
	OpenedDate  time.Time
	ClosedDate  time.Time
	Id          string
	Name        string
	Type        enum.AccountType
	Status      enum.AccountStatus
	AccessLevel enum.AccessLevel
}

func (a *Account) String() string {
	closedData := fmt.Sprint(a.ClosedDate)
	if a.ClosedDate.IsZero() || a.ClosedDate.Unix() == 0 {
		closedData = "n/a"
	}
	return fmt.Sprintf(`
Account(
	Id:          %s
	Name:        %s
	Type:        %s
	Status:      %s
	AccessLevel: %s
	
	OpenedDate:  %s
	ClosedDate:  %s
)`,
		a.Id,
		a.Name,
		a.Type,
		a.Status,
		a.AccessLevel,
		a.OpenedDate,
		closedData,
	)
}
