package common

const (
	DbTypeRestaurant = 1
	DbTypeFood       = 2
	DbTypeCategory   = 3
	DbTypeUser       = 4
	DbTypeUpload     = 5
	DbTypeFoodRating = 5
)

const CurrentUser = "user"

const (
	TopicUserLikeRestaurant    = "UserLikeRestaurant"
	TopicUserDislikeRestaurant = "UserDislikeRestaurant"
	TopicUserLikeFood          = "UserLikeFood"
	TopicUserDislikeFood       = "UserDislikeFood"
	TopicUserRatingFood        = "UserRatingFood"
	TopicUserUpdateRatingFood  = "UserUpadteRatingFood"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
