package filters

import (
	"fmt"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserFilter struct {
	UserID *string
}

func (f *UserFilter) WhereMods() []qm.QueryMod {
	var filter []qm.QueryMod
	if f == nil {
		return filter
	}
	if f.UserID != nil {
		filter = append(filter, qm.Where(fmt.Sprintf("%s = ?", dbmodels.UserColumns.UserID), f.UserID))
	}
	return filter
}
