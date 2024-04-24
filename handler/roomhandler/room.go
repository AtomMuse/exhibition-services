package roomhandler

import "atommuse/backend/exhibition-service/pkg/service/roomsvc"

// Handler is responsible for handling HTTP requests.
type Handler struct {
	RoomService roomsvc.IRoomServices
}
