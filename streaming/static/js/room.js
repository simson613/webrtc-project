function copyToClipboard(text) {
	if (window.clipboardData && window.clipboardData.setData) {
		clipboardData.setData("Text", text);
		return Swal.fire({
			position: 'top-end',
			text: "Copied",
			showConfirmButton: false,
			timer: 1000,
			width: '150px'
		})
	} else if (document.queryCommandSupported && document.queryCommandSupported("copy")) {
		var textarea = document.createElement("textarea");
		textarea.textContent = text;
		textarea.style.position = "fixed";
		document.body.appendChild(textarea);
		textarea.select();
		try {
			document.execCommand("copy");
			return Swal.fire({
				position: 'top-end',
				text: "Copied",
				showConfirmButton: false,
				timer: 1000,
				width: '150px'
			})
		} catch (ex) {
			console.warn("Copy to clipboard failed.", ex);
			return false;
		} finally {
			document.body.removeChild(textarea);
		}
	}
}

document.addEventListener('DOMContentLoaded', () => {
	(document.querySelectorAll('.notification .delete') || []).forEach(($delete) => {
		const $notification = $delete.parentNode;

		$delete.addEventListener('click', () => {
			$notification.style.display = 'none'
		});
	});
});

function connect(stream) {
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

	console.log("pc --> ",pc);

	stream.getTracks().forEach(track => pc.addTrack(track, stream))

	let ws = new WebSocket(RoomWebsocketAddr)
	pc.onicecandidate = e => {
		if (!e.candidate) {
			return
		}
		console.log("oncadidate and send cadidate to signaling server");

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
		while (pr.childElementCount > 3) {
			pr.lastChild.remove()
		}
		setTimeout(function () {
			connect(stream);
		}, 1000);
	}

	ws.onmessage = function (evt) {
		let msg = JSON.parse(evt.data)
		if (!msg) {
			return console.log('failed to parse msg')
		}

		console.log("receivce from signaling server --> ", msg.event);
		switch (msg.event) {
			case 'offer':
				let offer = JSON.parse(msg.data)
				if (!offer) {
					return console.log('failed to parse answer')
				}
				console.log("setRemoteDescription offer");
				console.log("create answer");
				console.log("setLocalDescription answer");
				console.log("send answer to signaling server");
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

navigator.mediaDevices.getUserMedia({
  video: {
    width: {
      max: 1280
    },
    height: {
      max: 720
    },
    frameRate: 60,
  },
  audio: {
    sampleSize: 16,
    channelCount: 2,
    echoCancellation: true
  }
})
.then(stream => {
  document.getElementById('localVideo').srcObject = stream;
  connect(stream);
}).catch(err => console.log(err))
