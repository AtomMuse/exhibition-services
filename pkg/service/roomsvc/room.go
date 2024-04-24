package roomsvc

import (
	"atommuse/backend/exhibition-service/pkg/model"
	roomrepo "atommuse/backend/exhibition-service/pkg/repositorty/Roomrepo"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RoomServices defines the interface for exhibition Room services.
type IRoomServices interface {
	CreateExhibitionRoom(ctx context.Context, Room *model.RequestCreateExhibitionRoom) (*primitive.ObjectID, error)
	DeleteExhibitionRoomByID(ctx context.Context, RoomID string) error
	GetExhibitionRoomByID(ctx context.Context, RoomID string) (*model.ResponseExhibitionRoom, error)
	GetAllExhibitionRooms(ctx context.Context) ([]model.ResponseExhibitionRoom, error)
	GetRoomsByExhibitionID(ctx context.Context, exhibitionID string) ([]model.Room, error)
	UpdateExhibitionRoom(ctx context.Context, RoomID string, updatedRoom *model.RequestUpdateExhibitionRoom) (*primitive.ObjectID, error)
}

// RoomServices is the implementation of the IExhibitionRoomServices interface.
type RoomServices struct {
	Repository roomrepo.IRoomRepository
}

func (service RoomServices) CreateExhibitionRoom(ctx context.Context, Room *model.RequestCreateExhibitionRoom) (*primitive.ObjectID, error) {
	return service.Repository.CreateExhibitionRoom(ctx, Room)
}

func (service RoomServices) DeleteExhibitionRoomByID(ctx context.Context, RoomID string) error {
	return service.Repository.DeleteExhibitionRoomByID(ctx, RoomID)
}

func (service RoomServices) GetExhibitionRoomByID(ctx context.Context, RoomID string) (*model.ResponseExhibitionRoom, error) {
	return service.Repository.GetExhibitionRoomByID(ctx, RoomID)
}

func (service RoomServices) GetAllExhibitionRooms(ctx context.Context) ([]model.ResponseExhibitionRoom, error) {
	return service.Repository.GetAllExhibitionRooms(ctx)
}

// GetRoomsByExhibitionID fetches Rooms for a given exhibition ID.
func (service RoomServices) GetRoomsByExhibitionID(ctx context.Context, exhibitionID string) ([]model.Room, error) {
	return service.Repository.GetRoomsByExhibitionID(ctx, exhibitionID)
}

func (service RoomServices) UpdateExhibitionRoom(ctx context.Context, RoomID string, updatedRoom *model.RequestUpdateExhibitionRoom) (*primitive.ObjectID, error) {
	return service.Repository.UpdateExhibitionRoom(ctx, RoomID, updatedRoom)
}
