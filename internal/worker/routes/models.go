package routes

import "github.com/osamikoyo/hrm-api/internal/worker/pb"

type Worker struct{
	UserID uint64 
	Firstname string 
	Secondname string
	Salary uint64
	Email string
	Post string
}

func ToModels(pbwork *pb.Worker) *Worker {
	work := &Worker{
		UserID: pbwork.UserID,
		Salary: pbwork.Salary,
		Firstname: pbwork.Firstname,
		Secondname: pbwork.Secondname,
		Post: pbwork.Post,
		Email: pbwork.Email,
	}

	return work
}


func ToPB(worker *Worker) *pb.Worker {
	work := &pb.Worker{
		UserID: worker.UserID,
		Email: worker.Email,
		Salary: worker.Salary,
		Firstname: worker.Firstname,
		Secondname: worker.Secondname,
		Post: worker.Post,
	}
	return work
}

type DeleteRequest struct{
	ID uint64 `json:"id"`
}

type GetRequest struct{
	ID uint64 `json:"id"`
}

type UpdateReq struct{
	Worker Worker `json:"worker"`
	ID uint64 `json:"id"`
}