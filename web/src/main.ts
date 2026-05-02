import "./styles.css";
import { bootApp } from "./app";

bootApp().catch((error) => {
  console.error(error);
  alert(error instanceof Error ? error.message : String(error));
});
