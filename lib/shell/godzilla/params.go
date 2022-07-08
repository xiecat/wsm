package godzilla

import (
	"errors"
	"fmt"
	"strings"
)

type ExecParams struct {
	Template    string `json:"template"`
	Command     string `json:"command"`
	CurrPath    string `json:"currPath"`
	RealCommand string `json:"realCommand"`
}

func (e *ExecParams) SetDefaultAndCheckValue() error {
	if len(e.Template) == 0 {
		e.Template = `cmd /c "{command}"`
	}
	if len(e.Command) == 0 {
		return errors.New("command is empty")
	}
	if len(e.CurrPath) == 0 {
		e.CurrPath = "C:/"
	}
	c := fmt.Sprintf(`cd /d "%s"&%s`, e.CurrPath, e.Command)
	if len(e.RealCommand) == 0 {
		e.RealCommand = strings.ReplaceAll(e.Template, "{command}", c)
	}
	return nil
}

//type UseMode string
//
//const (
//	NewFile         UseMode = "newFile"
//	DeleteFile      UseMode = "deleteFile"
//	GetFile         UseMode = "getFile"
//	DownloadFile    UseMode = "downloadFile"
//	UploadFile      UseMode = "uploadFile"
//	MoveFile        UseMode = "moveFile"
//	CopyFile        UseMode = "deleteFile"
//	NewDir          UseMode = "newDir"
//	BigFileUpload   UseMode = "bigFileUpload"
//	BigFileDownload UseMode = "bigFileDownload"
//	GetFileSize     UseMode = "getFileSize"
//	FileRemoteDown  UseMode = "fileRemoteDown"
//	SetFileAttr     UseMode = "setFileAttr"
//)

type GetFiles struct {
	DirName string `json:"dirName"`
}

func (g GetFiles) SetDefaultAndCheckValue() error {
	if len(g.DirName) == 0 {
		return errors.New("dir name is empty")
	}
	return nil
}

type NewDir struct {
	DirName string `json:"dirName"`
}

func (n NewDir) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type DownloadFile struct {
	FileName string `json:"fileName"`
}

func (d DownloadFile) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type UploadFile struct {
	FileName  string `json:"fileName"`
	FileValue string `json:"fileValue"`
}

func (u UploadFile) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type CopyFile struct {
	SrcFileName  string `json:"srcFileName"`
	DestFileName string `json:"destFileName"`
}

func (c CopyFile) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type MoveFile struct {
	SrcFileName  string `json:"srcFileName"`
	DestFileName string `json:"destFileName"`
}

func (m MoveFile) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type DeleteFile struct {
	FileName string `json:"fileName"`
}

func (d DeleteFile) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type NewFile struct {
	FileName string `json:"fileName"`
}

func (n NewFile) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type BigFileUpload struct {
	FileName     string `json:"fileName"`
	FileContents string `json:"fileContents"`
	Position     string `json:"position"`
}

func (b BigFileUpload) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type BigFileDownload struct {
	FileName    string `json:"fileName"`
	Position    string `json:"position"`
	ReadByteNum string `json:"readByteNum" default:"read"`
	Mode        string `json:"mode"`
}

func (b BigFileDownload) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type GetFileSize struct {
	FileName string `json:"fileName"`
	Mode     string `json:"mode" default:"fileSize"`
}

func (g GetFileSize) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type FileRemoteDown struct {
	Url      string `json:"url"`
	SaveFile string `json:"saveFile"`
}

func (f FileRemoteDown) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

type SetFileAttr struct {
	FileName string `json:"fileName"`
	Attr     string `json:"attr"`
}

func (s SetFileAttr) SetDefaultAndCheckValue() error {
	//TODO implement me
	panic("implement me")
}

//type FileOptParams struct {
//	//UseMode      UseMode `json:"useMode"`
//	DirName      string `json:"dirName"`
//	FileName     string `json:"fileName"`
//	FileValue    string `json:"fileValue"`
//	SrcFileName  string `json:"srcFileName"`
//	DestFileName string `json:"destFileName"`
//	FileContents string `json:"fileContents"`
//	Position     string `json:"position"`
//	ReadByteNum  string `json:"readByteNum"`
//	Mode         string `json:"mode"`
//	Url          string `json:"url"`
//	SaveFile     string `json:"saveFile"`
//	Attr         string `json:"attr"`
//}
//
//func (f FileOptParams) SetDefaultAndCheckValue() error {
//	//TODO implement me
//	panic("implement me")
//}

type SqlOptParams struct {
	dbType     string
	dbHost     string
	dbPort     string
	dbUsername string
	dbPassword string
	execType   string
	execSql    string
	dbCharset  string
	currentDb  string
}
