package handler

import (
	"context"

	"github.com/micro/go-log"

	file "bussinessenv/srv/file/proto/file"
)

type File struct {
}

func (f *File) UploadFile(ctx context.Context, req *file.UploadFileRequest, resp *file.Response) error {
	log.Log(req.Name)
	log.Log(req.File)
	return nil
}

func (f *File) DeleteFile(ctx context.Context, req *file.DeleteFileRequest, resp *file.Response) error {
	log.Log("received File.DeleteFile request, Fid is : ", req.Fid)
	resp.Status = 0
	resp.Name = "email.docx"
	resp.Msg = "deleted success"
	resp.Size = 0
	resp.Mime = ""
	return nil
}

func (f *File) UpdateFile(ctx context.Context, req *file.UpdateFileRequest, resp *file.Response) error {
	return nil
}
