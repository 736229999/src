package admin

import (
	"database/sql"
	"fmt"
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"golang.org/x/net/context"
	"log"
	"strings"
)

//获取usercenter 中user 的列表.
func (agt *DbAdminAgent) QueryUsercenterList(ctx context.Context, arg *dbproto.UsercenterList) (*dbproto.UsercenterList, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := `SELECT ui.account_id, ui.nickname, ui.sex, ui.invitation_code, ph.phone, vf.credits, vf.kxd, ih.id AS is_invitee  FROM userinfo AS ui LEFT JOIN phone_user AS ph ON ui.account_id = ph.account_id LEFT JOIN virtual_fund as vf ON ui.account_id = vf.account_id LEFT JOIN invite_history AS ih ON ui.account_id = ih.invitee`
	esql += ` WHERE 1 = 1 `

	totalSql := strings.Replace(esql, "ui.account_id, ui.nickname, ui.sex, ui.invitation_code, ph.phone, vf.credits, vf.kxd, ih.id AS is_invitee", "COUNT(*) AS num", -1)
	if len(arg.GetNickname()) > 0 {
		esql += fmt.Sprintf(` AND ui.nickname LIKE '%s' `, "%"+arg.GetNickname()+"%")
	}

	if len(arg.GetPhone()) > 0 {
		esql += fmt.Sprintf(` AND ph.phone LIKE '%s'`, "%"+arg.GetPhone()+"%")
	}

	offset := (arg.GetPage() - 1) * arg.GetSize()
	esql += fmt.Sprintf(` ORDER BY ui.account_id DESC offset %d LIMIT %d`, offset, arg.GetSize())

	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	rows, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	ucList := &dbproto.UsercenterList{}
	for rows.Next() {
		uc := &dbproto.Userinfo{}
		var nickname, invitation_code, phone sql.NullString
		var sex, credits, kxd sql.NullInt64
		var is_invitee sql.NullBool

		if err = rows.Scan(&uc.Id, &nickname, &sex, &invitation_code, &phone, &credits, &kxd, &is_invitee); err != nil {
			if err == sql.ErrNoRows {
				return &dbproto.UsercenterList{}, nil
			}
			log.Printf("%+v\n", err)
			return nil, err
		}

		if nickname.Valid {
			uc.Nickname = nickname.String
		}

		if invitation_code.Valid {
			uc.InvitationCode = invitation_code.String
		}

		if phone.Valid {
			uc.Phone = phone.String
		}

		if credits.Valid {
			uc.Credits = int32(credits.Int64)
		}

		if kxd.Valid {
			uc.Kxd = int32(kxd.Int64)
		}

		if is_invitee.Valid {
			uc.IsInvited = is_invitee.Bool
		}

		ucList.List = append(ucList.List, uc)
	}

	st, err = agt.ucDbConn.Prepare(totalSql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	if err = st.QueryRow().Scan(&ucList.Total); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return ucList, nil
}

//获取usercenter的详细用户数据.
func (agt *DbAdminAgent) QueryUsercenterDetail(ctx context.Context, arg *dbproto.IntValue) (*dbproto.Userinfo, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	fieldList := []string{
		"ui.account_id", "ui.nickname", "ui.sex", "ui.invitation_code", "ph.phone",
		"vf.credits", "vf.kxd", "ih.id AS is_invitee", "qq.id AS qq_id", "wx.id AS wx_id",
	}
	esql := fmt.Sprintf(`SELECT %s FROM userinfo AS ui `, strings.Join(fieldList, ", "))
	esql += ` LEFT JOIN phone_user AS ph ON ui.account_id = ph.account_id `
	esql += ` LEFT JOIN virtual_fund as vf ON ui.account_id = vf.account_id `
	esql += ` LEFT JOIN invite_history AS ih ON ui.account_id = ih.invitee `
	esql += ` LEFT JOIN qq_user AS qq ON ui.account_id = qq.account_id `
	esql += ` LEFT JOIN weixin_user AS wx ON ui.account_id = wx.account_id `
	esql += fmt.Sprintf(` WHERE ui.account_id = %d`, arg.GetValue())

	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	userInfo := &dbproto.Userinfo{}
	var invitation_code, phone sql.NullString
	var credits, kxd sql.NullInt64
	var is_invitee, qq_id, wx_id sql.NullBool

	values := []interface{}{
		&userInfo.Id, &userInfo.Nickname, &userInfo.Sex, &invitation_code,
		&phone, &credits, &kxd, &is_invitee, &qq_id, &wx_id,
	}

	if err = st.QueryRow().Scan(values...); err != nil {
		if err == sql.ErrNoRows {
			return &dbproto.Userinfo{}, nil
		}
		log.Printf("%+v\n", err)
		return nil, err
	}

	if invitation_code.Valid {
		userInfo.InvitationCode = invitation_code.String
	}

	if phone.Valid {
		userInfo.Phone = phone.String
		userInfo.BindPhone = true
	}

	if credits.Valid {
		userInfo.Credits = int32(credits.Int64)
	}

	if kxd.Valid {
		userInfo.Kxd = int32(kxd.Int64)
	}

	if is_invitee.Valid {
		userInfo.IsInvited = is_invitee.Bool
	}

	if qq_id.Valid {
		userInfo.BindQq = qq_id.Bool
	}

	if wx_id.Valid {
		userInfo.BindWechat = wx_id.Bool
	}

	return userInfo, nil
}

//根据userid 获取用户的金额统计.
func (agt *DbAdminAgent) QueryUsercenterFundById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.FundHistory, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	fieldList := []string{
		"SUM(r.money) AS recharge",
		fmt.Sprintf("(SELECT SUM(bvo.money) FROM buycai_vendor_order AS bvo WHERE bvo.account_id = %d) AS buycai", arg.GetValue()),
		fmt.Sprintf("(SELECT SUM(bvo.win_money) FROM buycai_vendor_order AS bvo WHERE bvo.account_id = %d) AS winning", arg.GetValue()),
		fmt.Sprintf("(SELECT SUM(w.amount) FROM withdraw AS w WHERE w.account_id = %d and status = %d) AS withdraw", arg.GetValue(), 1), //1为提现申请通过.
	}
	esql := fmt.Sprintf(`SELECT %s FROM recharge_history AS r WHERE r.account_id = %d `, strings.Join(fieldList, ","), arg.GetValue())

	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	var recharge, buycai, winning, withdraw sql.NullFloat64
	value := []interface{}{
		&recharge, &buycai, &winning, &withdraw,
	}
	if err = st.QueryRow().Scan(value...); err != nil {
		if err == sql.ErrNoRows {
			return &dbproto.FundHistory{}, nil
		}
		log.Printf("%+v\n", err)
		return nil, err
	}

	fund := &dbproto.FundHistory{}

	if recharge.Valid {
		fund.Recharge = fmt.Sprintf("%.2f", recharge.Float64)
	} else {
		fund.Recharge = "00.00"
	}

	if buycai.Valid {
		fund.Buycai = fmt.Sprintf("%.2f", buycai.Float64)
	} else {
		fund.Buycai = "00.00"
	}

	if withdraw.Valid {
		fund.Withdraw = fmt.Sprintf("%.2f", withdraw.Float64)
	} else {
		fund.Withdraw = "00.00"
	}

	if winning.Valid {
		fund.Winning = fmt.Sprintf("%.2f", winning.Float64)
	} else {
		fund.Winning = "00.00"
	}

	return fund, nil
}

//获取用户的充值列表数据.
func (agt *DbAdminAgent) QueryUsercenterRechargeById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.FundHistory, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string{
		"id", "account_id", "money", "recharge_time", "order_id", "source",
	}

	esql := fmt.Sprintf(`SELECT %s FROM recharge_history WHERE account_id = %d`, strings.Join(filedList, ", "), arg.GetValue())
	log.Println(esql)
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	rows, err := st.Query()
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	fund := &dbproto.FundHistory{}
	for rows.Next() {
		recharge := &dbproto.RechargeHistory{}
		var money float64
		value := []interface{}{
			&recharge.Id, &recharge.AccountId, &money, &recharge.RechargeTime, &recharge.OrderId, &recharge.Source,
		}
		if err = rows.Scan(value...); err != nil {
			if err == sql.ErrNoRows {
				return &dbproto.FundHistory{}, nil
			}
			log.Printf("%+v\n", err)
			return nil, err
		}

		recharge.Money = fmt.Sprintf("%.2f", money)

		fund.RechargeList = append(fund.RechargeList, recharge)
	}

	return fund, nil
}

//获取今日新增用户和用户总数.
func (agt *DbAdminAgent) QueryUserStatisticsNum(ctx context.Context, arg *dbproto.Nil) (*dbproto.UserStatisticsNum, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	filedList := []string{
		"COUNT(account_id) AS new_user_num",
		"(SELECT COUNT(account_id) FROM userinfo) AS total_user_num",
	}
	esql := fmt.Sprintf("SELECT %s  FROM userinfo GROUP BY to_char(to_timestamp(create_time),'yyyy-MM-DD')", strings.Join(filedList, ","))
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	defer st.Close()

	userStatisticsNum := &dbproto.UserStatisticsNum{}
	if err = st.QueryRow().Scan(&userStatisticsNum.NewUserNum, &userStatisticsNum.TotalUserNum); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	return userStatisticsNum, nil
}
