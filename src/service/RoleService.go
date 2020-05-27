package service

import (
	"RBAC_GO/src/dao"
	"RBAC_GO/src/models"
)

type RoleService interface {

	dao.RoleDao
}
type roleService struct{

	dao.RoleDao
}

func NewRoleService(roleDao dao.RoleDao) *roleService {
	return &roleService{RoleDao: roleDao}
}
func (r *roleService) PageQueryRoleData(maps map[string]string) ([]models.Role, error){
	return r.RoleDao.PageQueryRoleData(maps)
}
func (r *roleService) PageQueryRoleCount(maps map[string]string) (int64, error){
	return r.RoleDao.PageQueryRoleCount(maps)
}
func (r *roleService) QueryAllRole() ([]models.Role, error){
	return r.RoleDao.QueryAllRole()
}


