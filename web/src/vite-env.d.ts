export {};

declare global {
  interface Window {
    Go: any;
    senbayEncode: (payloadJson: string, compress: boolean) => string;
    senbayDecode: (senbayText: string) => string;
  }
}
