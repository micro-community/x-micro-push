package model

//CurrentUser stand for a user
type CurrentUser struct {
	Code         int32    `json:"code"`
	ErrorMessage string   `json:"errorMessage"`
	UserInfo     UserInfo `json:"userInfo"`
}

//UserInfo stand for user
type UserInfo struct {
	UserID       string `json:"userID"`
	UserName     string `json:"userName"`
	Tag          string `json:"tag"`
	UserType     int    `json:"userType"`
	UserPlatform int    `json:"userPlatform"`
	IsForbiden   int    `json:"isForbiden"`
}
