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

//添加权限.
func (agt *DbAdminAgent) InsertPrivilege(ctx context.Context, arg *dbproto.Privilege) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := fmt.Sprintf("INSERT INTO privilege (name, key, path, creator, create_time, p_id) VALUES ($1, $2, $3, $4, $5, $6)")
	st, err := agt.dbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}
	defer st.Close()

	_, err = st.Exec(arg.GetName(), arg.GetKey(), arg.GetPath(), arg.GetCreator(), time.Now().Unix(), arg.GetPId())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return &dbproto.Nil{}, nil
}

//获取权限列表.
func (agt *DbAdminAgent) QueryPrivilegeList(ctx context.Context, arg *dbproto.Nil) (*dbproto.PrivilegeList, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string{
		"id", "name", "key", "path",
		"creator", "create_time", "p_id",
	}
	esql := fmt.Sprintf("SELECT %s FROM privilege where p_id = 0", strings.Join(filedList, ", "))
	st, err := agt.dbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	defer st.Close()

	privilegeList := &dbproto.PrivilegeList{}

	rows, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	for rows.Next() {
		privilege := &dbproto.PrivilegeChildrenList{}
		value := []interface{}{
			&privilege.Id, &privilege.Name, &privilege.Key, &privilege.Path,
			&privilege.Creator, &privilege.CreateTime, &privilege.PId,
		}

		if err = rows.Scan(value...); err != nil {
			if err == sql.ErrNoRows {
				return &dbproto.PrivilegeList{}, nil
			}
			log.Printf("%+v\n", err)
			return nil, err
		}
		privilege.Label = privilege.Name
		//
		privilegeList.Children = append(privilegeList.Children, privilege)
		//
		//list := &dbproto.PrivilegeList{}
		////根据pid，查询子模块.
		//esql := fmt.Sprintf("SELECT %s FROM privilege where p_id = %d", strings.Join(filedList, ", "), privilege.GetId())
		//st, err = agt.dbConn.Prepare(esql)
		//if err != nil {
		//	log.Printf("%+v\n", err)
		//	return nil, err
		//}
		//
		//res, err := st.Query()
		//if err != nil {
		//	log.Printf("%+v\n", err)
		//	return nil, err
		//}
		//
		//for res.Next() {
		//	p := &dbproto.Privilege{}
		//	if err = res.Scan(value...); err != nil {
		//		if err == sql.ErrNoRows {
		//			privilege.Children = nil
		//		} else {
		//			log.Printf("%+v\n", err)
		//			return nil, err
		//		}
		//	}
		//	list.Children = append(list.Children, p)
		//}
	}

	for _, v := range privilegeList.Children {

		//list := &dbproto.PrivilegeList{}
		//根据pid，查询子模块.
		esql := fmt.Sprintf("SELECT %s FROM privilege where p_id = %d", strings.Join(filedList, ", "), v.GetId())
		stmt, err := agt.dbConn.Prepare(esql)
		if err != nil {
			log.Printf("%+v\n", err)
			return nil, err
		}

		defer stmt.Close()
		res, err := stmt.Query()
		if err != nil {
			log.Printf("%+v\n", err)
			return nil, err
		}

		for res.Next() {
			p := &dbproto.PrivilegeChildrenList{}
			val := []interface{}{
				&p.Id, &p.Name, &p.Key, &p.Path,
				&p.Creator, &p.CreateTime, &p.PId,
			}
			if err = res.Scan(val...); err != nil {
				if err == sql.ErrNoRows {
					v.Children = nil
					continue
				} else {
					log.Printf("%+v\n", err)
					return nil, err
				}
			}
			log.Printf("id:%+v\n", p.GetId())
			esq := fmt.Sprintf("SELECT %s FROM privilege where p_id = %d", strings.Join(filedList, ", "), p.GetId())
			log.Println("esq:", esq)
			stm, err := agt.dbConn.Prepare(esq)
			if err != nil {
				log.Printf("%+v\n", err)
				return nil, err
			}
			//
			defer stm.Close()
			//
			r, err := stm.Query()
			if err != nil {
				log.Printf("%+v\n", err)
				return nil, err
			}
			parents := p.Children
			for r.Next() {
				parent := &dbproto.PrivilegeChildrenList{}
				valu := []interface{}{
					&parent.Id, &parent.Name, &parent.Key, &parent.Path,
					&parent.Creator, &parent.CreateTime, &parent.PId,
				}
				if err = r.Scan(valu...); err != nil {
					if err == sql.ErrNoRows {
						parent.Children = nil
						continue
					} else {
						log.Printf("%+v\n", err)
						return nil, err
					}
				}
				log.Println("id:", parent.GetId())
				log.Println("name:", parent.GetName())
				parent.Label = parent.Name
				parents = append(parents, parent)
			}

			p.Children = parents
			p.Label = p.Name
			v.Children = append(v.Children, p)
		}

		//children := privilegeList.Children

		//for _, val := range children.Children {
		//	esql := fmt.Sprintf("SELECT %s FROM privilege where p_id = %d", strings.Join(filedList, ", "), val.GetId())
		//	stmt, err := agt.dbConn.Prepare(esql)
		//	if err != nil {
		//		log.Printf("%+v\n", err)
		//		return nil, err
		//	}
		//	//
		//	defer stmt.Close()
		//
		//	res, err := stmt.Query()
		//	if err != nil {
		//		log.Printf("%+v\n", err)
		//		return nil, err
		//	}
		//	//
		//	for res.Next() {
		//		p := &dbproto.PrivilegeChildrenList{}
		//		value := []interface{}{
		//			&p.Id, &p.Name, &p.Key, &p.Path,
		//			&p.Creator, &p.CreateTime, &p.PId,
		//		}
		//		if err = res.Scan(value...); err != nil {
		//			if err == sql.ErrNoRows {
		//				val.Children = nil
		//				continue
		//			} else {
		//				log.Printf("%+v\n", err)
		//				return nil, err
		//			}
		//		}
		//
		//		p.Label = p.Name
		//		val.Children = append(val.Children, p)
		//	}
		//}
	}

	//for _, val := range children.Children {
	//	val.Children
	//}



	//log.Printf("%+v\n", privilegeList)

	return privilegeList, nil
}
