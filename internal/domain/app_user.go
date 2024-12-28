package domain

// AppUser is a registered user of the app. An AppUser may or may not be verified - a process that
// requires a valid email.
type AppUser struct {
	ID    string
	Email string
}
