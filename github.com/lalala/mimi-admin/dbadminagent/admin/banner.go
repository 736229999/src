package admin

import (
	"database/sql"
	"fmt"
	dbproto "github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"golang.org/x/net/context"
	"log"
	"time"
)

const (
	TABLE_BANNER = "banner"
)

// CreateBanner 创建一条Banner信息
func (agt *DbAdminAgent) CreateBanner(ctx context.Context, arg *dbproto.Banner) (*dbproto.IntValue, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var rid int64
	esql := fmt.Sprintf("INSERT INTO %s(url, target_link, is_visible, description, created, sort, target_id, target_type, location, updated) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id", TABLE_BANNER)
	err := agt.optionsDbConn.QueryRow(esql, arg.GetUrl(), arg.GetTargetLink(), arg.GetIsVisible(), arg.GetDescription(), time.Now().Unix(), arg.GetSort(), arg.GetTargetId(), arg.GetTargetType(), arg.GetLocation(), time.Now().Unix()).Scan(&rid)
	if err != nil {
		log.Printf("error %v, arg: %+v\n", err, arg)
		log.Printf("esql is %v", esql)
		return nil, err
	}
	return &dbproto.IntValue{Value: rid}, nil
}

// QueryBannerList 后台查询Banner列表
func (agt *DbAdminAgent) QueryBannerList(ctx context.Context, arg *dbproto.QueryBannerArg) (*dbproto.BannerList, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		start                     int32
		total                     int64
		esql, where, order, pager string
	)
	if arg.GetPage() > 1 {
		start = (arg.GetPage() - 1) * arg.PageSize
	} else {
		start = 0
	}
	pager = fmt.Sprintf(" LIMIT %d OFFSET %d", arg.PageSize, start)
	// where = " WHERE is_visible = true"
	order = " ORDER BY is_visible DESC, sort DESC, created DESC"
	esql = fmt.Sprintf("SELECT id, url, target_type, target_link, is_visible, description, created, sort, target_id, location, updated FROM %s", TABLE_BANNER)
	// if arg.GetLocation() != 0 {
	where += fmt.Sprintf(" WHERE location = %d", int(arg.GetLocation()))
	// }
	rows, err := agt.optionsDbConn.Query(esql + where + order + pager)
	log.Println("SQL", esql+where+order+pager)
	if err != nil {
		log.Println(err, "SQL: ", esql+where+order)
		return nil, err
	}

	bannerList := make([]*dbproto.Banner, 0)

	for rows.Next() {
		banner := &dbproto.Banner{}
		err := rows.Scan(&banner.Id,
			&banner.Url,
			&banner.TargetType,
			&banner.TargetLink,
			&banner.IsVisible,
			&banner.Description,
			&banner.Created,
			&banner.Sort,
			&banner.TargetId,
			&banner.Location,
			&banner.Updated,
		)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		bannerList = append(bannerList, banner)
	}
	//todo
	totalEsql := fmt.Sprintf("SELECT count(*) FROM %s", TABLE_BANNER)
	row := agt.optionsDbConn.QueryRow(totalEsql + where)
	row.Scan(&total)
	log.Println("esql: ", totalEsql+where)
	// log.Printf("bannerList is %+v\n", bannerList)
	// esql := fmt.Sprintf("")
	log.Printf("total is %v\n", total)

	return &dbproto.BannerList{
		List:  bannerList,
		Total: total,
	}, nil
}

// QueryBannerById 查询一条Banner信息
func (agt *DbAdminAgent) QueryBannerById(ctx context.Context, arg *dbproto.BannerId) (*dbproto.Banner, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		esql string
	)
	esql = fmt.Sprintf("SELECT * FROM %s WHERE id = $1", TABLE_BANNER)

	log.Println("esql", esql, "id ", arg.GetId())
	st, err := agt.optionsDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	banner := &dbproto.Banner{}
	err = st.QueryRow(arg.GetId()).Scan(&banner.Id, &banner.Url, &banner.TargetType, &banner.TargetLink, &banner.IsVisible, &banner.Description, &banner.Created, &banner.Sort, &banner.TargetId, &banner.Location,
		&banner.Updated)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Banner不存在:%+v, %v\n", err, arg.GetId())
			return nil, err
		}
		log.Printf("%+v\n", err)
		return nil, err
	}
	log.Println("banner", banner)
	return banner, nil
}

// UpdateBanner 更新Banner信息
func (agt *DbAdminAgent) UpdateBanner(ctx context.Context, arg *dbproto.Banner) (*dbproto.IntValue, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()

	esql := fmt.Sprintf("UPDATE %s "+
		"SET url = $1, target_type = $2, target_link = $3, is_visible = $4, description = $5, sort = $6, target_id = $7, location=$8, updated=$9 WHERE id = $10", TABLE_BANNER)
	st, err := agt.optionsDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	log.Println("SQL", esql)
	res, err := st.Exec(arg.GetUrl(), arg.GetTargetType(), arg.GetTargetLink(), arg.GetIsVisible(), arg.GetDescription(), arg.GetSort(), arg.GetTargetId(), arg.GetLocation(), time.Now().Unix(), arg.GetId())
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		log.Printf("RowsAffected %v\n", err)
		return nil, err
	}
	return &dbproto.IntValue{
		Value: affect,
	}, nil
}

// QueryPreBanner 获取预览的Banner
func (agt *DbAdminAgent) QueryPreBanner(ctx context.Context, arg *dbproto.QueryBannerArg) (*dbproto.BannerList, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
	var (
		total              int64
		esql, where, order string
	)
	where = " WHERE is_visible = true"
	order = " ORDER BY sort DESC, created DESC"
	esql = fmt.Sprintf("SELECT id, url, target_type, target_link, description, target_id, location FROM %s", TABLE_BANNER)
	// if arg.GetLocation() != 0 {
	where += fmt.Sprintf(" AND location = %d", int(arg.GetLocation()))
	// }
	rows, err := agt.optionsDbConn.Query(esql + where + order)
	if err != nil {
		log.Println(err, "SQL: ", esql+where+order)
		return nil, err
	}

	bannerList := make([]*dbproto.Banner, 0)

	for rows.Next() {
		banner := &dbproto.Banner{}
		err := rows.Scan(&banner.Id,
			&banner.Url,
			&banner.TargetType,
			&banner.TargetLink,
			&banner.Description,
			&banner.TargetId,
			&banner.Location)
		if err != nil {
			log.Println(err, esql)
			return nil, err
		}
		bannerList = append(bannerList, banner)
	}
	//todo
	totalEsql := fmt.Sprintf("SELECT count(*) FROM %s", TABLE_BANNER)
	row := agt.optionsDbConn.QueryRow(totalEsql + where)
	row.Scan(&total)
	log.Println("esql: ", esql+where+order)
	log.Printf("bannerList is %+v\n", bannerList)
	// esql := fmt.Sprintf("")
	log.Printf("total is %v\n", total)

	return &dbproto.BannerList{
		List:  bannerList,
		Total: total,
	}, nil
}
