package subscriber

type HasRestaurantId interface {
	GetRestaurantId() int
}

type HasFoodId interface {
	GetFoodId() int
}
