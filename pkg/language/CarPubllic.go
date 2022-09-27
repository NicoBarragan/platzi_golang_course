package language

// carPublic Car with public access
type CarPublic struct {
	Brand string
	Year  int
}

// The declared vars initializing in upper case
// in the struct, means that they are public variables
// for other packages
