package repository

import "github.com/jackc/pgx/v5/pgtype"

// ToPgText converts string to pgtype.Text.
func ToPgText(val string) pgtype.Text {
	if val == "" {
		return pgtype.Text{}
	}

	return pgtype.Text{
		String: val,
		Valid:  true,
	}
}

// FromPgText converts pgtype.Text to string.
func FromPgText(val pgtype.Text) string {
	if !val.Valid {
		return ""
	}

	return val.String
}
