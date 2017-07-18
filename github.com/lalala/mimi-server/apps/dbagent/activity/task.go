/*************************用户需要完成的所有任务********************/
package activity

import (
	"log"
	"fmt"
)

//1、是否注册
func (agt *DbActivityAgent) HasRegister(account_id int64) bool{
	DeferFun()
	fmt.Println("进入了注册")
	sql := fmt.Sprintf(`SELECT id FROM %s WHERE id=$1`,TABLE_ACCOUNT)
	st,err := agt.dbConn.Prepare(sql)
	if err != nil {
		log.Println(err)
		return false
	}
	id := 0
	err = st.QueryRow(account_id).Scan(&id)
	if err != nil {
		log.Println("没有该条记录")
		return false
	}
	return true
}
//2、首次充值
func (agt *DbActivityAgent) FirstRecharge(account_id int64) bool  {
	DeferFun()
	//思路：recharg_history里面只有一条该account_id的记录就算是首次充值完成
	sql := fmt.Sprintf(`SELECT count(*) FROM %s WHERE id=$1`,TABLE_RECHARGE_HISTORY)
	st,err := agt.dbConn.Prepare(sql)
	if err != nil {
		log.Println(err)
		return false
	}
	var count int
	if err := st.QueryRow(account_id).Scan(&count);err != nil{
		log.Println(err)
		return false
	}
	if count == 1 {
		 return true
	}
	return false
}
//3、充值达到一定的金额
func (agt *DbActivityAgent) RechargeUpToQuota(activity_id,account_id int64,money float32) bool  {
	//要在活动有效时间内完成，查充值记录，在某个时间段内的总金额如果达到就算完成
	DeferFun()
	st,err := agt.dbConn.Prepare(`SELECT starttime,endtime FROM act_activity WHERE id=$1`)
	if err != nil {
		log.Printf("%+v\n",err)
		return false
	}
	var starttime int64
	var endtime int64
	err = st.QueryRow(activity_id).Scan(&starttime,&endtime)
	if err != nil {
		log.Printf("%+v\n",err)
		return false
	}
	//查询在一段时间内的充值总数
	st1,err := agt.dbConn.Prepare(`SELECT SUM(money)  FROM recharge_history WHERE account_id=$1 and recharge_time>$2 and recharge_time<$3`)
	if err != nil {
		log.Printf("%+v\n",err)
		return false
	}
	var totalMoney float32 = 0.0
	err = st1.QueryRow(account_id,starttime,endtime).Scan(&totalMoney)
	if err != nil {
		log.Printf("%+v\n",err)
		return false
	}
	if totalMoney > money {
		//说明充值达到了一定的金额
		return true
	}
	return false
}
//4、首次购彩
func (agt *DbActivityAgent) FirstBuyCai(account_id int64) bool  {
	DeferFun()
	//思路：购彩记录表里面只有一条该account_id的记录就算是首次充值完成
	sql := fmt.Sprintf(`SELECT count(*) FROM %s WHERE id=$1`,TABLE_RECHARGE_HISTORY)
	st,err := agt.dbConn.Prepare(sql)
	if err != nil {
		log.Println(err)
		return false
	}
	var count int
	if err := st.QueryRow(account_id).Scan(&count);err != nil{
		log.Println(err)
		return false
	}
	if count == 1 {
		return true
	}
	return false
}
//5、购彩是个大任务，里面包含了购买的各个彩种、购买次数、购买的总金额
func (agt *DbActivityAgent) BuyCai(account_id int64) bool  {
	return false
}
//6、分享
func (agt *DbActivityAgent) Share(account_id int64) bool  {
	return false
}


//是否完成了某个任务（以后可以考虑用反射，并且做成不定参数）
func(agt *DbActivityAgent) HasCompleteTask(activity_id,task_id,account_id int64,money ...float32)bool{
	switch task_id {
	case 1:
		return agt.HasRegister(account_id)
	case 2:
		return agt.FirstRecharge(account_id)
	case 3:
		return agt.RechargeUpToQuota(activity_id,account_id,money[0])
	case 4:
		return agt.FirstBuyCai(account_id)
	case 5:
		return agt.BuyCai(account_id)
	case 6:
		return agt.Share(account_id)
	}

	return false
}