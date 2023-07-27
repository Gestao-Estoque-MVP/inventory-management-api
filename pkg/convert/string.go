package convert

import "database/sql"

func ConvertString(input *string) sql.NullString {
	if input != nil {
		return sql.NullString{String: *input, Valid: true}
	}
	return sql.NullString{Valid: false}
}
