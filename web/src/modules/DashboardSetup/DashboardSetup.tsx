import { Environment } from "@@/common/libs/backendapi/client";
import { LAST_CHOSEN_ENVIRONMENT } from "@@/constants";
import { getPaths } from "@@/paths";
import { Spinner } from "@radix-ui/themes";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import styles from "./DashboardSetup.module.css";

interface Props {
  environments: Environment[];
}

export function DashboardSetup({ environments }: Props) {
  const navigate = useNavigate();

  useEffect(() => {
    const lastChosenEnvironment =
      localStorage.getItem(LAST_CHOSEN_ENVIRONMENT) || "";
    if (!lastChosenEnvironment) {
      localStorage.setItem(LAST_CHOSEN_ENVIRONMENT, environments[0].id);
    }

    navigate(
      getPaths(lastChosenEnvironment || environments[0].id).dashboardHomepage,
    );
  }, [navigate, environments]);

  return (
    <section className={styles.root}>
      <Spinner size="3" />
    </section>
  );
}
