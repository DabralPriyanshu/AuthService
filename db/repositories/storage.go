package db

type Storage struct {
	UserRepository UserRepository //using interface for loose coupling
}

func NewStorage() *Storage {
	return &Storage{UserRepository: &UserRepositoryImp{}} //dependency injection

}
