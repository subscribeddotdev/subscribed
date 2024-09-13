import { PropsWithChildren } from "react";
import styles from "./LayoutAuth.module.css";

interface Props extends PropsWithChildren {}

export function LayoutAuth({ children }: Props) {
  return <div className={styles.root}>{children}</div>;
}
