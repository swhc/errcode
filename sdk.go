// 所有sdk相关错误 板块 02开头，其余看具体code

package errors_map

// sdk : match
const (
	// ErrMatchNotExist - 400: match not found.
	ErrMatchNotExist = iota + 020101

	// ErrMatchLineupNotData - 400: lineup no data.
	ErrMatchLineupNotData

	// ErrDataNotExist - 400: no data.
	ErrDataNotExist
)

// sdk :  team
const (
	// ErrTeamNotExist - 400: team not found.
	ErrTeamNotExist = iota + 020201
)

// sdk :  player
const (
	// ErrPlayerNotExist - 400: player not found.
	ErrPlayerNotExist = iota + 020301
)

//sdk :  system
const (
	// ErrSystemDatabaseErr - 500: database err
	ErrSystemDatabaseErr = iota + 020401

	// ErrSystemRedisErr - 500: redis err
	ErrSystemRedisErr

	// ErrSystemJsonEncodeErr - 500: json encode err
	ErrSystemJsonEncodeErr

	// ErrSystemJsonDecodeErr - 500: json decode err
	ErrSystemJsonDecodeErr

	// ErrSystemProtoEncodeErr - 500: proto encode err
	ErrSystemProtoEncodeErr

	// ErrSystemProtoDecodeErr - 500: proto decode err
	ErrSystemProtoDecodeErr
)

