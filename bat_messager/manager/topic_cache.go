
package main

import (
	"bat_messager/storage/redis_store"
)

type TopicCacheCmd struct {
	CmdName string
	Args    []string
	AnyData *redis_store.TopicCacheData
}

func (self TopicCacheCmd)GetCmdName() string {
	return self.CmdName
}

func (self TopicCacheCmd)ChangeCmdName(newName string) {
	self.CmdName = newName
}

func (self TopicCacheCmd)GetArgs() []string {
	return self.Args
}

func (self TopicCacheCmd)AddArg(arg string) {
	self.Args = append(self.Args, arg)
}

func (self TopicCacheCmd)ParseCmd(msglist []string) {
	self.CmdName = msglist[1]
	self.Args = msglist[2:]
}

func (self TopicCacheCmd)GetAnyData() interface{} {
	return self.AnyData
}

