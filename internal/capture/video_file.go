package capture

import (
	"fmt"
	"gocv.io/x/gocv"
)

// 스캐너 객체 정의 
type VideoScanner struct {
	video *gocv.VideoCapture
	img gocv.Mat
}

// NewVideoScanner: 영상 파일 경로를 받아 객체 생성
func NewVideoScanner(filePath string) (*VideoScanner, error) {
	// TODO: 영상 파일 경로로 바꾸기
	video, err := gocv.OpenVideoCapture(filePath)
	// 에러처리
	if err != nil {
		return nil, fmt.Errorf("영상 파일을 열 수 없습니다: %v", err)
	}

	// 객체 반환
	return &VideoScanner{
		video: video, 
		img: gocv.NewMat(),
	}, nil
}

// GrabFrame: 영상에서 프레임 하나를 읽어 JPEG 바이트로 반환
func(v *VideoScanner) GrabFrame() ([]byte, error) {
	if ok := v.video.Read(&v.img); !ok {
		return nil, fmt.Errorf("영상이 끝났거나 읽기에 실패했습니다")
	}
	if v.img.Empty() {
		return nil, fmt.Errorf("빈 프레임입니다")
	}

	// gRPC로 보내기 위해 Mat 객체를 JPEG 바이너리로 인코딩
	buf, err := gocv.IMEncode(".jpg", v.img)
	if err != nil {
		return nil, fmt.Errorf("이미지 인코딩 실패: %v", err)
	}

	// 인코딩 된 버퍼의 바이트 반환
	return buf.GetBytes(), nil
}

// 메모리 누수 방지
func(v *VideoScanner) Close() {
	v.video.Close()
	v.img.Close()
}
