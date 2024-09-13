import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { CreateApplication } from "@@/modules/CreateApplication/CreateApplication";
import { ListApplications } from "@@/modules/ListApplications/ListApplications";
import { Flex, Heading } from "@radix-ui/themes";

export default function ApplicationsPage() {
  return (
    <>
      <PageMeta title="Applications" />

      <Flex justify="between" mb="4">
        <Heading>Applications</Heading>
        <CreateApplication />
      </Flex>

      <ListApplications />
    </>
  );
}
