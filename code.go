package errors_map

import (
	"github.com/mongofs/errors"
)

func init() {
	errors.Register(ErrMatchNotExist, 400, "match not found.")
	errors.Register(ErrMatchLineupNotData, 400, "lineup no data.")
	errors.Register(ErrDataNotExist, 400, "no data.")
	errors.Register(ErrTeamNotExist, 400, "team not found.")
	errors.Register(ErrPlayerNotExist, 400, "player not found.")

	errors.Register(ErrSystemDatabaseErr, 500, "database err")
	errors.Register(ErrSystemRedisErr, 500, "redis err")
	errors.Register(ErrSystemJsonEncodeErr, 500, "json encode err")
	errors.Register(ErrSystemJsonDecodeErr, 500, "json decode err")
	errors.Register(ErrSystemProtoEncodeErr, 500, "proto encode err")
	errors.Register(ErrSystemProtoDecodeErr, 500, "proto decode err")
}

