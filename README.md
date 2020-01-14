# Study Hub API server

## V1
### API
#### 회원
- [ ] SNS 회원 가입
- [ ] 로그인

#### 허브
- [x] 허브 리스트 조회
- [x] 허브 조회
- [x] 허브 등록
- [x] 정보 수정

#### 기타
- [ ] 이미지 저장

### Mobile
- [ ] google map or naver API
- [ ] 회원가입
- [ ] 로그인
- [ ] 허브 상세
- [ ] 리뷰 작성

### Backoffice
- [ ] 허브 관리

## Issues
1. go-bindata version
`go get -u github.com/jteeuwen/go-bindata/...`로 모듈 다운로드 시, 3.0.7 버전으로 받게 돼 `MustAssets`가 생성이 안됨.  
`go get -u github.com/jteeuwen/go-bindata/...@master`로 마스터 버전으로 설치 필수  
[github issues](https://github.com/jteeuwen/go-bindata/issues/13)

2. Go env 변경
`go env -w [key]=[value]`

3. dataloader
4. cusor based pagination
