package null

import (
	"database/sql"
)

// SQLString creates a new sql.NullString from a string pointer.
func SQLString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{
			String: *s,
			Valid:  true,
		}
	}
	return sql.NullString{
		Valid: false,
	}
}

// StringValue returns the value of the string pointer passed in.
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}
