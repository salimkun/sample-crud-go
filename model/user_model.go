package model

type Occupation struct {
	Name        string   `json:"company_name"`
	Position    string   `json:"occupation_position"`
	StartDate   string   `json:"occupation_start"`
	EndDate     string   `json:"occupation_end"`
	Status      string   `json:"occupation_status"`
	Achievement []string `json:"occupation_achievement"`
}

type Education struct {
	Name      string  `json:"education_name"`
	Degree    string  `json:"education_degree"`
	Faculty   string  `json:"education_faculty"`
	City      string  `json:"education_city"`
	StartDate string  `json:"education_start"`
	EndDate   string  `json:"education_end"`
	Score     float32 `json:"education_score"`
}

type User struct {
	ID            int64        `json:"id"`
	Name          string       `json:"name"`
	Email         string       `json:"email"`
	Phone         string       `json:"phone_number"`
	LinkedInUrl   string       `json:"linkedin_url"`
	PortofolioUrl string       `json:"portofolio_url"`
	Occupations   []Occupation `json:"occupations"`
	Educations    []Education  `json:"educations"`
	Achievement   []string     `json:"achievement"`
}
