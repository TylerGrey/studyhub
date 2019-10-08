# Study Hub API server

## Issues
1. go-bindata version
`go get -u github.com/jteeuwen/go-bindata/...`로 모듈 다운로드 시, 3.0.7 버전으로 받게 돼 `MustAssets`가 생성이 안됨.  
`go get -u github.com/jteeuwen/go-bindata/...@master`로 마스터 버전으로 설치 필수  
[github issues](https://github.com/jteeuwen/go-bindata/issues/13)