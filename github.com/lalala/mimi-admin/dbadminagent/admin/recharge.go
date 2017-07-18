package admin

import (
	"fmt"
	"log"
	"strings"

	"database/sql"

	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"golang.org/x/net/context"
	"time"
)

//查询充值列表.
func (agt *DbAdminAgent) QueryRechargeList(ctx context.Context, arg *dbproto.RechargeOrderList) (*dbproto.RechargeOrderList, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string{
		"to_char(to_timestamp(create_time),'yyyy-MM-dd') AS created_time",
		"COUNT(id) AS recharge_num",
		"SUM(money) AS recharge_total_amount",
	}
	//按照天数统计充值次数，和金额.
	esql := fmt.Sprintf("SELECT %s FROM recharge_order WHERE status = %d ", strings.Join(filedList, ", "), dbproto.RechargeOrder_SUCCESS)

	if arg.GetStartTime() > 0 {
		esql += fmt.Sprintf(" AND create_time >= %d ", arg.GetStartTime())
	}
	if arg.GetEndTime() > 0 {
		esql += fmt.Sprintf(" AND create_time <= %d ", arg.GetEndTime())
	}

	esql += " GROUP BY created_time ORDER BY created_time DESC"

	totalSql := fmt.Sprintf("SELECT COUNT(created_time) FROM (%s) AS temporary_table", esql)

	offset := (arg.GetPage() - 1) * arg.GetSize()
	esql += fmt.Sprintf(" OFFSET %d LIMIT %d", offset, arg.GetSize())

	rechargeList, err := agt.rechargeList(esql, 0)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	st, err := agt.rechargeDbConn.Prepare(totalSql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	if err = st.QueryRow().Scan(&rechargeList.Total); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	return rechargeList, nil
}

//根据os来查询充值的金额和数量.
func (agt *DbAdminAgent) getRechargeByOs(os string, recharge_time int32, source int) (recharge_num int32, recharge_amount string, err error) {

	var end_time int32
	if source == 0 {
		end_time = recharge_time + 24*3600 - 1
	} else if source == 1 {
		//获取这个月的最后一天.
		end_time = int32(time.Unix(int64(recharge_time), 0).AddDate(0, 1, -1).Unix() + 24*3600 - 1)
	} else if source == 2 {
		//年.
		end_time = int32(time.Unix(int64(recharge_time), 0).AddDate(1, 0, -1).Unix() + 24*3600 - 1)
	}
	whereList := []string{
		fmt.Sprintf("WHERE  status = %d  AND os = '%s'", dbproto.RechargeOrder_SUCCESS, os),
		fmt.Sprintf(" AND create_time BETWEEN %d", recharge_time),
		fmt.Sprintf(" AND %d", end_time),
	}
	filedList := []string{
		"COUNT(id) AS recharge_num",
		"SUM(money) AS recharge_amount",
	}
	esql := fmt.Sprintf("SELECT %s FROM recharge_order %s", strings.Join(filedList, ", "), strings.Join(whereList, ""))

	st, err := agt.rechargeDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return 0, "", err
	}

	var amount sql.NullFloat64
	if err = st.QueryRow().Scan(&recharge_num, &amount); err != nil {
		log.Printf("%+v\n", err)
		return 0, "", err
	}
	if amount.Valid {
		money := amount.Float64 / 100
		recharge_amount = fmt.Sprintf("%.2f", money)
	} else {
		recharge_amount = "00.00"
	}

	return recharge_num, recharge_amount, nil
}

//根据充值方式来获取充值的次数和金额.
func (agt *DbAdminAgent) getRechargeByPaymentMethod(method int32, recharge_time int32, source int) (recharge_num int32, recharge_amount string, err error) {

	var end_time int32
	if source == 0 {
		end_time = recharge_time + 24*3600 - 1
	} else if source == 1 {
		//获取这个月的最后一天.
		end_time = int32(time.Unix(int64(recharge_time), 0).AddDate(0, 1, -1).Unix() + 24*3600 - 1)
	} else if source == 2 {
		//年.
		end_time = int32(time.Unix(int64(recharge_time), 0).AddDate(1, 0, -1).Unix() + 24*3600 - 1)
	}
	fileldList := []string{
		"COUNT(id) AS recharge_num", "SUM(money) AS recharge_amount",
	}
	whereList := []string{
		fmt.Sprintf(" WHERE status = %d", dbproto.RechargeOrder_SUCCESS),
		fmt.Sprintf(" AND payment_method = %d", method),
		fmt.Sprintf(" AND create_time  BETWEEN %d AND %d", recharge_time, end_time),
	}
	esql := fmt.Sprintf("SELECT %s FROM recharge_order %s", strings.Join(fileldList, ", "), strings.Join(whereList, ""))
	st, err := agt.rechargeDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return 0, "", err
	}

	var amount sql.NullFloat64
	if err = st.QueryRow().Scan(&recharge_num, &amount); err != nil {
		log.Printf("%+v\n", err)
		return 0, "", err
	}
	if amount.Valid {
		money := amount.Float64 / 100
		recharge_amount = fmt.Sprintf("%.2f", money)
	} else {
		recharge_amount = "00.00"
	}

	return recharge_num, recharge_amount, nil
}

//按月来获取充值记录.
func (agt *DbAdminAgent) QueryRechargeListByMonth(ctx context.Context, arg *dbproto.Nil) (*dbproto.RechargeOrderList, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string{
		"to_char(to_timestamp(create_time),'yyyy-MM') AS created_time",
		"COUNT(id) AS recharge_num",
		"SUM(money) AS recharge_total_amount",
	}
	//按照天数统计充值次数，和金额.
	esql := fmt.Sprintf("SELECT %s FROM recharge_order WHERE status = %d ", strings.Join(filedList, ", "), dbproto.RechargeOrder_SUCCESS)

	esql += " GROUP BY created_time ORDER BY created_time DESC"

	rechrge, err := agt.rechargeList(esql, 1)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return rechrge, nil
}

func (agt *DbAdminAgent) rechargeList(esql string, source int) (*dbproto.RechargeOrderList, error) {
	st, err := agt.rechargeDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}

	rows, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	rechargeList := &dbproto.RechargeOrderList{}
	for rows.Next() {

		recharge := &dbproto.RechargeStatistics{}
		var recharge_time string
		var amount sql.NullFloat64
		value := []interface{}{
			&recharge_time, &recharge.RechargeNum, &amount,
		}
		if err = rows.Scan(value...); err != nil {
			if err == sql.ErrNoRows {
				return &dbproto.RechargeOrderList{}, nil
			}
			log.Printf("%+v\n", err)
			return nil, err
		}

		if amount.Valid {
			money := amount.Float64 / 100
			recharge.RechargeTotalAmount = fmt.Sprintf("%.2f", money)
		} else {
			recharge.RechargeTotalAmount = "00.00"
		}

		var t time.Time
		if source == 0 {
			t, err = time.ParseInLocation("2006-01-02", recharge_time, time.Local)
		} else if source == 1 {
			t, err = time.ParseInLocation("2006-01", recharge_time, time.Local)
		} else if source == 2 {
			t, err = time.ParseInLocation("2006", recharge_time, time.Local)
		}
		if err != nil {
			log.Printf("%+v\n", err)
			return nil, err
		}
		recharge.RechargeTime = int32(t.Unix())

		//查询安卓充值的次数和金额.
		recharge.AndroidRechargeNum, recharge.AndroidRechargeAmount, err = agt.getRechargeByOs("Android", recharge.RechargeTime, source)
		if err != nil {
			log.Printf("%+v\n", err)
			return nil, err
		}

		//查询iOS的充值次数和金额.
		recharge.IphoneRechargeNum, recharge.IphoneRechargeAmount, err = agt.getRechargeByOs("iOS", recharge.RechargeTime, source)
		if err != nil {
			log.Printf("%+v\n", err)
			return nil, err
		}

		//查询支付宝充值的次数和金额.
		recharge.AlipayRechargeNum, recharge.AlipayRechargeAmount, err = agt.getRechargeByPaymentMethod(int32(dbproto.RechargeOrder_ALIPAY), recharge.RechargeTime, source)
		if err != nil {
			log.Printf("%+v\n", err)
			return nil, err
		}

		//查询微信充值的次数和金额.
		recharge.WechatRechargeNum, recharge.WechatRechargeAmount, err = agt.getRechargeByPaymentMethod(int32(dbproto.RechargeOrder_WECHAT), recharge.RechargeTime, source)
		if err != nil {
			log.Printf("%+v\n", err)
			return nil, err
		}
		rechargeList.List = append(rechargeList.List, recharge)

	}

	return rechargeList, nil
}

//根据年份获取充值数据.
func (agt *DbAdminAgent) QueryRechangeListByYear(ctx context.Context, arg *dbproto.Nil) (*dbproto.RechargeOrderList, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string{
		"to_char(to_timestamp(create_time),'yyyy') AS created_time",
		"COUNT(id) AS recharge_num",
		"SUM(money) AS recharge_total_amount",
	}
	//按照天数统计充值次数，和金额.
	esql := fmt.Sprintf("SELECT %s FROM recharge_order WHERE status = %d ", strings.Join(filedList, ", "), dbproto.RechargeOrder_SUCCESS)

	esql += " GROUP BY created_time ORDER BY created_time DESC"

	rechrge, err := agt.rechargeList(esql, 2)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return rechrge, nil
}
