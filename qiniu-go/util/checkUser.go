package util

/**
   检查某张表的uid是否存在
 */
func CheckUserExist(table interface{},uid int) bool {
	hasUid,_ := Engine.Where("uid=?",uid).Get(table)
	return hasUid
}
