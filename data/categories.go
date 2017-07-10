package data

type ItemCategories struct {
	ItemCategories []struct {
		Category string `json:"key"`
		Value []struct {
			SubCategory string `json:"key"`
			Base []string `json:"value"`
		} `json:"value"`
	} `json:"itemCategories"`
	Items []string `json:"items"`
}
