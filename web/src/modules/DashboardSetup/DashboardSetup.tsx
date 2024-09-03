import { Environment } from "@@/common/libs/backendapi/client";
import { LAST_CHOSEN_ENVIRONMENT } from "@@/constants";
import { getPaths } from "@@/paths";
import { Spinner } from "@radix-ui/themes";
import { useRouter } from "next/router";
import { useEffect } from "react";
import styles from "./DashboardSetup.module.css";

interface Props {
  environments: Environment[];
}

export function DashboardSetup({ environments }: Props) {
  const router = useRouter();

  useEffect(() => {
    const lastChosenEnvironment = localStorage.getItem(LAST_CHOSEN_ENVIRONMENT) || "";
    if (!lastChosenEnvironment) {
      localStorage.setItem(LAST_CHOSEN_ENVIRONMENT, environments[0].id);
    }

    router.push(getPaths(lastChosenEnvironment || environments[0].id).dashboardHomepage);
  }, [router, environments]);

  return (
    <section className={styles.root}>
      <Spinner size="3" />
    </section>
  );
}
