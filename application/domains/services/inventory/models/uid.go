package models

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type UID string

func (u *UID) Scan(v interface{}) error {
	if v == nil {
		*u = ""
		return nil
	}

	switch x := v.(type) {
	case []byte:
		*u = UID(string(x))
	case string:
		*u = UID(x)
	case UID:
		*u = x
	case *UID:
		if x == nil {
			*u = ""
		} else {
			*u = *x
		}
	default:
		return fmt.Errorf("%T: %v:  %w", v, x, errors.New("failed to parse UID type, unknown type"))
	}
	return nil
}

func (UID) GormDataType() string {
	return "uid"
}

func (UID) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "TEXT"
}

// Value implements the driver.Valuer interface.
func (u UID) Value() (driver.Value, error) {
	return string(u), nil
}

func (u UID) String() string {
	return string(u)
}

func (u UID) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	if u.String() != "" {
		return clause.Expr{
			SQL:  `?`,
			Vars: []interface{}{u.String()},
		}
	} else {
		return clause.Expr{
			SQL: "concat('VN', lpad(nextval(uid_seq),11,'0'))",
		}
	}
}
