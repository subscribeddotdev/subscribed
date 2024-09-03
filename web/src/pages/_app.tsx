import "@@/common/styles/globals.css";
import { Theme } from "@radix-ui/themes";
import "@radix-ui/themes/styles.css";
import type { AppProps } from "next/app";

export default function App({ Component, pageProps }: AppProps) {
  return (
    <Theme accentColor="indigo" appearance="dark">
      <Component {...pageProps} />
    </Theme>
  );
}
