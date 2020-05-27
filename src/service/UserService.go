package service

import (
	"RBAC_GO/src/dao"
	"RBAC_GO/src/models"
)

type UserService interface {
	dao.UserDao
}

type userService struct {
	userDao     dao.UserDao
}

func NewUserService(userDao dao.UserDao) *userService {
	return &userService{userDao: userDao}
}

func (service *userService) QueryAll() (map[int64]models.User, error) {
	return service.userDao.QueryAll()
}
func (service *userService) QueryLogin(user *models.User) (models.User, error) {
	return service.userDao.QueryLogin(user)
}
func (service *userService) UpdateUser(user *models.User) int64 {
	return service.userDao.UpdateUser(user)
}
func (service *userService) QueryById(id int) (models.User, error) {
	return service.userDao.QueryById(id)
}
func (service *userService) DeleteUserById(id int) int64 {
	return service.userDao.DeleteUserById(id)
}
func (service *userService) DeleteUsers(users map[string]interface{}) int64 {
	return service.userDao.DeleteUsers(users)
}
func (service *userService) InsertUser(user *models.User) (int64, error) {
	return service.userDao.InsertUser(user)
}
func (service *userService) PageQueryUserData(maps map[string]interface{}) ([]models.User, error) {
	return service.userDao.PageQueryUserData(maps)
}
func (service *userService) PageQueryUserCount(maps map[string]interface{}) (int64, error) {
	return service.userDao.PageQueryUserCount(maps)
}

