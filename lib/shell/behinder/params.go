package behinder

import (
	"errors"
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

// SetDefaultAndCheckValue 检查是否赋值，没有就使用随机值
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

// ListFiles 列出指定 Path 的文件列表
type ListFiles struct {
	OnlyJavaParams
	Path string `json:"path"`
}

func (l *ListFiles) SetDefaultAndCheckValue() error {
	if len(l.Path) == 0 {
		return errors.New("path is empty")
	}
	return nil
}

// GetTimeStamp 指定文件获取时间戳
type GetTimeStamp struct {
	OnlyJavaParams
	Path string `json:"path"`
}

func (g GetTimeStamp) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

// UpdateTimeStamp 更新时间戳
type UpdateTimeStamp struct {
	OnlyJavaParams
	Path            string `json:"path"`
	CreateTimeStamp string `json:"createTimeStamp"`
	AccessTimeStamp string `json:"accessTimeStamp"`
	ModifyTimeStamp string `json:"modifyTimeStamp"`
}

func (u UpdateTimeStamp) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type DeleteFile struct {
	OnlyJavaParams
	Path string `json:"path"`
}

func (d DeleteFile) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type ShowFile struct {
	OnlyJavaParams
	Path    string `json:"path"`
	Charset string `json:"charset"`
}

func (s ShowFile) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type RenameFile struct {
	OnlyJavaParams
	Path    string `json:"path"`
	NewPath string `json:"newPath"`
}

func (r RenameFile) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type CreateFile struct {
	OnlyJavaParams

	Path string `json:"path"`
}

func (c CreateFile) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type CreateDirectory struct {
	OnlyJavaParams

	Path string `json:"path"`
}

func (c CreateDirectory) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type DownloadFile struct {
	OnlyJavaParams
	Path string `json:"path"`
}

func (d DownloadFile) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type UploadFile struct {
	OnlyJavaParams

	Path    string `json:"path"`
	Content []byte `json:"content"`
	IsChunk bool   `json:"isChunk"`
}

func (u UploadFile) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type AppendFile struct {
	OnlyJavaParams
	Path    string `json:"path"`
	Content []byte `json:"content"`
}

func (a AppendFile) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type DBManagerParams struct {
	OnlyJavaParams
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     int    `json:"port,string"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Database string `json:"database"`
	Sql      string `json:"sql"`
}

func (D DBManagerParams) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}
