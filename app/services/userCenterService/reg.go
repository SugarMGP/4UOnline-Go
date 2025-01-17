package userCenterService

import (
	"net/url"

	"4u-go/app/apiException"
	"4u-go/config/api/userCenterApi"
)

// RegWithoutVerify 用户中心不验证激活用户
func RegWithoutVerify(stuId string, pass string, iid string, email string, userType uint) error {
	userUrl, err := url.Parse(userCenterApi.UCRegWithoutVerify)
	if err != nil {
		return err
	}
	urlPath := userUrl.String()
	regMap := map[string]any{
		"stu_id":       stuId,
		"password":     pass,
		"iid":          iid,
		"email":        email,
		"type":         userType,
		"bound_system": 1,
	}
	resp, err := FetchHandleOfPost(regMap, urlPath)
	if err != nil {
		return err
	}
	return handleRegErrors(resp.Code)
}

// handleRegErrors 根据响应码处理不同的错误
func handleRegErrors(code int) error {
	switch code {
	case 400, 402:
		return apiException.StudentNumAndIidError
	case 401:
		return apiException.PwdError
	case 403:
		return apiException.ReactiveError
	default:
		return nil
	}
}
