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

