package endpoints

import (
	"context"
	"log"

	pb "github.com/coop-cmpers/what2do-backend/protos-gen/what2do/v1"
	"github.com/coop-cmpers/what2do-backend/src/helpers"
	"github.com/coop-cmpers/what2do-backend/src/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *What2doServer) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	if len(req.EventName) == 0 || !req.StartTime.IsValid() || !req.EndTime.IsValid() || len(req.Location) == 0 {
		log.Printf("CreateEvent: empty or invalid argument")
		return nil, status.Errorf(codes.InvalidArgument, "CreateEvent: empty or invalid argument")
	}

	db, err := GetStore(ctx)
	if err != nil {
		return nil, err
	}

	eventID := helpers.GenerateUUID()
	event := &store.Event{
		ID:        eventID,
		Name:      req.EventName,
		StartTime: req.StartTime.AsTime(),
		EndTime:   req.EndTime.AsTime(),
		Location:  req.Location,
	}

	err = db.CreateEvent(ctx, event)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "CreateEvent: failed database transaction")
	}

	return &pb.CreateEventResponse{Id: eventID.String()}, nil
}

func (s *What2doServer) GetEvent(ctx context.Context, req *pb.GetEventRequest) (*pb.GetEventResponse, error) {
	if len(req.Id) == 0 {
		log.Printf("GetEvent: empty or invalid argument")
		return nil, status.Errorf(codes.InvalidArgument, "GetEvent: empty or invalid argument")
	}

	db, err := GetStore(ctx)
	if err != nil {
		return nil, err
	}

	event, err := db.GetEvent(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetEvent: failed to get event from database")
	}
	if event == nil {
		return &pb.GetEventResponse{}, nil
	}

	eventProto := &pb.Event{
		Id:        event.ID.String(),
		EventName: event.Name,
		StartTime: timestamppb.New(event.StartTime),
		EndTime:   timestamppb.New(event.EndTime),
		Location:  event.Location,
	}

	return &pb.GetEventResponse{Event: eventProto}, nil
}
