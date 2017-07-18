package admin
							/***********订单管理***********/
import (
	"golang.org/x/net/context"
	dbproto "github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"fmt"
	"log"
	"database/sql"
)




//查询用户订单列表
func (agt *DbAdminAgent) QueryUserOrderList(ctx context.Context, arg *dbproto.UserOrderList) (*dbproto.UserOrderList, error) {
	DeferFunc()
	if arg.GetPage() < 1 {
		arg.Page = 1
	}
	offset := (arg.GetPage() - 1) * arg.GetSize()
	esql := fmt.Sprintf(`SELECT c.*,g.nickname FROM buycai_user_order
	 AS c LEFT JOIN userinfo AS g ON c.account_id = g.account_id
	 ORDER BY id DESC OFFSET %d LIMIT %d`,offset,arg.GetSize())

	totalSql := `SELECT COUNT(*) FROM buycai_user_order WHERE 1=1`
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return arg, err
	}

	rows, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err)
		return arg,err
	}
	userOrderList := &dbproto.UserOrderList{}
	for rows.Next() {
		userOrder := &dbproto.BuycaiUserOrder{}
		if err = rows.Scan(&userOrder.Id, &userOrder.AccountId, &userOrder.LotteryId,
			&userOrder.IssueNum, &userOrder.ChaseNo,&userOrder.Cai,&userOrder.Balance,
			&userOrder.SumMoney,&userOrder.Issues,&userOrder.SchemeList,&userOrder.OrderTime,
			&userOrder.TicketSubMoney,&userOrder.Status,&userOrder.IsWinStop,&userOrder.CostCai,
			&userOrder.CostBalance,&userOrder.UserName); err != nil {
			if err == sql.ErrNoRows {
				log.Printf("%+v\n", err)
				return arg, err
			}
			log.Printf("%+v\n", err)
			return arg, err

		}
		//查询对应的购彩券信息
		flag := 0
		st,err = agt.ucDbConn.Prepare(`SELECT COUNT(*) FROM ticket WHERE order_id=$1`)
		if err != nil {
			log.Printf("%+v\n", err)
			return arg,err
		}
		if err := st.QueryRow(userOrder.GetId()).Scan(&flag);err != nil{
			log.Printf("%+v\n", err)
			return arg,err
		}
		if flag == 0 {
			//说明没有使用购彩券

		}else {
			st,err := agt.ucDbConn.Prepare(`SELECT use_base,use_sub,valid_start,valid_end,restrict_id,restrict_type FROM ticket WHERE order_id=$1`)
			if err != nil {
				log.Printf("%+v\n", err)
				return arg,err
			}
			userOrder.Ticket = &dbproto.Ticket{}
			if err = st.QueryRow(userOrder.GetId()).Scan(&userOrder.Ticket.UseBase,&userOrder.Ticket.UseSub,
				&userOrder.Ticket.ValidStart,&userOrder.Ticket.ValidEnd,
				&userOrder.Ticket.RestrictId,&userOrder.Ticket.RestrictType); err != nil {
				//说明没使用购彩券，程序继续执行
				log.Printf("%+v\n", err)
			}
		}

		userOrderList.List = append(userOrderList.List,userOrder)
	}

	st, err = agt.ucDbConn.Prepare(totalSql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	if err = st.QueryRow().Scan(&userOrderList.Total); err != nil {
		log.Printf("%+v\n", err)
		return arg, err
	}
	return userOrderList,nil

}

//带条件查询用户列表
func (agt *DbAdminAgent) QueryUserOrderListWithCondition(ctx context.Context, arg *dbproto.UserOrderList) (*dbproto.UserOrderList, error) {
	DeferFunc()
	if arg.GetPage() < 1 {
		arg.Page = 1
	}
	offset := (arg.GetPage() - 1) * arg.GetSize()
	userOrderList := &dbproto.UserOrderList{}
	//只查id 和 nickname
	esql := fmt.Sprintf(`SELECT c.*,g.nickname FROM buycai_user_order
	 AS c LEFT JOIN userinfo AS g ON c.account_id = g.account_id
	 WHERE   g.nickname=$1 OR CAST(c.id as VARCHAR(10))=$2
	 ORDER BY id DESC  OFFSET %d LIMIT %d`,offset,arg.GetSize())
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return userOrderList, err
	}

	rows, err := st.Query(arg.GetCondition(),arg.GetCondition())
	if err != nil {
		log.Printf("%+v\n", err)
		return userOrderList,err
	}

	for rows.Next() {
		userOrder := &dbproto.BuycaiUserOrder{}
		if err = rows.Scan(&userOrder.Id, &userOrder.AccountId, &userOrder.LotteryId,
			&userOrder.IssueNum, &userOrder.ChaseNo,&userOrder.Cai,&userOrder.Balance,
			&userOrder.SumMoney,&userOrder.Issues,&userOrder.SchemeList,&userOrder.OrderTime,
			&userOrder.TicketSubMoney,&userOrder.Status,&userOrder.IsWinStop,&userOrder.CostCai,
			&userOrder.CostBalance,&userOrder.UserName); err != nil {
			if err == sql.ErrNoRows {
				log.Printf("%+v\n", err)
				return userOrderList, err
			}
			log.Printf("%+v\n", err)
			return userOrderList, err

		}
		//查询对应的购彩券信息
		st,err := agt.ucDbConn.Prepare(`SELECT use_base,use_sub,valid_start,valid_end,restrict_id,restrict_type FROM ticket WHERE order_id=$1`)
		if err != nil {
			log.Printf("%+v\n", err)
			return userOrderList,err
		}
		userOrder.Ticket = &dbproto.Ticket{}
		if err = st.QueryRow(userOrder.GetId()).Scan(&userOrder.Ticket.UseBase,&userOrder.Ticket.UseSub,
			&userOrder.Ticket.ValidStart,&userOrder.Ticket.ValidEnd,
			&userOrder.Ticket.RestrictId,&userOrder.Ticket.RestrictType); err != nil {
			//说明没使用购彩券，程序继续执行
			log.Printf("%+v\n", err)
		}
		userOrderList.List = append(userOrderList.List,userOrder)
	}

	totalSql := fmt.Sprintf(`SELECT COUNT(*) FROM buycai_user_order
	 AS c LEFT JOIN userinfo AS g ON c.account_id = g.account_id
	 WHERE   g.nickname=$1 OR CAST(c.id as VARCHAR(10))=$2
	  `)
	st, err = agt.ucDbConn.Prepare(totalSql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	if err = st.QueryRow(arg.GetCondition(),arg.GetCondition()).Scan(&userOrderList.Total); err != nil {
		log.Printf("%+v\n", err)
		return userOrderList, err
	}
	return userOrderList,nil
}

