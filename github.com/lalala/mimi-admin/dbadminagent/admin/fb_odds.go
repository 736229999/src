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

//赔率相关
func (agt *DbAdminAgent) QueryOddsById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.PlayOdds, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	Info := &dbproto.PlayOdds{}
	esql := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", TABLE_PLAYODDS)
	err := agt.footballDbConn.QueryRow(esql, arg.Value).Scan(&Info.Id, &Info.Spf, &Info.Rqspf, &Info.Zjqs, &Info.Bf,
		&Info.Bqc, &Info.Spfdg, &Info.Rqspfdg, &Info.Zjqsdg, &Info.Bfdg, &Info.Bqcdg)
	log.Println("SQL", esql)
	if err != nil {
		log.Println(err, "SQL: ", esql)
		if err == sql.ErrNoRows {
			log.Println("No rows are requeried")
			return Info, nil
		}
		return nil, err
	}
	return Info, nil
}

func (agt *DbAdminAgent) CreatePlayOdds(ctx context.Context, arg *dbproto.PlayOdds) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		esql   string
		eid    int64
		column = []string{"id", "spf", "rqspf", "zjqs", "bf", "bqc", "spf_dg", "rqspf_dg", "zjqs_dg", "bf_dg", "bqc_dg"}
	)
	id := arg.GetId()
	spf := arg.GetSpf()
	rqspf := arg.GetRqspf()
	zjqs := arg.GetZjqs()
	bf := arg.GetBf()
	bqc := arg.GetBqc()
	spfDg := arg.GetSpfdg()
	rqspfDg := arg.GetRqspfdg()
	zjqsDg := arg.GetZjqsdg()
	bfDg := arg.GetBfdg()
	bqcDg := arg.GetBqcdg()

	esql = fmt.Sprintf("INSERT INTO %s(%s) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id;", TABLE_PLAYODDS, strings.Join(column, ","))
	err := agt.footballDbConn.QueryRow(esql, id, spf, rqspf, zjqs, bf, bqc, spfDg, rqspfDg, zjqsDg, bfDg, bqcDg).Scan(&eid)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	log.Printf("Create Playbet id=%v successfully", eid)

	return &dbproto.IntValue{Value: eid}, nil
}

func (agt *DbAdminAgent) UpdatePlayOdds(ctx context.Context, arg *dbproto.PlayOdds) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		esql string
		eid  int64
	)
	id := arg.GetId()
	spf := arg.GetSpf()
	rqspf := arg.GetRqspf()
	zjqs := arg.GetZjqs()
	bf := arg.GetBf()
	bqc := arg.GetBqc()
	spfDg := arg.GetSpfdg()
	rqspfDg := arg.GetRqspfdg()
	zjqsDg := arg.GetZjqsdg()
	bfDg := arg.GetBfdg()
	bqcDg := arg.GetBqcdg()

	valueArg := fmt.Sprintf("spf='%s',rqspf='%s',zjqs='%s',bf='%s',bqc='%s',spf_dg='%v',rqspf_dg='%v',zjqs_dg='%v',bf_dg='%v',bqc_dg='%v'",
		spf, rqspf, zjqs, bf, bqc, spfDg, rqspfDg, zjqsDg, bfDg, bqcDg)
	esql = fmt.Sprintf("UPDATE %s SET %s WHERE id=$1 RETURNING id;", TABLE_PLAYODDS, valueArg)
	err := agt.footballDbConn.QueryRow(esql, id).Scan(&eid)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	log.Printf("Playodds id=%v update successfuly", eid)

	return &dbproto.IntValue{Value: eid}, nil
}

//开奖相关
func (agt *DbAdminAgent) QueryFbResult(ctx context.Context, arg *dbproto.QueryOpencaiArg) (*dbproto.FbOpencaiList, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		homeId, guestId                                      int
		total, start                                         int64
		esql, order, pager, where                            string
		resultId, homeBall, guestBall, homeHBall, guestHBall sql.NullInt64
		rqspf, bqc                                           sql.NullString
		ifOpen                                               sql.NullBool
	)
	if arg.GetPage() > 1 {
		start = (arg.GetPage() - 1) * arg.PageSize
	} else {
		start = 0
	}

	esql = fmt.Sprintf("SELECT n.id,n.game_no,n.open_time,n.home_team,n.guest_team,pn.* FROM %s AS n LEFT JOIN %s AS pn ON n.id=pn.id", TABLE_GAMES, TABLE_OPENCAI)
	order = " ORDER BY n.id DESC"
	pager = fmt.Sprintf(" LIMIT %d OFFSET %d", arg.PageSize, start)

	if arg.GetDate() != 0 {
		queryStart := arg.GetDate()
		queryEnd := time.Unix(queryStart, 0).AddDate(0, 0, 1).Unix()

		log.Println(arg.GetDate(), time.Unix(queryStart, 0), time.Unix(queryEnd, 0))
		if strings.Contains(where, "WHERE") {
			where = fmt.Sprintf("%s AND start_time>'%d' AND start_time<'%d'", where, queryStart, queryEnd)
		} else {
			where = fmt.Sprintf(" WHERE start_time>'%d' AND start_time<'%d'", queryStart, queryEnd)
		}
	}

	rows, err := agt.footballDbConn.Query(esql + where + order + pager)
	log.Println("SQL", esql+where+order+pager)
	if err != nil {
		log.Println(err, "SQL: ", esql+order)
		return nil, err
	}

	List := make([]*dbproto.FbOpencai, 0)
	for rows.Next() {
		Info := &dbproto.FbOpencai{Game: &dbproto.GameInfo{}, Result: &dbproto.FbGameresult{}}
		Infoless := &dbproto.FbOpencai{Game: &dbproto.GameInfo{}} //如果只有前一半数据
		err := rows.Scan(&Info.Game.Id, &Info.Game.GameNo, &Info.Game.OpenTime, &homeId, &guestId,
			&resultId, &homeBall, &guestBall, &homeHBall, &guestHBall, &rqspf, &bqc, &ifOpen)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		//查球队
		err = agt.QueryTeamById(homeId, guestId, Info.Game)
		if err != nil {
			log.Println(err, "球队查询失败")
			return nil, err
		}
		if resultId.Valid {
			Info.Result.Homeball = homeBall.Int64
			Info.Result.Guestball = guestBall.Int64
			Info.Result.HomeHball = homeHBall.Int64
			Info.Result.GuestHball = guestHBall.Int64
			Info.Result.Rqspf = rqspf.String
			Info.Result.Bqc = bqc.String
			Info.Result.Ifopen = ifOpen.Bool
			List = append(List, Info)
		} else {
			Infoless.Game.Id = Info.Game.Id
			Infoless.Game.GameNo = Info.Game.GameNo
			Infoless.Game.OpenTime = Info.Game.OpenTime
			Infoless.Game.HomeTeam = Info.Game.HomeTeam
			Infoless.Game.GuestTeam = Info.Game.GuestTeam
			List = append(List, Infoless)
		}
	}

	totalEsql := fmt.Sprintf("SELECT count(*) FROM %s", TABLE_GAMES)
	row := agt.footballDbConn.QueryRow(totalEsql + where)
	row.Scan(&total)

	log.Printf("total is %v\n", total)

	return &dbproto.FbOpencaiList{List: List, Total: total}, nil
}

func (agt *DbAdminAgent) QueryFbResultById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.FbGameresult, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	Info := &dbproto.FbGameresult{}
	esql := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", TABLE_OPENCAI)
	err := agt.footballDbConn.QueryRow(esql, arg.Value).Scan(&Info.Id, &Info.Homeball, &Info.Guestball, &Info.HomeHball, &Info.GuestHball, &Info.Rqspf, &Info.Bqc, &Info.Ifopen)
	log.Println("SQL", esql)
	if err != nil {
		log.Println(err, "SQL: ", esql)
		if err == sql.ErrNoRows {
			log.Println("No rows are requeried")
			return Info, nil
		}
		return nil, err
	}

	return Info, nil
}

func (agt *DbAdminAgent) CreateFbResult(ctx context.Context, arg *dbproto.FbGameresult) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		esql string
		eid  int64
	)
	id := arg.GetId()
	homeBall := arg.GetHomeball()
	guestBall := arg.GetGuestball()
	homeHBall := arg.GetHomeHball()
	guestHBall := arg.GetGuestHball()
	rqspf := arg.GetRqspf()
	bqc := arg.GetBqc()
	ifOpen := arg.GetIfopen()

	flag := agt.IsIdExist(id, TABLE_OPENCAI)
	var column = []string{"id", "homeball", "guestball", "home_hball", "guest_hball", "rqspf", "bqc", "ifopen"}
	if flag == 0 {
		esql = fmt.Sprintf("INSERT INTO %s(%s) VALUES($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id;", TABLE_OPENCAI, strings.Join(column, ","))
		err := agt.footballDbConn.QueryRow(esql, id, homeBall, guestBall, homeHBall, guestHBall, rqspf, bqc, ifOpen).Scan(&eid)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		log.Printf("Create Opencai id=%v successfully", eid)
	} else {
		valueArg := fmt.Sprintf("homeball=%d,guestball=%d,home_hball=%d,guest_hball=%d,rqspf='%s',bqc='%s',ifopen=%v",
			homeBall, guestBall, homeHBall, guestHBall, rqspf, bqc, ifOpen)
		esql = fmt.Sprintf("UPDATE %s SET %s WHERE id=$1 RETURNING id;", TABLE_OPENCAI, valueArg)
		err := agt.footballDbConn.QueryRow(esql, id).Scan(&eid)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		log.Printf("Opencai id=%v update successfuly", eid)
	}
	return &dbproto.IntValue{Value: eid}, nil
}
