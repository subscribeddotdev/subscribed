import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { PageTitle } from "@@/common/components/PageTitle/PageTitle";
import { LayoutDashboard } from "@@/common/layouts/LayoutDashboard/LayoutDashboard";
import { EventType, Pagination } from "@@/common/libs/backendapi/client";
import { createApiClientsSRR } from "@@/common/libs/backendapi/ssr";
import { logger } from "@@/common/libs/logger";
import { ListEventTypes } from "@@/modules/ListEventTypes/ListEventTypes";
import { Flex } from "@radix-ui/themes";

import { GetServerSideProps } from "next";

interface Props {
  eventTypes: EventType[];
  pagination: Pagination;
}

export default function EventTypesPage({ eventTypes, pagination }: Props) {
  return (
    <LayoutDashboard>
      <PageMeta title="Event types" />
      <Flex justify="between" mb="4">
        <PageTitle title="Event types" />
      </Flex>

      <ListEventTypes data={eventTypes} pagination={pagination} />
    </LayoutDashboard>
  );
}

export const getServerSideProps: GetServerSideProps<Props> = async ({ req, params }) => {
  try {
    const apiClients = await createApiClientsSRR(req);
    const { data } = await apiClients.EventTypes.getEventTypes();

    return {
      props: { eventTypes: data.data, pagination: data.pagination },
    };
  } catch (error) {
    logger.error(error);
    return { redirect: { destination: "/500", permanent: false } };
  }
};
