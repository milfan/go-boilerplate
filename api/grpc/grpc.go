package api_grpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpc_employee "github.com/milfan/go-boilerplate/internal/grpc/employee"
	pkg_grpc_employee "github.com/milfan/go-boilerplate/pkg/grpc/employee"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer  *grpc.Server
	netListener net.Listener
	logger      *logrus.Logger
}

func New(
	server *grpc.Server,
	netListener net.Listener,
	logger *logrus.Logger,
) *Server {

	employeeRpc := grpc_employee.New(logger)
	pkg_grpc_employee.RegisterEmployeeGrpcServer(server, employeeRpc)

	return &Server{
		grpcServer:  server,
		netListener: netListener,
		logger:      logger,
	}
}

func (s *Server) Start() {
	// Start serving
	logrus.Info("success to serve rpc ", s.netListener.Addr().String())

	if err := s.grpcServer.Serve(s.netListener); err != nil {
		s.logger.Fatal(fmt.Errorf("failed to serve: %v", err))
	}

	s.gracefullShutdown()
}

func (s *Server) gracefullShutdown() {
	// Handle graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("shutting down gracefully...")
		s.grpcServer.GracefulStop()
		s.netListener.Close()
	}()
}
