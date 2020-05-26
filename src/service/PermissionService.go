package service

import (
	"RBAC_GO/src/dao"
	"RBAC_GO/src/models"
)

type PermissionService interface {
	dao.PermissionDao
	dao.RolePermissionDao
}
type permissionService struct{
	dao.PermissionDao
	dao.RolePermissionDao
}

func NewPermissionService(permissionDao dao.PermissionDao,
	rolePermissionDao dao.RolePermissionDao) *permissionService {
	return &permissionService{PermissionDao: permissionDao, RolePermissionDao: rolePermissionDao}
}


func (ps *permissionService) QueryRootPermission() (models.Permission, error) {
	return ps.PermissionDao.QueryRootPermission()
}
func (ps *permissionService) QueryChildPermissions(pid int) ([]models.Permission, error)        {
	return ps.PermissionDao.QueryChildPermissions(pid)
}
func (ps *permissionService) QueryAllPermission() ([]models.Permission, error)                  {
	return ps.PermissionDao.QueryAllPermission()
}
func (ps *permissionService) InsertPermission(permission models.Permission) (int64, error)      {
	return ps.PermissionDao.InsertPermission(permission)
}
func (ps *permissionService) QueryPermissionById(id int) (models.Permission, error)             {
	return ps.PermissionDao.QueryPermissionById(id)
}
func (ps *permissionService) UpdatePermission(permission models.Permission) (int64, error)      {
	return ps.PermissionDao.UpdatePermission(permission)
}
func (ps *permissionService) DeletePermission(permission models.Permission) (int64, error) {
	return ps.PermissionDao.DeletePermission(permission)
}
func (ps *permissionService) QueryPermissionIdsByRoleId(roleId int) ([]int, error)              {
	return ps.RolePermissionDao.QueryPermissionIdsByRoleId(roleId)
}
func (ps *permissionService) QueryPermissionsByUser(userRole models.UserRole) ([]models.Permission, error) {
	return ps.RolePermissionDao.QueryPermissionsByUser(userRole)
}
