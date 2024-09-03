import { EnvironmentSelector } from "@@/modules/EnvironmentSelector/EnvironmentSelector";
import { UserMenu } from "@@/modules/UserMenu/UserMenu";
import { Heading } from "@radix-ui/themes";
import { PropsWithChildren } from "react";
import styles from "./LayoutDashboard.module.css";
import { MenuList } from "./Menu";

interface Props extends PropsWithChildren {}

export function LayoutDashboard({ children }: Props) {
  return (
    <>
      <aside className={styles.sidebar}>
        <Heading className={styles.sidebarHeading} size="4" mb="4">
          Subscribed
        </Heading>

        <EnvironmentSelector />

        <MenuList />
      </aside>

      <main className={styles.main}>
        <header className={styles.header}>
          <div></div>
          <UserMenu />
        </header>

        <div className={styles.content}>{children}</div>
      </main>
    </>
  );
}
