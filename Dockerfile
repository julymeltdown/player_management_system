# 베이스 이미지 설정 (Go 1.23 사용)
FROM golang:1.23

# 작업 디렉토리 설정
WORKDIR /app

# 호스트의 소스 코드를 컨테이너의 /app 디렉토리에 복사
COPY . /app

# Go 모듈 다운로드
RUN go mod download

# 테스트에 필요한 패키지 설치 (libpq-dev)
RUN apt-get update && apt-get install -y libpq-dev

# 테스트 실행
CMD ["go", "test", "./...", "-v"]