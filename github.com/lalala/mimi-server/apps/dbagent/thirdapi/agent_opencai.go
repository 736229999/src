package thirdapi

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	. "github.com/caojunxyz/mimi-server/apps/dbagent/helper"
	"golang.org/x/net/context"

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
)

const TIME_MINUTE_LAYOUT = "2006-01-02 15:04:00"
const TIME_SECOND_LAYOUT = "2006-01-02 15:04:05"

func tableName(code string) string {
	return fmt.Sprintf("opencai_%s", code)
}

func (agt *DbThirdApiAgent) OpencaiQueryByOpendate(arg *dbproto.OpencaiQueryArg, stream dbproto.DbThirdApiAgent_OpencaiQueryByOpendateServer) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	args := arg.GetArgs()
	if len(args) != 1 {
		return ErrInvalidArgs
	}
	tbl := tableName(arg.GetCode())
	openDate, err := time.Parse("2006-01-02", args[0])
	if err != nil {
		log.Println(err, args[0])
		return ErrInvalidArgs
	}
	start := openDate.Format(TIME_SECOND_LAYOUT)
	end := time.Date(openDate.Year(), openDate.Month(), openDate.Day()+1, 0, 0, 0, 0, time.Local).Format(TIME_SECOND_LAYOUT)
	esql := fmt.Sprintf("SELECT issue, opentime, balls, grabtime, grabsource, detail FROM %s WHERE opentime>'%s' AND opentime <='%s'", tbl, start, end)
	rows, err := agt.dbConn.Query(esql)
	if err != nil {
		log.Println(err, esql)
		return err
	}

	for rows.Next() {
		var issue string
		var openTime time.Time
		var balls string
		var grabTime time.Time
		var grabSource string
		var data []byte
		if err = rows.Scan(&issue, &openTime, &balls, &grabTime, &grabSource, &data); err != nil {
			log.Println(err, esql)
			break
		} else {
			openInfo := &dbproto.OpenInfo{
				Issue:      issue,
				OpenTime:   openTime.Unix(),
				Balls:      balls,
				GrabTime:   grabTime.Unix(),
				GrabSource: grabSource,
			}

			if data != nil {
				detail := &dbproto.OpenDetail{}
				if err := json.Unmarshal(data, detail); err != nil {
					log.Println(err)
				} else {
					openInfo.Detail = detail
				}
			}
			stream.Send(openInfo)
		}
	}

	if err = rows.Err(); err != nil {
		log.Println(err, esql)
	}
	return err
}

func (agt *DbThirdApiAgent) OpencaiQueryByIssue(ctx context.Context, arg *dbproto.OpencaiQueryArg) (*dbproto.OpenInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	args := arg.GetArgs()
	if len(args) != 1 {
		return nil, ErrInvalidArgs
	}
	tbl := tableName(arg.GetCode())
	var issue string
	var openTime time.Time
	var balls string
	var grabTime time.Time
	var grabSource string
	var data []byte

	esql := fmt.Sprintf("SELECT issue, opentime, balls, grabtime, grabsource, detail FROM %s WHERE issue='%s'", tbl, args[0])
	if err := agt.dbConn.QueryRow(esql).Scan(&issue, &openTime, &balls, &grabTime, &grabSource, &data); err != nil {
		// log.Println(err, esql)
		return nil, err
	}

	openInfo := &dbproto.OpenInfo{
		Issue:      issue,
		OpenTime:   openTime.Unix(),
		Balls:      balls,
		GrabTime:   grabTime.Unix(),
		GrabSource: grabSource,
	}

	if data != nil {
		detail := &dbproto.OpenDetail{}
		if err := json.Unmarshal(data, detail); err != nil {
			log.Println(err)
		} else {
			openInfo.Detail = detail
		}
	}
	return openInfo, nil
}

func (agt *DbThirdApiAgent) OpencaiQueryByLatestNum(arg *dbproto.OpencaiQueryArg, stream dbproto.DbThirdApiAgent_OpencaiQueryByLatestNumServer) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	args := arg.GetArgs()
	if len(args) != 1 {
		return ErrInvalidArgs
	}
	tbl := tableName(arg.GetCode())
	num, err := strconv.Atoi(args[0])
	if err != nil {
		log.Println(err)
		return ErrInvalidArgs
	}
	if num <= 0 {
		log.Println(num)
		return ErrInvalidArgs
	}
	esql := fmt.Sprintf("SELECT issue, opentime, balls, grabtime, grabsource, detail FROM %s ORDER BY opentime DESC LIMIT %d", tbl, num)
	rows, err := agt.dbConn.Query(esql)
	if err != nil {
		log.Println(err, esql)
		return err
	}

	for rows.Next() {
		var issue string
		var openTime time.Time
		var balls string
		var grabTime time.Time
		var grabSource string
		var data []byte
		if err = rows.Scan(&issue, &openTime, &balls, &grabTime, &grabSource, &data); err != nil {
			log.Println(err, esql)
			break
		} else {
			openInfo := &dbproto.OpenInfo{
				Issue:      issue,
				OpenTime:   openTime.Unix(),
				Balls:      balls,
				GrabTime:   grabTime.Unix(),
				GrabSource: grabSource,
			}

			if data != nil {
				detail := &dbproto.OpenDetail{}
				if err := json.Unmarshal(data, detail); err != nil {
					log.Println(err)
				} else {
					openInfo.Detail = detail
				}
			}
			stream.Send(openInfo)
		}
	}

	if err = rows.Err(); err != nil {
		log.Println(err, esql)
	}
	return err
}

func (agt *DbThirdApiAgent) OpencaiInsert(ctx context.Context, arg *dbproto.OpencaiInsertArg) (*dbproto.Nil, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	tbl := tableName(arg.GetCode())
	info := arg.GetInfo()
	issue := info.GetIssue()
	openTime := time.Unix(info.GetOpenTime(), 0)
	balls := info.GetBalls()
	grabTime := time.Unix(info.GetGrabTime(), 0)
	grabSource := info.GetGrabSource()
	var esql string
	var err error
	var data []byte
	if detail := info.GetDetail(); detail != nil {
		data, err = json.Marshal(detail)
		if err != nil {
			log.Println(err)
		}
	}

	if data == nil {
		esql = fmt.Sprintf("INSERT INTO %s(issue, opentime, balls, grabtime, grabsource) VALUES($1, $2, $3, $4, $5) ON CONFLICT (issue) DO NOTHING", tbl)
		_, err = agt.dbConn.Exec(esql, issue, openTime.Format(TIME_SECOND_LAYOUT), balls, grabTime.Format(TIME_SECOND_LAYOUT), grabSource)
	} else {
		esql = fmt.Sprintf("INSERT INTO %s(issue, opentime, balls, grabtime, grabsource, detail) VALUES($1, $2, $3, $4, $5, $6) ON CONFLICT (issue) DO NOTHING", tbl)
		_, err = agt.dbConn.Exec(esql, issue, openTime.Format(TIME_SECOND_LAYOUT), balls, grabTime.Format(TIME_SECOND_LAYOUT), grabSource, data)
	}
	if err != nil {
		log.Println(err, esql)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

func (agt *DbThirdApiAgent) OpencaiUpsertDetail(ctx context.Context, arg *dbproto.OpencaiUpsertDetailArg) (*dbproto.Nil, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	tbl := tableName(arg.GetCode())
	issue := arg.GetIssue()
	var err error
	var data []byte
	if detail := arg.GetDetail(); detail != nil {
		data, err = json.Marshal(detail)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	if data != nil {
		esql := fmt.Sprintf("UPDATE %s SET detail=%v WHERE issue=%s", tbl, data, issue)
		_, err = agt.dbConn.Exec(esql)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
	}
	return &dbproto.Nil{}, nil
}

func (agt *DbThirdApiAgent) OpencaiQueryLatestIssue(ctx context.Context, arg *dbproto.StringValue) (*dbproto.OpenInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	tbl := tableName(arg.GetValue())
	esql := fmt.Sprintf("SELECT issue, opentime, balls, grabtime, grabsource, detail FROM %s ORDER BY opentime DESC LIMIT 1", tbl)
	openInfo := &dbproto.OpenInfo{}
	var openTime time.Time
	var grabTime time.Time
	var data []byte
	if err := agt.dbConn.QueryRow(esql).Scan(&openInfo.Issue, &openTime, &openInfo.Balls, &grabTime, &openInfo.GrabSource, &data); err != nil {
		log.Println(err, esql)
		return nil, err
	}

	if data != nil {
		detail := &dbproto.OpenDetail{}
		if err := json.Unmarshal(data, detail); err != nil {
			log.Println(err)
		} else {
			openInfo.Detail = detail
		}
	}
	return openInfo, nil
}
