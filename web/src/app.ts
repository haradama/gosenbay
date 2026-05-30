import {
  startCamera,
  fitCanvasToDisplayOrientation,
  drawVideoCover,
} from "./camera/camera";
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

type CanvasSize = {
  width: number;
  height: number;
};

const QR_SIZE = 180;
const QR_MARGIN = 16;
const QR_UPDATE_INTERVAL_MS = 100;

export async function bootApp(): Promise<void> {
  await loadSenbayWasm();

  const video = requireElement<HTMLVideoElement>("#cameraVideo");
  const canvas = requireElement<HTMLCanvasElement>("#outputCanvas");
  const payloadView = requireElement<HTMLPreElement>("#payloadView");
  const encodedView = requireElement<HTMLPreElement>("#encodedView");

  const startCameraButton = requireElement<HTMLButtonElement>("#startCamera");
  const enableSensorsButton =
    requireElement<HTMLButtonElement>("#enableSensors");
  const startRecordingButton =
    requireElement<HTMLButtonElement>("#startRecording");
  const stopRecordingButton =
    requireElement<HTMLButtonElement>("#stopRecording");
  const downloadLink = requireElement<HTMLAnchorElement>("#downloadLink");
  const recordingStatus = requireElement<HTMLSpanElement>("#recordingStatus");
  const recordingElapsed = requireElement<HTMLSpanElement>("#recordingElapsed");
  const qrMoveHint = requireElement<HTMLDivElement>("#qrMoveHint");

  const drawCtx = require2dContext(canvas);

  const recorder = new CanvasRecorder();

  let cameraStarted = false;
  let isRecording = false;
  let recordingStartedAt = 0;

  let lastQrUpdate = 0;
  let lastEncoded = "";
  let lastDownloadUrl: string | null = null;

  let hasCanvasSizeInitialized = false;

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
    startCameraButton.disabled = true;
    render();
  });

  enableSensorsButton.addEventListener("click", async () => {
    await enableSensors();
  });

  startRecordingButton.addEventListener("click", () => {
    recorder.start(canvas, 30);

    isRecording = true;
    recordingStartedAt = Date.now();

    recordingElapsed.textContent = "00:00";
    recordingStatus.hidden = false;

    startRecordingButton.disabled = true;
    stopRecordingButton.disabled = false;
    downloadLink.hidden = true;
  });

  stopRecordingButton.addEventListener("click", async () => {
    isRecording = false;
    recordingStatus.hidden = true;

    const blob = await recorder.stop();

    if (lastDownloadUrl) URL.revokeObjectURL(lastDownloadUrl);
    lastDownloadUrl = URL.createObjectURL(blob);

    downloadLink.href = lastDownloadUrl;
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

    qrMoveHint.hidden = true;

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

    const previousCanvasSize: CanvasSize = {
      width: canvas.width,
      height: canvas.height,
    };

    const canvasResized = fitCanvasToDisplayOrientation(canvas);

    if (canvasResized && hasCanvasSizeInitialized) {
      qrRect = scaleQrRectForCanvasChange(qrRect, previousCanvasSize, canvas);
    }

    hasCanvasSizeInitialized = true;
    qrRect = clampQrRect(qrRect, canvas);

    drawVideoCover(drawCtx, video, canvas);

    const now = performance.now();

    if (now - lastQrUpdate > QR_UPDATE_INTERVAL_MS) {
      const payload = getSensorSnapshot();
      lastEncoded = encodeSenbay(payload, true);

      payloadView.textContent = JSON.stringify(payload, null, 2);
      encodedView.textContent = lastEncoded;

      lastQrUpdate = now;
    }

    if (lastEncoded) {
      await drawQrOverlay(
        drawCtx,
        lastEncoded,
        qrRect.x,
        qrRect.y,
        qrRect.size,
      );
    }

    if (isRecording) {
      recordingElapsed.textContent = formatElapsed(
        Date.now() - recordingStartedAt,
      );
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

function scaleQrRectForCanvasChange(
  qrRect: Rect,
  previousCanvasSize: CanvasSize,
  canvas: HTMLCanvasElement,
): Rect {
  if (previousCanvasSize.width <= 0 || previousCanvasSize.height <= 0) {
    return qrRect;
  }

  return {
    ...qrRect,
    x: qrRect.x * (canvas.width / previousCanvasSize.width),
    y: qrRect.y * (canvas.height / previousCanvasSize.height),
  };
}

function requireElement<T extends Element>(selector: string): T {
  const element = document.querySelector<T>(selector);
  if (!element) {
    throw new Error(`Required DOM element not found: ${selector}`);
  }
  return element;
}

function require2dContext(canvas: HTMLCanvasElement): CanvasRenderingContext2D {
  const ctx = canvas.getContext("2d");
  if (!ctx) throw new Error("Canvas 2D context is not available");
  return ctx;
}

function formatElapsed(ms: number): string {
  const totalSeconds = Math.floor(ms / 1000);
  const minutes = Math.floor(totalSeconds / 60);
  const seconds = totalSeconds % 60;

  return `${String(minutes).padStart(2, "0")}:${String(seconds).padStart(2, "0")}`;
}
