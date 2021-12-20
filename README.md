# BIoT-data control
A blockchain-based IoT system for management the personal data from smart watch

I studied blockchain from https://github.com/adhavpavan/BasicNetwork-2.0.git
And get repository from above URL.
Latest Docker and Visual Studio Code was used to manage this works.


# 1. Version check

- node js: 12.16.0
- go lang: 1.17
- python: 2.7.0
- hyperledger fabric: 2.1.1
- you can download the hyperledger 'bin' folder through cli
- curl -sSL https://bit.ly/2ysbOFE | bash -s -- 2.1.1 1.4.6 0.4.18

# 2. Set environment variables and PATH
**Is 'GO' installed?**
- export GOPATH=$HOME/go


**Is 'Hyperledger' installed?**
- export PATH=$PATH:$GOPATH/bin
- export PATH=$PATH:/home/jaewon/fabric-samples/bin

**Change the permission**
- sudo mount --make-shared /
- sudo chmod -R 755 ./

# 3. Create organization and crypto materials
3-1.  Configtx File: Create Crypto Materials

- use 'cyrpto-config.yaml' in the 'artifacts/channel/crypto-config.yaml'

<img src="https://user-images.githubusercontent.com/96426723/146774752-d1bd0613-572a-4be2-85e6-36759a211f64.png" width="40%">


**cryto-config.yaml** 파일에는 네트워크를 구성하는 org와 peer의 정보가 담겨있습니다.
- Oragnization 1  template count 가 2개라는 건 피어(peer)가 두개라 certifacate를 2개 만든다는 의미입니다.
- Users count가 1개라는건 default user을 의미합니다.

**create-artifacts.sh**에는 여러가지 CLI 명령어가 입력되어 있습니다.
- **CA**: 하이퍼레저 패브릭의 에플리케이션의 일종이며 허락받은 참여자만 블록체인 네트워크에 참가가 가능함을 보여줍니다. (블록체인용)
- **MSP**: 이러한 정보들은 모여서 MSP (Membership Service Provider)를 구성합니다. MSP를 통해서 채널의 관리 권한이나 접근 권한을 관리할 수 있습니다.
- 간단히 말해서 특정 채널에 피어가 참여하고 체인코드를 호출하기 위해 필요한 권한을 여기서 관리합니다.
- MSP도 마찬가지로 PKI (공개키 구조 기반)을 사용합니다.
- 여기서, tlsca는 보안통신용을 말합니다.

3-2. Configtx File: Create Network Artifacts

앞서서 Organization을 구성했다면 이제는 network를 구성해볼 차례입니다.

<img src="https://user-images.githubusercontent.com/96426723/146775513-70c16d1f-39f1-4bd4-b601-6b948eeb9461.png" width="40%">

여기서는 genesisblock과 같은 네트워킹 요소를 설정할 수 있습니다.

**Oragnization 설정**
- Oraganization은 crypto-config 파일에서 2개의 organization을 만들었기 때문에 "OrdererOrg", "Org1", "Org2" 로 구성되어있습니다.
- Name: 오더러org의 이름을 적어주면 됩니다.
- ID: 여기서는 오더러 MSP 불러올 때의 ID를 의미합니다.
- MSPDir: crypto-config 파일에서 생성된 orderer 의 MSP 폴더 위치를 적어줍니다.
- Policies: policies에는 종류가 2가지 있습니다. (공식홈페이지를 참고합시다.)

create-artifacts.sh 파일에는 OrdererGenesis와 BasicChannel을 만들 수 있는 cli 명령어가 있습니다.

두 명령어를 각각 치면 폴더에 genesis.block 파일과 mychannel. Tx 트랜잭션 파일 그리고 OrgMSPanchors.tx 파일이 생기는걸 확인할 수 있습니다.

permission denied error가 뜬다면, "chmod -R 0755 /home/bin/configtxgen" 명령어를 입력합시다.

<img src="https://user-images.githubusercontent.com/96426723/146776081-fc191725-e14e-4414-936f-6eb8677bddb7.png" width="80%">

# 4. Network Up
저는 hyperledger fabric 2.1.1을 사용했습니다.

**ca-orgs1**
docker compose에는 앞서서 설정한 두개 organization과 organization에 해당하는 2개의 peer 정보 그리고 각각의 peer에 해당하는 4개의 couchdb를 장착하였습니다.


**orderer.example.com**
오더러를 보면 volumes에서 우리는 이미 생성한 genesis 블록이 있기 때문에 이를 마운팅해주기 위해 다음과 같이 적어줍니다.

<img src="https://user-images.githubusercontent.com/96426723/146777119-1b3d307e-7816-4b1e-8bfc-9d5f01a745e4.png" width="80%">

orderer_operations_listenaddress가 설정되어있지만 이는 크게 중요하지 중요하지 않다고 합니다.

<img src="https://user-images.githubusercontent.com/96426723/146777207-0e88b197-31e9-404b-8a84-c4d4f5820a98.png" width="50%">

7050포트는 orderer가 작동되는 곳이며 8443 포트는 listenport가 작동하는 곳입니다.

**couchdb0**
environment 변수는 아직 공백으로 두고 port는 5984로 열며 network는 같은 네트워크에서 작동하도록 합니다.

<img src="https://user-images.githubusercontent.com/96426723/146777325-5bd7e2c1-c53a-4f9b-b145-8ef1c85d2765.png" width="30%">

couchDB는 피어별로 포트만 다르게 해서 설정해줍니다. 6984:5984 , 7984:5984, 8984:5984

**peer0.org1.example.com**
base yaml에 적혀있는 환경변수와 연결시켜 줍니다.

<img src="https://user-images.githubusercontent.com/96426723/146777433-1a58bfe4-3dae-4d62-a484-1c7b4480909e.png" width="40%">

volumes은 기존에 있는 폴더를 연결해주도록 합니다.

<img src="https://user-images.githubusercontent.com/96426723/146777488-9880de25-b156-4b96-b608-09fd412f0f5f.png" width="80%">


**network up!**
터미널에서 docker-compose.yaml있는 폴더로 가서 네트워크를 실행해보도록 합니다.
$ docker-compose up -d

<img src="https://user-images.githubusercontent.com/96426723/146777535-98c7307f-a443-4657-a93e-82213793b2ca.png" width="20%">

PORTS와 NAMES를 확인해봅니다.
$ docker ps

<img src="https://user-images.githubusercontent.com/96426723/146777616-1608883b-1298-46df-a6a1-4df0461d701b.png" width="100%">

* 에러처리

Error response from daemon: path /home/jaewon/BasicNetwork-2.0/artifacts/channel is mounted on / but it is not a shared mount.

$ sudo mount --make-shared /

# 5. Channel up

5.1 Create channel

**-o**: 오더러를 의미합니다. 오더러의 포트를 써줍니다. docker ps 를 쳤을 때 oderer의 포트는 다음과 같습니다.

 0.0.0.0:7050->7050/tcp, :::7050->7050/tcp, 0.0.0.0:8443->8443/tcp, :::8443->8443/tcp
 
 orderer의 주소는 7050이므로 localhost:7050dmf 여기에 적어줍니다.
 
 **--ordererTLSHostnameOverride**는 orderer.example.com으로 작성해줍니다.
 
 **-f**는 mychannel.tx의 파일이 있는 곳의 경로를 쳐줍시다.
 
 그리고 그 뒤에는 결과로 나올 block의 경로를 적어줍니다. channel-artifacts 폴더에 블록을 생성해주도록 합니다.
 
**--tls**는 통신 여부를 묻는 것입니다. -> TRUE

**--cafile**는 인증서의 위치를 묻는 것입니다. ->  ORDERER_CA=${PWD}/artifacts/channel/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

<img src="https://user-images.githubusercontent.com/96426723/146779736-c2e9d0a2-f5b1-4584-a83e-57dfa393954b.png" width="80%">

5.2 joinChannel

createChannel()에서 만든 ./artifacts/channel/${CHANNEL_NAME}.tx에다가 미쳐 추가하지 못한 Peer를 추가해줍니다.

<img src="https://user-images.githubusercontent.com/96426723/146779789-bc9eff3f-dd18-4dd7-a58c-32167375b3ab.png" width="45%">

5.3 updateAnchorPeers

peer0가 각 organization에서 anchor peer을 맞고 있기 때문에 update 해주도록 합니다.

<img src="https://user-images.githubusercontent.com/96426723/146779803-7d04bf6a-6a96-43a9-8fd5-5125727a7b0e.png" width="75%">

# 6. Chaincode up

deployChaincode.sh에서 체인코드를 설치하고 불러오고 지지고 볶을 수 있습니다.

6.1 presetup

체인코드가 적혀있는 go 파일을 복사해서 붙여넣기를 해주고  잘 작동할지 on go mod vendor로 확인합니다. go mod vendor로 모듈관리가 가능합니다.

<img src="https://user-images.githubusercontent.com/96426723/146780272-6961d40f-f308-4d79-8373-701ea035f73c.png" width="25%">

※ go sumdb오류가 나는 경우

export GOSUMDB=off

go의 위치를 /home/jaewon으로 해서 다시 설치해줍니다.

6.2 packageChaincode

fabcar.tar.gz파일이 생기는걸 확인가능합니다.
그 후, 도커 컨테이너에 피어들이 작동하도록 설정합니다.

cd artifacts

docker-compose up -d

<img src="https://user-images.githubusercontent.com/96426723/146780429-2a41c0d4-dd46-4add-806b-2bff279352ba.png" width="75%">


6.3 installChaincode

모든 peer에 chaincode를 설치합니다. 앞서 생성한 체인코드가 각 피어에 설치됩니다.

<img src="ttps://user-images.githubusercontent.com/96426723/146780498-c373a427-0372-4ae0-b24c-c73eeaa091c6.png" width="75%">

6.4 queryInstalled & approveForMyOrg1

queryInstalled 함수를 통해 잘 설치되었는지 확인할 수 있습니다.

<img src="https://user-images.githubusercontent.com/96426723/146780604-deada050-72f9-4bd4-9987-59f5ceacd5ef.png" width="85%">

6.5 checkCommitReadyness

Organization 1의 MSP가 승인되었음을 알 수 있습니다. Organization 2의 MSP도 이제 설정해주도록합니다.

<img src="https://user-images.githubusercontent.com/96426723/146780703-9bdee530-b927-4f37-b82c-ba5c0e9017f6.png" width="75%">

<img src="https://user-images.githubusercontent.com/96426723/146780759-c37bcd85-de12-4ec9-a1f8-0b196ace3cc0.png" width="35%">

6.6 commitChaincodeDefination

localhost 7051, localhost 9051에 대해서 chaincode가 설치외어있음을 확인할 수 있습니다.

<img src="https://user-images.githubusercontent.com/96426723/146780831-43748a89-32d0-4270-9804-98bb4eb53d03.png" width="100%">

6.7 chaincodeInvokeInit

체인코드를 invoke하는 코드입니다.

docker ps를 보면 각 peer에 해당하는 status 뒤에 체인코드가 등록되었음을 확인할 수 있습니다.

<img src="https://user-images.githubusercontent.com/96426723/146780931-814d6056-feeb-4081-9482-ec7b14ab242c.png" width="55%">

fabcar.go에 적혀있는 initLedger code chaincode가 시작되고 initLedger을 발동시켜서 데이터베이스에 저장시키는걸 확인할 수 있습니다.

그 후, 크롬을 키고 'localhost:5984/_utils'에 접속해서 data를 확인해봅시다.

peer0.org1.example.com에 연결되어 있는 couchDB 0가 localhost:5984입니다.

peer0Org1에 chaincode를 invoke했기 때문에 couchDB 0에서 해당 자료를 확인할 수 있습니다.

![image](https://user-images.githubusercontent.com/96426723/146781149-441531d1-3169-4096-92ec-76a0bdbe3e29.png)

당연히, 모든 peer에서 distributed ledger에 담긴 정보를 확인할 수 있어야 되기 때문에 5984가 아닌 6984, 7984, 8984에서도 정보를 알 수 있습니다.




