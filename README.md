# golang_ECDSA

비트코인 주소 체계를 기반으로 임의의 데이터의 전자서명 검증로직 개발


1. ECDSA 기반 keypair 생성 (Private Key & Public Key)

2. Bitcoin Address 생성[1/3] (Public key -> SHA256 -> RIPEMD160 = Public Key Hash)

3. Bitcoin Address 생성[2/3] (Public Key -> SHA256 -> SHA256 = Checksum)

4. Bitcoin Address 생성[3/3] (Version + Public Key Hash + CheckSum -> Base58Encode = Bitcoin Address)
