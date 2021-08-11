# golang_ECDSA

비트코인 주소 체계를 기반으로 임의의 데이터의 전자서명 검증로직 개발

## 실행방법

$ go mod tidy
$ go mod init
$ go build

위 명령어를 입력 후, build된, "ecdsa" 파일을 실행해서, 사용법 설명서에 맞추어 0, 1, 2, 3, 4를 입력하여 실행한다. (영상 참조)

## 실행영상



## 로직설명

### 비트코인 주소생성

1. ECDSA 기반 keypair 생성 (Private Key & Public Key)

2. Bitcoin Address 생성[1/3] (Public key -> SHA256 -> RIPEMD160 = Public Key Hash)

3. Bitcoin Address 생성[2/3] (Public Key -> SHA256 -> SHA256 = Checksum)

4. Bitcoin Address 생성[3/3] (Version + Public Key Hash + CheckSum -> Base58Encode = Bitcoin Address)

### Signing & Verification

1. Data를 ECDSA SIGN 을 한다. (Private Key + Data = Signature)

2. 생성된 Signature값과, 해싱된 데이터 값 그리고 퍼블릭키를 통해 검증 수행. (Public Key + SignHash + Signature = True OR False)
