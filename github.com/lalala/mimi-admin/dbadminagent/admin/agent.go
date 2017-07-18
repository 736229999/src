package admin

import (
	"database/sql"
	"fmt"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"golang.org/x/net/context"
	"log"
	"strings"
	"time"
)

const (
	USER_STATUS_OK     = 0
	USER_STATUS_DELETE = 1

	TABLE_USER = "user"

	TIME_FORMRT = "2006-01-02 15:04"

	GIFT_PACKAGE_EXCHANGE = 0
	GIFT_PACKAGE_INVITE   = 1
)

func (agt *DbAdminAgent) QueryUserInfoByEmail(ctx context.Context, arg *dbproto.StringValue) (*dbproto.AdminUserInfoArg, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	st, err := agt.dbConn.Prepare(`SELECT id, email, username, password, salt, status, mobile, create_time, creator, register_ip FROM "user" WHERE email = $1 AND status = $2`)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	user := &dbproto.AdminUserInfoArg{}
	err = st.QueryRow(arg.GetValue(), USER_STATUS_OK).Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Salt, &user.Status, &user.Mobile, &user.CreateTime, &user.Creator, &user.RegisterIp)
	if err != nil {
		if err == sql.ErrNoRows {

			log.Printf("用户不存在:%+v\n", err, arg.GetValue())
			return nil, err
		}
		log.Printf("%+v\n", err)
		return nil, err
	}
	return user, nil
}

//修改用户信息.
func (agt *DbAdminAgent) SetUserInfoById(ctx context.Context, arg *dbproto.AdminUserInfoArg) (*dbproto.AdminUserInfoArg, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	st, err := agt.dbConn.Prepare(`UPDATE "user" SET mobile = $1, username = $2 where id = $3`)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	_, err = st.Exec(arg.GetMobile(), arg.GetUsername(), arg.GetId())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	userInfo, err := queryUserInfoById(agt, &dbproto.IntValue{Value: arg.GetId()})
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return userInfo, nil
}

//根据id查询用户.
func queryUserInfoById(agt *DbAdminAgent, arg *dbproto.IntValue) (*dbproto.AdminUserInfoArg, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	st, err := agt.dbConn.Prepare(`SELECT id, email, username, password, salt, status, mobile, create_time, creator, register_ip FROM "user" WHERE id = $1 AND status = $2`)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	user := &dbproto.AdminUserInfoArg{}
	err = st.QueryRow(arg.GetValue(), USER_STATUS_OK).Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Salt, &user.Status, &user.Mobile, &user.CreateTime, &user.Creator, &user.RegisterIp)
	if err != nil {
		if err == sql.ErrNoRows {

			log.Printf("用户不存在:%+v\n", err, arg.GetValue())
			return nil, err
		}
		log.Printf("%+v\n", err)
		return nil, err
	}
	return user, nil
}

//用户列表.
func (agt *DbAdminAgent) QueryUserList(ctx context.Context, arg *dbproto.Nil) (*dbproto.AdminUserList, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	st, err := agt.dbConn.Prepare(`SELECT u1.id, u1.email, u1.username, u1.password, u1.salt, u1.status, u1.mobile, u1.create_time, (SELECT u2.username FROM "user" AS u2 WHERE id = u1.creator) AS creator , register_ip FROM "user" AS u1`)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	rows, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	var userList []*dbproto.AdminUserInfoReply

	for rows.Next() {

		user := &dbproto.AdminUserInfoReply{}
		var createTime, status int64
		err = rows.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Salt, &status, &user.Mobile, &createTime, &user.Creator, &user.RegisterIp)
		if err != nil {
			log.Printf("%+v\n", err)
			return nil, err
		}
		t := time.Unix(createTime, 0).Format(TIME_FORMRT)
		user.CreateTime = t

		if status == USER_STATUS_OK {
			user.Status = "正常"
		} else {
			user.Status = "已删除"
		}

		log.Printf("%+v\n", user)
		userList = append(userList, user)
	}
	return &dbproto.AdminUserList{UserList: userList}, nil
}

//角色列表.
func (agt *DbAdminAgent) QueryRoleList(ctx context.Context, arg *dbproto.Nil) (*dbproto.RoleList, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	fieldList := []string{
		"r.id", "r.role_name", "r.remarks", "u.username as creator_name", "r.create_time",
	}
	esql := fmt.Sprintf(`SELECT %s FROM role AS r LEFT JOIN "user" AS u ON r.creator = u.id ORDER BY id DESC`, strings.Join(fieldList, ", "))
	st, err := agt.dbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	rows, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	roleList := &dbproto.RoleList{}
	for rows.Next() {
		role := &dbproto.Role{}
		var remarks sql.NullString
		value := []interface{}{
			&role.Id, &role.RoleName, &remarks, &role.CreatorName, &role.CreateTime,
		}
		if err := rows.Scan(value...); err != nil {
			if err == sql.ErrNoRows {
				return &dbproto.RoleList{}, nil
			}
			log.Printf("%+v\n", err)
			return nil, err
		}

		if remarks.Valid {
			role.Remarks = remarks.String
		}
		roleList.List = append(roleList.List, role)
	}

	return roleList, nil
}

//权限列表.
func (agt *DbAdminAgent) QueryPrivilegesList(ctx context.Context, arg *dbproto.Nil) (*dbproto.AdminPrivilegesList, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	st, err := agt.dbConn.Prepare(`SELECT p.id, p.name, p.key, p.path, (SELECT u.username FROM "user" AS u  WHERE u.id = p.creator) AS creator, p.create_time FROM privilege AS p ORDER BY id DESC`)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	rows, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	var privilegeList []*dbproto.AdminPrivilegesReply

	for rows.Next() {

		privilege := &dbproto.AdminPrivilegesReply{}
		var createTime int64
		if err := rows.Scan(&privilege.Id, &privilege.Name, &privilege.Key, &privilege.Path, &privilege.Creator, &createTime); err != nil {
			log.Printf("%+v\n", err)
			return nil, err
		}
		t := time.Unix(createTime, 0).Format(TIME_FORMRT)
		privilege.CreateTime = t

		privilegeList = append(privilegeList, privilege)
	}

	return &dbproto.AdminPrivilegesList{PrivilegesList: privilegeList}, nil
}

//添加权限.
//func (agt *DbAdminAgent) InsertPrivileges (ctx context.Context, arg *dbproto.AdminPrivileges) (*dbproto.Nil, error) {
//
//	defer func() {
//		if err := recover(); err != nil {
//			log.Println("recover from panic:", err)
//		}
//	}()
//
//	st, err := agt.dbConn.Prepare(`INSERT INTO privilege (name, key, path, creator, create_time) VALUES ($1, $2, $3, $4, $5)`)
//	if err != nil {
//		log.Printf("%+v\n", err, arg)
//		return nil, err
//	}
//	_, err = st.Exec(arg.GetName(), arg.GetKey(), arg.GetPath(), arg.GetCreator(), arg.GetCreateTime())
//	if err != nil {
//		log.Printf("%+v\n", err, arg)
//		return nil, err
//	}
//	return &dbproto.Nil{}, nil
//}

//编辑权限.
func (agt *DbAdminAgent) SetPrivileges(ctx context.Context, arg *dbproto.AdminPrivileges) (*dbproto.Nil, error) {

	st, err := agt.dbConn.Prepare(`update "privilege" set name = $1, key = $2, path = $3 where id = $4`)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	_, err = st.Exec(arg.GetName(), arg.GetKey(), arg.GetPath(), arg.GetId())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

//删除权限.
func (agt *DbAdminAgent) DeletePrivileges(ctx context.Context, arg *dbproto.IntValue) (*dbproto.Nil, error) {

	st, err := agt.dbConn.Prepare(`DELETE from privilege WHERE id = $1`)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	_, err = st.Exec(arg.GetValue())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	return &dbproto.Nil{}, err
}
