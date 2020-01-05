package response

type Movie struct {
	ID          int    `json:"id"`
	ImageName   string `json:"imageName"`
	Format      string `json:"format"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DirectLink  string `json:"homepage"`
}
