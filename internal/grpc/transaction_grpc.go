package grpc

import (
	"context"
	"errors"
	"mygo/internal/grpc/pb"
	"mygo/internal/pkg/common"
	"mygo/internal/server/model"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type transactionService struct{}

func (s *transactionService) GetAllTransactions(
	req *pb.GetAllTransactionsRequest,
	stream pb.TransactionService_GetAllTransactionsServer,
) error {

	transactions, err := model.GetAllTransactions()
	if err != nil {
		return err
	}

	for _, t := range transactions {
		stream.Send(&pb.TransactionResponse{
			Transaction: newTransactionMessage(t),
		})
	}

	log.Infof("[GRPC] >> Get all transactions, request from %s\n", stream.Context().Value("user"))

	return nil
}

func (s *transactionService) HandleTransaction(ctx context.Context, req *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	id := req.GetId()
	action := req.GetAction()

	t, err := model.GetTransactionById(int(id))
	if err != nil {
		return nil, err
	}

	switch action {
	case pb.TransactionRequest_GET:
	case pb.TransactionRequest_PASS:
		t.Status = common.StatusPassed
		t.Update()
	case pb.TransactionRequest_REJECT:
		t.Status = common.StatusRejected
		t.Update()
	case pb.TransactionRequest_DELETE:
		model.DeleteTransaction(t.TransactionID)
	default:
		return nil, errors.New("invalid action")
	}

	return &pb.TransactionResponse{
		Transaction: newTransactionMessage(t),
	}, nil
}

var statusMap = map[common.Status]pb.TransactionMessage_Status{
	common.StatusDraft:     pb.TransactionMessage_DRAFT,
	common.StatusPassed:    pb.TransactionMessage_PASSED,
	common.StatusRejected:  pb.TransactionMessage_REJECTED,
	common.StatusCensoring: pb.TransactionMessage_CENSORING,
}

func newTransactionMessage(t *model.Transaction) *pb.TransactionMessage {
	return &pb.TransactionMessage{
		Id:          int32(t.TransactionID),
		UserId:      int32(t.UserID),
		Title:       t.Title,
		Description: t.Description,
		Value:       int64(t.Value),
		Status:      statusMap[t.Status],
		CreatedAt:   timestamppb.New(t.CreatedAt),
		UpdatedAt:   timestamppb.New(t.UpdatedAt),
	}
}
