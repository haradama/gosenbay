import { startCamera, fitCanvasToVideo } from "./camera/camera";
import { CanvasRecorder } from "./recorder/recorder";
import { drawQrOverlay } from "./qr/qrcode";
import { enableSensors, getSensorSnapshot } from "./sensors/sensor-source";
import { encodeSenbay, loadSenbayWasm } from "./wasm/senbay";

export async function bootApp(): Promise<void> {
  await loadSenbayWasm();

  const video = document.querySelector<HTMLVideoElement>("#cameraVideo");
  const canvas = document.querySelector<HTMLCanvasElement>("#outputCanvas");
  const payloadView = document.querySelector<HTMLPreElement>("#payloadView");
  const encodedView = document.querySelector<HTMLPreElement>("#encodedView");

  const startCameraButton =
    document.querySelector<HTMLButtonElement>("#startCamera");
  const enableSensorsButton =
    document.querySelector<HTMLButtonElement>("#enableSensors");
  const startRecordingButton =
    document.querySelector<HTMLButtonElement>("#startRecording");
  const stopRecordingButton =
    document.querySelector<HTMLButtonElement>("#stopRecording");
  const downloadLink =
    document.querySelector<HTMLAnchorElement>("#downloadLink");

  if (
    !video ||
    !canvas ||
    !payloadView ||
    !encodedView ||
    !startCameraButton ||
    !enableSensorsButton ||
    !startRecordingButton ||
    !stopRecordingButton ||
    !downloadLink
  ) {
    throw new Error("Required DOM elements were not found");
  }

  const ctx = canvas.getContext("2d");
  if (!ctx) throw new Error("Canvas 2D context is not available");

  const recorder = new CanvasRecorder();

  let cameraStarted = false;
  let lastQrUpdate = 0;
  let lastEncoded = "";

  startCameraButton.addEventListener("click", async () => {
    await startCamera(video);
    cameraStarted = true;
    render();
  });

  enableSensorsButton.addEventListener("click", async () => {
    await enableSensors();
  });

  startRecordingButton.addEventListener("click", () => {
    recorder.start(canvas, 30);
    startRecordingButton.disabled = true;
    stopRecordingButton.disabled = false;
  });

  stopRecordingButton.addEventListener("click", async () => {
    const blob = await recorder.stop();
    const url = URL.createObjectURL(blob);

    downloadLink.href = url;
    downloadLink.download = `senbay-${Date.now()}.webm`;
    downloadLink.hidden = false;
    downloadLink.textContent = "Download Video";

    startRecordingButton.disabled = false;
    stopRecordingButton.disabled = true;
  });

  async function render(): Promise<void> {
    if (!cameraStarted) return;

    fitCanvasToVideo(canvas, video);
    ctx.drawImage(video, 0, 0, canvas.width, canvas.height);

    const now = performance.now();

    if (now - lastQrUpdate > 100) {
      const payload = getSensorSnapshot();
      lastEncoded = encodeSenbay(payload, true);

      payloadView.textContent = JSON.stringify(payload, null, 2);
      encodedView.textContent = lastEncoded;

      lastQrUpdate = now;
    }

    if (lastEncoded) {
      await drawQrOverlay(ctx, lastEncoded, 16, 16, 180);
    }

    requestAnimationFrame(render);
  }
}
