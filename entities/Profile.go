package entities

type ProfileStatus string

type Profile struct {
	Username string						`json:"username"`
	Email string						`json:"email"`
	Password string						`json:"password"`
}