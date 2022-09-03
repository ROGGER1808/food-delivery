package common

const (
	DbTypeRestaurant = 1
	DbTypeFood       = 2
	DbTypeCategory   = 3
	DbTypeUser       = 4
	DbTypeUpload     = 5
)

const CurrentUser = "user"

const (
	TopicUserLikeRestaurant    = "UserLikeRestaurant"
	TopicUserDislikeRestaurant = "UserDislikeRestaurant"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
