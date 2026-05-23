package lib

import (
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
)

func ToPgText(value *string) pgtype.Text {
	if value == nil {
		return pgtype.Text{}
	}

	return pgtype.Text{String: strings.TrimSpace(*value), Valid: true}
}

func FromPgText(value pgtype.Text) *string {
	if !value.Valid {
		return nil
	}

	bio := value.String
	return &bio
}
