package buycai

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
)

var buycaiTableFileds = "id, issue, start_time, end_time, open_time, open_balls"

func (agt *DbBuycaiAgent) BuycaiQuerySaleList(ctx context.Context, arg *dbproto.StringValue) (*dbproto.BuycaiSaleList, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	result := &dbproto.BuycaiSaleList{}
	tbl := arg.GetValue()
	now := time.Now()
	var startTime, endTime, openTime time.Time

	// 获取最近10期停售的
	esql := fmt.Sprintf("SELECT %s FROM %s WHERE end_time<='%s' ORDER BY id DESC LIMIT 10", buycaiTableFileds, tbl, now.Format("2006-01-02 15:04:05"))
	v := &dbproto.BuycaiSaleIssue{}
	var openBalls sql.NullString
	if err := agt.dbConn.QueryRow(esql).Scan(&v.Id, &v.Issue, &startTime, &endTime, &openTime, &openBalls); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err, esql)
			return nil, err
		}
	}
	if v.Id > 0 {
		v.StartTime = startTime.Unix()
		v.EndTime = endTime.Unix()
		v.OpenTime = openTime.Unix()
		if openBalls.Valid {
			v.OpenBalls = openBalls.String
		}
		result.List = append(result.List, v)
	}

	// 获取还在销售的
	esql = fmt.Sprintf("SELECT %s FROM %s WHERE end_time>'%s' ORDER BY id ASC", buycaiTableFileds, tbl, now.Format("2006-01-02 15:04:05"))
	rows, err := agt.dbConn.Query(esql)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err, esql)
			return nil, err
		}
	}

	for rows.Next() {
		v := &dbproto.BuycaiSaleIssue{}
		var openBalls sql.NullString
		if err = rows.Scan(&v.Id, &v.Issue, &startTime, &endTime, &openTime, &openBalls); err != nil {
			log.Println(err, esql)
			return nil, err
		}
		v.StartTime = startTime.Unix()
		v.EndTime = endTime.Unix()
		v.OpenTime = openTime.Unix()
		if openBalls.Valid {
			v.OpenBalls = openBalls.String
		}
		result.List = append(result.List, v)
	}
	if err = rows.Err(); err != nil {
		log.Println(err, esql)
	}
	return result, err
}

func (agt *DbBuycaiAgent) BuycaiQueryIssue(ctx context.Context, arg *dbproto.BuycaiQueryIssueArg) (*dbproto.BuycaiSaleIssue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	tbl := arg.GetCode()
	issue := arg.GetIssue()
	var startTime, endTime, openTime time.Time

	esql := fmt.Sprintf("SELECT %s FROM %s WHERE issue='%s'", buycaiTableFileds, tbl, issue)
	v := &dbproto.BuycaiSaleIssue{}
	var openBalls sql.NullString
	if err := agt.dbConn.QueryRow(esql).Scan(&v.Id, &v.Issue, &startTime, &endTime, &openTime, &openBalls); err != nil {
		log.Println(err, esql)
		return nil, err
	}
	v.StartTime = startTime.Unix()
	v.EndTime = endTime.Unix()
	v.OpenTime = openTime.Unix()
	if openBalls.Valid {
		v.OpenBalls = openBalls.String
	}
	return v, nil
}

func (agt *DbBuycaiAgent) BuycaiUpsertIssue(ctx context.Context, arg *dbproto.BuycaiUpsertIssueArg) (*dbproto.Nil, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	tbl := arg.GetCode()
	saleIssue := arg.GetSaleIssue()
	id := saleIssue.GetId()
	issue := saleIssue.GetIssue()
	startTime := time.Unix(saleIssue.GetStartTime(), 0).Format("2006-01-02 15:04:05")
	endTime := time.Unix(saleIssue.GetEndTime(), 0).Format("2006-01-02 15:04:05")
	openTime := time.Unix(saleIssue.GetOpenTime(), 0).Format("2006-01-02 15:04:05")
	openBalls := saleIssue.GetOpenBalls()
	log.Println(tbl, id, issue, startTime, endTime, openTime, openBalls)

	var esql string
	if id == 0 {
		esql = fmt.Sprintf("INSERT INTO %s(issue, start_time, end_time, open_time) VALUES('%s', '%s', '%s', '%s')",
			tbl, issue, startTime, endTime, openTime)
	} else {
		esql = fmt.Sprintf("UPDATE %s SET start_time='%s', end_time='%s', open_time='%s', open_balls='%s' WHERE issue='%s'",
			tbl, startTime, endTime, openTime, openBalls, issue)
	}

	log.Println(esql)
	_, err := agt.dbConn.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

func (agt *DbBuycaiAgent) BuycaiUpdateOpenBalls(ctx context.Context, arg *dbproto.BuycaiUpsertIssueArg) (*dbproto.Nil, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	tbl := arg.GetCode()
	saleIssue := arg.GetSaleIssue()
	issue := saleIssue.GetIssue()
	openBalls := saleIssue.GetOpenBalls()
	log.Println(tbl, issue, openBalls)

	esql := fmt.Sprintf("UPDATE %s SET open_balls='%s' WHERE issue='%s'", tbl, openBalls, issue)
	_, err := agt.dbConn.Exec(esql)
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}
