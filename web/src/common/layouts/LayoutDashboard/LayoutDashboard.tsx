import { Breadcrumbs } from "@@/common/components/Breadcrumbs/Breadcrumbs";
import { EnvironmentSelector } from "@@/modules/EnvironmentSelector/EnvironmentSelector";
import { UserMenu } from "@@/modules/UserMenu/UserMenu";
import { Heading } from "@radix-ui/themes";
import { PropsWithChildren } from "react";
import styles from "./LayoutDashboard.module.css";
import { MenuList } from "./Menu";

interface Props extends PropsWithChildren {
  breadcrumbs?: {
    variables: Record<string, string>;
  };
}

export function LayoutDashboard({ children, breadcrumbs }: Props) {
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
        <div className={styles.content}>
          {!!breadcrumbs && (
            <Breadcrumbs
              renderFromIndex={2}
              variables={breadcrumbs.variables}
            />
          )}
          {children}
        </div>
      </main>
    </>
  );
}
