import QRCode from "qrcode";

const qrCanvas = document.createElement("canvas");

let lastRenderedText = "";
let lastRenderedSize = 0;

export async function drawQrOverlay(
  ctx: CanvasRenderingContext2D,
  text: string,
  x: number,
  y: number,
  size: number,
): Promise<void> {
  if (text !== lastRenderedText || size !== lastRenderedSize) {
    await QRCode.toCanvas(qrCanvas, text, {
      width: size,
      margin: 1,
      errorCorrectionLevel: "M",
    });

    lastRenderedText = text;
    lastRenderedSize = size;
  }

  ctx.drawImage(qrCanvas, x, y, size, size);
}
