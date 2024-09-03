import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { Heading, Text } from "@radix-ui/themes";
import styles from "./index.module.css";

export default function InternalServerErrorPage() {
  return (
    <>
      <PageMeta isHome title="Internal Server Error" />

      <main className={styles.main}>
        <div className={styles.container}>
          <Heading>500</Heading>
          <Text>Internal Server Error</Text>
        </div>
      </main>
    </>
  );
}
