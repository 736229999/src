package service

import "qiniu-go/model/store"

/******************管理model，包括存储一些model到内存和删除model*****************/

/**
	将userinfo存到内存
 */
func StoreUserInfo(uid int,user *store.StoreUser){
	if user == nil{
		user = new(store.StoreUser)
	}
	store.StoreUserMap[uid] = user
}

func DeleteUserInfo(uid int ,user *store.StoreUser) bool{
	if user == nil{
		return false
	}
	store.StoreUserMap[uid] = nil
	return true
}
