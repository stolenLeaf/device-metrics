package user

type UserService struct {
	Repo *UserRepo
}

func NewUserService(repo *UserRepo) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetUsers() ([]User, error) {
	return s.Repo.GetAll()
}

func (s *UserService) CreateUser(u *User) error {
	return s.Repo.Create(u)
}
