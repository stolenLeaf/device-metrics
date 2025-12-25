package user

type userService struct {
	Repo *UserRepo
}

func NewUserService(repo *UserRepo) *userService {
	return &userService{Repo: repo}
}

func (s *userService) GetUsers() ([]User, error) {
	return s.Repo.GetAll()
}

func (s *userService) CreateUser(u *User) error {
	return s.Repo.Create(u)
}
