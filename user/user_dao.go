package user

// DAO is DAO of User information
type DAO struct{}

// NewUserDAO instantiates UserDAO
func NewUserDAO() *DAO {
	return &DAO{}
}

// Register creates new user
func (dao DAO) Register() User {
	return User{
		ID:    1,
		Name:  "Brown",
		Email: "brown@cafe.com",
	}
}
