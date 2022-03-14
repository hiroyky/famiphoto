package repositories

import "github.com/volatiletech/sqlboiler/v4/boil"

type SQLExecutor interface {
	boil.ContextExecutor
	boil.ContextBeginner
}
