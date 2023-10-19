# froxy

HTTP, HTTPS 상에서 다양한 기능을 제공하는 프록시 서버 애플리케이션.
다음 기능들을 커맨드와 옵션 플래그 한 두 개로 간단히 사용할 수 있다.
* 포워드 프록시
* 리버스 프록시
* 로드 밸런서

## Installation
```
  go install github.com/lifthus/froxy/cmd/froxy@latest
```


## Usage

### 1. 포워드 프록시

```
froxy
forxy -p 8542
```

-p 옵션을 통해 포트를 지정할 수 있으며, 지정하지 않으면 8542 포트가 기본으로 설정된다. 이렇게 **포트만 설정된 경우 froxy는 포워드 프록시로 동작**한다.
froxy는 수신한 모든 HTTP 요청을 각 요청의 타겟 URL로 포워딩하고 그 응답을 다시 클라이언트에게 되돌려 준다.
HTTPS의 경우 클라이언트로부터 CONNECT 메소드를 수신했을 때 클라이언트와 오리진 서버 사이의 TCP 터널을 연결해 제공한다.

개인용 서버에서 간단히 실행 후 운영체제에서 해당 서버로 프록시 설정만 해주면 개인용 HTTP 프록시 서버로 사용하는데 있어 충분하다.

### 2. 리버스 프록시
   
```
froxy -t 12.123.123.12
froxy -p 8542 -t example.com

froxy -t http://example.com -cert cert.pem -key key.pem
```

포트 외에 -t 옵션을 통해 **타겟 URL을 지정해주면 froxy는 리버스 프록시로 동작**한다.
froxy는 그것이 실행된 서버의 지정된 포트로의 모든 HTTP 요청을 지정된 타겟으로 포워딩하고 응답을 클라이언트에게 되돌려준다.
-cert, -key 옵션을 통해 자격증명과 키 파일을 제공하면 HTTPS 서버로 동작한다.

간단한 리버스 프록시로 사용하거나, 개발 서버가 따로 배포되어 있을 때 로컬에서 쿠키를 사용하기 위해서, 원격 서버에 대한 로컬 리버스 프록시로 활용하면 쿠키가 포트를 구분하지 않는 점을 이용해 로컬에서 원격 서버의 쿠키를 활용할 수 있다.

### 3. 로드 밸런서
```
froxy -lb lblist
froxy -p 8542 -lb lblist

froxy -t http://example.com -cert cert.pem -lb lb
```

-lb 옵션을 통해 개행으로 구분된 **URL 리스트를 제공하면 froxy는 해당 URL들에 대한 라운드-로빈식 로드 밸런서로 동작**한다.
-cert, -key 옵션을 통해 자격증명과 키 파일을 제공하면 HTTPS 서버로 동작한다.

분산 서버를 손쉽게 테스트 해보는 수준에 적합하다.
