package models

// Movie for modeling data dummy
type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

//A model sets a blueprint of API. it sets the data and its values. This adds the data
// tyopes to each movie property that the API will use.
