package dao

import (
	"RBAC_GO/configs"
	"RBAC_GO/src/models"
	"strconv"
)

// 分页查找Role数据
func PageQueryRoleData(maps map[string]string) ([]models.Role, error) {
	roles := make([]models.Role, 0)
	if _, ok := maps["queryText"]; ok {
		sqlStr1 := "select * from role  where name like concat('%', ?, '%')  limit ?, ?"

		start, err := strconv.Atoi(maps["start"])
		size, err := strconv.Atoi(maps["size"])
		rows, err := configs.Engine.SQL(sqlStr1, start, size).Rows(new(models.Role))
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			role := models.Role{}
			err = rows.Scan(&role)
			roles = append(roles, role)
			//...
		}
		return roles, nil
	} else {
		sqlStr2 := "select * from role  limit ?,?"
		start, err := strconv.Atoi(maps["start"])
		size, err := strconv.Atoi(maps["size"])
		rows, err := configs.Engine.SQL(sqlStr2, start, size).Rows(new(models.Role))
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			role := models.Role{}
			err = rows.Scan(&role)
			roles = append(roles, role)
			//...
		}
		return roles, nil
	}
}

// 查找role 的数量
func PageQueryRoleCount(maps map[string]string) (int64, error) {
	if _, ok := maps["queryText"]; ok {
		sqlStr1 := "select count(*) from role  where name like concat('%', ?, '%')"

		count, err := configs.Engine.SQL(sqlStr1, maps["queryText"]).Count(new(map[string]models.Role))
		if err != nil {
			return -1, err
		}
		return count, nil
	} else {
		sqlStr2 := "select count(*) from role"
		count, err := configs.Engine.SQL(sqlStr2).Count(new(map[string]models.Role))
		if err != nil {
			return -1, err
		}

		return count, nil
	}
}

// 查找所有数据
func QueryAllRole() ([]models.Role, error) {
	sqlStr := "select * from role"

	roles := make([]models.Role, 0)
	rows, err := configs.Engine.SQL(sqlStr).Rows(new(models.Role))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		role := models.Role{}
		err = rows.Scan(&role)
		roles = append(roles, role)
		//...
	}
	return roles, nil
}

