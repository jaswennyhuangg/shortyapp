package objects
import "time"

type V1ShortyObjectRequest struct {
	Url string `json:"url"`
	ShortCode string `json:"shortcode"`
	
}

type V1ShortyObjectResponse struct {
	Url string `json:"url"`
	ShortCode string `json:"shortcode"` //the sequence in postman depends on this
	
}

type V1ShortyObjectStatusResponse struct {
	CreatedAt *time.Time `json:"startDate"`
	LastSeen *time.Time `json:"lastSeenDate"`
	RedirectCount int `json:"redirectCount"`
	
	
}
