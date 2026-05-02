import QRCode from "qrcode";

const qrCanvas = document.createElement("canvas");

export async function drawQrOverlay(
  ctx: CanvasRenderingContext2D,
  text: string,
  x: number,
  y: number,
  size: number,
): Promise<void> {
  await QRCode.toCanvas(qrCanvas, text, {
    width: size,
    margin: 1,
    errorCorrectionLevel: "M",
  });

  ctx.drawImage(qrCanvas, x, y, size, size);
}
