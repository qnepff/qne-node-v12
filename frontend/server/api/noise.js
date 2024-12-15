import { useNoiseConnection } from '../composables/useNoiseConnection';

export default defineEventHandler(async (event) => {
  const { initiateHandshake, messageLog } = await useNoiseConnection(true); // Initiator role

  // Example: Remote peer details and STUN server for hole punching
  const remoteAddress = '203.0.113.5';
  const remotePort = 12345;
  const stunServerAddress = 'stun.example.com';
  const stunServerPort = 3478;

  // Call the server-side UDP logic to initiate the Noise handshake
  await initiateHandshake(remoteAddress, remotePort, stunServerAddress, stunServerPort);

  return {
    status: 'Handshake initiated',
    logs: messageLog.value,
  };
});
