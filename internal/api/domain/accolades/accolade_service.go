package accolades

type AccoladeService interface {
	GetUserAccolades(userID string) ([]Accolade, error)
}
