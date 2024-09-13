import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { PageTitle } from "@@/common/components/PageTitle/PageTitle";
import { ListEventTypes } from "@@/modules/ListEventTypes/ListEventTypes";
import { Flex } from "@radix-ui/themes";

export default function EventTypesPage() {
  return (
    <>
      <PageMeta title="Event types" />
      <Flex justify="between" mb="4">
        <PageTitle title="Event types" />
      </Flex>

      <ListEventTypes />
    </>
  );
}
