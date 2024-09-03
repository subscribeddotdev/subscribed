import { Environment } from "@@/common/libs/backendapi/client";
import { createApiClientsSRR } from "@@/common/libs/backendapi/ssr";
import { logger } from "@@/common/libs/logger";
import { DashboardSetup } from "@@/modules/DashboardSetup/DashboardSetup";
import { GetServerSideProps } from "next";

interface Props {
  environments: Environment[];
}

export default function DashboardSetupPage({ environments }: Props) {
  return <DashboardSetup environments={environments} />;
}

export const getServerSideProps: GetServerSideProps = async ({ req }) => {
  try {
    const clients = await createApiClientsSRR(req);
    const { data } = await clients.Environments.getEnvironments();

    return {
      props: {
        environments: data.data,
      },
    };
  } catch (error) {
    logger.error(error, "error rendering /dashboard page");
    return { redirect: { permanent: false, destination: "/500" } };
  }
};
