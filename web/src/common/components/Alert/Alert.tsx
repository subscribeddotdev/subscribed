import { Button, Callout } from "@radix-ui/themes";
import { RiCloseLine } from "@remixicon/react";
import { PropsWithChildren, ReactNode } from "react";
import styles from "./Alert.module.css";

interface Props extends PropsWithChildren {
  onClose?(): void;
  Icon?: ReactNode;
  size?: "1" | "2" | "3";
  mb?: "1" | "2" | "3" | "4" | "5";
  mt?: "1" | "2" | "3" | "4" | "5";
  variant?: "soft" | "surface" | "outline";
  color?: "red" | "amber" | "gray" | "green";
}

export function Alert({ mb, mt, children, Icon, size, variant = "soft", color = "gray", onClose }: Props) {
  return (
    <Callout.Root className={styles.root} mb={mb} size={size} variant={variant} color={color} mt={mt}>
      {Icon && <Callout.Icon>{Icon}</Callout.Icon>}
      <Callout.Text className={styles.text}>
        {children}
        {onClose && (
          <Button variant="soft" onClick={onClose} className={styles.closeButton}>
            <RiCloseLine />
          </Button>
        )}
      </Callout.Text>
    </Callout.Root>
  );
}
