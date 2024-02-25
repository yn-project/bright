//nolint:nolintlint,dupl
package datafin

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Vigo-Tea/go-ethereum-ant/accounts/abi/bind"
	"github.com/Vigo-Tea/go-ethereum-ant/common"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	accountmgr "yun.tea/block/bright/account/pkg/mgr"
	converter "yun.tea/block/bright/datafin/pkg/converter/datafin"
	crud "yun.tea/block/bright/datafin/pkg/crud/datafin"

	data_fin "yun.tea/block/bright/common/chains/eth/datafin"
	"yun.tea/block/bright/common/cruder"
	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/common/utils"
	"yun.tea/block/bright/proto/bright"
	proto "yun.tea/block/bright/proto/bright/datafin"
	"yun.tea/block/bright/proto/bright/topic"
)

func (s *DataFinServer) CreateDataFin(ctx context.Context, in *proto.CreateDataFinRequest) (*proto.CreateDataFinResponse, error) {
	topicServer := &TopicServer{}
	_, err := topicServer.GetTopic(ctx, &topic.GetTopicRequest{TopicID: in.TopicID})
	if err != nil {
		return &proto.CreateDataFinResponse{}, err
	}

	reqList := []*proto.DataFinReq{}
	retries := uint32(0)

	needCompact := false
	if in.Type == proto.DataType_JsonType {
		needCompact = true
	}

	for _, v := range in.Infos {
		reqItem := &proto.DataFinReq{
			DataID:  &v.DataID,
			TopicID: &in.TopicID,
			Retries: &retries,
			State:   proto.DataFinState_DataFinStateDefault.Enum(),
		}
		dfhash, err := utils.SumSha256String(v.Data, needCompact)
		if err != nil {
			remark := fmt.Sprintf("parse json string failed, err: %v", err)
			reqItem.State = proto.DataFinState_DataFinStateFailed.Enum()
			reqItem.Remark = &remark
		} else {
			_dfhash := dfhash.ToHexString()
			reqItem.DataFin = &_dfhash
		}
		reqList = append(reqList, reqItem)

	}

	infos, err := crud.CreateBulk(ctx, reqList)
	if err != nil {
		return &proto.CreateDataFinResponse{}, err
	}

	logger.Sugar().Infof("success to create datafin,topic-id: %v,type: %v", in.TopicID, in.Type)
	return &proto.CreateDataFinResponse{
		Infos: converter.Ent2GrpcMany(infos),
	}, nil
}

func (s *DataFinServer) GetDataFins(ctx context.Context, in *proto.GetDataFinsRequest) (*proto.GetDataFinsResponse, error) {
	topicServer := &TopicServer{}
	_, err := topicServer.GetTopic(ctx, &topic.GetTopicRequest{TopicID: in.TopicID})
	if err != nil {
		return &proto.GetDataFinsResponse{}, err
	}

	conds := &proto.Conds{
		IDs: &bright.StringSliceVal{
			Op:    cruder.IN,
			Value: in.DataFinIDs,
		},
	}
	rows, total, err := crud.Rows(ctx, conds, 0, 0)
	if err != nil {
		return &proto.GetDataFinsResponse{}, err
	}

	return &proto.GetDataFinsResponse{Infos: converter.Ent2GrpcMany(rows), Total: uint32(total)}, nil
}

func (s *DataFinServer) GetAllDataFins(ctx context.Context, in *proto.GetAllDataFinsRequest) (*proto.GetAllDataFinsResponse, error) {
	rows, total, err := crud.AllRows(ctx, in.Desc, int(in.Offset), int(in.Limit))
	if err != nil {
		return &proto.GetAllDataFinsResponse{}, err
	}
	return &proto.GetAllDataFinsResponse{Infos: converter.Ent2GrpcMany(rows), Total: uint32(total)}, nil
}

func (s *DataFinServer) CheckDataFin(ctx context.Context, in *proto.CheckDataFinRequest) (*proto.CheckDataFinResponse, error) {
	topicServer := &TopicServer{}
	_, err := topicServer.GetTopic(ctx, &topic.GetTopicRequest{TopicID: in.TopicID})
	if err != nil || len(in.DataFins) == 0 {
		return &proto.CheckDataFinResponse{TopicID: in.TopicID}, err
	}

	vals := []*big.Int{}
	for _, v := range in.DataFins {
		fin256, err := utils.FromHexString(v)
		if err != nil {
			vals = append(vals, &big.Int{})
		} else {
			vals = append(vals, fin256.ToBigInt())
		}

	}

	rets := []*big.Int{}
	err = accountmgr.WithReadContract(ctx, false, func(ctx context.Context, from common.Address, contract *data_fin.DataFin, cli *ethclient.Client) error {
		rets, err = contract.VerifyItems(&bind.CallOpts{Pending: true, From: from, Context: ctx}, in.TopicID, vals)
		return err
	})
	if err != nil {
		logger.Sugar().Error(err)
		return &proto.CheckDataFinResponse{}, err
	}

	infos := []*proto.CheckDataFinResp{}
	for i, v := range in.DataFins {
		infos = append(infos, &proto.CheckDataFinResp{
			DataFin: v,
			TxTime:  uint32(rets[i].Uint64()),
			Passed:  rets[i].Uint64() > 0,
		})
	}

	return &proto.CheckDataFinResponse{
		TopicID: in.TopicID,
		Infos:   infos,
	}, nil
}

func (s *DataFinServer) CheckIDDataFin(ctx context.Context, in *proto.CheckIDDataFinRequest) (*proto.CheckIDDataFinResponse, error) {
	topicServer := &TopicServer{}
	_, err := topicServer.GetTopic(ctx, &topic.GetTopicRequest{TopicID: in.TopicID})
	if err != nil || len(in.Infos) == 0 {
		return &proto.CheckIDDataFinResponse{TopicID: in.TopicID}, err
	}

	vals := []*big.Int{}
	ids := []string{}
	nullDataFin := ""
	for _, v := range in.Infos {
		if v.DataFin == nil {
			v.DataFin = &nullDataFin
		}
		fin256, err := utils.FromHexString(*v.DataFin)
		if err != nil {
			vals = append(vals, &big.Int{})
		} else {
			vals = append(vals, fin256.ToBigInt())
		}
		ids = append(ids, v.DataID)
	}

	rets := []*big.Int{}
	err = accountmgr.WithReadContract(ctx, false, func(ctx context.Context, from common.Address, contract *data_fin.DataFin, cli *ethclient.Client) error {
		rets, err = contract.VerifyIDItems(&bind.CallOpts{Pending: true, From: from, Context: ctx}, in.TopicID, ids, vals)
		return err
	})
	if err != nil {
		logger.Sugar().Error(err)
		return &proto.CheckIDDataFinResponse{}, err
	}
	fmt.Println(utils.PrettyStruct(ids))
	fmt.Println(utils.PrettyStruct(rets))
	infos := []*proto.CheckIDDataFinResp{}
	for i, v := range in.Infos {
		infos = append(infos, &proto.CheckIDDataFinResp{
			DataID:  v.DataID,
			DataFin: *v.DataFin,
			TxTime:  uint32(rets[i].Uint64()),
			Passed:  rets[i].Uint64() > 0,
		})
	}

	return &proto.CheckIDDataFinResponse{
		TopicID: in.TopicID,
		Infos:   infos,
	}, nil
}

func (s *DataFinServer) CheckDataFinWithData(ctx context.Context, in *proto.CheckDataFinWithDataRequest) (*proto.CheckDataFinResponse, error) {
	topicServer := &TopicServer{}
	_, err := topicServer.GetTopic(ctx, &topic.GetTopicRequest{TopicID: in.TopicID})
	if err != nil || len(in.Datas) == 0 {
		return &proto.CheckDataFinResponse{TopicID: in.TopicID}, err
	}

	needCompact := false
	if in.Type == proto.DataType_JsonType {
		needCompact = true
	}

	vals := []*big.Int{}
	dfhash := []string{}
	for _, v := range in.Datas {
		fin256, err := utils.SumSha256String(v, needCompact)
		if err != nil {
			vals = append(vals, &big.Int{})
		} else {
			vals = append(vals, fin256.ToBigInt())
		}
		dfhash = append(dfhash, fin256.ToHexString())
	}

	rets := []*big.Int{}
	err = accountmgr.WithReadContract(ctx, false, func(ctx context.Context, from common.Address, contract *data_fin.DataFin, cli *ethclient.Client) error {
		rets, err = contract.VerifyItems(&bind.CallOpts{Pending: true, From: from, Context: ctx}, in.TopicID, vals)
		return err
	})
	if err != nil {
		logger.Sugar().Error(err)
		return &proto.CheckDataFinResponse{}, err
	}

	infos := []*proto.CheckDataFinResp{}
	for i, v := range dfhash {
		infos = append(infos, &proto.CheckDataFinResp{
			DataFin: v,
			TxTime:  uint32(rets[i].Uint64()),
			Passed:  rets[i].Uint64() > 0,
		})
	}

	return &proto.CheckDataFinResponse{
		TopicID: in.TopicID,
		Infos:   infos,
	}, nil
}

func (s *DataFinServer) CheckIDDataFinWithData(ctx context.Context, in *proto.CheckIDDataFinWithDataRequest) (*proto.CheckIDDataFinResponse, error) {
	topicServer := &TopicServer{}
	_, err := topicServer.GetTopic(ctx, &topic.GetTopicRequest{TopicID: in.TopicID})
	if err != nil || len(in.Infos) == 0 {
		return &proto.CheckIDDataFinResponse{TopicID: in.TopicID}, err
	}

	needCompact := false
	if in.Type == proto.DataType_JsonType {
		needCompact = true
	}

	vals := []*big.Int{}
	ids := []string{}
	dfhash := []string{}
	nullDataFin := ""
	for _, v := range in.Infos {
		if v.Data == nil {
			v.Data = &nullDataFin
		}
		fin256, err := utils.SumSha256String(*v.Data, needCompact)
		if err != nil {
			vals = append(vals, &big.Int{})
		} else {
			vals = append(vals, fin256.ToBigInt())
		}
		ids = append(ids, v.DataID)
		dfhash = append(dfhash, fin256.ToHexString())
	}

	rets := []*big.Int{}
	err = accountmgr.WithReadContract(ctx, false, func(ctx context.Context, from common.Address, contract *data_fin.DataFin, cli *ethclient.Client) error {
		rets, err = contract.VerifyIDItems(&bind.CallOpts{Pending: true, From: from, Context: ctx}, in.TopicID, ids, vals)
		return err
	})
	if err != nil {
		logger.Sugar().Error(err)
		return &proto.CheckIDDataFinResponse{}, err
	}

	infos := []*proto.CheckIDDataFinResp{}
	for i, v := range in.Infos {
		infos = append(infos, &proto.CheckIDDataFinResp{
			DataID:  v.DataID,
			DataFin: dfhash[i],
			TxTime:  uint32(rets[i].Uint64()),
			Passed:  rets[i].Uint64() > 0,
		})
	}
	return &proto.CheckIDDataFinResponse{
		TopicID: in.TopicID,
		Infos:   infos,
	}, nil
}

func (s *DataFinServer) QRCheckDefaultParms(ctx context.Context, in *proto.QRCheckDefaultParmsRequest) (*proto.QRCheckDefaultParmsResponse, error) {
	topicServer := &TopicServer{}
	resp, err := topicServer.GetTopics(ctx, &topic.GetTopicsRequest{
		Offset: 0,
		Limit:  0,
	})

	if err != nil {
		return &proto.QRCheckDefaultParmsResponse{}, err
	}

	topicIDs := []string{}
	for _, v := range resp.Infos {
		topicIDs = append(topicIDs, v.TopicID)
	}

	types := []string{}
	for _, v := range proto.DataFinState_name {
		types = append(types, v)
	}

	return &proto.QRCheckDefaultParmsResponse{
		PrefixUrl: fmt.Sprintf("%v/%v", in.HostAddr, "api/qr/check"),
		Types:     types,
		TopicIDs:  topicIDs,
	}, nil
}

func (s *DataFinServer) GetQRCheckUrl(ctx context.Context, in *proto.GetQRCheckUrlRequest) (*proto.GetQRCheckUrlResponse, error) {
	topicServer := &TopicServer{}
	resp, err := topicServer.GetTopics(ctx, &topic.GetTopicsRequest{
		Offset: 0,
		Limit:  0,
	})

	if err != nil {
		return &proto.GetQRCheckUrlResponse{}, err
	}

	topicIDs := []string{}
	for _, v := range resp.Infos {
		topicIDs = append(topicIDs, v.TopicID)
	}

	types := []string{}
	for _, v := range proto.DataFinState_name {
		types = append(types, v)
	}

	return &proto.GetQRCheckUrlResponse{
		Url: "",
	}, nil
}
