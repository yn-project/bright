package mgr

import (
	"bufio"
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"yun.tea/block/bright/common/logger"
	datafinclient "yun.tea/block/bright/datafin/pkg/client/datafin"
	crud "yun.tea/block/bright/datafin/pkg/crud/filerecord"
	datafinproto "yun.tea/block/bright/proto/bright/datafin"
	proto "yun.tea/block/bright/proto/bright/filerecord"

	"yun.tea/block/bright/config"
)

func ParseFileTask(ctx context.Context) {
	for {
		select {
		case <-time.NewTimer(time.Second * 2).C:
			parseFileTask(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func parseFileTask(ctx context.Context) {
	filesDir := config.GetConfig().DataFin.DataDir
	files, _ := os.ReadDir(filesDir)
	for _, f := range files {
		filePath := fmt.Sprintf("%v/%v", filesDir, f.Name())
		if !strings.Contains(f.Name(), "done-") {
			continue
		}
		if f.IsDir() {
			continue
		}

		allBytes, err := os.ReadFile(filePath)
		if err != nil {
			logger.Sugar().Error(err)
			os.Remove(filePath)
			continue
		}

		sha1sum := fmt.Sprintf("%x", sha1.Sum(allBytes))

		ff, err := os.Open(filePath)
		if err != nil {
			logger.Sugar().Error(err)
			os.Remove(filePath)
			continue
		}
		defer ff.Close()

		r := bufio.NewReader(ff)
		bytes, _, err := r.ReadLine()
		if err != nil {
			logger.Sugar().Error(err)
			os.Remove(filePath)
			continue
		}

		req := &proto.CreateFileRecordRequest{}
		err = json.Unmarshal(bytes, req)
		if err != nil {
			logger.Sugar().Error(err)
			os.Remove(filePath)
			continue
		}

		infos := []*datafinproto.DataItemReq{}

		for {
			bytes, _, err := r.ReadLine()
			if err == io.EOF {
				break
			}

			if err != nil {
				logger.Sugar().Error(err)
				break
			}

			info := &datafinproto.DataItemReq{}
			err = json.Unmarshal(bytes, info)
			if err != nil {
				logger.Sugar().Error(err)
				continue
			}
			infos = append(infos, info)
		}

		os.Remove(filePath)

		_, err = datafinclient.CreateDataFin(ctx, &datafinproto.CreateDataFinRequest{
			TopicID: req.TopicID,
			Type:    datafinproto.DataType(datafinproto.DataType_value[req.Type]),
			Infos:   infos,
		})

		state := proto.FileRecordState_FileRecordSuccess
		if err != nil {
			logger.Sugar().Error(err)
			req.Remark = err.Error()
			state = proto.FileRecordState_FileRecordFailed
		}

		recordNum := uint32(len(infos))
		_, err = crud.Create(ctx, &proto.FileRecordReq{
			FileName:  &req.File,
			TopicID:   &req.TopicID,
			RecordNum: &recordNum,
			Sha1Sum:   &sha1sum,
			State:     &state,
			Remark:    &req.Remark,
		})
		if err != nil {
			logger.Sugar().Error(err)
		}
	}
}
