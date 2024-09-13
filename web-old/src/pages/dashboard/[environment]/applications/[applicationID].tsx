import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { PageTitle } from "@@/common/components/PageTitle/PageTitle";
import { LayoutDashboard } from "@@/common/layouts/LayoutDashboard/LayoutDashboard";
import { Application } from "@@/common/libs/backendapi/client";
import { createApiClientsSRR } from "@@/common/libs/backendapi/ssr";
import { dates } from "@@/common/libs/dates";
import { logger } from "@@/common/libs/logger";
import { Flex, Text } from "@radix-ui/themes";
import { GetServerSideProps } from "next";

interface Props {
  application: Application;
}

export default function ApplicationPage({ application }: Props) {
  return (
    <LayoutDashboard breadcrumbs={{ variables: { [application.id]: application.name } }}>
      <PageMeta title={application.name} />

      <Flex justify="between" mb="4">
        <PageTitle title={application.name} label="Applications" />
      </Flex>

      <Flex>
        <Flex direction="column">
          <Text color="gray" size="2">
            Created on
          </Text>
          <Text size="2">{dates(application.created_at).format("LL")}</Text>
        </Flex>
      </Flex>
    </LayoutDashboard>
  );
}

export const getServerSideProps: GetServerSideProps<Props> = async ({ req, params }) => {
  try {
    const apiClients = await createApiClientsSRR(req);
    const { data } = await apiClients.Applications.getApplicationById(params?.applicationID as string);

    return {
      props: { application: data.data },
    };
  } catch (error) {
    logger.error(error);
    return { redirect: { destination: "/500", permanent: false } };
  }
};
