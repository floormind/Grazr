package structs

type User struct {
	ID           int32
	Name         string
	Age          int32
	Gender       string
	Location_Lat float64
	Location_Lng float64
	DietType     string
	Preferences  Preferences // This could either be loaded into the user or be a seperate entity, but if thats's the case we would need to fetch the preferences by userId, for now lets assume we load user preferences when we fetch the user.
}
