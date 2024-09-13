import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { PageTitle } from "@@/common/components/PageTitle/PageTitle";
import { apiClients } from "@@/common/libs/backendapi/browser";
import { Application } from "@@/common/libs/backendapi/client";
import { dates } from "@@/common/libs/dates";
import { Flex, Spinner, Text } from "@radix-ui/themes";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

export default function ApplicationPage() {
  const [app, setApp] = useState<Application>();
  const [loading, setLoading] = useState(true);
  const params = useParams();

  useEffect(() => {
    (async () => {
      try {
        const { data } = await apiClients().Applications.getApplicationById(
          params.appId || "",
        );

        setApp(data.data);
      } catch (error) {
        console.log(error);
      } finally {
        setLoading(false);
      }
    })();
  }, [params]);

  if (loading) {
    return <Spinner />;
  }

  const application = app as Application;

  return (
    <>
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
    </>
  );
}
