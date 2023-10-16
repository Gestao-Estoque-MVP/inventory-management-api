package convert

import (
	"database/sql"
	"fmt"
	"math/big"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ConvertString(input *string) sql.NullString {
	if input != nil {
		return sql.NullString{String: *input, Valid: true}
	}
	return sql.NullString{Valid: false}
}

func UUIDToString(uuidBytes pgtype.UUID) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuidBytes.Bytes[0:4], uuidBytes.Bytes[4:6], uuidBytes.Bytes[6:8], uuidBytes.Bytes[8:10], uuidBytes.Bytes[10:16])
}

func ConvertUUIDs(uuidStrings []string) ([][16]byte, error) {
	var uuidBytesArray [][16]byte

	for _, uuidStr := range uuidStrings {
		id, err := uuid.FromString(uuidStr)
		if err != nil {
			return nil, fmt.Errorf("Erro ao converter string para UUID: %v", err)
		}
		uuidBytesArray = append(uuidBytesArray, id)
	}

	return uuidBytesArray, nil
}

func StringToByte16(s string) [16]byte {
	var b [16]byte
	copy(b[:], s)
	return b
}

func StringToPgUUID(s string) pgtype.UUID {
	var uuid pgtype.UUID
	if err := uuid.Scan(s); err != nil {
		return uuid
	}
	return uuid
}

func ConvertFloat64ToNumeric(value float64) pgtype.Numeric {
	numeric := pgtype.Numeric{}
	f := new(big.Float).SetFloat64(value)
	i, _ := f.Int(nil)
	numeric.Int.Set(i)
	return numeric
}
