package util

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func Int32(val string) int32 {
	v, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return int32(v)
}

func ConvertUuid(val uuid.UUID) pgtype.UUID {
	var uuid pgtype.UUID
	uuid.Scan(val.String())
	return uuid
}

func StringToPgUuid(val string) pgtype.UUID {
	var uuid pgtype.UUID
	uuid.Scan(val)
	return uuid
}
