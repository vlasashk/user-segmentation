package api

import (
	"context"
	"github.com/vlasashk/user-segmentation/internal/model/storage"
	"log/slog"
)

type Storage interface {
	CascadeDeleteSegment(context.Context, storage.Segment, *slog.Logger) error
	DeleteUserFromSegments(context.Context, storage.UserSegments, *slog.Logger) error
	GetUserSegmentsInfo(context.Context, storage.User, *slog.Logger) ([]string, error)
	GetSegmentUsersInfo(context.Context, storage.Segment, *slog.Logger) ([]uint64, error)
	AddUserToSegments(context.Context, storage.UserSegments, *slog.Logger) error
	AddUser(context.Context, storage.User, *slog.Logger) (uint64, error)
	AddSegment(context.Context, storage.Segment, *slog.Logger) (uint64, error)
	CsvHistoryReport(context.Context, storage.CsvReport, *slog.Logger) error
}

type ServerAPI struct {
	ListenAddr string
	Store      Storage
	Log        *slog.Logger
}

type UserRequest struct {
	storage.User
}

type SegmentRequest struct {
	storage.Segment
}

type UserSegmentRequest struct {
	storage.UserSegments
}

type UserResponse struct {
	ResponseStatus
	storage.User
}

type SegmentResponse struct {
	ResponseStatus
	storage.Segment
}

type UserSegmentResponse struct {
	ResponseStatus
	storage.UserSegments
}

type GetSegmentsResponse struct {
	ResponseStatus
	UserID      uint64   `json:"user_id" validate:"required"`
	SegmentSlug []string `json:"user_segments" validate:"required"`
}

type GetUsersResponse struct {
	ResponseStatus
	SegmentSlug string   `json:"user_segment" validate:"required"`
	UserIDs     []uint64 `json:"user_ids" validate:"required"`
}

type CsvReportRequest struct {
	storage.CsvReport
}

type CsvReportResponse struct {
	ResponseStatus
	CsvUrl string `json:"csv_url" validate:"required"`
}
