package constants

// ProfileStatus is a type for profile status
type ProfileStatus string

// ProfileStatus constants
const (
	ProfileActive   ProfileStatus = "Active"
	ProfileInactive ProfileStatus = "Inactive"
	ProfilePending  ProfileStatus = "Pending"
)