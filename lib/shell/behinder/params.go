package behinder

import (
	"errors"
	"github.com/xiecat/wsm/lib/utils"
)

// OnlyJavaParams 只针对 JavaScript 类型的参数
type OnlyJavaParams struct {
	// print 模式，java class 改 String 类型比较好改，所以用 bool 类型传递字符串
	ForcePrint bool `json:"forcePrint,string"`
	// 对结果不进行加密
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
	if len(g.Path) == 0 {
		return errors.New("path is empty")
	}
	return nil
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
	if len(u.Path) == 0 {
		return errors.New("path is empty")
	}
	if len(u.CreateTimeStamp) == 0 {
		return errors.New("createTimeStamp is empty")
	}
	if len(u.AccessTimeStamp) == 0 {
		return errors.New("accessTimeStamp is empty")
	}
	return nil
}

// DeleteFile 删除指定文件
type DeleteFile struct {
	OnlyJavaParams
	Path string `json:"path"`
}

func (d DeleteFile) SetDefaultAndCheckValue() error {
	if len(d.Path) == 0 {
		return errors.New("path is empty")
	}
	return nil
}

type ShowFile struct {
	OnlyJavaParams
	Path    string `json:"path"`
	Charset string `json:"charset"`
}

func (s ShowFile) SetDefaultAndCheckValue() error {
	if len(s.Path) == 0 {
		return errors.New("path is empty")
	}
	return nil
}

type RenameFile struct {
	OnlyJavaParams
	Path    string `json:"path"`
	NewPath string `json:"newPath"`
}

func (r RenameFile) SetDefaultAndCheckValue() error {
	if len(r.Path) == 0 {
		return errors.New("path is empty")
	}
	return nil
}

type CreateFile struct {
	OnlyJavaParams

	Path string `json:"path"`
}

func (c CreateFile) SetDefaultAndCheckValue() error {
	if len(c.Path) == 0 {
		return errors.New("path is empty")
	}
	return nil
}

type CreateDirectory struct {
	OnlyJavaParams

	Path string `json:"path"`
}

func (c CreateDirectory) SetDefaultAndCheckValue() error {
	if len(c.Path) == 0 {
		return errors.New("path is empty")
	}
	return nil
}

type DownloadFile struct {
	OnlyJavaParams
	Path string `json:"path"`
}

func (d DownloadFile) SetDefaultAndCheckValue() error {
	if len(d.Path) == 0 {
		return errors.New("path is empty")
	}
	return nil
}

type UploadFile struct {
	OnlyJavaParams

	Path    string `json:"path"`
	Content []byte `json:"content"`
	IsChunk bool   `json:"isChunk"`
}

func (u UploadFile) SetDefaultAndCheckValue() error {
	if len(u.Path) == 0 {
		return errors.New("path is empty")
	}
	if len(u.Content) == 0 {
		return errors.New("content is empty")
	}
	return nil
}

// AppendFile 对指定文件进行内容追加
type AppendFile struct {
	OnlyJavaParams
	Path    string `json:"path"`
	Content []byte `json:"content"`
}

func (a AppendFile) SetDefaultAndCheckValue() error {
	if len(a.Path) == 0 {
		return errors.New("path is empty")
	}
	return nil
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

func (d DBManagerParams) SetDefaultAndCheckValue() error {
	if len(d.Type) == 0 {
		return errors.New("db type is empty")
	}
	if len(d.Host) == 0 {
		return errors.New("db host is empty")
	}
	if d.Port == 0 {
		return errors.New("db port is error")
	}
	if len(d.Sql) == 0 {
		return errors.New("db sql is empty")
	}
	return nil
}
