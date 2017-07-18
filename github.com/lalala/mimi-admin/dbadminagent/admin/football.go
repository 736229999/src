package admin

import (
	"fmt"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"golang.org/x/net/context"
	"log"
	"strings"
)

const TABLE_GAMES string = "games"       //赛事表
const TABLE_TEAMS string = "teams"       //球队表
const TABLE_LEAGUE string = "league"     //联赛表
const TABLE_OPENCAI string = "opencai"   //开奖表
const TABLE_PLAYODDS string = "playodds" //玩法投注表

func (agt *DbAdminAgent) CreateFbGame(ctx context.Context, arg *dbproto.GameInfo) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		esql                      string
		id                        int64
		leagueId, homeId, guestId int
		err                       error
	)
	gameId := arg.GetId()
	gameNo := arg.GetGameNo()
	gameType := arg.GetGameType()
	openTime := arg.GetOpenTime()
	homeTeam := arg.GetHomeTeam()
	guestTeam := arg.GetGuestTeam()
	giveBall := arg.GetGiveball()
	startTime := arg.GetStartTime()
	endTime := arg.GetEndTime()

	agt.QueryLeagueByName(gameType, &leagueId)
	if err != nil {
		log.Println(err, "赛事查询失败")
		return nil, err
	}
	//查球队
	agt.QueryTeamByName(homeTeam, guestTeam, &homeId, &guestId)
	if err != nil {
		log.Println(err, "球队查询失败")
		return nil, err
	}

	var column = []string{
		"game_no", "league_id", "open_time", "home_team", "guest_team", "ball_give", "start_time", "end_time",
	}

	if gameId == 0 {
		esql = fmt.Sprintf("INSERT INTO %s(%s) VALUES($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id;", TABLE_GAMES, strings.Join(column, ","))
		err := agt.footballDbConn.QueryRow(esql, gameNo, leagueId, openTime, homeId, guestId, giveBall, startTime, endTime).Scan(&id)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		log.Printf("Game 添加成功，GameId is %v", id)
	} else {
		valueArg := fmt.Sprintf("game_no='%s',league_id=%d,open_time='%d', home_team=%d,guest_team=%d,ball_give=%d,start_time=%d,end_time=%d",
			gameNo, leagueId, openTime, homeId, guestId, giveBall, startTime, endTime)
		esql = fmt.Sprintf("UPDATE %s SET %s WHERE id=$1 RETURNING id;", TABLE_GAMES, valueArg)
		err := agt.footballDbConn.QueryRow(esql, gameId).Scan(&id)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		log.Printf("Games %v update succeed:%s", id, esql)
	}

	return &dbproto.IntValue{Value: id}, nil
}

func (agt *DbAdminAgent) QueryFbGame(ctx context.Context, arg *dbproto.QueryFbGameArg) (*dbproto.FbGameList, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		homeId, guestId, leagueId int
		total, start              int64
		esql, order, pager, where string
	)
	if arg.GetPage() > 1 {
		start = (arg.GetPage() - 1) * arg.PageSize
	} else {
		start = 0
	}

	esql = fmt.Sprintf("SELECT * FROM %s", TABLE_GAMES)
	order = " ORDER BY id DESC"
	pager = fmt.Sprintf(" LIMIT %d OFFSET %d", arg.PageSize, start)

	if arg.GetStart() != 0 && arg.GetEnd() != 0 {
		Querystart := arg.GetStart()
		Queryend := arg.GetEnd()
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND open_time>'%d' AND open_time<'%d'", where, Querystart, Queryend)
		} else {
			where = fmt.Sprintf(" WHERE open_time>'%d' AND open_time<'%d'", Querystart, Queryend)
		}
	}

	if arg.GetTeam() != 0 {
		Queryteam := arg.GetTeam()
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND (home_team=%d OR guest_team=%d)", where, Queryteam, Queryteam)
		} else {
			where = fmt.Sprintf(" WHERE (home_team=%d OR guest_team=%d)", Queryteam, Queryteam)
		}
	}

	rows, err := agt.footballDbConn.Query(esql + where + order + pager)
	log.Println("SQL", esql+where+order+pager)
	if err != nil {
		log.Println(err, "SQL: ", esql+order)
		return nil, err
	}

	gameList := make([]*dbproto.GameInfo, 0)

	for rows.Next() {
		Info := &dbproto.GameInfo{}
		err := rows.Scan(&Info.Id, &Info.GameNo, &leagueId, &Info.OpenTime, &homeId, &guestId,
			&Info.Giveball, &Info.StartTime, &Info.EndTime,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		//查赛事
		err = agt.QueryLeagueById(leagueId, Info)
		if err != nil {
			log.Println(err, "赛事查询失败")
			return nil, err
		}
		//查球队
		err = agt.QueryTeamById(homeId, guestId, Info)
		if err != nil {
			log.Println(err, "球队查询失败")
			return nil, err
		}

		gameList = append(gameList, Info)
	}

	totalEsql := fmt.Sprintf("SELECT count(*) FROM %s", TABLE_GAMES)
	row := agt.footballDbConn.QueryRow(totalEsql + where)
	row.Scan(&total)

	log.Printf("total is %v\n", total)

	return &dbproto.FbGameList{List: gameList, Total: total}, nil
}

func (agt *DbAdminAgent) QueryFbGameById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.GameInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		homeId, guestId, leagueId int
		esql                      string
	)

	esql = fmt.Sprintf("SELECT * FROM %s WHERE id=$1", TABLE_GAMES)
	rows, err := agt.footballDbConn.Query(esql, arg.Value)
	log.Println("SQL", esql)
	if err != nil {
		log.Println(err, "SQL: ", esql)
		return nil, err
	}
	Info := &dbproto.GameInfo{}
	rows.Next()
	err = rows.Scan(&Info.Id, &Info.GameNo, &leagueId, &Info.OpenTime, &homeId, &guestId,
		&Info.Giveball, &Info.StartTime, &Info.EndTime,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//查赛事
	err = agt.QueryLeagueById(leagueId, Info)
	if err != nil {
		log.Println(err, "赛事查询失败")
		return nil, err
	}
	//查球队
	err = agt.QueryTeamById(homeId, guestId, Info)
	if err != nil {
		log.Println(err, "球队查询失败")
		return nil, err
	}

	return Info, nil
}

//
//func (agt *DbAdminAgent) QuerySelectOfLeague(ctx context.Context, arg *dbproto.QueryFbTeamArg) (*dbproto.FbLeagueList, error) {
//	defer func() {
//		if err := recover(); err != nil {
//			log.Println("recover from panic:", err)
//		}
//	}()
//	var (
//		esql, where, order string
//		total              int64
//	)
//	where = fmt.Sprintf(" WHERE name LIKE '%%%s%%'", arg.GetName())
//	order = " ORDER BY id DESC"
//	esql = fmt.Sprintf("SELECT id,name FROM %s", TABLE_LEAGUE)
//	rows, err := agt.footballDbConn.Query(esql + where + order)
//	if err != nil {
//		log.Println(err, esql+where+order)
//		return nil, err
//	}
//	list := make([]*dbproto.FbLeagueInfo, 0)
//	for rows.Next() {
//		league := &dbproto.FbLeagueInfo{}
//		err := rows.Scan(&league.Id, &league.Name)
//		if err != nil {
//			log.Println(err, esql)
//			return nil, err
//		}
//		list = append(list, league)
//	}
//	totalSql := fmt.Sprintf("SELECT count(*) FROM %s", TABLE_LEAGUE)
//	row := agt.optionsDbConn.QueryRow(totalSql + where)
//	row.Scan(&total)
//	log.Println("esql: ", esql+where+order)
//	log.Printf("total is %v", total)
//	return &dbproto.FbLeagueList{List: list, Total: total}, nil
//}
