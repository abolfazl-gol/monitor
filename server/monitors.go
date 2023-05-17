package monitor

import (
	"context"
	"fmt"
	"strings"

	"../api-monitor/models"
	monitoring "../api-monitor/monitor/v1"
	"../api-monitor/pkg/convert"
	"../api-monitor/pkg/errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rezam90/goutils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) ListMonitors(ctx context.Context, req *monitoring.ListMonitorsRequest) (*monitoring.ListMonitorsResponse, error) {
	monitors, err := models.FindMonitors()
	if err != nil {
		fmt.Println("[ListMonitors] can't FindMonitors:", err)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	protoToMonitor := make([]*monitoring.Monitor, 0)
	for _, monitor := range monitors {
		protoToMonitor = append(protoToMonitor, convert.MonitorToProto(monitor))
	}

	return &monitoring.ListMonitorsResponse{Monitor: protoToMonitor}, nil
}

func (s *Service) GetMonitor(ctx context.Context, req *monitoring.GetMonitorRequest) (*monitoring.Monitor, error) {
	monitor, err := models.GetMonitorByID(req.Id)
	if err != nil {
		fmt.Println("[GetMonitor] can't GetMonitorByID:", err)
		if errors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return convert.MonitorToProto(monitor), nil
}

func (s *Service) CreateMonitor(ctx context.Context, req *monitoring.CreateMonitorRequest) (*monitoring.Monitor, error) {
	monitor := &models.Monitor{
		ID:             req.Monitor.Id,
		UserID:         req.Monitor.UserId,
		ProtocolKind:   req.Monitor.ProtocolKind,
		Title:          req.Monitor.Title,
		Method:         req.Monitor.Method,
		URL:            req.Monitor.Url,
		RequestTime:    req.Monitor.RequestTime,
		ServerLocation: req.Monitor.ServerLocation,
		Enabled:        req.Monitor.Enabled,
	}

	if err := models.CreateMonitor(monitor); err != nil {
		fmt.Println("[CreateMonitor] can't CreateMonitor:", err)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return convert.MonitorToProto(monitor), nil

}

func (s *Service) UpdateMonitor(ctx context.Context, req *monitoring.UpdateMonitorRequest) (*monitoring.Monitor, error) {
	monitor, err := models.GetMonitorByID(req.Monitor.Id)
	if err != nil {
		fmt.Println("[UpdateMonitor] can't GetMonitorByID:", err)
		if errors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	forbiddens := []string{"id", "user_id"}
	cols := make([]string, 0)
	for _, path := range req.UpdateMask.GetPaths() {
		if strings.HasPrefix(path, "monitor.") {
			colum := path[strings.Index(path, ".")+1:]
			if goutils.SliceExists(forbiddens, colum) {
				continue
			}
			cols = append(cols, colum)
		}
	}

	monitor.URL = req.Monitor.Url
	monitor.Title = req.Monitor.Title
	monitor.Method = req.Monitor.Method
	monitor.Enabled = req.Monitor.Enabled
	monitor.RequestTime = req.Monitor.RequestTime
	monitor.ProtocolKind = req.Monitor.ProtocolKind
	monitor.ServerLocation = req.Monitor.ServerLocation

	if err := models.UpdateMonitor(monitor, cols...); err != nil {
		fmt.Println("[UpdateMonitor] can't UpdateMonitor:", err)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return convert.MonitorToProto(monitor), nil
}

func (s *Service) DeleteMonitor(ctx context.Context, req *monitoring.DeleteMonitorRequest) (*empty.Empty, error) {
	_, err := models.GetMonitorByID(req.Id)
	if err != nil {
		fmt.Println("[DeleteMonitor] can't GetMonitorByID:", err)
		if errors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	if err := models.DeleteMonitor(req.Id); err != nil {
		fmt.Println("[DeleteMonitor] can't DeleteMonitor:", err)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &empty.Empty{}, nil
}
