export async function startCamera(
  video: HTMLVideoElement,
): Promise<MediaStream> {
  const isPortrait = isPortraitOrientation();

  const stream = await navigator.mediaDevices.getUserMedia({
    video: {
      facingMode: { ideal: "environment" },
      width: { ideal: isPortrait ? 720 : 1280 },
      height: { ideal: isPortrait ? 1280 : 720 },
    },
    audio: false,
  });

  video.srcObject = stream;
  await video.play();

  return stream;
}

export function fitCanvasToDisplayOrientation(
  canvas: HTMLCanvasElement,
): boolean {
  const isPortrait = isPortraitOrientation();

  const width = isPortrait ? 720 : 1280;
  const height = isPortrait ? 1280 : 720;

  const changed = canvas.width !== width || canvas.height !== height;

  if (changed) {
    canvas.width = width;
    canvas.height = height;
  }

  return changed;
}

export function drawVideoCover(
  ctx: CanvasRenderingContext2D,
  video: HTMLVideoElement,
  canvas: HTMLCanvasElement,
): void {
  const sourceWidth = video.videoWidth || canvas.width;
  const sourceHeight = video.videoHeight || canvas.height;

  const scale = Math.max(
    canvas.width / sourceWidth,
    canvas.height / sourceHeight,
  );

  const drawWidth = sourceWidth * scale;
  const drawHeight = sourceHeight * scale;

  const dx = (canvas.width - drawWidth) / 2;
  const dy = (canvas.height - drawHeight) / 2;

  ctx.drawImage(video, dx, dy, drawWidth, drawHeight);
}

function isPortraitOrientation(): boolean {
  const orientationType = screen.orientation?.type;

  if (orientationType?.startsWith("portrait")) return true;
  if (orientationType?.startsWith("landscape")) return false;

  if (window.matchMedia) {
    return window.matchMedia("(orientation: portrait)").matches;
  }

  return window.innerHeight >= window.innerWidth;
}
