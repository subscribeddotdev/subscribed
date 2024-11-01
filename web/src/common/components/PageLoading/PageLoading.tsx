import { Box, Spinner } from "@radix-ui/themes";
import styles from "./PageLoading.module.css";

export function PageLoading() {
  return (
    <Box className={styles.root}>
      <Spinner />
    </Box>
  );
}
