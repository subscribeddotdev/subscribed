import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { LayoutDashboard } from "@@/common/layouts/LayoutDashboard/LayoutDashboard";
import { Application, Pagination } from "@@/common/libs/backendapi/client";
import { createApiClientsSRR } from "@@/common/libs/backendapi/ssr";
import { logger } from "@@/common/libs/logger";
import { CreateApplication } from "@@/modules/CreateApplication/CreateApplication";
import { ListApplications } from "@@/modules/ListApplications/ListApplications";
import { Flex, Heading } from "@radix-ui/themes";
import { GetServerSideProps } from "next";

interface Props {
  pagination: Pagination;
  applications: Application[];
}

export default function ApplicationsPage({ applications }: Props) {
  return (
    <LayoutDashboard>
      <PageMeta title="Applications" />

      <Flex justify="between" mb="4">
        <Heading>Applications</Heading>
        <CreateApplication />
      </Flex>

      <ListApplications data={applications} />
    </LayoutDashboard>
  );
}

export const getServerSideProps: GetServerSideProps<Props> = async ({ req, params }) => {
  const clients = await createApiClientsSRR(req);
  const environmentId = params?.environment as string;

  try {
    const { data } = await clients.Applications.getApplications(environmentId);
    return {
      props: { applications: data.data, pagination: data.pagination },
    };
  } catch (error) {
    logger.error(error);
    return {
      props: {
        applications: [],
        pagination: {
          total: 0,
          per_page: 0,
          total_pages: 0,
          current_page: 0,
        },
      },
    };
  }
};
