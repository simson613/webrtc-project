# WebRTC


## 목차
* [WebRTC(Web Real-Time Communications)란 ](#WebRTC(Web-Real-Time-Communications)란)
* [WebRTC 주요 개념 정리](#WebRTC-주요-개념-정리)


## WebRTC(Web Real-Time Communications)란
Web에서 플러그인 없이 실시간으로 동영상, 음성 등 실시간 멀티미디어 데이터를 전송할 수 있도록 설계된 API


## WebRTC 주요 개념 정리

**Signaling 서버**

피어간 연결을 설정하고 관리하는 중계 역할의 서버로 RTCPeerConnection 사용하여 피어간 연결을 관리하는 서버. 시그널링을 할 때 WebSocket 또는 XMLHttpRequest등을 사용할 수 있으며, 해당 프로젝트에서는 WebSocket을 사용

**SDP**

P2P 연결을 설명하는 표준으로 해상도, 형식, 코덱, 암호화 등의 멀티미디어 컨텐츠의 연결을 위한 정보가 포함됨

**ICE**

ICE는 브라우저가 피어를 통한 연결이 가능하도록 하게 해주는 프레임워크로 공인 IP가 없는경우, 방화벽 정책 등의 문제로 피어간 연결이 실패하는 경우 등의 상황에서 STUN, TURN 서버를 선택하여 사용합니다

**STUN 서버**

클라이언트가 자신의 공인 IP 주소와 NAT 타입 등을 얻을 수 있도록 도와주는 서버

![stun](https://github.com/user-attachments/assets/856cb33b-1b65-46c0-a61e-016a78e54d5e)


**TURN 서버** 

클라이언트 간의 직접 통신이 불가능하여 STUN 서버를 사용할 수 없는 경우에 데이터를 중계하는 서버

![turn](https://github.com/user-attachments/assets/15f32390-ed50-46c3-95c9-521ef75904f2)


   
## **WebRTC 구현방식**

**Mesh**

![mesh](https://github.com/user-attachments/assets/876eeb57-8583-40b8-93df-0e8bb57e782d)

클라이언트(peer)간의 연결을 위한 정보를 중계한다. </br>
따라서 처음 WebRTC가 peer간의 정보를 중계할 때만 서버에 부하가 발생하여 서버에 부하가 적다.</br>
클라이언트 Uplink, Downlink를 각각 하나씩 가지며 위 이미지처럼 클라이언트가 5명일 경우 link가 8개로 늘어나 연결된 클라이언트가 많아질수록 클라이언트의 부하가 커진다.</br> 
1:1 연결에 적합하다.


**SFU**

![sfu](https://github.com/user-attachments/assets/1a57e32f-a67e-40a3-acd5-028fae4b403b)


서버에서 미디어 트래픽을 중계하는 방식이다.</br>
클라이언트(peer)간의 연결이 아닌, 서버와 클라이언트 간의 연결한다.</br>
Mesh방식과 마찬가지로 클라이언트는 수신받을 클라이언트 만큼의 Downlink를 갖는다.</br>
하지만 송신은 서버로만하기 때문에 하나의 Uplink만 가져 Mesh방식에 비해 클라이언트의 부하가 적다.</br>
1:N 혹은 소규모 N:M 스트리밍 형식에 적합하다.



**MCU**

![mcu](https://github.com/user-attachments/assets/88acecf6-6bb1-4219-b895-5bafffac61df)

SFU 방식과 비슷하게 서버가 미디어 데이터를 중계하지만, MCU의 경우 미디어 데이터를 서버에서 혼합, 가공하여 클라이언트에 전송하는 방식이다. <br/>
따라서 클라이언트는 하나의 Uplink, Downlink를 가지므로 부하가 매우 적으나, 서버의 부하가 매우 커진다.
<br/><br/>


**출처**
<br/>
https://developer.mozilla.org/ko/docs/Web/API/WebRTC_API/Protocols