
# 프로젝트 목적 (WIP)

## 1. WebRTC를 활용한 스트리밍 구현
    
[WebRTC 주요 개념 정리](https://github.com/simson613/webrtc-project/tree/master/streaming/README.md) <br/>

이 프로젝트에서는 SUTN/TURN 서버를 사용하며, SFU 방식으로 signaling 서버를 구현


## 2. 마이크로서비스 구현
이 프로젝트에서는 auth, user, streaming, chat 이렇게 4가지 서비스로 나누어 구현하며, 아래 작성된 내용을 기반으로 구현됩니다.

### CQRS
CQRS 패턴을 사용하여 Command와 Query를 분리하였으며, <br/>
Command는 생성, 수정, 삭제를 비동기 이벤트를 사용하여 다른 서비스에 전달하여 Materialized View를 생성하는 등의 비즈니스 로직을 수행한다. 이 때 서비스간 트랜잭션이 달라 일시적으로 일관성을 잃지만 결과적으로는 일관성이 유지되는 형태로 일관성을 구현하였다. <br/>
Query는 미리 생성된 Materialized View를 사용하여 빠르게 응답한다.

## 3. 클린아키텍쳐 구현
![CleanArchitecture](https://github.com/user-attachments/assets/6b1953d9-4104-4aea-acf4-00417dd96f0d)

### Domain 
이 프로젝트에서 domain은 entity와 usecase로 구성됩니다.
<br/>
entity는 도메인에서 사용되는 struct 모델로 어느 계층에도 의존하지 않는 최상위 계층입니다.<br/>
usecase는 핵심 비즈니스 로직을 포함합니다.


### Interface Adapter
이 프로젝트에서 adapter는 controller, event, repository로 구성됩니다.<br/>
controller는 사용자의 요청의 진입점 역할을 합니다.<br/>
event는 kafka consumer와 producer가 정의되어 service간 통신의 진입점 역할을 합니다. <br/>
repository는 MariaDB, MongoDB의 진입점 역할을 합니다.
<br/>


## WIP
* 이벤트 장애 처리
* 채팅 서비스
* CI/CD
* 모니터링
* EKS 구축
* 회고 작성

<!-- ## 회고
202 -->