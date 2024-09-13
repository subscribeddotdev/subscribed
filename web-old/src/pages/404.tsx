import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { Heading, Text } from "@radix-ui/themes";
import styles from "./index.module.css";

export default function InternalServerErrorPage() {
  return (
    <>
      <PageMeta isHome title="Page Not Found" />

      <main className={styles.main}>
        <div className={styles.container}>
          <Heading>404</Heading>
          <Text>Page Not Found</Text>
        </div>
      </main>
    </>
  );
}
