package behinder

import (
	"github.com/go0p/wsm/lib/utils"
)

type OnlyJavaParams struct {
	// print 模式，java class 改 String 类型比较好改，所以用 bool 类型传递字符串
	ForcePrint bool `json:"forcePrint,string"`
	// 不采用加密
	NotEncrypt bool `json:"notEncrypt,string"`
}

type PingParams struct {
	OnlyJavaParams
	Content string `json:"content"`
}

// Check 检查是否赋值，没有就使用随机值
func (p *PingParams) SetDefaultAndCheckValue() error {
	if len(p.Content) == 0 {
		p.Content = utils.RandomRangeString(50, 200)
	}
	return nil
}

type BasicInfoParams struct {
	OnlyJavaParams
	WhatEver string `json:"whatever"`
}

func (b *BasicInfoParams) SetDefaultAndCheckValue() error {
	if len(b.WhatEver) == 0 {
		b.WhatEver = utils.RandomRangeString(50, 200)
	}
	return nil
}

type ExecParams struct {
	OnlyJavaParams
	Cmd  string `json:"cmd"`
	Path string `json:"path"`
}

func (e *ExecParams) SetDefaultAndCheckValue() error {
	return nil
}

type FileOptParams struct {
	OnlyJavaParams
	Mode            string `json:"mode"`
	CurrentPath     string `json:"current_path"`
	NewFileName     string `json:"new_file_name"`
	OldFileName     string `json:"old_file_name"`
	DirName         string `json:"dir_name"`
	RemotePath      string `json:"remote_path"`
	LocalPath       string `json:"local_path"`
	Charset         string `json:"charset"`
	CreateTimeStamp string `json:"create_time_stamp"`
	ModifyTimeStamp string `json:"modify_time_stamp"`
	AccessTimeStamp string `json:"access_time_stamp"`
	IsChunk         bool   `json:"is_chunk"`
}
