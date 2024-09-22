package activity

type ActivityType string

const (
	GAME     = "GAME"
	ACTIVITY = "ACTIVITY"
)

type Activity struct {
	Typ  ActivityType `firestore:"type"`
	Name string       `firestore:"name"`
}

func NewActivity(typ ActivityType, name string) *Activity {
	return &Activity{
		Typ:  typ,
		Name: name,
	}
}