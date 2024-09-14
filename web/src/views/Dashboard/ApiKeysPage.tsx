import { apiClients } from "@@/common/libs/backendapi/browser";
import { ApiKey } from "@@/common/libs/backendapi/client";
import { ApiKeysProvider } from "@@/modules/ApiKeys/ApiKeysContext";
import ApiKeysPage from "@@/modules/ApiKeys/ApiKeysPage/ApiKeysPage";
import { Spinner } from "@radix-ui/themes";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

export default function ApiKeysHome() {
  const params = useParams();
  const [apiKeys, setApiKeys] = useState<ApiKey[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    (async () => {
      try {
        const { data } = await apiClients().ApiKeys.getAllApiKeys(
          params.environment as string,
        );
        setApiKeys(data.data);
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

  return (
    <ApiKeysProvider
      initialState={{ apiKeys, environmentId: params.environment as string }}
    >
      <ApiKeysPage />
    </ApiKeysProvider>
  );
}
