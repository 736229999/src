package admin

import (
	"golang.org/x/net/context"
	dbproto "github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"fmt"
	"log"
	"database/sql"
	"time"
)
const TABLE_ACTIVITY string = "act_activity"
const TABLE_ACTIVITY_TASK string = "act_activity_task"
const TABLE_GIFT_TEMPLATE string = "gift_template"
const TABLE_TASK string = "act_task"

//添加任务
func (agt *DbAdminAgent) InsertTask(ctx context.Context, arg *dbproto.Task) (*dbproto.Nil, error) {
	DeferFunc()
	st,err := agt.ucDbConn.Prepare(`INSERT INTO act_task(name,des,addtime,is_finish,type,money) VALUES($1,$2,$3,$4,$5,$6)`)
	if err != nil {
		log.Println(err)
		return &dbproto.Nil{},err
	}
	_,err = st.Exec(arg.GetName(),arg.GetDes(),time.Now().Unix(),0,arg.GetType(),arg.GetMoney())
	if err != nil {
		log.Println(err)
		return &dbproto.Nil{},err
	}
	return &dbproto.Nil{},nil
}
//查询所有任务
func (agt *DbAdminAgent) QueryAllTaskList(ctx context.Context, arg *dbproto.Nil) (*dbproto.TaskList, error) {
	DeferFunc()
	st,err := agt.ucDbConn.Prepare(`SELECT id,name,des,addtime,is_finish FROM act_task`)
	if err != nil {
		log.Println(err)
		return &dbproto.TaskList{},err
	}
	rows ,err := st.Query()
	if err != nil {
		log.Println(err)
		return &dbproto.TaskList{},err
	}
	taskList := &dbproto.TaskList{}
	for rows.Next() {
		task := &dbproto.Task{}
		if err = rows.Scan(&task.Id, &task.Name, &task.Des, &task.Addtime,&task.IsFinish); err != nil {
			if err == sql.ErrNoRows {
				log.Println(err)
				return &dbproto.TaskList{}, err
			}
			log.Printf("%+v\n", err)
			return nil, err
		}
		taskList.List = append(taskList.List,task)
	}
	return taskList ,nil
}

//查询任务列表
func (agt *DbAdminAgent) QueryTaskList(ctx context.Context, arg *dbproto.TaskReplyList) (*dbproto.TaskReplyList, error) {
	DeferFunc()
	offset := (arg.GetPage() - 1) * arg.GetSize()
	esql := fmt.Sprintf(`SELECT * FROM act_task ORDER BY id DESC OFFSET %d LIMIT %d`,offset,arg.GetSize())
	totalSql := `SELECT COUNT(*) FROM act_task WHERE 1=1`
	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}

	rows, err := st.Query()
	if err != nil {
		log.Println(err)
		return nil,err
	}
	taskReplyList := &dbproto.TaskReplyList{}
	for rows.Next() {
		task := &dbproto.Task{}
		if err = rows.Scan(&task.Id, &task.Name, &task.Des, &task.Addtime, &task.IsFinish,&task.Type,&task.Money); err != nil {
			if err == sql.ErrNoRows {
				log.Printf("%+v\n", err, esql)
				return &dbproto.TaskReplyList{}, err
			}
			log.Printf("%+v\n", err)
			return nil, err

		}
		taskReplyList.List = append(taskReplyList.List,task)
	}

	st, err = agt.ucDbConn.Prepare(totalSql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	if err = st.QueryRow().Scan(&taskReplyList.Total); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	};
	return taskReplyList,nil

}
//查询所有任务类型
func (agt *DbAdminAgent) QueryAllTaskType(ctx context.Context, arg *dbproto.Nil) (*dbproto.TaskTypeList, error) {
	DeferFunc()
	typeList := &dbproto.TaskTypeList{}
	sql := fmt.Sprintf(`SELECT type FROM act_task `)
	st,err := agt.ucDbConn.Prepare(sql)
	if err != nil{
		log.Println(err)
		return typeList ,err
	}
	rows,err := st.Query()
	if err != nil{
		log.Println(err)
		return typeList ,err
	}
	for rows.Next() {
		var str string = ""
		if err := rows.Scan(&str);err !=nil{
			log.Println(err)
			return typeList,err
		}
		typeList.TypeList = append(typeList.TypeList,str)
	}
	return typeList,nil
}

//删除任务
func (agt *DbAdminAgent) DeleteTask(ctx context.Context, arg *dbproto.IntValue) (*dbproto.Nil, error) {
	DeferFunc()
	sql := fmt.Sprintf(`DELETE FROM act_task WHERE id=$1`)
	st,err := agt.ucDbConn.Prepare(sql)
	if err != nil{
		log.Println(err)
		return &dbproto.Nil{} ,err
	}
	_,err = st.Exec(arg.GetValue())
	if err != nil{
		log.Println(err)
		return &dbproto.Nil{} ,err
	}
	return &dbproto.Nil{},nil

}
//获取单个任务详情
func (agt *DbAdminAgent) QueryTaskById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.Task, error) {
	sql := fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`,TABLE_TASK)
	st,err := agt.ucDbConn.Prepare(sql)
	if err != nil{
		log.Println(err)
		return &dbproto.Task{} ,err
	}
	row  := st.QueryRow(arg.GetValue())
	task := &dbproto.Task{}
	if err = row.Scan(&task.Id,&task.Name,&task.Des,&task.Addtime,&task.IsFinish,&task.Type,&task.Money);err != nil{
		log.Println(err)
		return &dbproto.Task{} ,err
	}
	return task,nil
}
//更新任务
func (agt *DbAdminAgent) UpdateTask(ctx context.Context, arg *dbproto.Task) (*dbproto.Nil, error) {
	DeferFunc()
	sql := fmt.Sprintf(`UPDATE %s SET name = $1, des = $2 WHERE id = $3`,TABLE_TASK)
	st,err := agt.ucDbConn.Prepare(sql)
	if err != nil{
		log.Println(err)
		return &dbproto.Nil{} ,err
	}
	_,err = st.Exec(arg.GetName(),arg.GetDes(),arg.GetId())
	if err != nil{
		log.Println(err)
		return &dbproto.Nil{} ,err
	}
	return &dbproto.Nil{},nil
}
//添加活动
func (agt *DbAdminAgent) InsertActivity(ctx context.Context, arg *dbproto.Activity) (*dbproto.Nil, error) {
	DeferFunc()
	tx ,err := agt.ucDbConn.Begin()
	if err != nil {
		log.Println(err)
		return &dbproto.Nil{},err

	}
	sql1 := fmt.Sprintf(`INSERT INTO %s(title,des,logo,num,starttime,endtime,package_id,create_time,left_num) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`,TABLE_ACTIVITY)
	var lastId int64
	err = tx.QueryRow(sql1,arg.GetTitle(),arg.GetDes(),arg.GetLogo(),arg.GetNum(),arg.GetStarttime(),arg.GetEndtime(),arg.GetPackageId(),time.Now().Unix(),arg.GetNum()).Scan(&lastId)
	if err != nil {
		log.Println(err)
		return &dbproto.Nil{},err
	}
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return &dbproto.Nil{}, err
	}
	//现在再往中间表添加
	if len(arg.TaskLists) == 0 {
		//没有给活动添加任务，这样是允许的
		return &dbproto.Nil{}, tx.Commit()
	}
	for _,v := range arg.TaskLists {
		sql2 := fmt.Sprintf(`INSERT INTO %s(act_id,task_id,addtime) VALUES($1,$2,$3)`,TABLE_ACTIVITY_TASK)
		st2 , err := tx.Prepare(sql2)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return &dbproto.Nil{}, err
		}
		res2 ,err := st2.Exec(lastId,v,time.Now().Unix())
		if res2 == nil {
			tx.Rollback()
			return &dbproto.Nil{},err
		}
	}

	return &dbproto.Nil{}, tx.Commit()
}

//查询所有活动
func (agt *DbAdminAgent) QueryActivityList(ctx context.Context, arg *dbproto.ActivityReplyList) (*dbproto.ActivityReplyList, error) {
	DeferFunc()
	//思路：用分页查询，除了活动表，还需要返回对应的礼包名称和操作人员名称（操作人员名称暂时不查，应该是存的session还没看）
	esql := fmt.Sprintf(`SELECT c.id,c.title,c.des,c.logo,c.num,c.starttime,c.endtime,c.create_time,g.title FROM act_activity AS c LEFT JOIN gift_template AS g ON c.package_id = g.id `)

	totalSql := `SELECT COUNT(*) FROM act_activity AS c LEFT JOIN gift_package AS g ON  c.package_id = g.id `

	offset := (arg.GetPage() - 1) * arg.GetSize()
	esql += fmt.Sprintf(` ORDER BY c.id DESC OFFSET %d LIMIT %d`, offset, arg.GetSize())

	st, err := agt.ucDbConn.Prepare(esql)
	if err != nil {
		log.Printf("%+v\n", err, esql)
		return nil, err
	}

	rows, err := st.Query()
	if err != nil {
		log.Println(err)
		return nil,err
	}

	activityReplyList := &dbproto.ActivityReplyList{}
	for rows.Next() {
		activityReply := &dbproto.ActivityReply{}

		if err = rows.Scan(&activityReply.Id, &activityReply.Title, &activityReply.Des, &activityReply.Logo, &activityReply.Num, &activityReply.Starttime,&activityReply.Endtime,&activityReply.CreateTime,&activityReply.PackageName); err != nil {
			if err == sql.ErrNoRows {
				log.Printf("%+v\n", err, esql)
				return &dbproto.ActivityReplyList{}, err
			}
			log.Printf("%+v\n", err)
			return nil, err

		}
		activityReply.CreateAdmin = "admin1" //这儿暂时先没查管理员姓名
		activityReplyList.List = append(activityReplyList.List, activityReply)
	}
	st, err = agt.ucDbConn.Prepare(totalSql)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}
	if err = st.QueryRow().Scan(&activityReplyList.Total); err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	};
	return activityReplyList, nil
}


//查询所有礼包模板
func (agt *DbAdminAgent) QueryAllGiftTemplateList(ctx context.Context, arg *dbproto.Nil) (*dbproto.TemplateList, error) {
	DeferFunc()
	st,err := agt.ucDbConn.Prepare(`SELECT id,title FROM gift_template`)
	if err != nil {
		log.Println(err)
		return &dbproto.TemplateList{},err
	}
	rows ,err := st.Query()
	if err != nil {
		log.Println(err)
		return &dbproto.TemplateList{},err
	}
	giftTemplateList := &dbproto.TemplateList{}
	for rows.Next() {
		giftTemplate := &dbproto.GiftTemplate{}
		if err = rows.Scan(&giftTemplate.Id, &giftTemplate.Title); err != nil {
			if err == sql.ErrNoRows {
				log.Println(err)
				return &dbproto.TemplateList{}, err
			}
			log.Printf("%+v\n", err)
			return nil, err
		}
		giftTemplateList.List = append(giftTemplateList.List,giftTemplate)
	}
	return giftTemplateList ,nil
}


//删除活动
func (agt *DbAdminAgent) DeleteActivity(ctx context.Context, arg *dbproto.IntValue) (*dbproto.Nil, error) {
	DeferFunc()
	sql := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`,TABLE_ACTIVITY)
	st,err := agt.ucDbConn.Prepare(sql)
	if err != nil{
		log.Println(err)
		return &dbproto.Nil{} ,err
	}
	_,err = st.Exec(arg.GetValue())
	if err != nil{
		log.Println(err)
		return &dbproto.Nil{} ,err
	}
	return &dbproto.Nil{},nil
}

//获取单个活动详情
func (agt *DbAdminAgent) QueryActivityById(ctx context.Context, arg *dbproto.IntValue) (*dbproto.Activity, error) {
	DeferFunc()
	sql := fmt.Sprintf(`SELECT c.id,c.title,c.des,c.logo,c.num,c.starttime,c.endtime,g.id,g.title FROM %s AS c LEFT JOIN gift_template AS g ON c.package_id = g.id where c.id=$1`,TABLE_ACTIVITY)
	st,err := agt.ucDbConn.Prepare(sql)
	activity := &dbproto.Activity{}
	if err != nil{
		log.Println(err)
		return activity ,err
	}
	row := st.QueryRow(arg.GetValue())

	err = row.Scan(&activity.Id, &activity.Title, &activity.Des, &activity.Logo, &activity.Num, &activity.Starttime,&activity.Endtime,&activity.PackageId,&activity.PackageName)
	if err != nil {
		log.Println(err)
		return activity ,err
	}
	st2,err := agt.ucDbConn.Prepare(`SELECT task_id FROM act_activity_task WHERE act_id=$1`)
	if err != nil {
		log.Println(err)
		return activity ,err
	}
	rows,err := st2.Query(arg.GetValue())
	if err != nil {
		log.Println(err)
		return activity ,err
	}
	for rows.Next() {
		var id int64
		if err = rows.Scan(&id);err != nil{
			log.Println(err)
			return activity ,err
		}
		activity.TaskLists = append(activity.TaskLists,id)
	}
	return activity,nil
}
//更新活动
func (agt *DbAdminAgent) UpdateActivity(ctx context.Context, arg *dbproto.Activity) (*dbproto.Nil, error) {
	DeferFunc()
	//要更新act_activity和act_activity_task
	tx ,err := agt.ucDbConn.Begin()
	sql1 := fmt.Sprintf(`UPDATE %s SET title = $1, des = $2, num = $3, package_id = $4,starttime=$5,endtime=$6,logo=$7 WHERE id = $8 `,TABLE_ACTIVITY)
	st,err := tx.Prepare(sql1)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return &dbproto.Nil{} ,err
	}
	_,err = st.Exec(arg.GetTitle(),arg.GetDes(),arg.GetNum(),arg.GetPackageId(),arg.GetStarttime(),arg.GetEndtime(),arg.GetLogo(),arg.GetId())
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return &dbproto.Nil{} ,err
	}
	//更新中间表，先删后加
	sql2 := fmt.Sprintf(`DELETE FROM %s WHERE act_id=$1`,TABLE_ACTIVITY_TASK)
	st2,err := tx.Prepare(sql2)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return &dbproto.Nil{} ,err
	}
	_,err = st2.Exec(arg.Id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return &dbproto.Nil{},nil
	}
	for _,v := range arg.TaskLists{
		esql := fmt.Sprintf(`INSERT INTO %s (act_id,task_id,addtime) VALUES($1,$2,$3)`,TABLE_ACTIVITY_TASK)
		st,err := tx.Prepare(esql)
		if err !=nil {
			log.Println(err)
			tx.Rollback()
			return &dbproto.Nil{},nil
		}
		_,err = st.Exec(arg.GetId(),v,time.Now().Unix())
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return &dbproto.Nil{},nil
		}
	}
	return &dbproto.Nil{},tx.Commit()
}


func DeferFunc() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
}
