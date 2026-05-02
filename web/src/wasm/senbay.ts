let readyPromise: Promise<void> | null = null;

export async function loadSenbayWasm(): Promise<void> {
  if (readyPromise) return readyPromise;

  readyPromise = (async () => {
    const go = new window.Go();
    const wasmUrl = `${import.meta.env.BASE_URL}senbay.wasm`;

    try {
      const result = await WebAssembly.instantiateStreaming(
        fetch(wasmUrl),
        go.importObject,
      );

      go.run(result.instance);
      return;
    } catch (error) {
      console.warn(
        "instantiateStreaming failed. Falling back to ArrayBuffer.",
        error,
      );
    }

    const response = await fetch(wasmUrl);
    const bytes = await response.arrayBuffer();
    const result = await WebAssembly.instantiate(bytes, go.importObject);

    go.run(result.instance);
  })();

  return readyPromise;
}

export function encodeSenbay(
  payload: Record<string, number | string | boolean | null | undefined>,
  compress = true,
): string {
  const cleanPayload: Record<string, number | string | boolean> = {};

  for (const [key, value] of Object.entries(payload)) {
    if (value === null || value === undefined) continue;
    cleanPayload[key] = value;
  }

  return window.senbayEncode(JSON.stringify(cleanPayload), compress);
}

export function decodeSenbay(text: string): Record<string, string> {
  return JSON.parse(window.senbayDecode(text));
}
