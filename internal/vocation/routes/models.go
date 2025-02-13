package routes

import (
	"time"

	"github.com/osamikoyo/hrm-api/internal/vocation/pb"
)

type Vocation struct{
	VocId uint64 `gorm:"primaryKey;autoIncrement"`
	UserID uint64
	UserEmail string
	DateStart time.Time
	DateEnd time.Time
}

const TIME_TAMPLATE = "2006.01.02"

type GetReq struct{
	UserID uint64 `json:"userid"`
}

type DeleteReq struct{
	UserID uint64 `json:"userid"`
}

func ToModels(voc *pb.Vocation) (*Vocation, error) {
	startTime, err := time.Parse(TIME_TAMPLATE, voc.StartTime)
	if err != nil{
		return nil, err
	}
	endTime, err := time.Parse(TIME_TAMPLATE, voc.EndTime)
	if err != nil{
		return nil, err
	}

	return &Vocation{
		DateStart: startTime,
		DateEnd: endTime,
		UserEmail: voc.Email,
		UserID: voc.UserID,
		VocId: voc.VocID,
	}, nil
}

func ToPB(voc *Vocation) *pb.Vocation {
	return &pb.Vocation{
		StartTime: voc.DateStart.Format(TIME_TAMPLATE),
		EndTime: voc.DateStart.Format(TIME_TAMPLATE),
		UserID: voc.UserID,
		Email: voc.UserEmail,
		VocID: voc.VocId,
	}
}