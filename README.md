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
- ![image](https://user-images.githubusercontent.com/96426723/146774752-d1bd0613-572a-4be2-85e6-36759a211f64.png)

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
![image](https://user-images.githubusercontent.com/96426723/146775513-70c16d1f-39f1-4bd4-b601-6b948eeb9461.png)

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
![image](https://user-images.githubusercontent.com/96426723/146776081-fc191725-e14e-4414-936f-6eb8677bddb7.png)

# 4. Network Up
저는 hyperledger fabric 2.1.1을 사용했습니다.

**ca-orgs1**
docker compose에는 앞서서 설정한 두개 organization과 organization에 해당하는 2개의 peer 정보 그리고 각각의 peer에 해당하는 4개의 couchdb를 장착하였습니다.
![image](https://user-images.githubusercontent.com/96426723/146777023-288eaa33-346d-4b05-b526-9b019d6e6331.png)

**orderer.example.com**
오더러를 보면 volumes에서 우리는 이미 생성한 genesis 블록이 있기 때문에 이를 마운팅해주기 위해 다음과 같이 적어줍니다.
![image](https://user-images.githubusercontent.com/96426723/146777119-1b3d307e-7816-4b1e-8bfc-9d5f01a745e4.png)
orderer_operations_listenaddress가 설정되어있지만 이는 크게 중요하지 중요하지 않다고 합니다.
![image](https://user-images.githubusercontent.com/96426723/146777207-0e88b197-31e9-404b-8a84-c4d4f5820a98.png)
7050포트는 orderer가 작동되는 곳이며 8443 포트는 listenport가 작동하는 곳입니다.

**couchdb0**
environment 변수는 아직 공백으로 두고 port는 5984로 열며 network는 같은 네트워크에서 작동하도록 gkqslek.
![image](https://user-images.githubusercontent.com/96426723/146777325-5bd7e2c1-c53a-4f9b-b145-8ef1c85d2765.png)
couchDB는 피어별로 포트만 다르게 해서 설정해줍니다. 6984:5984 , 7984:5984, 8984:5984

**peer0.org1.example.com**
base yaml에 적혀있는 환경변수와 연결시켜 줍니다.
![image](https://user-images.githubusercontent.com/96426723/146777433-1a58bfe4-3dae-4d62-a484-1c7b4480909e.png)
volumes은 기존에 있는 폴더를 연결해주도록 합니다.
![image](https://user-images.githubusercontent.com/96426723/146777488-9880de25-b156-4b96-b608-09fd412f0f5f.png)

**network up!**
터미널에서 docker-compose.yaml있는 폴더로 가서 네트워크를 실행해보도록 합니다.
$ docker-compose up -d
![image](https://user-images.githubusercontent.com/96426723/146777535-98c7307f-a443-4657-a93e-82213793b2ca.png)

PORTS와 NAMES를 확인해봅니다.
$ docker ps
![image](https://user-images.githubusercontent.com/96426723/146777616-1608883b-1298-46df-a6a1-4df0461d701b.png)

* 에러처리
Error response from daemon: path /home/jaewon/BasicNetwork-2.0/artifacts/channel is mounted on / but it is not a shared mount.
$ sudo mount --make-shared /
