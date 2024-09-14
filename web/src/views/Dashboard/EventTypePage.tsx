import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { PageTitle } from "@@/common/components/PageTitle/PageTitle";
import { Flex } from "@radix-ui/themes";

export default function EventTypePage() {
  return (
    <>
      <PageMeta title="Event types" />
      <Flex justify="between" mb="4">
        <PageTitle title="Single" label="Event types" />
      </Flex>
    </>
  );
}
