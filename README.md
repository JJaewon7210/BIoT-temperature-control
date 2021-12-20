# BIoT-data control
A blockchain-based IoT system for management the personal data from smart watch

I studied blockchain from https://github.com/adhavpavan/BasicNetwork-2.0.git
And get repository from above URL.
Latest Docker and Visual Studio Code was used to manage this works.


# 1. version check

  - node js: 12.16.0
  - go lang: 1.17
  - python: 2.7.0
  - hyperledger fabric: 2.1.1
  - you can download the hyperledger 'bin' folder through cli
  - curl -sSL https://bit.ly/2ysbOFE | bash -s -- 2.1.1 1.4.6 0.4.18

# 2. Set environment variables and PATH
   **Is it 'GO' installed?**
  - export GOPATH=$HOME/go


   **is it 'Hyperledger' installed?**
  - export PATH=$PATH:$GOPATH/bin
  - export PATH=$PATH:/home/jaewon/fabric-samples/bin

   **change the permission**
  - sudo mount --make-shared /
  - sudo chmod -R 755 ./

