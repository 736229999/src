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
	TIME_FORMAT = "2006-01-02 15:04:05"
)

//获取购彩配置.
func (agt *DbAdminAgent) QueryBuycaiOptions(ctx context.Context, arg *dbproto.BuycaiOptionsReply) (*dbproto.BuycaiOptionsReply, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	offset := (arg.GetPage() - 1) * arg.GetSize()
	log.Printf("arg:%+v\n", arg)
	esql := fmt.Sprintf(`SELECT id, issue, start_time, end_time, open_time, open_balls FROM %s ORDER BY id DESC OFFSET %d LIMIT %d`, arg.GetLottery(), offset, arg.GetSize())

	st, err := agt.buycaiDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}
	rows, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	var buycai = &dbproto.BuycaiOptionsReply{}

	for rows.Next() {
		by := &dbproto.BuycaiOptions{}
		var start_time, end_time, open_time time.Time
		var openBalls sql.NullString
		if err = rows.Scan(&by.Id, &by.Issue, &start_time, &end_time, &open_time, &openBalls); err != nil {
			log.Printf("%+v\n", err)
			return nil, err
		}

		if openBalls.Valid {
			by.OpenBalls = openBalls.String
		}

		by.StartTime = start_time.Unix()
		by.EndTime = end_time.Unix()
		by.OpenTime = open_time.Unix()

		buycai.Buycai = append(buycai.Buycai, by)
	}

	//获取总的分页条数.
	buycai.Total, err = agt.queryBuycaiOptionsTotal(arg.GetLottery())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return buycai, nil
}

//获取总条数.
func (agt *DbAdminAgent) queryBuycaiOptionsTotal(tableName string) (int64, error) {

	esql := fmt.Sprintf(`SELECT count(*) as num FROM %s`, tableName)
	st, err := agt.buycaiDbConn.Prepare(esql)
	if err != nil {
		return 0, err
	}
	var num int64 = 0
	if err = st.QueryRow().Scan(&num); err != nil {
		return 0, err
	}

	return num, nil
}

//获取期次信息.
func (agt *DbAdminAgent) QueryLotteryIssue(ctx context.Context, arg *dbproto.StringValue) (*dbproto.BuycaiOptions, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := fmt.Sprintf(`SELECT id, issue, start_time, end_time, open_time, open_balls FROM %s ORDER BY id DESC LIMIT 1`, arg.GetValue())
	st, err := agt.buycaiDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}

	lottery := &dbproto.BuycaiOptions{}

	var start_time, end_time, open_time time.Time
	var openBalls sql.NullString
	if err = st.QueryRow().Scan(&lottery.Id, &lottery.Issue, &start_time, &end_time, &open_time, &openBalls); err != nil {
		if err == sql.ErrNoRows {
			return &dbproto.BuycaiOptions{}, nil
		}
		log.Printf("%+v\n", err)
		return nil, err
	}

	lottery.StartTime = start_time.Unix()
	lottery.EndTime = end_time.Unix()
	lottery.OpenTime = open_time.Unix()

	if openBalls.Valid {
		lottery.OpenBalls = openBalls.String
	}
	return lottery, nil
}

//添加开奖期次配置数据.
func (agt *DbAdminAgent) InsertLotteryOptions(ctx context.Context, arg *dbproto.BuycaiOptionsIssue) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	tx, err := agt.buycaiDbConn.Begin()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	esql := fmt.Sprintf(`INSERT INTO %s (issue, start_time, end_time, open_time) VALUES ($1, $2, $3, $4)`, arg.GetLottery())
	st, err := tx.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		tx.Rollback()
		return nil, err
	}
	for _, v := range arg.GetBuycai() {
		var start_time, end_time, open_time time.Time
		start_time = time.Unix(v.GetStartTime(), 0)
		end_time = time.Unix(v.GetEndTime(), 0)
		open_time = time.Unix(v.GetOpenTime(), 0)
		_, err := st.Exec(v.GetIssue(), start_time, end_time, open_time)
		if err != nil {
			log.Printf("%+v\n", err)
			tx.Rollback()
			return nil, err
		}
	}

	return &dbproto.Nil{}, tx.Commit()
}

//获取玩法时间列表.
func (agt *DbAdminAgent) QueryPlayTimeSettingList(ctx context.Context, arg *dbproto.Nil) (*dbproto.PlayTimeSettingsList, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	st, err := agt.buycaiDbConn.Prepare(`SELECT id, lottery_id, start_time, end_time, chase_start_time FROM play_time_settings`)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	rows, err := st.Query()

	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	playTimeSettingList := &dbproto.PlayTimeSettingsList{}

	for rows.Next() {
		list := &dbproto.PlayTimeSettings{}
		if err = rows.Scan(&list.Id, &list.LotteryId, &list.StartTime, &list.EndTime, &list.ChaseStartTime); err != nil {
			log.Printf("%+v\n", err)
			return nil, err
		}
		playTimeSettingList.List = append(playTimeSettingList.List, list)
	}

	return playTimeSettingList, nil
}

//添加彩票的玩法时间设置.
func (agt *DbAdminAgent) InsertPlayTimeSettings(ctx context.Context, arg *dbproto.PlayTimeSettings) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	st, err := agt.buycaiDbConn.Prepare(`INSERT INTO play_time_settings (lottery_id, start_time, end_time, chase_start_time) VALUES ($1, $2, $3, $4)`)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	_, err = st.Exec(arg.GetLotteryId(), arg.GetStartTime(), arg.GetEndTime(), arg.GetChaseStartTime())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

//更新玩法时间间隔设置.
func (agt *DbAdminAgent) UpdatePlayTimeSettings(ctx context.Context, arg *dbproto.PlayTimeSettings) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	st, err := agt.buycaiDbConn.Prepare(`UPDATE play_time_settings SET start_time = $1, end_time = $2, chase_start_time = $3 WHERE lottery_id = $4`)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	_, err = st.Exec(arg.GetStartTime(), arg.GetEndTime(), arg.GetChaseStartTime(), arg.GetLotteryId())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

//根据id删除期号.
func (agt *DbAdminAgent) DeleteIssueById(ctx context.Context, arg *dbproto.BuycaiOptionsIssue) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := fmt.Sprintf(`DELETE FROM %s WHERE id = %d`, arg.GetLottery(), arg.GetId())
	st, err := agt.buycaiDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}
	defer st.Close()

	_, err = st.Exec()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return &dbproto.Nil{}, nil
}

//根据id获取期号.
func (agt *DbAdminAgent) QueryLotteryBuycaiOptionsById(ctx context.Context, arg *dbproto.BuycaiOptionsIssue) (*dbproto.BuycaiOptions, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string{
		"id", "issue", "start_time", "end_time", "open_time", "open_balls",
	}
	esql := fmt.Sprintf(`SELECT %s FROM %s WHERE id = %d`, strings.Join(filedList, ", "), arg.GetLottery(), arg.GetId())

	log.Println(esql)
	st, err := agt.buycaiDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	defer st.Close()

	issue := &dbproto.BuycaiOptions{}
	var open_balls sql.NullString
	var start_time, end_time, open_time time.Time
	value := []interface{}{
		&issue.Id, &issue.Issue, &start_time, &end_time, &open_time, &open_balls,
	}
	if err = st.QueryRow().Scan(value...); err != nil {
		if err == sql.ErrNoRows {
			return &dbproto.BuycaiOptions{}, nil
		}
		log.Printf("%+v\n", err)
		return nil, err
	}

	if open_balls.Valid {
		issue.OpenBalls = open_balls.String
	}

	issue.StartTime = start_time.Unix()
	issue.EndTime = end_time.Unix()
	issue.OpenTime = open_time.Unix()

	return issue, nil
}

//更新期号.
func (agt *DbAdminAgent) UpdateLotteryBuycaiOptionsById(ctx context.Context, arg *dbproto.BuycaiOptionsUpdateIssue) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	buycai := arg.GetBuycai()
	fieldList := []string{
		fmt.Sprintf("issue = '%s'", buycai.GetIssue()),
		fmt.Sprintf("start_time = '%s'", time.Unix(buycai.GetStartTime(), 0).Format(TIME_FORMRT)),
		fmt.Sprintf("end_time = '%s'", time.Unix(buycai.GetEndTime(), 0).Format(TIME_FORMRT)),
		fmt.Sprintf("open_time = '%s'", time.Unix(buycai.GetOpenTime(), 0).Format(TIME_FORMRT)),
		fmt.Sprintf("open_balls = '%s'", buycai.GetOpenBalls()),
	}
	esql := fmt.Sprintf(`UPDATE %s SET %s WHERE id = %d`, arg.GetLottery(), strings.Join(fieldList, ", "), arg.GetId())
	st, err := agt.buycaiDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}
	defer st.Close()

	_, err = st.Exec()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	return &dbproto.Nil{}, nil
}

//添加初始化期号数据.
func (agt *DbAdminAgent) InsertLotteryBuycaiOptions(ctx context.Context, arg *dbproto.BuycaiOptionsUpdateIssue) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string{
		"issue", "start_time", "end_time", "open_time", "open_balls",
	}
	buycai := arg.GetBuycai()
	valueList := []string{
		fmt.Sprintf(`'%s'`, buycai.GetIssue()),
		fmt.Sprintf(`'%s'`, time.Unix(buycai.GetStartTime(), 0).Format(TIME_FORMRT)),
		fmt.Sprintf(`'%s'`, time.Unix(buycai.GetEndTime(), 0).Format(TIME_FORMRT)),
		fmt.Sprintf(`'%s'`, time.Unix(buycai.GetOpenTime(), 0).Format(TIME_FORMRT)),
		fmt.Sprintf(`'%s'`, buycai.GetOpenBalls()),
	}
	esql := fmt.Sprintf(`INSERT INTO %s(%s) VALUES (%s)`, arg.GetLottery(), strings.Join(filedList, ", "), strings.Join(valueList, ", "))
	st, err := agt.buycaiDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}
	defer st.Close()

	_, err = st.Exec()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return &dbproto.Nil{}, err
}

//获取每天的购彩数据.
func (agt *DbAdminAgent) QueryBuycaiListByDay(ctx context.Context, arg *dbproto.BuycaiStatisticsList) (*dbproto.BuycaiStatisticsList, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	return &dbproto.BuycaiStatisticsList{}, nil
}
