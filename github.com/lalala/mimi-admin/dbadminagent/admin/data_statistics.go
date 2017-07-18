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

func (agt *DbAdminAgent) QueryOrderAndIncome(ctx context.Context, arg *dbproto.Nil) (*dbproto.StatisticsOrderAndIncome, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	sta := &dbproto.StatisticsOrderAndIncome{}
	//今日新增收入和新增的购彩订单.
	filedList := []string{
		"SUM(money) AS new_income",
		"COUNT(id) AS new_buycai_order",
	}

	esql := fmt.Sprintf(`SELECT %s FROM buycai_vendor_order WHERE vendor_req_time BETWEEN %d AND %d`, strings.Join(filedList, ", "), time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).Unix(), time.Now().Unix())
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}

	var new_income, income_total sql.NullFloat64
	value := []interface{}{
		&new_income, &sta.NewBuycaiOrder,
	}
	if err := st.QueryRow().Scan(value...); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		log.Printf("%+v\n", err)
		return nil, err
	}

	if new_income.Valid {
		sta.NewIncome = fmt.Sprintf("%.2f", new_income.Float64)
	} else {
		sta.NewIncome = "00.00"
	}

	//查询历史总收入.
	history_income_sql := fmt.Sprintf(`SELECT SUM(money) AS income_total FROM buycai_vendor_order`)
	st, err = agt.ucDbConn.Prepare(history_income_sql)
	if err != nil {
		log.Printf("%+v\n", err, history_income_sql)
		return nil, err
	}
	if err = st.QueryRow().Scan(&income_total); err != nil {
		if err == sql.ErrNoRows {
			income_total = sql.NullFloat64{}
		} else {
			log.Printf("%+v\n", err)
			return nil, err
		}
	}
	if income_total.Valid {
		sta.IncomeTotal = fmt.Sprintf("%.2f", income_total.Float64)
	} else {
		sta.IncomeTotal = "00.00"
	}

	//查询用户新增订单.
	new_user_order := fmt.Sprintf(`SELECT COUNT(id) AS new_user_order FROM buycai_user_order WHERE order_time BETWEEN %d AND %d`, time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).Unix(), time.Now().Unix())
	st, err = agt.ucDbConn.Prepare(new_user_order)
	if err != nil {
		log.Printf("%+v\n", err, new_user_order)
		return nil, err
	}
	if err = st.QueryRow().Scan(&sta.NewUserOrder); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	//查询用户历史总订单.
	history_user_order := `SELECT COUNT(id) AS user_order_total FROM buycai_user_order`
	st, err = agt.ucDbConn.Prepare(history_user_order)
	if err != nil {
		log.Printf("%+v\n", err, history_user_order)
		return nil, err
	}
	if err = st.QueryRow().Scan(&sta.UserOrderTotal); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	//查询历史总购彩订单.
	history_buycai_order := `SELECT COUNT(id) AS buycai_order_total FROM buycai_vendor_order`
	st, err = agt.ucDbConn.Prepare(history_buycai_order)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	if err = st.QueryRow().Scan(&sta.BuycaiOrderTotal); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return sta, nil
}
