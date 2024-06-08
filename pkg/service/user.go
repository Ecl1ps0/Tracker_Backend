package service

import (
	"Proctor/models"
	"Proctor/models/DTO"
	"Proctor/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]DTO.UserDTO, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	userDTOs := s.ParseUsersToDTOs(users)

	return userDTOs, nil
}

func (s *UserService) GetProfile(userId uint) (DTO.UserDTO, error) {
	user, err := s.repo.GetProfile(userId)
	if err != nil {
		return DTO.UserDTO{}, err
	}

	return s.UserToDTO(user), nil
}

func (s *UserService) GetAllStudents() ([]DTO.UserDTO, error) {
	students, err := s.repo.GetAllStudents()
	if err != nil {
		return nil, err
	}

	studentsDTOs := s.ParseUsersToDTOs(students)

	return studentsDTOs, nil
}

func (s *UserService) GetStudentBySolutionID(id uint) (DTO.UserDTO, error) {
	student, err := s.repo.GetStudentBySolutionID(id)
	if err != nil {
		return DTO.UserDTO{}, err
	}

	return s.UserToDTO(student), nil
}

func (s *UserService) GetStudentsByTeacherID(id uint) ([]DTO.UserDTO, error) {
	students, err := s.repo.GetStudentsByTeacherID(id)
	if err != nil {
		return nil, err
	}

	studentsDTOs := s.ParseUsersToDTOs(students)

	return studentsDTOs, nil
}

func (s *UserService) AddStudentToTask(studentId, taskId uint) error {
	return s.repo.AddStudentToTask(models.StudentTask{
		StudentID: studentId,
		TaskID:    taskId,
	})
}

func (s *UserService) GetRoleByUserID(id uint) (uint, error) {
	return s.repo.GetRoleByID(id)
}

func (s *UserService) UserToDTO(user models.User) DTO.UserDTO {
	return DTO.UserDTO{
		ID:      user.ID,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		RoleID:  user.RoleID,
	}
}

func (s *UserService) ParseUsersToDTOs(users []models.User) []DTO.UserDTO {
	var userDTOs []DTO.UserDTO
	for _, user := range users {
		userDTOs = append(userDTOs, s.UserToDTO(user))
	}

	return userDTOs
}
