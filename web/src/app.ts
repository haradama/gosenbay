import { startCamera, fitCanvasToVideo } from "./camera/camera";
import { CanvasRecorder } from "./recorder/recorder";
import { drawQrOverlay } from "./qr/qrcode";
import { enableSensors, getSensorSnapshot } from "./sensors/sensor-source";
import { encodeSenbay, loadSenbayWasm } from "./wasm/senbay";

type Point = {
  x: number;
  y: number;
};

type Rect = {
  x: number;
  y: number;
  size: number;
};

const QR_SIZE = 180;
const QR_MARGIN = 16;
const QR_UPDATE_INTERVAL_MS = 100;

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
  const recordingStatus =
    document.querySelector<HTMLSpanElement>("#recordingStatus");

  if (
    !video ||
    !canvas ||
    !payloadView ||
    !encodedView ||
    !startCameraButton ||
    !enableSensorsButton ||
    !startRecordingButton ||
    !stopRecordingButton ||
    !downloadLink ||
    !recordingStatus
  ) {
    throw new Error("Required DOM elements were not found");
  }

  const ctx = canvas.getContext("2d");
  if (!ctx) throw new Error("Canvas 2D context is not available");

  const recorder = new CanvasRecorder();

  let cameraStarted = false;
  let isRecording = false;
  let recordingStartedAt = 0;

  let lastQrUpdate = 0;
  let lastEncoded = "";

  let qrRect: Rect = {
    x: QR_MARGIN,
    y: QR_MARGIN,
    size: QR_SIZE,
  };

  let isDraggingQr = false;
  let dragOffset: Point = { x: 0, y: 0 };

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

    isRecording = true;
    recordingStartedAt = Date.now();

    recordingStatus.hidden = false;
    startRecordingButton.disabled = true;
    stopRecordingButton.disabled = false;
    downloadLink.hidden = true;
  });

  stopRecordingButton.addEventListener("click", async () => {
    isRecording = false;
    recordingStatus.hidden = true;

    const blob = await recorder.stop();
    const url = URL.createObjectURL(blob);

    downloadLink.href = url;
    downloadLink.download = `senbay-${Date.now()}.webm`;
    downloadLink.hidden = false;
    downloadLink.textContent = "Download Video";

    startRecordingButton.disabled = false;
    stopRecordingButton.disabled = true;
  });

  canvas.addEventListener("pointerdown", (event) => {
    const point = toCanvasPoint(canvas, event);

    if (!isInsideQr(point, qrRect)) return;

    isDraggingQr = true;
    dragOffset = {
      x: point.x - qrRect.x,
      y: point.y - qrRect.y,
    };

    canvas.classList.add("dragging");
    canvas.setPointerCapture(event.pointerId);
    event.preventDefault();
  });

  canvas.addEventListener("pointermove", (event) => {
    const point = toCanvasPoint(canvas, event);

    if (!isDraggingQr) {
      canvas.classList.toggle("qr-hover", isInsideQr(point, qrRect));
      return;
    }

    qrRect = clampQrRect(
      {
        ...qrRect,
        x: point.x - dragOffset.x,
        y: point.y - dragOffset.y,
      },
      canvas,
    );

    event.preventDefault();
  });

  canvas.addEventListener("pointerup", (event) => {
    isDraggingQr = false;
    canvas.classList.remove("dragging");
    canvas.releasePointerCapture(event.pointerId);
  });

  canvas.addEventListener("pointercancel", () => {
    isDraggingQr = false;
    canvas.classList.remove("dragging");
  });

  async function render(): Promise<void> {
    if (!cameraStarted) return;

    fitCanvasToVideo(canvas, video);

    qrRect = clampQrRect(qrRect, canvas);

    ctx.drawImage(video, 0, 0, canvas.width, canvas.height);

    const now = performance.now();

    if (now - lastQrUpdate > QR_UPDATE_INTERVAL_MS) {
      const payload = getSensorSnapshot();
      lastEncoded = encodeSenbay(payload, true);

      payloadView.textContent = JSON.stringify(payload, null, 2);
      encodedView.textContent = lastEncoded;

      lastQrUpdate = now;
    }

    if (lastEncoded) {
      await drawQrOverlay(ctx, lastEncoded, qrRect.x, qrRect.y, qrRect.size);
      drawQrHandle(ctx, qrRect);
    }

    if (isRecording) {
      const elapsedMs = Date.now() - recordingStartedAt;
      const elapsedText = formatElapsed(elapsedMs);

      recordingStatus.innerHTML = `
        <span class="recording-dot"></span>
        Recording ${elapsedText}
      `;

      drawRecordingOverlay(ctx, elapsedText);
    }

    requestAnimationFrame(render);
  }
}

function toCanvasPoint(canvas: HTMLCanvasElement, event: PointerEvent): Point {
  const rect = canvas.getBoundingClientRect();

  return {
    x: ((event.clientX - rect.left) / rect.width) * canvas.width,
    y: ((event.clientY - rect.top) / rect.height) * canvas.height,
  };
}

function isInsideQr(point: Point, qrRect: Rect): boolean {
  return (
    point.x >= qrRect.x &&
    point.x <= qrRect.x + qrRect.size &&
    point.y >= qrRect.y &&
    point.y <= qrRect.y + qrRect.size
  );
}

function clampQrRect(qrRect: Rect, canvas: HTMLCanvasElement): Rect {
  const maxX = Math.max(0, canvas.width - qrRect.size);
  const maxY = Math.max(0, canvas.height - qrRect.size);

  return {
    ...qrRect,
    x: Math.min(Math.max(qrRect.x, 0), maxX),
    y: Math.min(Math.max(qrRect.y, 0), maxY),
  };
}

function drawQrHandle(ctx: CanvasRenderingContext2D, qrRect: Rect): void {
  ctx.save();

  ctx.lineWidth = 4;
  ctx.strokeStyle = "rgba(255, 255, 255, 0.9)";
  ctx.strokeRect(qrRect.x, qrRect.y, qrRect.size, qrRect.size);

  ctx.fillStyle = "rgba(0, 0, 0, 0.65)";
  ctx.fillRect(qrRect.x, qrRect.y + qrRect.size - 32, qrRect.size, 32);

  ctx.fillStyle = "#ffffff";
  ctx.font = "bold 18px system-ui, sans-serif";
  ctx.textBaseline = "middle";
  ctx.fillText("Drag QR", qrRect.x + 12, qrRect.y + qrRect.size - 16);

  ctx.restore();
}

function drawRecordingOverlay(
  ctx: CanvasRenderingContext2D,
  elapsedText: string,
): void {
  const padding = 16;
  const width = 190;
  const height = 52;
  const x = ctx.canvas.width - width - padding;
  const y = padding;

  ctx.save();

  ctx.fillStyle = "rgba(0, 0, 0, 0.68)";
  roundRect(ctx, x, y, width, height, 18);
  ctx.fill();

  ctx.fillStyle = "#ff3434";
  ctx.beginPath();
  ctx.arc(x + 28, y + height / 2, 9, 0, Math.PI * 2);
  ctx.fill();

  ctx.fillStyle = "#ffffff";
  ctx.font = "bold 22px system-ui, sans-serif";
  ctx.textBaseline = "middle";
  ctx.fillText(`REC ${elapsedText}`, x + 48, y + height / 2);

  ctx.restore();
}

function roundRect(
  ctx: CanvasRenderingContext2D,
  x: number,
  y: number,
  width: number,
  height: number,
  radius: number,
): void {
  const r = Math.min(radius, width / 2, height / 2);

  ctx.beginPath();
  ctx.moveTo(x + r, y);
  ctx.arcTo(x + width, y, x + width, y + height, r);
  ctx.arcTo(x + width, y + height, x, y + height, r);
  ctx.arcTo(x, y + height, x, y, r);
  ctx.arcTo(x, y, x + width, y, r);
  ctx.closePath();
}

function formatElapsed(ms: number): string {
  const totalSeconds = Math.floor(ms / 1000);
  const minutes = Math.floor(totalSeconds / 60);
  const seconds = totalSeconds % 60;

  return `${String(minutes).padStart(2, "0")}:${String(seconds).padStart(2, "0")}`;
}
