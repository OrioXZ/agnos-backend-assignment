package dto

type PatientSearchInput struct {
	NationalID  string `form:"national_id" json:"national_id"`
	PassportID  string `form:"passport_id" json:"passport_id"`
	FirstName   string `form:"first_name" json:"first_name"`
	MiddleName  string `form:"middle_name" json:"middle_name"`
	LastName    string `form:"last_name" json:"last_name"`
	DateOfBirth string `form:"date_of_birth" json:"date_of_birth"` // เก็บเป็น string ก่อนก็ได้
	PhoneNumber string `form:"phone_number" json:"phone_number"`
	Email       string `form:"email" json:"email"`
}

type PatientSearchQuery struct {
	HospitalID uint
	Input      PatientSearchInput
}
