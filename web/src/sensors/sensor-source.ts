export type SensorSnapshot = {
  TIME: number;

  LONG?: number;
  LATI?: number;
  ALTI?: number;
  SPEE?: number;

  ACCX?: number;
  ACCY?: number;
  ACCZ?: number;

  YAW?: number;
  ROLL?: number;
  PITC?: number;
};

let lastGeo: GeolocationCoordinates | null = null;
let lastMotion: DeviceMotionEvent | null = null;
let lastOrientation: DeviceOrientationEvent | null = null;

export async function enableSensors(): Promise<void> {
  if ("geolocation" in navigator) {
    navigator.geolocation.watchPosition(
      (position) => {
        lastGeo = position.coords;
      },
      (error) => {
        console.warn("Geolocation error:", error);
      },
      {
        enableHighAccuracy: true,
        maximumAge: 500,
        timeout: 10000,
      },
    );
  }

  const DeviceOrientation = window.DeviceOrientationEvent as any;
  if (
    DeviceOrientation &&
    typeof DeviceOrientation.requestPermission === "function"
  ) {
    const result = await DeviceOrientation.requestPermission();
    if (result !== "granted") {
      console.warn("DeviceOrientation permission was not granted");
    }
  }

  window.addEventListener("devicemotion", (event) => {
    lastMotion = event;
  });

  window.addEventListener("deviceorientation", (event) => {
    lastOrientation = event;
  });
}

export function getSensorSnapshot(): SensorSnapshot {
  const acceleration = lastMotion?.accelerationIncludingGravity;

  return {
    TIME: Date.now(),

    LONG: lastGeo?.longitude,
    LATI: lastGeo?.latitude,
    ALTI: lastGeo?.altitude ?? undefined,
    SPEE: lastGeo?.speed ?? undefined,

    ACCX: acceleration?.x ?? undefined,
    ACCY: acceleration?.y ?? undefined,
    ACCZ: acceleration?.z ?? undefined,

    YAW: lastOrientation?.alpha ?? undefined,
    PITC: lastOrientation?.beta ?? undefined,
    ROLL: lastOrientation?.gamma ?? undefined,
  };
}
