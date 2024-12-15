import { ref } from 'vue'

export const useWebRTC = () => {
  const localStream = ref<MediaStream | null>(null)
  const remoteStream = ref<MediaStream | null>(null)
  const peerConnection = ref<RTCPeerConnection | null>(null)
  const ws = ref<WebSocket | null>(null)
  const error = ref<string>('')
  const isCallEnabled = ref(false)

  const config = {
    iceServers: [
      { urls: 'stun:stun.l.google.com:19302' }
    ]
  }

  const connectSignaling = () => {
    const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
    const wsUrl = `${protocol}//${window.location.host}/ws`
    ws.value = new WebSocket(wsUrl)

    ws.value.onopen = () => {
      console.log('Connected to signaling server')
    }

    ws.value.onmessage = async (event) => {
      const message = JSON.parse(event.data)
      
      switch (message.type) {
        case 'offer':
          await handleOffer(message)
          break
        case 'answer':
          await handleAnswer(message)
          break
        case 'ice-candidate':
          await handleIceCandidate(message)
          break
      }
    }
  }

  const createPeerConnection = async () => {
    peerConnection.value = new RTCPeerConnection(config)
    
    // Add local stream tracks to peer connection
    if (localStream.value) {
      localStream.value.getTracks().forEach(track => {
        peerConnection.value?.addTrack(track, localStream.value!)
      })
    }

    // Handle incoming streams
    peerConnection.value.ontrack = (event) => {
      remoteStream.value = event.streams[0]
    }

    // Handle ICE candidates
    peerConnection.value.onicecandidate = (event) => {
      if (event.candidate && ws.value) {
        ws.value.send(JSON.stringify({
          type: 'ice-candidate',
          candidate: event.candidate
        }))
      }
    }

    return peerConnection.value
  }

  const init = async () => {
    try {
      console.log('Requesting camera and microphone permissions...')
      
      const constraints = {
        audio: {
          echoCancellation: true,
          noiseSuppression: true,
          autoGainControl: true
        },
        video: {
          width: { ideal: 1280 },
          height: { ideal: 720 },
          facingMode: 'user'
        }
      }

      localStream.value = await navigator.mediaDevices.getUserMedia(constraints)
      console.log('Got media stream:', localStream.value.getTracks().map(track => track.kind))
      
      // Connect to signaling server
      connectSignaling()
      isCallEnabled.value = true
    } catch (e: any) {
      console.error('Error getting user media:', e)
      if (e.name === 'NotAllowedError' || e.name === 'PermissionDeniedError') {
        error.value = 'Please allow camera and microphone access to use this app.'
      } else if (e.name === 'NotFoundError' || e.name === 'DevicesNotFoundError') {
        error.value = 'No camera or microphone found. Please connect a device and try again.'
      } else if (e.name === 'NotReadableError' || e.name === 'TrackStartError') {
        error.value = 'Your camera or microphone may be in use by another application.'
      } else {
        error.value = 'Error accessing camera and microphone: ' + e.message
      }
    }
  }

  const startCall = async () => {
    if (!peerConnection.value) {
      await createPeerConnection()
    }

    try {
      const offer = await peerConnection.value?.createOffer()
      await peerConnection.value?.setLocalDescription(offer)
      
      ws.value?.send(JSON.stringify({
        type: 'offer',
        offer: offer
      }))
    } catch (e) {
      console.error('Error creating offer:', e)
    }
  }

  const handleOffer = async (message: any) => {
    if (!peerConnection.value) {
      await createPeerConnection()
    }

    try {
      await peerConnection.value?.setRemoteDescription(new RTCSessionDescription(message.offer))
      const answer = await peerConnection.value?.createAnswer()
      await peerConnection.value?.setLocalDescription(answer)
      
      ws.value?.send(JSON.stringify({
        type: 'answer',
        answer: answer
      }))
    } catch (e) {
      console.error('Error handling offer:', e)
    }
  }

  const handleAnswer = async (message: any) => {
    try {
      await peerConnection.value?.setRemoteDescription(new RTCSessionDescription(message.answer))
    } catch (e) {
      console.error('Error handling answer:', e)
    }
  }

  const handleIceCandidate = async (message: any) => {
    try {
      if (message.candidate) {
        await peerConnection.value?.addIceCandidate(new RTCIceCandidate(message.candidate))
      }
    } catch (e) {
      console.error('Error handling ICE candidate:', e)
    }
  }

  const cleanup = () => {
    if (localStream.value) {
      localStream.value.getTracks().forEach(track => track.stop())
    }
    if (peerConnection.value) {
      peerConnection.value.close()
    }
    if (ws.value) {
      ws.value.close()
    }
  }

  return {
    localStream,
    remoteStream,
    error,
    isCallEnabled,
    init,
    startCall,
    cleanup
  }
}
