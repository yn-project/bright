//nolint:nolintlint,dupl
package datafin

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"yun.tea/block/bright/config"
	converter "yun.tea/block/bright/datafin/pkg/converter/filerecord"
	crud "yun.tea/block/bright/datafin/pkg/crud/filerecord"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/common/servermux"
	"yun.tea/block/bright/proto/bright/datafin"
	proto "yun.tea/block/bright/proto/bright/filerecord"
	"yun.tea/block/bright/proto/bright/topic"

	"github.com/google/uuid"
)

const (
	MaxUploadFileSize = 1 << 10 << 10 * 8
	UploadFileFeild   = "File"
	TopicIDFeild      = "TopicID"
	TypeFeild         = "Type"
	RemarkFeild       = "Remark"
)

func init() {
	mux := servermux.AppServerMux()
	mux.HandleFunc("/upload/file", UploadDataFile)
}

func UploadDataFile(w http.ResponseWriter, r *http.Request) {
	startT := time.Now()
	respBody := []byte{}
	var err error
	var errMsg string
	defer func() {
		respBody = []byte("successed to upload file to server")
		if errMsg != "" {
			logger.Sugar().Errorf("failed to upload file to server, err: %v", errMsg)
			w.WriteHeader(http.StatusBadRequest)
			respBody = []byte(errMsg)
		}

		_, err = w.Write(respBody)
		if err != nil {
			logger.Sugar().Errorf("failed to write response,err %v", err)
		}
	}()

	req, err := parseCreateFileRecordRequest(r.Context(), r)
	if err != nil {
		errMsg = fmt.Sprintf("failed to parse request feild,err: %v", err)
		return
	}

	// judge weather filesize exceed max-size
	err = r.ParseMultipartForm(MaxUploadFileSize)
	if err != nil {
		errMsg = fmt.Sprintf("read file failed %v, %v", MaxUploadFileSize, err)
		return
	}

	inT := time.Now()
	logger.Sugar().Infof("check params %v ms", inT.UnixMilli()-startT.UnixMilli())

	uploadFile, fHandle, err := r.FormFile(UploadFileFeild)
	if err != nil {
		errMsg = fmt.Sprintf("read file failed %v, %v", MaxUploadFileSize, err)
		return
	}
	defer uploadFile.Close()

	req.File = fHandle.Filename
	metaData, err := json.Marshal(req)
	if err != nil {
		errMsg = fmt.Sprintf("failed to parse request feild,err: %v", err)
		return
	}

	filePath := fmt.Sprintf("%v/%v", config.GetConfig().DataFin.DataDir, uuid.NewString())
	donePath := fmt.Sprintf("%v/done-%v", config.GetConfig().DataFin.DataDir, uuid.NewString())
	saveFile, err := os.Create(filePath)
	if err != nil {
		errMsg = fmt.Sprintf("failed to create file,err: %v", err)
		return
	}
	defer func() {
		saveFile.Close()
		os.Rename(filePath, donePath)
	}()

	_, err = saveFile.Write(metaData)
	if err != nil {
		errMsg = fmt.Sprintf("failed to write file,err: %v", err)
		return
	}
	saveFile.WriteString("\n")

	_, err = io.Copy(saveFile, uploadFile)
	if err != nil {
		errMsg = fmt.Sprintf("failed to read upload file to server,err: %v", err)
		return
	}
}

func parseCreateFileRecordRequest(ctx context.Context, r *http.Request) (*proto.CreateFileRecordRequest, error) {
	req := &proto.CreateFileRecordRequest{}
	req.TopicID = r.FormValue(TopicIDFeild)
	topicServer := &TopicServer{}
	_, err := topicServer.GetTopic(ctx, &topic.GetTopicRequest{TopicID: req.TopicID})
	if err != nil {
		return &proto.CreateFileRecordRequest{}, err
	}

	_dataType := r.FormValue(TypeFeild)
	if dataType, ok := datafin.DataType_value[_dataType]; ok {
		req.Type = datafin.DataType_name[dataType]
	} else {

		return &proto.CreateFileRecordRequest{}, fmt.Errorf("wrong type feild")
	}

	req.Remark = r.FormValue(RemarkFeild)
	return req, nil
}

func (s *FileRecordServer) CreateFileRecord(ctx context.Context, in *proto.CreateFileRecordRequest) (*proto.CreateFileRecordResponse, error) {
	return &proto.CreateFileRecordResponse{}, nil
}

func (s *FileRecordServer) GetFileRecord(ctx context.Context, in *proto.GetFileRecordRequest) (*proto.GetFileRecordResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetFileRecord", "ID", in.GetID(), "error", err)
		return &proto.GetFileRecordResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetFileRecord", "ID", in.GetID(), "error", err)
		return &proto.GetFileRecordResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to get filerecord,id: %v,sha1sum: %v,file name: %v", info.ID, info.Sha1Sum, info.FileName)
	return &proto.GetFileRecordResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *FileRecordServer) GetFileRecords(ctx context.Context, in *proto.GetFileRecordsRequest) (*proto.GetFileRecordsResponse, error) {
	var err error
	rows, total, err := crud.Rows(ctx, in.GetInfo(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetFileRecords", "error", err)
		return &proto.GetFileRecordsResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to get filerecords,offset-limit: %v-%v,total: %v", in.GetOffset(), in.GetLimit(), total)
	return &proto.GetFileRecordsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}
