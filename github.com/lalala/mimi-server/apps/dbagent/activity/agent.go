package activity

import (

	dbproto "github.com/caojunxyz/mimi-server/apps/dbagent/proto"
	"golang.org/x/net/context"
	"log"
	"fmt"
	"database/sql"
	"time"
)

const TABLE_USERINFO = "userinfo"
const TABLE_ACTIVITY  = "act_activity"
const TABLE_ACCOUNT  = "account"
const TABLE_RECHARGE_HISTORY  = "recharge_history"
const TABLE_ACTIVITY_TASK  = "act_activity_task"
const TABLE_TASK  = "act_task"

//所有的活动列表
func (agt *DbActivityAgent) QueryActivityList(ctx context.Context, arg *dbproto.IntValue) (*dbproto.ActivityList, error) {
	DeferFun()
	activityList := &dbproto.ActivityList{}
	sql := fmt.Sprintf(`SELECT id,title,des,logo,num FROM %s`,TABLE_ACTIVITY)
	st,err := agt.dbConn.Prepare(sql)
	if err != nil {
		log.Printf("%+v\n",err)
		return activityList,err
	}
	rows ,err := st.Query()
	if err != nil {
		log.Printf("%+v\n",err)
		return activityList,err
	}
	for rows.Next() {
		activity := &dbproto.Activity{}
		if err := rows.Scan(&activity.Id,&activity.Title,&activity.Des,&activity.Logo,&activity.Num);err != nil{
			log.Println(err)
			return activityList,err
		}
		//判断活动是否已经参加
		st2,err := agt.dbConn.Prepare(`SELECT id FROM act_user_activity WHERE act_id=$1 and account_id=$2`)
		if err != nil {
			log.Printf("%+v\n",err)
			return activityList,err
		}
		id := 0
		flag := false
		err = st2.QueryRow(activity.GetId(),arg.GetValue()).Scan(&id)
		if err !=nil {
			//说明没有参加
			flag = false
		}else {
			flag = true
		}
		activity.HasJoin = flag
		activity.Status = 0	//默认正在进行中
		//再看活动状态
		st3 ,err := agt.dbConn.Prepare(`SELECT left_num ,starttime,endtime FROM act_activity WHERE id=$1`)
		if err != nil {
			log.Printf("%+v\n",err)
			return activityList,err
		}
		var leftNum int
		var starttime int64
		var endtime int64
		if err := st3.QueryRow(activity.GetId()).Scan(&leftNum,&starttime,&endtime);err != nil{
			log.Printf("%+v\n",err)
			return activityList,err
		}
		if leftNum == 0 {
			activity.Status = 2	//已领完
		}
		if starttime > time.Now().Unix() {
			activity.Status = 3	//活动还没开始
		}
		if endtime < time.Now().Unix() {
			activity.Status = 1	//已过期
		}

		activityList.List = append(activityList.List,activity)
	}

	return activityList,nil
}
//某个活动详情
func (agt *DbActivityAgent) ActivityDetail(ctx context.Context, arg *dbproto.ActivityAccount) (*dbproto.ActivtyDetail, error) {
	//思路：先查activity的一些信息，再查对应的任务、是否完成、已经参与了活动的用户
	DeferFun()
	activityDetail := &dbproto.ActivtyDetail{}
	sql1 := fmt.Sprintf(`SELECT id,title,des,logo,num,left_num,starttime,endtime FROM %s WHERE id=$1`,TABLE_ACTIVITY)
	st,err := agt.dbConn.Prepare(sql1)
	if err != nil {
		log.Println(err)
		return activityDetail,err
	}
	activity := &dbproto.Activity{}
	if err := st.QueryRow(arg.GetActivityId()).Scan(&activity.Id,&activity.Title,&activity.Des,&activity.Logo,&activity.Num,&activity.LeftNum,&activity.Starttime,&activity.Endtime);err != nil{
		log.Println(err)
		return activityDetail,err
	}
	activityDetail.Activity = activity	//要返回的活动详情
	sql2 := fmt.Sprintf(`SELECT c.task_id,g.name,g.des,g.type,g.money FROM %s AS c LEFT JOIN %s AS g ON c.task_id = g.id WHERE c.act_id=$1`,TABLE_ACTIVITY_TASK,TABLE_TASK)
	st2 ,err := agt.dbConn.Prepare(sql2)
	rows ,err := st2.Query(arg.GetActivityId())
	if err != nil {
		log.Println(err)
		return activityDetail,err
	}
	for rows.Next() {
		task := &dbproto.Task{}
		if err := rows.Scan(&task.Id,&task.Name,&task.Des,&task.Type,&task.Money);err != nil{
			if err == sql.ErrNoRows{
				log.Println(err)
				return activityDetail,err
			}
		}
		//再看用户是否完成了该任务
		if task.GetType() == "充值" {
			task.IsComplete = agt.HasCompleteTask(arg.GetActivityId(),3,arg.GetAccountId(),task.GetMoney())		//3表示充值的相关任务
		}else {
			task.IsComplete = agt.HasCompleteTask(arg.GetActivityId(),task.GetId(),arg.GetAccountId())
		}


		activityDetail.TaskList = append(activityDetail.TaskList,task)
	}
	//最后再查所有参与了活动的用户
	st3,err := agt.dbConn.Prepare(`SELECT c.account_id,g.nickname,g.icon FROM act_user_activity AS c LEFT JOIN userinfo AS g ON c.account_id = g.account_id WHERE c.act_id=$1 and c.account_id=$2`)
	if err != nil {
		log.Println(err)
		return activityDetail,err
	}
	row3,err := st3.Query(arg.GetActivityId(),arg.GetAccountId())
	if row3.Next() {
		gainer := &dbproto.UserBaseInfo{}
		if err:= row3.Scan(&gainer.AccountId,&gainer.Nickname,&gainer.Icon);err != nil{
			log.Println(err)
			return activityDetail,err
		}
		activityDetail.Gainers = append(activityDetail.Gainers,gainer)
	}
	fmt.Println(activityDetail.TaskList)
	return activityDetail,nil
}



func DeferFun()  {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic:", err)
		}
	}()
}
