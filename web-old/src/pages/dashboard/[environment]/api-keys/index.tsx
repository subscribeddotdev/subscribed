import { ApiKey } from "@@/common/libs/backendapi/client";
import { createApiClientsSRR } from "@@/common/libs/backendapi/ssr";
import { ApiKeysProvider } from "@@/modules/ApiKeys/ApiKeysContext";
import ApiKeysPage from "@@/modules/ApiKeys/ApiKeysPage/ApiKeysPage";
import { GetServerSideProps } from "next";

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

export const getServerSideProps: GetServerSideProps<Props> = async ({ req, params }) => {
  const clients = await createApiClientsSRR(req);
  const environmentId = params?.environment as string;

  try {
    const { data } = await clients.ApiKeys.getAllApiKeys(environmentId);

    return {
      props: { apiKeys: data.data, environmentId },
    };
  } catch (error) {
    return {
      props: { apiKeys: [], environmentId },
    };
  }
};
