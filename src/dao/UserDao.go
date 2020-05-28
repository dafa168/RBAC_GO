package dao

import (
	"RBAC_GO/configs"
	"RBAC_GO/src/models"
	"bytes"
	"database/sql"
	"fmt"
)

// 查询所有user 信息
func QueryAll() (map[int64]models.User, error) {
	users := make(map[int64]models.User, 0)
	err := configs.Engine.Find(users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// 用户名loginacct 和密码判断用户名可以登录
func QueryLogin(user *models.User) (models.User, error) {
	//sqlStr := "select * from user where loginacct = ? and usepswd = ?"
	resultUser := models.User{}
	_, err := configs.Engine.Table("user").Where("login_acct = ? and user_ps_wd = ?",user.LoginAcct,user.UserPsWd).Get(&resultUser)

	if err != nil {
		fmt.Println("出现错误：", err)
		return resultUser, err
	}

	return resultUser, nil
}

// 更新用户信息
func UpdateUser(user *models.User) int64 {
	//sqlStr := "update t_user set loginacct = ?,?,?,? where id = ?"
	update, err := configs.Engine.ID(user.Id).Update(user)
	if err != nil {
		fmt.Printf("UpdateUser 出现异常：%e", err)
	}
	return update
}

//根据id 查找用户user
func QueryById(id int) (models.User, error) {
	resultUser := models.User{}
	err := configs.Engine.ID(id).Find(resultUser)
	if err != nil {
		return models.User{}, err
	}
	return resultUser, nil
}

// 通过user 的id 删除用户
func DeleteUserById(id int) int64 {

	//sqlStr := "delete from t_user where id = ?"

	results, err := configs.Engine.ID(id).Delete(models.User{})
	if err != nil {
		fmt.Printf("DeleteUserById 出现异常：%e", err)
	}
	return results
}

// 批量删除用户
func DeleteUsers(users map[string][]string) sql.Result {

	sqlstr := "delete from user where id in ("
	//定义Buffer类型
	var bt bytes.Buffer
	//向bt中写入字符串
	bt.WriteString(sqlstr)
	for n, i := range users["usersid"] {
		bt.WriteString(i)
		if n != len(users["usersid"])-1 {
			bt.WriteString(",")
		}
	}
	bt.WriteString(")")
	//获得拼接后的字符串
	sql1 := bt.String()
	fmt.Println(sql1)
	exec, err := configs.Engine.Exec(sqlstr)
	if err != nil {
		fmt.Println(err)
	}
	return exec
}

// 添加新用户
func InsertUser(user *models.User) (int64, error) {
	//sqlstr = "insert into user login_acct ,username,user_ps_wd,email,create_time values (?,?,?,?,?)"
	result, err := configs.Engine.InsertOne(user)
	if err != nil {
		return -1, err
	}
	return result, nil
}

// 分页查找user的数据
func PageQueryUserData(maps map[string]interface{}) ([]models.User, error) {

	users := make([]models.User, 0)
	if _, ok := maps["queryText"]; ok {
		sqlStr1 := "select * from user where login_acct like concat('%', ?, '%')  order by create_time desc limit ?,?"

		start := maps["start"].(int64)
		size := maps["size"].(int64)
		rows, err := configs.Engine.SQL(sqlStr1, maps["queryText"].(string), start, size).Rows(new(models.Role))
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			user := models.User{}
			err = rows.Scan(&user)
			users = append(users, user)
			//...
		}
		return users, nil
	} else {
		sqlStr2 := "select * from role  limit ?,?"
		start := maps["start"].(int64)
		size := maps["size"].(int64)
		rows, err := configs.Engine.SQL(sqlStr2, start, size).Rows(new(models.Role))
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			user := models.User{}
			err = rows.Scan(&user)
			users = append(users, user)
			//...
		}
		return users, nil
	}
}

// 分页查找user 的count
func PageQueryUserCount(maps map[string]interface{}) (int64, error) {
	if _, ok := maps["queryText"]; ok {
		sqlStr1 := "select count(*) from user  where name like concat('%', ?, '%')"

		count, err := configs.Engine.SQL(sqlStr1, maps["queryText"]).Count(new(map[string]models.User))
		if err != nil {
			return -1, err
		}
		return count, nil
	} else {
		sqlStr2 := "select count(*) from user"
		count, err := configs.Engine.SQL(sqlStr2).Count(new(map[string]models.User))
		if err != nil {
			return -1, err
		}

		return count, nil
	}
}
