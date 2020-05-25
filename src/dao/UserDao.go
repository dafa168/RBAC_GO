package dao

import (
	"RBAC_GO/configs"
	"RBAC_GO/src/models"
	"bytes"
	"fmt"
)

type UserDao interface {
	QueryAll() (map[int64]models.User, error)
	QueryLogin(user *models.User) (models.User,error)
	UpdateUser(user *models.User) int64
	QueryById(id int)(models.User,error)
	DeleteUserById(id int) int64
	DeleteUsers(users map[string]interface{}) int64
	InsertUser(user *models.User) (int64, error)
	PageQueryUserData(maps map[string]interface{})([]models.User,error)
	PageQueryUserCount(maps map[string]interface{})(int64,error)
}
// sql采用工具里， 这里不需要在结构体中重新定义
type userDao struct {}
// 查询所有user 信息
func (dao *userDao)QueryAll() (map[int64]models.User, error) {
	users := make(map[int64]models.User, 0)
	err := configs.Engine.Find(users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
// 用户名loginacct 和密码判断用户名可以登录
func  (dao *userDao)QueryLogin(user *models.User) (models.User,error) {
	//sqlStr := "select * from user where loginacct = ? and usepswd = ?"
	resultUser := models.User{}
	err := configs.Engine.Where("login_acct=", user.LoginAcct).Where("user_ps_wd=", user.Username).Find(resultUser)
	if err != nil{
		fmt.Println("出现错误：",err)
		return resultUser,err
	}
	return resultUser,nil
}
// 更新用户信息
func  (dao *userDao)UpdateUser(user *models.User) int64 {
	//sqlStr := "update t_user set loginacct = ?,?,?,? where id = ?"
	update, err := configs.Engine.ID(user.Id).Update(user)
	if err != nil {
		fmt.Printf("UpdateUser 出现异常：%e", err)
	}
	return update
}
//根据id 查找用户user
func  (dao *userDao)QueryById(id int)(models.User,error){
	resultUser := models.User{}
	err := configs.Engine.ID(id).Find(resultUser)
	if err != nil{
		return models.User{}, err
	}
	return resultUser,nil
}
// 通过user 的id 删除用户
func  (dao *userDao)DeleteUserById(id int) int64 {

	//sqlStr := "delete from t_user where id = ?"

	results, err := configs.Engine.ID(id).Delete(models.User{})
	if err != nil {
		fmt.Printf("DeleteUserById 出现异常：%e", err)
	}
	return results
}
// 批量删除用户
func (dao *userDao) DeleteUsers(users map[string]interface{}) int64 {

	sqlstr := "delete from user where id in ("
	//定义Buffer类型
	var bt bytes.Buffer
	//向bt中写入字符串
	bt.WriteString(sqlstr)
	usersIds := users["usersId"].([]int)
	fmt.Println(usersIds)
	for n,i:= range []string{"1","2"}{
		bt.WriteString(i)
		if n != len([]string{"1","2"})-1{
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
	affected, _ := exec.RowsAffected()

	return affected
}
// 添加新用户
func  (dao *userDao)InsertUser(user *models.User) (int64, error) {
	//sqlstr = "insert into user login_acct ,username,user_ps_wd,email,create_time values (?,?,?,?,?)"
	result, err := configs.Engine.InsertOne(user)
	if err !=nil{
		return -1, err
	}
	return result,nil
}

// 分页查找user的数据
func  (dao *userDao)PageQueryUserData(maps map[string]interface{})([]models.User,error){

	users := make([]models.User, 0)
	if _, ok := maps["queryText"]; ok {
		sqlStr1 := "select * from user where login_acct like concat('%', ?, '%')  order by create_time desc limit ?,?"

		start := maps["start"].(int64)
		size := maps["size"].(int64)
		rows, err := configs.Engine.SQL(sqlStr1,maps["queryText"].(string), start, size).Rows(new(models.Role))
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
func  (dao *userDao)PageQueryUserCount(maps map[string]interface{})(int64,error){
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