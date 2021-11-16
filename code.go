package errors_map

import (
	"github.com/mongofs/errors"
)

func init() {
	errors.Register(ErrSdkMatchNotExist, 400, "match not found.")
	errors.Register(ErrSdkMatchLineupNotData, 400, "lineup no data.")
	errors.Register(ErrSdkDataNotExist, 400, "no data.")
	errors.Register(ErrSdkTeamNotExist, 400, "team not found.")
	errors.Register(ErrSdkPlayerNotExist, 400, "player not found.")

	errors.Register(ErrSdkSystemDatabaseErr, 500, "database err")
	errors.Register(ErrSdkSystemRedisErr, 500, "redis err")
	errors.Register(ErrSdkSystemJsonEncodeErr, 500, "json encode err")
	errors.Register(ErrSdkSystemJsonDecodeErr, 500, "json decode err")
	errors.Register(ErrSdkSystemProtoEncodeErr, 500, "proto encode err")
	errors.Register(ErrSdkSystemProtoDecodeErr, 500, "proto decode err")
}
