package users

import "github.com/9thDuck/ecommerce-api.git/common"

type userResponse common.Response

func successSignupResponse(data *User) userResponse {
	return userResponse{Data: data}
}

func failedSignupResopnse(errorStr string) userResponse {
	return userResponse{nil, errorStr}
}
