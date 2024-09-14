import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { PageTitle } from "@@/common/components/PageTitle/PageTitle";

import { ListEventTypes } from "@@/modules/ListEventTypes/ListEventTypes";
import { usePaths } from "@@/paths";
import { Button, Flex } from "@radix-ui/themes";
import { useNavigate } from "react-router-dom";

export default function EventTypesPage() {
  const navigate = useNavigate();
  const paths = usePaths();
  return (
    <>
      <PageMeta title="Event types" />
      <Flex justify="between" mb="4">
        <PageTitle title="Event types" />
        <Button onClick={() => navigate(paths.urls.createEventType)}>
          Create event type
        </Button>
      </Flex>

      <ListEventTypes />
    </>
  );
}
