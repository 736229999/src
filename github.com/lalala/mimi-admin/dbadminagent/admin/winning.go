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

//import (
//	"golang.org/x/net/context"
//	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
//	"log"
//	"fmt"
//	"github.com/caojunxyz/mimi-admin/dbadminagent/helper"
//)

//获取中奖记录列表.
//func (agt *DbAdminAgent) QueryWinningList(ctx context.Context, arg *dbproto.WinningList) (*dbproto.WinningList, error) {
//
//	defer func() {
//		if err := recover(); err != nil {
//			log.Println("recover from panic:", err)
//		}
//	}()
//
//	esql := "select w.id, w.user_id, w.username, w.info, w.status, w.create_time from winning as w left join ("
//	esql += fmt.Sprintf(` select * from dblink('dbname=%s', 'select account_id,phone from phone_user') as phone_user (account_id int, phone text) where 1 = 1 `, helper.DB_UC)
//	if arg.GetMobile() == "" {
//		esql += fmt.Sprintf(` and phone = '%s'`, arg.GetMobile())
//	}
//	esql += fmt.Sprintf(` ) as p on w.user_id = p.account_id`)
//
//	//分页sql.
//	totalSql := esql
//
//	offset := (arg.GetPage() - 1) * arg.GetSize()
//	esql += fmt.Sprintf(` order by w.id desc offset %d limit %d`, offset, arg.GetSize())
//
//	st, err := agt.opencaiDbConn.Prepare(esql)
//	if err != nil {
//		log.Printf("%+v\n", err, esql)
//		return nil, err
//	}
//	rows, err := st.Query()
//	if err != nil {
//		log.Printf("%+v\n", err)
//		return nil, err
//	}
//
//	for rows.Next() {
//
//	}
//
//}

//或者每种彩票最新一期的中奖数据.
func (agt *DbAdminAgent) QueryWinningLotteryList(ctx context.Context, arg *dbproto.WinningLotteryList) (*dbproto.LotteryWinningNoList, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	winningList := &dbproto.LotteryWinningNoList{}
	for _, v := range arg.GetList() {
		fieldList := []string{
			"issue", "open_time", "open_balls",
		}
		esql := fmt.Sprintf(`SELECT %s FROM %s WHERE open_balls IS NOT NULl ORDER BY id DESC LIMIT 1`, strings.Join(fieldList, ", "), v.GetLottery())
		log.Println(esql)
		st, err := agt.buycaiDbConn.Prepare(esql)
		if err != nil {
			log.Printf("%+v\n", err)
			return nil, err
		}
		defer st.Close()

		winning := &dbproto.LotteryWinningNo{}
		var open_time time.Time
		value := []interface{}{
			&winning.Issue, &open_time, &winning.OpenBalls,
		}

		if err = st.QueryRow().Scan(value...); err != nil {
			if err == sql.ErrNoRows {
				winning.OpenTime = 0
				winning.OpenBalls = ""
				winning.Lottery = v.GetLottery()
				winningList.List = append(winningList.List, winning)
				continue
			}
			log.Printf("%+v\n", err)
			return nil, err
		}

		winning.OpenTime = open_time.Unix()
		winning.Lottery = v.GetLottery()

		winningList.List = append(winningList.List, winning)
	}

	log.Printf("%+v\n", winningList)

	return winningList, nil
}

//根据期号查询开奖信息.
func (agt *DbAdminAgent) QueryWinningByIssue(ctx context.Context, arg *dbproto.LotteryWinningNo) (*dbproto.LotteryWinningNoList, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string{
		"id", "issue", "start_time", "end_time", "open_time", "open_balls",
	}
	esql := fmt.Sprintf(`SELECT %s FROM %s WHERE issue = '%s' AND open_balls IS NOT NULL`, strings.Join(filedList, ","), arg.GetLottery(), arg.GetIssue())
	st, err := agt.buycaiDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}
	defer st.Close()

	winningList := &dbproto.LotteryWinningNoList{}
	var start_time, end_time, open_time time.Time

	rows, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	for rows.Next() {
		winning := &dbproto.LotteryWinningNo{}
		value := []interface{}{
			&winning.Issue, &winning.Issue, &start_time, &end_time, &open_time, &winning.OpenBalls,
		}

		if err = rows.Scan(value...); err != nil {
			if err == sql.ErrNoRows {
				return &dbproto.LotteryWinningNoList{}, nil
			}
			log.Printf("%+v\n", err)
			return nil, err
		}

		winning.StartTime = start_time.Unix()
		winning.EndTime = end_time.Unix()
		winning.OpenTime = open_time.Unix()
		winning.Lottery = arg.GetLottery()

		winningList.List = append(winningList.List, winning)

	}

	return winningList, nil
}

//获取中奖的历史数据.
func (agt *DbAdminAgent) QueryWinningHistoryListByLottery(ctx context.Context, arg *dbproto.WinningListByLottery) (*dbproto.WinningListByLottery, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string{
		"id", "issue", "start_time", "end_time", "open_time", "open_balls",
	}
	esql := fmt.Sprintf(`SELECT %s FROM %s WHERE open_balls IS NOT NULL`, strings.Join(filedList, ", "), arg.GetLottery())

	if arg.GetStartTime() > 0 {
		str := fmt.Sprintf("%s", time.Unix(int64(arg.GetStartTime()), 0))
		esql += fmt.Sprintf(` AND start_time >= '%s'`, strings.TrimSuffix(str, "00 CST"))
	}

	if arg.GetEndTime() > 0 {
		str := fmt.Sprintf("%s", time.Unix(int64(arg.GetEndTime()), 0))
		esql += fmt.Sprintf(` AND end_time <= '%s'`, strings.TrimSuffix(str, "00 CST"))
	}

	if len(arg.GetIssue()) > 0 {
		esql += fmt.Sprintf(` AND issue LIKE '%s'`, "%"+arg.GetIssue()+"%")
	}

	offset := (arg.GetPage() - 1) * arg.GetSize()
	totalSql := strings.Replace(esql, strings.Join(filedList, ", "), "COUNT(*) AS num", -1)

	esql += fmt.Sprintf(` ORDER BY id DESC OFFSET %d LIMIT %d`, offset, arg.GetSize())

	log.Println("esql:", esql)
	st, err := agt.buycaiDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	list := &dbproto.WinningListByLottery{}
	for rows.Next() {
		winning := &dbproto.LotteryWinningNo{}
		var start_time, end_time, open_time time.Time
		value := []interface{}{
			&winning.Id, &winning.Issue, &start_time, &end_time, &open_time, &winning.OpenBalls,
		}
		if err = rows.Scan(value...); err != nil {
			if err == sql.ErrNoRows {
				return &dbproto.WinningListByLottery{}, nil
			}
			log.Printf("%+v\n", err)
			return nil, err
		}

		winning.StartTime = start_time.Unix()
		winning.EndTime = end_time.Unix()
		winning.OpenTime = open_time.Unix()
		winning.Lottery = arg.GetLottery()

		list.List = append(list.List, winning)
	}

	st, err = agt.buycaiDbConn.Prepare(totalSql)
	if err != nil {
		log.Printf("%+v\n", err, totalSql)
		return nil, err
	}

	if err = st.QueryRow().Scan(&list.Total); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return list, nil
}
