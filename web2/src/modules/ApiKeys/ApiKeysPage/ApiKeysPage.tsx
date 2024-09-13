import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { LayoutDashboard } from "@@/common/layouts/LayoutDashboard/LayoutDashboard";
import { apiClients } from "@@/common/libs/backendapi/browser";
import { Flex, Heading } from "@radix-ui/themes";
import { useCallback } from "react";
import { useApiKeysContext, useApiKeysDispatch } from "../ApiKeysContext";
import { CreateApiKey } from "../CreateApiKey/CreateApiKey";
import { ListApiKeys } from "../ListApiKeys/ListApiKeys";

export default function ApiKeysPage() {
  const state = useApiKeysContext();
  const dispatch = useApiKeysDispatch();

  const refetchAll = useCallback(async () => {
    const { data } = await apiClients().ApiKeys.getAllApiKeys(state.environmentId);
    dispatch({ type: "set", payload: data.data });
  }, [state.environmentId, dispatch]);

  return (
    <LayoutDashboard>
      <PageMeta title="API Keys" />
      <Flex justify="between" mb="4">
        <Heading>API Keys</Heading>
        <CreateApiKey onSuccess={refetchAll} />
      </Flex>

      <ListApiKeys apiKeys={state.apiKeys} />
    </LayoutDashboard>
  );
}
