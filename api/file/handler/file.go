package handler

import (
	proto "bussinessenv/api/file/proto/file"
	file "bussinessenv/srv/file/proto/file"
	"context"
	"log"
)

type File struct {
	Client file.FileService
}

func (f *File) Upload(ctx context.Context, req *proto.UploadFileRequest, resp *proto.Response) error {
	log.Print("Received FileDAO.UploadFile API request")
	response, err := f.Client.UploadFile(ctx, &file.UploadFileRequest{
		Name: req.Name,
		File: req.File,
	})
	if err != nil {
		return err
	}
	// set api response
	resp.Msg, resp.Status, resp.Size, resp.Mime = response.Msg, response.Status, response.Size, response.Mime
	return nil
}

func (f *File) Delete(ctx context.Context, req *proto.DeleteFileRequest, resp *proto.Response) error {
	log.Print("Received FileDAO.UploadFile API request")
	response, err := f.Client.DeleteFile(ctx, &file.DeleteFileRequest{
		Fid: req.Fid,
	})
	if err != nil {
		return err
	}
	resp.Size, resp.Mime, resp.Size, resp.Name, resp.Msg = response.Size, response.Mime, response.Size, response.Name, response.Msg
	return nil
}

func (f *File) Update(ctx context.Context, req *proto.UpdateFileRequest, resp *proto.Response) error {
	return nil
}
