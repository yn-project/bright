package mgr

import (
	"bufio"
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"yun.tea/block/bright/common/utils"
	datafinclient "yun.tea/block/bright/datafin/pkg/client/datafin"
	crud "yun.tea/block/bright/datafin/pkg/crud/filerecord"
	datafinproto "yun.tea/block/bright/proto/bright/datafin"
	proto "yun.tea/block/bright/proto/bright/filerecord"

	"yun.tea/block/bright/config"
)

func ParseFileTask(ctx context.Context) {
	filesDir := config.GetConfig().DataFin.DataDir
	files, _ := os.ReadDir(filesDir)
	for _, f := range files {
		filePath := fmt.Sprintf("%v/%v", filesDir, f.Name())
		if f.IsDir() {
			continue
		}

		allBytes, err := os.ReadFile(filePath)

		ff, err := os.Open(filePath)
		if err != nil {
			os.Remove(f.Name())
			continue
		}
		defer ff.Close()

		r := bufio.NewReader(ff)
		bytes, _, err := r.ReadLine()
		if err != nil {
			os.Remove(f.Name())
			continue
		}

		req := &proto.CreateFileRecordRequest{}
		err = json.Unmarshal(bytes, req)
		if err != nil {
			fmt.Println(err)
			os.Remove(f.Name())
			continue
		}

		fmt.Println(utils.PrettyStruct(req))
		infos := []*datafinproto.DataItemReq{}

		for {
			// ReadLine is a low-level line-reading primitive.
			// Most callers should use ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
			bytes, _, err := r.ReadLine()
			if err == io.EOF {
				break
			}

			if err != nil {
				break
			}

			info := &datafinproto.DataItemReq{}
			err = json.Unmarshal(bytes, info)
			if err != nil {
				continue
			}
			infos = append(infos, info)
		}

		recordNum := uint32(len(infos))
		_, err = datafinclient.CreateDataFin(ctx, &datafinproto.CreateDataFinRequest{
			TopicID: req.TopicID,
			Type:    datafinproto.DataType(datafinproto.DataType_value[req.Type]),
			Infos:   infos,
		})

		if err != nil {
			req.Remark = err.Error()

		}

		checkSum := sha1.Sum()
		crud.Create(ctx, &proto.FileRecordReq{
			FileName:  &req.File,
			TopicID:   &req.TopicID,
			RecordNum: &recordNum,
			Remark:    &req.Remark,
		})

	}
}
