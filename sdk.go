// 所有sdk相关错误 板块 02开头，其余看具体code

package errors_map

// sdk : match
const (
	// ErrSdkMatchNotExist - 500: match not found.
	ErrSdkMatchNotExist = iota + 020101

	// ErrSdkMatchLineupNotData - 500: lineup no data.
	ErrSdkMatchLineupNotData

	// ErrSdkDataNotExist - 500: no data.
	ErrSdkDataNotExist
)

// sdk :  team
const (
	// ErrSdkTeamNotExist - 500: team not found.
	ErrSdkTeamNotExist = iota + 020201
)

// sdk :  player
const (
	// ErrSdkPlayerNotExist - 500: player not found.
	ErrSdkPlayerNotExist = iota + 020301
)

// sdk :  params
const (
	// ErrSdkParams - 400: params error.
	ErrSdkParams = iota + 020501
)

//sdk :  system
const (
	// ErrSdkSystemDatabaseErr - 500: database err
	ErrSdkSystemDatabaseErr = iota + 020401

	// ErrSdkSystemRedisErr - 500: redis err
	ErrSdkSystemRedisErr

	// ErrSdkSystemJsonEncodeErr - 500: json encode err
	ErrSdkSystemJsonEncodeErr

	// ErrSdkSystemJsonDecodeErr - 500: json decode err
	ErrSdkSystemJsonDecodeErr

	// ErrSdkSystemProtoEncodeErr - 500: proto encode err
	ErrSdkSystemProtoEncodeErr

	// ErrSdkSystemProtoDecodeErr - 500: proto decode err
	ErrSdkSystemProtoDecodeErr
)
