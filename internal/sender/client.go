package sender

// 서버로 데이터를 쏘는 로직

import (
	"context"
	"time"

	"github.com/Joseng8908/image-preprocessor/api/gen/seat"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc"
)

// 기본 정보들
// grpc 통신에 필요한 연결 정보를 가지는 데이터
type GrpcSender struct {
	client seat.SeatAnalyzerClient
	conn *grpc.ClientConn
}

// 보낼 새로운 데이터 셋 만드는 로직
func NewGrpcSender(address string) (*GrpcSender, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// 오류면 바로 리턴
	if err != nil {
		return nil, err
	}
	// 오류가 아니면 정상 반환
	// conn, 즉 연결 통로를 이요하여 gRPC클라이언트 객체를 생성해서 반환
	return &GrpcSender{
		client: seat.NewSeatAnalyzerClient(conn),
		conn: conn,
	}, nil
}

// 실제 보내는 로직
// 위에서 생성된 객체를 이용해 
func (s *GrpcSender) Send(image []byte, id string) error{
	// 5초 안에 응답 없으면 포기
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	// cancel() 보장, memory leak 방지
	defer cancel()

	// 위에서 만든 grpc객체에(택배 상자) 아래의 데이터들을 넣어서 보냄
	// .proto파일에서 정의한 구조대로 보내는 거임
	_, err := s.client.AnalyzeFrame(ctx, &seat.FrameRequest{
		DeviceId: id,
		ImagePayload: image,
		Timestamp: time.Now().Unix(),
	})
	return err
	
}

func (s *GrpcSender) Close() {
	if s.conn != nil {
		s.conn.Close()
	}
}

