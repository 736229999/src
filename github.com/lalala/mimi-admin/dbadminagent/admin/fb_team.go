package admin

import (
	"fmt"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"golang.org/x/net/context"
	"log"
)

//球队相关
func (agt *DbAdminAgent) CreateFbTeam(ctx context.Context, arg *dbproto.FbTeamInfo) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var esql string
	id := arg.GetId()
	name := arg.GetName()
	url := arg.GetUrl()

	if id == 0 {
		esql = fmt.Sprintf("INSERT INTO %s(name,url) VALUES($1,$2) RETURNING id;", TABLE_TEAMS)
		err := agt.footballDbConn.QueryRow(esql, name, url).Scan(&id)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		log.Printf("Create Team id=%v successfully", id)
	} else {
		esql = fmt.Sprintf("UPDATE %s SET name=$2,url=$3 WHERE id=$1 RETURNING id;", TABLE_TEAMS)
		err := agt.footballDbConn.QueryRow(esql, id, name, url).Scan(&id)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		log.Printf("Team id=%v update successfuly", id)
	}
	return &dbproto.IntValue{Value: id}, nil
}

func (agt *DbAdminAgent) QueryFbTeamList(ctx context.Context, arg *dbproto.QueryFbTeamArg) (*dbproto.FbTeamList, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		esql, where, pager, order string
		total, start              int64
	)

	if arg.GetPage() > 1 {
		start = (arg.GetPage() - 1) * arg.PageSize
	} else {
		start = 0
	}

	if arg.GetPageSize() > 0 {
		pager = fmt.Sprintf(" LIMIT %d OFFSET %d", arg.PageSize, start)
	}

	if len(arg.GetName()) != 0 {
		where = fmt.Sprintf(" WHERE name LIKE '%%%s%%'", arg.Name)
	}
	order = " ORDER BY id DESC"
	esql = fmt.Sprintf("SELECT id, name,url FROM %s", TABLE_TEAMS)
	rows, err := agt.footballDbConn.Query(esql + where + order + pager)
	if err != nil {
		log.Println(err, esql+where+order)
		return nil, err
	}
	list := make([]*dbproto.FbTeamInfo, 0)
	for rows.Next() {
		team := &dbproto.FbTeamInfo{}
		err := rows.Scan(&team.Id, &team.Name, &team.Url)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		list = append(list, team)
	}
	totalSql := fmt.Sprintf("SELECT count(*) FROM %s", TABLE_TEAMS)
	row := agt.footballDbConn.QueryRow(totalSql + where)
	row.Scan(&total)
	log.Println("esql: ", esql+where+order)
	log.Printf("total is %v", total)
	return &dbproto.FbTeamList{List: list, Total: total}, nil
}

func (agt *DbAdminAgent) QueryFbTeamById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.FbTeamInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		esql string
	)

	esql = fmt.Sprintf("SELECT name,url FROM %s WHERE id=$1", TABLE_TEAMS)
	rows, err := agt.footballDbConn.Query(esql, arg.Value)
	log.Println("SQL", esql)
	if err != nil {
		log.Println(err, "SQL: ", esql)
		return nil, err
	}
	Info := &dbproto.FbTeamInfo{}
	rows.Next()
	err = rows.Scan(&Info.Name, &Info.Url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return Info, nil
}

//联赛相关
func (agt *DbAdminAgent) CreateFbLeague(ctx context.Context, arg *dbproto.FbLeagueInfo) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var esql string
	id := arg.GetId()
	name := arg.GetName()

	if id == 0 {
		esql = fmt.Sprintf("INSERT INTO %s(name) VALUES($1) RETURNING id;", TABLE_LEAGUE)
		err := agt.footballDbConn.QueryRow(esql, name).Scan(&id)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		log.Printf("Create League id=%v successfully", id)
	} else {
		esql = fmt.Sprintf("UPDATE %s SET name=$2 WHERE id=$1 RETURNING id;", TABLE_LEAGUE)
		err := agt.footballDbConn.QueryRow(esql, id, name).Scan(&id)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		log.Printf("League id=%v update successfuly", id)
	}
	return &dbproto.IntValue{Value: id}, nil
}

func (agt *DbAdminAgent) QueryFbLeague(ctx context.Context, arg *dbproto.QueryFbLeagueArg) (*dbproto.FbLeagueList, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		esql, where, pager, order string
		total, start              int64
	)

	if arg.GetPage() > 1 {
		start = (arg.GetPage() - 1) * arg.PageSize
	} else {
		start = 0
	}

	if arg.GetPageSize() > 0 {
		pager = fmt.Sprintf(" LIMIT %d OFFSET %d", arg.PageSize, start)
	}

	if len(arg.GetName()) != 0 {
		where = fmt.Sprintf(" WHERE name LIKE '%%%s%%'", arg.Name)
	}
	order = " ORDER BY id DESC"
	esql = fmt.Sprintf("SELECT * FROM %s", TABLE_LEAGUE)
	rows, err := agt.footballDbConn.Query(esql + where + order + pager)
	if err != nil {
		log.Println(err, esql+where+order)
		return nil, err
	}
	list := make([]*dbproto.FbLeagueInfo, 0)
	for rows.Next() {
		team := &dbproto.FbLeagueInfo{}
		err := rows.Scan(&team.Id, &team.Name)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		list = append(list, team)
	}

	totalSql := fmt.Sprintf("SELECT count(*) FROM %s", TABLE_LEAGUE)
	row := agt.footballDbConn.QueryRow(totalSql + where)
	row.Scan(&total)
	log.Println("esql: ", esql+where+order)
	log.Printf("total is %v", total)

	return &dbproto.FbLeagueList{List: list, Total: total}, nil
}
