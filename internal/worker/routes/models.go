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