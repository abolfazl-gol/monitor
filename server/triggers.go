package monitor

import (
	"context"
	"fmt"
	"strings"

	"../api-monitor/models"
	monitornig "../api-monitor/monitor/v1"
	"../api-monitor/pkg/convert"
	"../api-monitor/pkg/errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rezam90/goutils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) ListTriggers(ctx context.Context, req *monitornig.ListTriggersRequest) (*monitornig.ListTriggersResponse, error) {
	triggers, err := models.FindTriggers()
	if err != nil {
		fmt.Println("[ListTriggers] can't FindTriggers:", err)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	protoToTriggers := make([]*monitornig.Trigger, 0)
	for _, trigger := range triggers {
		protoToTriggers = append(protoToTriggers, convert.TriggerToProto(trigger))
	}

	return &monitornig.ListTriggersResponse{Triggers: protoToTriggers}, nil
}

func (s *Service) GetTrigger(ctx context.Context, req *monitornig.GetTriggerRequest) (*monitornig.Trigger, error) {
	trigger, err := models.GetTriggerByID(req.Id)
	if err != nil {
		fmt.Println("[GetTrigger] can't GetTriggerByID:", err)
		if errors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return convert.TriggerToProto(trigger), nil
}

func (s *Service) CreateTrigger(ctx context.Context, req *monitornig.CreateTriggerRequest) (*monitornig.Trigger, error) {
	trigger := &models.Trigger{
		ID:               req.Trigger.Id,
		UserID:           req.Trigger.UserId,
		MonitorID:        req.Trigger.MonitorId,
		TriggerKind:      req.Trigger.TriggerKind,
		NotificationKind: req.Trigger.NotificationKind,
		ErrTolerance:     req.Trigger.ErrTolerance,
		Warning:          req.Trigger.Warning,
		Enabled:          req.Trigger.Enabled,
	}

	if err := models.CreateTrigger(trigger); err != nil {
		fmt.Println("[CreateTrigger] can't CreateTrigger:", err)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return convert.TriggerToProto(trigger), nil
}

func (s *Service) UpdateTrigger(ctx context.Context, req *monitornig.UpdateTriggerRequest) (*monitornig.Trigger, error) {
	trigger, err := models.GetTriggerByID(req.Trigger.Id)
	if err != nil {
		fmt.Println("[UpdateTrigger] can't GetTriggerByID:", err)
		if errors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	forbiddens := []string{"id", "user_id", "monitor_id"}
	cols := make([]string, 0)
	for _, path := range req.UpdateMask.GetPaths() {
		if strings.HasPrefix(path, "trigger.") {
			colum := path[strings.Index(path, ".")+1:]
			if goutils.SliceExists(forbiddens, colum) {
				continue
			}
			cols = append(cols, colum)
		}
	}

	trigger.ID = req.Trigger.Id
	trigger.Warning = req.Trigger.Warning
	trigger.UserID = req.Trigger.UserId
	trigger.Enabled = req.Trigger.Enabled
	trigger.MonitorID = req.Trigger.MonitorId
	trigger.TriggerKind = req.Trigger.TriggerKind
	trigger.ErrTolerance = req.Trigger.ErrTolerance
	trigger.NotificationKind = req.Trigger.NotificationKind

	if err := models.UpdateTrigger(trigger, cols...); err != nil {
		fmt.Println("[UpdateTrigger] can't UpdateTrigger:", err)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return convert.TriggerToProto(trigger), nil
}

func (s *Service) DeleteTrigger(ctx context.Context, req *monitornig.DeleteTriggerRequest) (*empty.Empty, error) {
	_, err := models.GetTriggerByID(req.Id)
	if err != nil {
		fmt.Println("[DeleteTrigger] can't GetTriggerByID:", err)
		if errors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	if err := models.DeleteTrigger(req.Id); err != nil {
		fmt.Println("[DeleteTrigger] can't DeleteTrigger:", err)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &empty.Empty{}, nil
}
