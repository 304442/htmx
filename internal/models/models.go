package models

// blog
type Blog struct {
	ID          string `db:"id" json:"id"  form:"id"`
	Title       string `db:"title" json:"title"  form:"title"`
	Content     string `db:"content" json:"content" form:"content"`
	Author      string `db:"author" json:"author" form:"author"`
	Description string `db:"description" json:"description" form:"description"`
	Image       string `db:"image" json:"image" form:"image"`
}

// service
type Service struct {
	ID          string  `db:"id" json:"id"  form:"id"`
	Title       string  `db:"title" json:"title"  form:"title"`
	Description string  `db:"description" json:"decription" form:"decription"`
	Details     string  `db:"details" json:"details" form:"details"`
	Price       float32 `db:"price" json:"price" form:"price"`
	Image       string  `db:"image" json:"image" form:"image"`
	AvgRating   float32 `db:"avgrating"`
	Reviews     []Review
}

// faqs
type FAQ struct {
	Question string `db:"question" json:"question" form:"question"`
	Answer   string `db:"answer" json:"answer" form:"answer"`
}

// question
type Question struct {
	User    string `db:"user" json:"user" form:"user"`
	Subject string `db:"subject" json:"subject" form:"subject"`
}

// appointment
type Appointment struct {
	Name        string `db:"name" json:"name" form:"name"`
	Email       string `db:"email" json:"email" form:"email"`
	Phone       string `db:"phone" json:"phone" form:"phone"`
	JobTittle   string `db:"title" json:"title" form:"title"`
	Company     string `db:"company" json:"company" form:"company"`
	CompanyType string `db:"companytype" json:"companytype" form:"companytype"`
	CompanySize int    `db:"companysizer" json:"ucompanysize" form:"ucompanysize"`
	Industry    string `db:"industry" json:"industry" form:"industry"`
	Country     string `db:"country" json:"country" form:"country"`
	UseCase     string `db:"usecase" json:"usecase" form:"usecase"`
	Description string `db:"description" json:"description" form:"description"`
}

// country
type Country struct {
	Name string `db:"name" json:"name" form:"name"`
}

type Review struct {
	ServiceId string  `db:"service_id" json:"service_id" form:"service_id"`
	User      string  `param:"user" query:"user" form:"user" json:"user" xml:"user" db:"user"`
	Review    string  `param:"review" query:"review" form:"review" json:"review" xml:"review" db:"review"`
	Rating    float32 `param:"rating" query:"rating" form:"rating" json:"rating" xml:"rating" db:"rating"`
}

type ProductReviews struct {
	Reviews   []Review
	AvgRating float32 `db:"avgrating"`
}

type ActiveHours struct {
	ID string `db:"id" json:"id" form:"id"`
	Name string `db:"name" json:"name" form:"name"`
	Email string `db:"email" json:"email" form:"email"`
	Active    string `db:"active" json:"active" form:"active"`
	Taken   bool   `db:"taken" json:"taken" form:"taken"`
}



type User struct {
	ID string `db:"id" json:"id" form:"id"`
	Name string `db:"name" json:"name" form:"name"`
	Email string `db:"email" json:"email" form:"email"`
	Username string `db:"username" json:"username" form:"username"`
	Verified bool `db:"verified" json:"verified" form:"verified"`
}