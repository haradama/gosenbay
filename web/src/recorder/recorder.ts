export class CanvasRecorder {
  private recorder: MediaRecorder | null = null;
  private chunks: Blob[] = [];

  start(canvas: HTMLCanvasElement, fps = 30): void {
    this.chunks = [];

    const stream = canvas.captureStream(fps);
    const mimeType = pickMimeType();

    this.recorder = new MediaRecorder(
      stream,
      mimeType ? { mimeType } : undefined,
    );

    this.recorder.addEventListener("dataavailable", (event) => {
      if (event.data.size > 0) {
        this.chunks.push(event.data);
      }
    });

    this.recorder.start();
  }

  stop(): Promise<Blob> {
    return new Promise((resolve, reject) => {
      if (!this.recorder) {
        reject(new Error("Recorder is not started"));
        return;
      }

      this.recorder.addEventListener(
        "stop",
        () => {
          resolve(new Blob(this.chunks, { type: this.recorder?.mimeType }));
        },
        { once: true },
      );

      this.recorder.stop();
    });
  }
}

function pickMimeType(): string {
  const candidates = [
    "video/webm;codecs=vp9",
    "video/webm;codecs=vp8",
    "video/webm",
  ];

  if (typeof MediaRecorder === "undefined") return "";
  return candidates.find((type) => MediaRecorder.isTypeSupported(type)) ?? "";
}
