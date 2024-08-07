function connectStream() {
	let pc = new RTCPeerConnection({
		iceServers: [
			{
				'urls': 'stun:stun.webrtc-sim.kro.kr:3478',
			},
			{
				'urls': 'turn:webrtc-sim.kro.kr:3478',
				'username': 'simson',
				'credential': 'simson',
			}
		]
	})

	pc.ontrack = function (event) {
		if (event.track.kind === 'audio') {
			return
		}
		col = document.createElement("div")
		let el = document.createElement(event.track.kind)
		el.srcObject = event.streams[0]
		el.setAttribute("controls", "true")
		el.setAttribute("autoplay", "true")
		el.setAttribute("playsinline", "true")
    el.setAttribute("id", "localVideo")

		let playAttempt = setInterval(() => {
			el.play()
				.then(() => {
					clearInterval(playAttempt);
				})
				.catch(error => {
					console.log('unable to play the video, user has not interacted yet');
				});
		}, 3000);

		col.appendChild(el)
		document.getElementById('videos').appendChild(col)

		event.track.onmute = function (event) {
			el.play()
		}

		event.streams[0].onremovetrack = ({
			track
		}) => {
			if (el.parentNode) {
				el.parentNode.remove()
			}
		}
	}

	console.log("StreamWebsocketAddr --> ", StreamWebsocketAddr);
	let ws = new WebSocket(StreamWebsocketAddr)
	pc.onicecandidate = e => {
		if (!e.candidate) {
			return
		}

		console.log("send cadidate");
		ws.send(JSON.stringify({
			event: 'candidate',
			data: JSON.stringify(e.candidate)
		}))
	}

	ws.addEventListener('error', function (event) {
		console.log('error: ', event)
	})

	ws.onclose = function (evt) {
		console.log("websocket has closed")
		pc.close();
		pc = null;
		pr = document.getElementById('videos')
		while (pr.childElementCount > 2) {
			pr.lastChild.remove()
		}
		setTimeout(function () {
			connectStream();
		}, 1000);
	}

	ws.onmessage = function (evt) {
		let msg = JSON.parse(evt.data)
		if (!msg) {
			return console.log('failed to parse msg')
		}

		console.log("onmessage --> ", msg.event);
		switch (msg.event) {
			case 'offer':
				let offer = JSON.parse(msg.data)
				if (!offer) {
					return console.log('failed to parse answer')
				}
				pc.setRemoteDescription(offer)
				pc.createAnswer().then(answer => {
					pc.setLocalDescription(answer)
					ws.send(JSON.stringify({
						event: 'answer',
						data: JSON.stringify(answer)
					}))
				})
				return

			case 'candidate':
				let candidate = JSON.parse(msg.data)
				if (!candidate) {
					return console.log('failed to parse candidate')
				}

				pc.addIceCandidate(candidate)
		}
	}

	ws.onerror = function (evt) {
		console.log("error: " + evt.data)
	}
}

connectStream();