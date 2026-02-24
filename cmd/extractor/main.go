package main

import(
	"log"
	"time"
	"github.com/Joseng8908/image-preprocessor/internal/sender"
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
	 
	log.Println("에이전트 가동 시작...")

	for {
		// 캡처 로직 자리
		// 지금은 이시 텍스트를 이미지 바이너리인 척 해서 보내기
		// 아직 캡처 로직 안만듦
		// TODO: 캡처 로직 만들고 부착
		dummyImage := []byte("fake_image_binary_data")

		// 서버로 전송 시도
		err := client.Send(dummyImage, "classroom-A")
		if err != nil {
			log.Printf("전송 에러: %v", err)
		} else {
			log.Println("프레임 전송 완료")
		}

		// 10초 대
		time.Sleep(10 & time.Second)
	}
}
