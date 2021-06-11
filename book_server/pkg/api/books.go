package api

import (
	"context"
	"github.com/PutskouDzmitry/golang-training-library-grpc/book_server/pkg/data"
	pb "github.com/PutskouDzmitry/golang-training-library-grpc/proto/go_proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sirupsen/logrus"
)

type BookServer struct {
	data *data.BookData
}

var log = logrus.New()

func NewReadServer(data *data.BookData) *BookServer {
	return &BookServer{data: data}
}

func (b BookServer) ReadAll(ctx context.Context, request *pb.ReadAllBooksRequest) (*pb.ReadAllBooksResponse, error) {
	log.Debug("Package books, method ReadAll")
	books, err := b.data.ReadAll()
	if err != nil {
		s := status.Newf(codes.PermissionDenied, "got an error when tried to read all books", err)
		errWithDetails, err := s.WithDetails(request)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return nil, errWithDetails.Err()
	}
	log.Debug("Package books, method ReadAll, Data users", books)
	var respBook []*pb.Book
	for _, currentBooks := range books {
		book := &pb.Book{
			BookId:            currentBooks.BookId,
			AuthorId:          currentBooks.AuthorId,
			PublisherId:       currentBooks.PublisherId,
			NameOfBook:        currentBooks.NameOfBook,
			YearOfPublication: currentBooks.YearOfPublication,
			BookVolume:        currentBooks.BookVolume,
			Number:            currentBooks.Number,
		}
		respBook = append(respBook, book)
	}
	log.Debug("Package books, method ReadAll, final data", respBook)
	return &pb.ReadAllBooksResponse{Book: respBook}, nil
}

func (b BookServer) Read(ctx context.Context, request *pb.ReadBookRequest) (*pb.ReadBookResponse, error) {
	log.Debug("Package books, method Read. Start info id: ", request.Id)
	id := request.Id
	book, err := b.data.Read(id)
	log.Debug("Package books, method Read. Book info: ", book)
	respBook := &pb.Result{
		BookId:          book.BookId,
		NameOfBook:      book.NameOfBook,
		NameOfPublisher: book.NameOfPublisher,
	}
	if err != nil {
		s := status.Newf(codes.PermissionDenied, "got an error when tried to read book", err)
		errWithDetails, err := s.WithDetails(request)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return nil, errWithDetails.Err()
	}
	log.Debug("Package books, method Read. Final book info: ", respBook)
	return &pb.ReadBookResponse{Book: respBook}, nil
}

func (b BookServer) Add(ctx context.Context, request *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	log.Debug("Package books, method Add. Start info book: ", request.Book)
	requestBook := request.Book
	book := data.Book{
		BookId:            requestBook.BookId,
		AuthorId:          requestBook.AuthorId,
		PublisherId:       requestBook.PublisherId,
		NameOfBook:        requestBook.NameOfBook,
		YearOfPublication: requestBook.YearOfPublication,
		BookVolume:        requestBook.BookVolume,
		Number:            requestBook.Number,
	}
	id, err := b.data.Add(book)
	log.Debug("Package books, method Add. Get id from database: ", id)
	if err != nil {
		s := status.Newf(codes.PermissionDenied, "got an error when tried to add book", err)
		errWithDetails, err := s.WithDetails(request)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return nil, errWithDetails.Err()
	}
	log.Debug("Package books, method Add. Final id: ", id)
	return &pb.AddBookResponse{Id: id}, nil
}

func (b BookServer) Update(ctx context.Context, request *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	log.Debug("Package books, method Update. Start data id and price: ", request.Id, request.UnitPrice)
	requestId := request.Id
	requestPrice := request.UnitPrice
	err := b.data.Update(requestId, requestPrice)
	if err != nil {
		s := status.Newf(codes.PermissionDenied, "got an error when tried to update book", err)
		errWithDetails, err := s.WithDetails(request)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return nil, errWithDetails.Err()
	}
	log.Debug("Package books, method Update")
	return &pb.UpdateBookResponse{}, nil
}

func (b BookServer) Delete(ctx context.Context, request *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	log.Debug("Package books, method Delete. Start data id", request.Id)
	requestId := request.Id
	err := b.data.Delete(requestId)
	if err != nil {
		s := status.Newf(codes.PermissionDenied, "got an error when tried to delete book", err)
		errWithDetails, err := s.WithDetails(request)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return nil, errWithDetails.Err()
	}
	log.Debug("Package books, method Delete")
	return &pb.DeleteBookResponse{}, nil
}
