import { ApiKey } from "@@/common/libs/backendapi/client";
import { ApiKeysProvider } from "@@/modules/ApiKeys/ApiKeysContext";
import ApiKeysPage from "@@/modules/ApiKeys/ApiKeysPage/ApiKeysPage";

interface Props {
  apiKeys: ApiKey[];
  environmentId: string;
}

export default function ApiKeysHome({ apiKeys, environmentId }: Props) {
  return (
    <ApiKeysProvider initialState={{ apiKeys, environmentId }}>
      <ApiKeysPage />;
    </ApiKeysProvider>
  );
}
