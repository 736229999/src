package admin

import (
	"database/sql"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"golang.org/x/net/context"
	"log"
	"time"
)

func (agt *DbAdminAgent) InsertLotteryOption(ctx context.Context, arg *dbproto.LotteryOptions) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := `INSERT INTO lottery_options (id, lottery_name, is_plus_award, info, stop_sale, create_time, update_time) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	st, err := agt.optionsDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	unix := time.Now().Unix()
	_, err = st.Exec(arg.GetId(), arg.GetLotteryName(), arg.GetIsPlusAward(), arg.GetInfo(), arg.GetStopSale(), unix, unix)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return &dbproto.Nil{}, nil
}

//获取彩票配置的列表.
func (agt *DbAdminAgent) QueryLotteryOptionsList(ctx context.Context, arg *dbproto.Nil) (*dbproto.LotteryOptionsList, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := `SELECT id, lottery_name, is_plus_award, info, stop_sale, create_time, update_time  FROM lottery_options`
	st, err := agt.optionsDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	rows, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	list := &dbproto.LotteryOptionsList{}
	for rows.Next() {
		options := &dbproto.LotteryOptions{}
		var boolValue sql.NullBool
		if err = rows.Scan(&options.Id, &options.LotteryName, &boolValue, &options.Info, &options.StopSale, &options.CreateTime, &options.UpdateTime); err != nil {
			if err == sql.ErrNoRows {
				return &dbproto.LotteryOptionsList{}, nil
			}
			log.Printf("%+v\n", err)
			return nil, err
		}

		if boolValue.Valid {
			options.IsPlusAward = boolValue.Bool
		}

		list.List = append(list.List, options)
	}

	return list, nil
}

//根据id获取彩种配置.
func (agt *DbAdminAgent) GetLotteryOptionsById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.LotteryOptions, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := `SELECT id, lottery_name, is_plus_award, info, stop_sale, create_time, update_time FROM lottery_options WHERE id = $1`
	st, err := agt.optionsDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	option := &dbproto.LotteryOptions{}
	var boolValue sql.NullBool
	if err = st.QueryRow(arg.GetValue()).Scan(&option.Id, &option.LotteryName, &boolValue, &option.Info, &option.StopSale, &option.CreateTime, &option.UpdateTime); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	if boolValue.Valid {
		option.IsPlusAward = boolValue.Bool
	}
	return option, nil
}

//更新彩票种类的配置.
func (agt *DbAdminAgent) UpdateLotteryOptionsById(ctx context.Context, arg *dbproto.LotteryOptions) (*dbproto.Nil, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := `UPDATE lottery_options SET lottery_name = $1, is_plus_award = $2, info = $3, stop_sale = $4, update_time = $5 where id = $6`
	log.Println("sql:", esql)
	st, err := agt.optionsDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	_, err = st.Exec(arg.GetLotteryName(), arg.GetIsPlusAward(), arg.GetInfo(), arg.GetStopSale(), time.Now().Unix(), arg.GetId())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	return &dbproto.Nil{}, err
}
