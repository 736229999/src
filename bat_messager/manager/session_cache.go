
package main

import (
	"bat_messager/storage/redis_store"
)

type SessionCacheCmd struct {
	CmdName string
	Args    []string
	AnyData *redis_store.SessionCacheData
}

func (self SessionCacheCmd)GetCmdName() string {
	return self.CmdName
}

func (self SessionCacheCmd)ChangeCmdName(newName string) {
	self.CmdName = newName
}

func (self SessionCacheCmd)GetArgs() []string {
	return self.Args
}

func (self SessionCacheCmd)AddArg(arg string) {
	self.Args = append(self.Args, arg)
}

func (self SessionCacheCmd)ParseCmd(msglist []string) {
	self.CmdName = msglist[1]
	self.Args = msglist[2:]
}

func (self SessionCacheCmd)GetAnyData() interface{} {
	return self.AnyData
}