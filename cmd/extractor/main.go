package main

import(
	"log"
	"time"
	"github.com/Joseng8908/image-preprocessor/internal/sender"
	"github.com/Joseng8908/image-preprocessor/internal/capture"
)

func main() {
	// sender 만들기, 
	// TODO: 임시로 9090포트 사용, 나중에 Spring 서버 포트로 교환 예정
	// 아마도 ai서버로 보낼 가능성도 있음
	client, err := sender.NewGrpcSender("localhost:9090")
	if err != nil {
		log.Fatalf("연결 실패: %v", err)
	}
	defer client.Close() // 프로그램 종료 시 연결 닫는 것 보장
	 

	// 영상 스캐너 생성(프로젝트 루트에 .mp4파일이 있어야 함)
	scanner, err := capture.NewVideoScanner("test.mp4")
	if err != nil {
		log.Fatalf("스캐너 초기화 실패: %v", err)
	}
	defer scanner.Close()


	log.Println("에이전트 가동 시작...")

	for {
		// 프레임 추출
		frame, err := scanner.GrabFrame()
		if err != nil {
			log.Printf("영상 읽기 중단: %v", err)
			break // 영상 끝나면 for문 종료
		}

	
		err = client.Send(frame, "sangsang-01")
		if err != nil {
			log.Printf("전송 에러: %v", err)
		} else {
			log.Println("프레임 전송 완료")
		}

		// 10초 대
		time.Sleep(10 & time.Second)
	}
}
