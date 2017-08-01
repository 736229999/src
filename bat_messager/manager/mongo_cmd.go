
package main

import (
	"bat_messager/storage/mongo_store"
)

type SessionStoreCmd struct {
	CmdName   string
	Args      []string
	AnyData   *mongo_store.SessionStoreData
}

func (self SessionStoreCmd)GetCmdName() string {
	return self.CmdName
}

func (self SessionStoreCmd)ChangeCmdName(newName string) {
	self.CmdName = newName
}

func (self SessionStoreCmd)GetArgs() []string {
	return self.Args
}

func (self SessionStoreCmd)AddArg(arg string) {
	self.Args = append(self.Args, arg)
}

func (self SessionStoreCmd)ParseCmd(msglist []string) {
	self.CmdName = msglist[1]
	self.Args = msglist[2:]
}

func (self SessionStoreCmd)GetAnyData() interface{} {
	return self.AnyData
}


type TopicStoreCmd struct {
	CmdName string
	Args    []string
	AnyData *mongo_store.TopicStoreData
}

func (self TopicStoreCmd)GetCmdName() string {
	return self.CmdName
}

func (self TopicStoreCmd)ChangeCmdName(newName string) {
	self.CmdName = newName
}

func (self TopicStoreCmd)GetArgs() []string {
	return self.Args
}

func (self TopicStoreCmd)AddArg(arg string) {
	self.Args = append(self.Args, arg)
}

func (self TopicStoreCmd)ParseCmd(msglist []string) {
	self.CmdName = msglist[1]
	self.Args = msglist[2:]
}

func (self TopicStoreCmd)GetAnyData() interface{} {
	return self.AnyData
}
