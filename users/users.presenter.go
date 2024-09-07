package users

import "github.com/9thDuck/ecommerce-api.git/common"

type userResponse common.Response

func successResponse(message string, data *User) userResponse {
	return userResponse{message, data}
}

func failedResopnse(errorStr string) userResponse {
	return userResponse{errorStr, nil}
}
