import { ref } from 'vue';
import dgram from 'dgram'; // Node.js UDP support, server-side only

// Dynamically import the CommonJS module for Noise
const NoiseStatePromise = import('noise-handshake').then(mod => mod.NoiseState);

export async function useNoiseConnection(isInitiator) {
  if (typeof window !== "undefined") {
    throw new Error("UDP logic can only run on the server-side");
  }

  const NoiseState = await NoiseStatePromise;

  // Create UDP socket (server-side only)
  const socket = dgram.createSocket('udp4');
  const noiseConnections = {}; // Store NoiseState for each peer (IP + port)
  const messageLog = ref([]); // Store received messages for display
  const handshakeCompleted = ref(false); // Track if the handshake is done
  const externalPort = ref(null); // Store discovered external port for hole punching

  // Get or create NoiseState for a given peer (IP + port)
  function getNoiseStateForPeer(remoteInfo) {
    const peerKey = `${remoteInfo.address}:${remoteInfo.port}`;
    if (!noiseConnections[peerKey]) {
      const noiseState = new NoiseState('XX', isInitiator);
      noiseState.initialise(null, null); // Initialize without static keys
      noiseConnections[peerKey] = noiseState;
    }
    return noiseConnections[peerKey];
  }

  // Function to send an encrypted message to a peer
  function sendMessage(message, remoteAddress, remotePort) {
    const noiseState = getNoiseStateForPeer({ address: remoteAddress, port: remotePort });
    const serializedMessage = Buffer.from(message);
    const encryptedMessage = new Uint8Array(serializedMessage.length + 16); // Add space for Noise encryption tag

    // Encrypt the message using the Noise protocol
    noiseState.writeMessage(encryptedMessage, serializedMessage);

    // Send the encrypted message over UDP
    socket.send(encryptedMessage, remotePort, remoteAddress, (err) => {
      if (err) {
        console.error('Error sending message:', err);
      } else {
        console.log(`Message sent to ${remoteAddress}:${remotePort}`);
      }
    });
  }

  // Function to handle incoming messages (decrypt them)
  function handleIncomingMessage(msg, rinfo) {
    const noiseState = getNoiseStateForPeer(rinfo);
    try {
      const decryptedMessage = new Uint8Array(msg.length - 16);
      noiseState.readMessage(decryptedMessage, msg);

      const messageText = decryptedMessage.toString();
      messageLog.value.push(`Received from ${rinfo.address}:${rinfo.port}: ${messageText}`);
      console.log(`Received message from ${rinfo.address}:${rinfo.port}: ${messageText}`);
    } catch (err) {
      console.error('Error decrypting message:', err);
    }
  }

  // Start the UDP socket and listen for incoming messages
  function startSocket(localPort) {
    return new Promise((resolve, reject) => {
      socket.on('message', handleIncomingMessage);
      socket.bind(localPort, () => {
        const address = socket.address(); // Get the internal address and port
        console.log(`Socket bound to ${address.address}:${address.port}`);
        resolve(address.port);
      });
    });
  }

  // Discover the external port for hole punching (via a STUN-like server)
  function discoverExternalPort(stunServerAddress, stunServerPort, localPort) {
    return new Promise((resolve, reject) => {
      // Send a "dummy" message to the STUN server to discover external port
      socket.send('DISCOVER', stunServerPort, stunServerAddress, (err) => {
        if (err) {
          return reject(err);
        }

        socket.on('message', (msg, rinfo) => {
          // STUN server responds with our external port
          const externalPortInfo = JSON.parse(msg.toString());
          externalPort.value = externalPortInfo.port;
          console.log(`Discovered external port: ${externalPort.value}`);
          resolve(externalPort.value);
        });
      });
    });
  }

  // Initiate a Noise handshake
  async function initiateHandshake(remoteAddress, remotePort, stunServerAddress, stunServerPort) {
    // Bind the socket to a known internal port
    const localPort = await startSocket(50000);

    // Use hole punching if necessary
    await discoverExternalPort(stunServerAddress, stunServerPort, localPort);

    // Proceed with Noise handshake
    const noiseState = new NoiseState('XX', isInitiator);
    noiseState.initialise(null, null);
    const handshakeMessage = new Uint8Array(32);

    noiseState.writeMessage(handshakeMessage);

    // Send the handshake message to the peer
    socket.send(handshakeMessage, remotePort, remoteAddress, (err) => {
      if (err) {
        console.error('Error sending handshake message:', err);
      } else {
        console.log(`Handshake initiated with ${remoteAddress}:${remotePort}`);
        handshakeCompleted.value = true;
      }
    });
  }

  return {
    messageLog,
    handshakeCompleted,
    sendMessage,
    initiateHandshake,
    startSocket,
    externalPort,
  };
}
