import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { Heading, Link, Text } from "@radix-ui/themes";
import styles from "./index.module.css";

export default function Home() {
  return (
    <>
      <PageMeta isHome title="Subscribed" />

      <main className={styles.main}>
        <div className={styles.container}>
          <Heading>Subscribed</Heading>
          <Text>
            In early development, please head to our <Link href="https://github.com/subscribeddotdev">GitHub</Link>.
          </Text>
        </div>
      </main>
    </>
  );
}
