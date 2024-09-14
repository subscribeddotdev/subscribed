import { PropsWithChildren } from "react";
import styles from "./LayoutAuth.module.css";

export function LayoutAuth({ children }: PropsWithChildren) {
  return <div className={styles.root}>{children}</div>;
}
