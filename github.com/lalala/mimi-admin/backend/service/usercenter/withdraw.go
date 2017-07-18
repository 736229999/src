package usercenter

import (
	"github.com/caojunxyz/mimi-admin/dbadminagent/proto"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"strconv"
	"github.com/caojunxyz/mimi-admin/backend/core"
)

// HandleWithdrawApplyList 获取提现申请列表
func (srv *UsercenterService) HandleWithdrawApplyList(c *gin.Context) {
	log.Println("Handle WithdrawApply List ")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pagesize", "10")
	start := c.DefaultQuery("start", "0")
	end := c.DefaultQuery("end", "0")

	pageDb, err := strconv.Atoi(page)
	if err != nil {
		log.Println("error", err)
		srv.Json("页码参数错误", http.StatusForbidden, c)
		return
	}
	pageSizeDb, err := strconv.Atoi(pageSize)
	if err != nil {
		log.Println("error", err)
		srv.Json("每页尺寸参数错误", http.StatusForbidden, c)
		return
	}

	startDb, err := strconv.ParseInt(start, 10, 0)
	if err != nil {
		log.Println("error", err)
		srv.Json("开始时间参数错误", http.StatusForbidden, c)
		return
	}
	endDb, err := strconv.ParseInt(end, 10, 0)
	if err != nil {
		log.Println("error", err)
		srv.Json("结束时间参数错误", http.StatusForbidden, c)
		return
	}

	isSuccessStr := c.DefaultQuery("is_success", "false") 
	isSuccess, err := strconv.ParseBool(isSuccessStr)
	if err != nil {
		log.Println("error", err)
		srv.Json("参数错误", http.StatusForbidden, c)
		return
	}

	stepStr := c.DefaultQuery("step", strconv.Itoa(int(dbproto.WithdrawApply_STEP_ALL)))
	step, err := strconv.Atoi(stepStr)
	if err != nil {
		log.Println("statusStr", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}

	msg := dbproto.QueryWithdrawApplyArg{
		Step:   dbproto.WithdrawApply_WAStepBackend(step),
		IsSuccess: isSuccess,
		Realname: c.Query("realname"),
		Phone:    c.Query("phone"),
		Start:    startDb,
		End:      endDb,
		Page:     int32(pageDb),
		PageSize: int32(pageSizeDb),
		InNo:     c.Query("inNo"),
		// Auditor: c.Query("")
	}

	res, err := srv.Db().QueryWithdrawApplyList(context.Background(), &msg)
	if err != nil {
		log.Println("QueryWithdrawApplyList error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	auditor, ok := srv.GetUserInfo(c, "username").(string)
	if !ok {
		// log.Println("GetUserInfo error")
		// msg.Auditor = "测试"
		srv.Json("禁止操作", http.StatusInternalServerError, c)
		return
	}
	for _, wa := range res.List {
		if wa.Auditor == "" || wa.Auditor == auditor {
			wa.CanOperate = true
		}
	}
	log.Printf("res is %+v", res)

	// srv.Log(c, msg, "查询提现申请列表", core.QUERY_OPERATION)
	srv.Json(res, http.StatusOK, c)
}

// HandleGetWithdrawApplyById 获取某个提现申请
func (srv *UsercenterService) HandleGetWithdrawApplyById(c *gin.Context) {
	log.Println("Handle Get WithdrawApply")
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	// var ok bool
	auditor, ok := srv.GetUserInfo(c, "username").(string)
	if !ok {
		log.Println("GetUserInfo error")
		auditor = "测试"
		// srv.Json("禁止操作", http.StatusInternalServerError, c)
		// return
	}

	checkRes, err := srv.Db().CheckWithdrawApply(context.Background(), &dbproto.CheckWAArg{WithdrawApplyId: id, Auditor: auditor})
	// 只能认领后才能直接获取该提现申请详情，审核人不是自己都无法查看
	if (checkRes.GetExistAuditor() && !checkRes.GetIsOwn()) || !checkRes.GetExistAuditor() {
		srv.Json("无权查看", http.StatusForbidden, c)
		return
	}

	res, err := srv.Db().QueryWithdrawApplyById(context.Background(), &dbproto.WithdrawApplyId{Id: id})
	if err != nil {
		log.Println("QueryWithdrawApplyById error", err)
		srv.Json(err, http.StatusInternalServerError, c)
		return
	}
	log.Printf("res is %+v", res)
	// srv.Log(c, id, "获取提现申请的详细信息", core.QUERY_OPERATION)
	srv.Json(res, http.StatusOK, c)
}

// HandleUpdateWithdrawApplyStaus 修改提现申请的状态
func (srv *UsercenterService) HandleUpdateWithdrawApplyStaus(c *gin.Context) {
	msg := dbproto.UpdateWAStatusArg{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}
	log.Printf("msg is %+v\n", msg)
	rid, err := srv.Db().UpdateWithdrawApplyStatus(context.Background(), &msg)
	if err != nil {
		log.Println("UpdateWithdrawApplyStatus error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	log.Printf("Id is %v\n", rid.GetValue())
	// srv.Log(c, msg, "更新提现申请的状态", core.UPDATE_OPERATION)
	srv.Json("", http.StatusOK, c)
}

// HandleClaimWithdrawApply 认领提现申请
func (srv *UsercenterService) HandleClaimWithdrawApply(c *gin.Context) {
	msg := dbproto.ClaimWAArg{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		return
	}
	log.Printf("msg is %+v\n", msg)
	var ok bool
	msg.Auditor, ok = srv.GetUserInfo(c, "username").(string)
	if !ok {
		log.Println("GetUserInfo error")
		msg.Auditor = "测试"
		// srv.Json("禁止操作", http.StatusInternalServerError, c)
		// return
	}

	checkRes, err := srv.Db().CheckWithdrawApply(context.Background(), &dbproto.CheckWAArg{WithdrawApplyId: msg.GetWithdrawApplyId(), Auditor: msg.GetAuditor()})
	if checkRes.GetExistAuditor() {
		if !checkRes.GetIsOwn() {
			srv.Json("该申请已被他人认领", http.StatusForbidden, c)
			return
		}
		srv.Json("你已经认领过", http.StatusForbidden, c)
		return
	}

	res, err := srv.Db().ClaimWithdrawApply(context.Background(), &msg)
	if err != nil {
		log.Println("ClaimWithdrawApply error", err)
		srv.Json("禁止操作", http.StatusInternalServerError, c)
		return
	}
	log.Printf("res is %+v\n", res)
	// srv.GetUserInfo(c *gin.Context, key string)
	srv.Json(res, http.StatusOK, c)
}

// HandleCheckWithdrawApply 检查
func (srv *UsercenterService) HandleCheckWithdrawApply(c *gin.Context) {
	log.Println("Handle Check Withdraw Apply")
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	msg := dbproto.CheckWAArg{
		WithdrawApplyId: id,
	}
	var ok bool
	msg.Auditor, ok = srv.GetUserInfo(c, "username").(string)
	if !ok {
		log.Println("GetUserInfo error")
		msg.Auditor = "测试"
		// srv.Json("禁止操作", http.StatusInternalServerError, c)
		// return
	}
	res, err := srv.Db().CheckWithdrawApply(context.Background(), &msg)
	if err != nil {
		log.Println("CheckWithdrawApply error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	log.Printf("res is %+v\n", res)
	// srv.GetUserInfo(c *gin.Context, key string)
	srv.Json(res, http.StatusOK, c)
}

// HandleWithdrawAuditAuthAdd 添加一条审核权限信息
func (srv *UsercenterService) HandleWithdrawAuditAuthAdd(c *gin.Context)  {
	log.Println("Handle WithdrawAuditAuth Add")
	msg := dbproto.WithdrawAuditAuth{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}
	log.Printf("msg is %+v\n", msg)
	createId, ok := srv.GetUserInfo(c, "id").(float64)
	if !ok {
		log.Println("GetUserInfo error")
		// auditor = "测试"
		srv.Json("禁止操作", http.StatusInternalServerError, c)
		return
	}
	msg.Creator = int64(createId)
	rid, err := srv.Db().CreateWithdrawAuditAuth(context.Background(), &msg)
	if err != nil {
		log.Println("CreateWithdrawAuditAuth error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	log.Printf("Id is %v\n", rid)
	srv.Log(c, msg, "添加提现申请权限信息", core.ADD_OPERATION)
	srv.Json("", http.StatusOK, c)
}

// HanldeWithdrawAuditAuthList 获取提现申请权限列表
func (srv *UsercenterService) HanldeWithdrawAuditAuthList(c *gin.Context)  {
	log.Println("Handle Withdraw Audit Auth List ")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pagesize", "10")
	userId := c.DefaultQuery("userId", "0")

	pageDb, err := strconv.Atoi(page)
	if err != nil {
		log.Println("error", err)
		srv.Json("页码参数错误", http.StatusForbidden, c)
		return
	}
	pageSizeDb, err := strconv.Atoi(pageSize)
	if err != nil {
		log.Println("error", err)
		srv.Json("每页尺寸参数错误", http.StatusForbidden, c)
		return
	}

	userIdDb, err := strconv.ParseInt(userId, 10, 0)
	if err != nil {
		log.Println("error", err)
		srv.Json("用户ID错误", http.StatusForbidden, c)
		return
	}

	waaTypeStr := c.DefaultQuery("type", strconv.Itoa(int(dbproto.WithdrawAuditAuth_TYPE_ALL)))
	waaType, err := strconv.Atoi(waaTypeStr)
	if err != nil {
		log.Println("statusStr", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}

	msg := dbproto.QueryWithdrawAuditAuthArg{
		Page:     int32(pageDb),
		PageSize: int32(pageSizeDb),
		Type: dbproto.WithdrawAuditAuth_TypeBackend(waaType),
		UserId: userIdDb,
		// Auditor: c.Query("")
	}

	res, err := srv.Db().QueryWithdrawAuditAuthList(context.Background(), &msg)
	if err != nil {
		log.Println("QuerYWithdrawAuditAuthList error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	// srv.Log(c, msg, "查询提现申请列表", core.QUERY_OPERATION)
	srv.Json(res, http.StatusOK, c)
}

// HandleWithdrawAuditAuth 获取提现审核权限信息
func (srv *UsercenterService) HandleWithdrawAuditAuth(c *gin.Context)  {
	log.Println("Handle WithdrawAuditAuth")
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	res, err := srv.Db().QueryWithdrawAuditAuthById(context.Background(), &dbproto.IntValue{Value: id})
	if err != nil {
		log.Println("QueryWithdrawAuditAuthById error", err)
		srv.Json(err, http.StatusInternalServerError, c)
		return
	}
	log.Printf("res is %+v", res)
	// srv.Log(c, id, "获取提现申请的详细信息", core.QUERY_OPERATION)
	srv.Json(res, http.StatusOK, c)
}

// HandleUpdateWithdrawAuditAuth 更新提现权限
func (srv *UsercenterService) HandleUpdateWithdrawAuditAuth(c *gin.Context)  {
	log.Println("Handle Update Withdraw Audit Auth")
	msg := dbproto.WithdrawAuditAuth{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}
	log.Printf("msg is %+v\n", msg)
	rid, err := srv.Db().UpdateWithdrawAuditAuth(context.Background(), &msg)
	if err != nil {
		log.Println("UpdateWithdrawAuditAuth error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	log.Printf("Id is %v\n", rid.GetValue())
	srv.Log(c, msg, "更新提现权限", core.UPDATE_OPERATION)
	srv.Json("", http.StatusOK, c)
}

// HandleWithdrawTransferList 获取提现转账列表
func (srv *UsercenterService) HandleWithdrawTransferList(c *gin.Context) {
	log.Println("Handle WithdrawTransfer List ")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pagesize", "10")
	start := c.DefaultQuery("start", "0")
	end := c.DefaultQuery("end", "0")

	pageDb, err := strconv.Atoi(page)
	if err != nil {
		log.Println("error", err)
		srv.Json("页码参数错误", http.StatusForbidden, c)
		return
	}
	pageSizeDb, err := strconv.Atoi(pageSize)
	if err != nil {
		log.Println("error", err)
		srv.Json("每页尺寸参数错误", http.StatusForbidden, c)
		return
	}

	startDb, err := strconv.ParseInt(start, 10, 0)
	if err != nil {
		log.Println("error", err)
		srv.Json("开始时间参数错误", http.StatusForbidden, c)
		return
	}
	endDb, err := strconv.ParseInt(end, 10, 0)
	if err != nil {
		log.Println("error", err)
		srv.Json("结束时间参数错误", http.StatusForbidden, c)
		return
	}

	isSuccessStr := c.DefaultQuery("is_success", "false") 
	isSuccess, err := strconv.ParseBool(isSuccessStr)
	if err != nil {
		log.Println("error", err)
		srv.Json("参数错误", http.StatusForbidden, c)
		return
	}

	stepStr := c.DefaultQuery("step", strconv.Itoa(int(dbproto.WithdrawApply_STEP_ALL)))
	step, err := strconv.Atoi(stepStr)
	if err != nil {
		log.Println("statusStr", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}

	msg := dbproto.QueryWithdrawTransferArg{
		Step:   dbproto.WithdrawTransfer_WTStepBackend(step),
		IsSuccess: isSuccess,
		Realname: c.Query("realname"),
		Phone:    c.Query("phone"),
		Start:    startDb,
		End:      endDb,
		Page:     int32(pageDb),
		PageSize: int32(pageSizeDb),
		InNo:     c.Query("inNo"),
		Operator: c.Query("operator"),
		// Auditor: c.Query("")
	}

	res, err := srv.Db().QueryWithdrawTransferList(context.Background(), &msg)
	if err != nil {
		log.Println("QueryWithdrawTransferList error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	operator, ok := srv.GetUserInfo(c, "username").(string)
	if !ok {
		log.Println("GetUserInfo error")
		// msg.Auditor = "测试"
		// srv.Json("禁止操作", http.StatusInternalServerError, c)
		// return
	}
	for _, wa := range res.List {
		if wa.Operator == "" || wa.Operator == operator {
			wa.CanOperate = true
		}
	}
	log.Printf("res is %+v", res)

	// srv.Log(c, msg, "查询提现申请列表", core.QUERY_OPERATION)
	srv.Json(res, http.StatusOK, c)
}


// HandleGetWithdrawTransfer 获取提现转账
func (srv *UsercenterService) HandleGetWithdrawTransfer(c *gin.Context)  {
	log.Println("Handle WithdrawTransfer")
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		log.Printf("%+v\n", err)
		srv.Json("获取失败", http.StatusForbidden, c)
		return
	}

	res, err := srv.Db().QueryWithdrawTransferById(context.Background(), &dbproto.IntValue{Value: id})
	if err != nil {
		log.Println("QueryWithdrawTransferById error", err)
		srv.Json(err, http.StatusInternalServerError, c)
		return
	}
	log.Printf("res is %+v", res)
	// srv.Log(c, id, "获取提现申请的详细信息", core.QUERY_OPERATION)
	srv.Json(res, http.StatusOK, c)
}


// HandleUpdateWithdrawTransfer 修改提现转账
func (srv *UsercenterService) HandleUpdateWithdrawTransfer(c *gin.Context)  {
	log.Println("Handle Update WithdrawTransfer")
	msg := dbproto.WithdrawTransfer{}
	err := c.BindJSON(&msg)
	if err != nil {
		log.Println("BindJSON", err)
		srv.Json("数据格式错误", http.StatusBadRequest, c)
		return
	}
	log.Printf("msg is %+v\n", msg)
	rid, err := srv.Db().UpdateWithdrawTransfer(context.Background(), &msg)
	if err != nil {
		log.Println("UpdateWithdrawTransfer error", err)
		srv.Json("服务器异常", http.StatusInternalServerError, c)
		return
	}
	log.Printf("Id is %v\n", rid.GetValue())
	srv.Log(c, msg, "修改提现转账", core.UPDATE_OPERATION)
	srv.Json("", http.StatusOK, c)	
}


