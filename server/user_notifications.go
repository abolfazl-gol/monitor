package monitor

import (
	"context"
	"fmt"
	"strings"

	"git.tulz.ir/abolfazl/fing-monitor/models"
	monitornig "git.tulz.ir/abolfazl/fing-monitor/monitor/v1"
	"git.tulz.ir/abolfazl/fing-monitor/pkg/convert"
	"git.tulz.ir/abolfazl/fing-monitor/pkg/errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rezam90/goutils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) ListUserNotifications(ctx context.Context, req *monitornig.ListUserNotificationsRequest) (*monitornig.ListUserNotificationsResponse, error) {
	userNotifications, err := models.FindUserNotifications()
	if err != nil {
		fmt.Println("[ListUserNotifications] can't FindUserNotifications:", err)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	protoToUserNotifications := make([]*monitornig.UserNotification, 0)
	for _, userNotifi := range userNotifications {
		protoToUserNotifications = append(protoToUserNotifications, convert.UserNotificationToProto(userNotifi))
	}

	return &monitornig.ListUserNotificationsResponse{UserNotifications: protoToUserNotifications}, nil
}

func (s *Service) GetUserNotification(ctx context.Context, req *monitornig.GetUserNotificationRequest) (*monitornig.UserNotification, error) {
	UserNotfication, err := models.GetUserNotificationByID(req.Id)
	if err != nil {
		fmt.Println("[GetUseerNotfication] can't GetUserNotificationByID:", err)
		if errors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return convert.UserNotificationToProto(UserNotfication), nil
}

func (s *Service) CreateUserNotification(ctx context.Context, req *monitornig.CreateUserNotificationRequest) (*monitornig.UserNotification, error) {
	UserNotfication := &models.UserNotification{
		ID:             req.UserNotification.Id,
		UserID:         req.UserNotification.UserId,
		NotificationID: req.UserNotification.NotificationId,
		Title:          req.UserNotification.Title,
		URL:            req.UserNotification.Url,
		MethodKind:     req.UserNotification.MethodKind,
		Phone:          req.UserNotification.Phone,
		Email:          req.UserNotification.Email,
		Code:           req.UserNotification.Code,
		APIKey:         req.UserNotification.ApiKey,
		Enabled:        req.UserNotification.Enabled,
	}

	if err := models.CreateUserNotification(UserNotfication); err != nil {
		fmt.Println("[CreateUserNotification] can't CreateUserNotification:", err)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return convert.UserNotificationToProto(UserNotfication), nil
}

func (s *Service) UpdateUserNotification(ctx context.Context, req *monitornig.UpdateUserNotificationRequest) (*monitornig.UserNotification, error) {
	userNotification, err := models.GetUserNotificationByID(req.UserNotification.Id)
	if err != nil {
		fmt.Println("[UpdateUserNotification] can't GetUserNotificationByID:", err)
		if errors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	forbiddens := []string{"id", "user_id", "notification_id"}
	cols := make([]string, 0)
	for _, path := range req.UpdateMask.GetPaths() {
		if strings.HasPrefix(path, "user_notification.") {
			colum := path[strings.Index(path, ".")+1:]
			if goutils.SliceExists(forbiddens, colum) {
				continue
			}
			cols = append(cols, colum)
		}
	}

	userNotification.URL = req.UserNotification.Url
	userNotification.Title = req.UserNotification.Title
	userNotification.MethodKind = req.UserNotification.MethodKind
	userNotification.Phone = req.UserNotification.Phone
	userNotification.Email = req.UserNotification.Email
	userNotification.Code = req.UserNotification.Code
	userNotification.APIKey = req.UserNotification.ApiKey
	userNotification.Enabled = req.UserNotification.Enabled

	if err := models.UpdateUserNotification(userNotification, cols...); err != nil {
		fmt.Println("[UpdateUserNotification] can't UpdateUserNotification:", err)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return convert.UserNotificationToProto(userNotification), nil
}

func (s *Service) DeleteUserNotification(ctx context.Context, req *monitornig.DeleteUserNotificationRequest) (*empty.Empty, error) {
	_, err := models.GetUserNotificationByID(req.Id)
	if err != nil {
		fmt.Println("[DeleteUserNotification] can't GetUserNotificationByID:", err)
		if errors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	if err := models.DeleteUserNotification(req.Id); err != nil {
		fmt.Println("[DeleteUserNotification] can't DeleteUserNotification:", err)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &empty.Empty{}, nil
}
