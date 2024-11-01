import { Alert } from "@@/common/components/Alert/Alert";
import { PageLoading } from "@@/common/components/PageLoading/PageLoading";
import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { PageTitle } from "@@/common/components/PageTitle/PageTitle";
import { apiClients, getApiError } from "@@/common/libs/backendapi/browser";
import { EventType } from "@@/common/libs/backendapi/client";
import { dates } from "@@/common/libs/dates";
import { Flex, Text } from "@radix-ui/themes";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

export default function EventTypePage() {
  const { eventType, loading, error } = useEventType();

  if (loading) {
    return <PageLoading />;
  }

  if (error) {
    return <Alert color="red">{error}</Alert>;
  }

  return (
    <>
      <PageMeta title="Event types" />
      <Flex justify="between" mb="4">
        <PageTitle title={eventType?.name || ""} label="Event types" />
        {/* <DropdownMenu.Root>
          <DropdownMenu.Trigger>
            <Button variant="ghost" size="1" color="gray">
              <RiMoreFill size="16px" />
            </Button>
          </DropdownMenu.Trigger>
          <DropdownMenu.Content>
            <DropdownMenu.Item onClick={() => {}}>Archive</DropdownMenu.Item>
          </DropdownMenu.Content>
        </DropdownMenu.Root> */}
      </Flex>

      <Flex direction="column" mb="4">
        <Text color="gray" size="2">
          Description
        </Text>
        <Text size="2">{eventType?.description}</Text>
      </Flex>

      <Flex direction="column" mb="4">
        <Text color="gray" size="2">
          Schema
        </Text>
        <Text size="2">{eventType?.schema}</Text>
      </Flex>

      <Flex>
        <Flex direction="column">
          <Text color="gray" size="2">
            Created on
          </Text>
          <Text size="2">{dates(eventType?.created_at).format("LL")}</Text>
        </Flex>
      </Flex>
    </>
  );
}

function useEventType() {
  const params = useParams();
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(true);
  const [eventType, setEventType] = useState<EventType>();

  useEffect(() => {
    (async () => {
      try {
        const { data } = await apiClients().EventTypes.getEventTypeById(
          params.eventTypeId as string
        );
        setEventType(data.data);
      } catch (error) {
        setError(getApiError(error));
      } finally {
        setLoading(false);
      }
    })();
  }, [params]);

  return { loading, error, eventType };
}
