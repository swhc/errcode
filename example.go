package errcode

//go:generate codegen -type=int

// 基础服务 : user errors.
const (
	// ErrUserNotFound - 404: User not found.
	ErrUserNotFound int = iota + 110001

	// ErrUserAlreadyExist - 400: User already exist.
	ErrUserAlreadyExist
)

// 基础服务: secret errors.
const (
	// ErrEncrypt - 400: Secret reach the max count.
	ErrReachMaxCount int = iota + 110101

	//  ErrSecretNotFound - 404: Secret not found.
	ErrSecretNotFound
)

// 基础服务: policy errors.
const (
	// ErrPolicyNotFound - 404: Policy not found.
	ErrPolicyNotFound int = iota + 110201
)
