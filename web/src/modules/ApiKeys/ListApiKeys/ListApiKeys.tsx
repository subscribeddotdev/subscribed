import { Alert } from "@@/common/components/Alert/Alert";
import { ConfirmPrompt } from "@@/common/components/ConfirmPrompt/ConfirmPrompt";
import { apiClients, getApiError } from "@@/common/libs/backendapi/browser";
import { ApiKey } from "@@/common/libs/backendapi/client";
import { dates } from "@@/common/libs/dates";
import { Badge, Button, DropdownMenu, Table } from "@radix-ui/themes";
import { RiMoreFill } from "@remixicon/react";
import { useCallback, useState } from "react";
import { useApiKeysDispatch } from "../ApiKeysContext";

interface Props {
  apiKeys: ApiKey[];
}

export function ListApiKeys({ apiKeys }: Props) {
  const dispatch = useApiKeysDispatch();
  const [error, setError] = useState("");
  const [showConfirm, setShowConfirm] = useState(false);
  const [apiKeyToRemove, setApiKeyToRemove] = useState<ApiKey | null>(null);

  const destroyApiKey = useCallback(
    async (id: string) => {
      try {
        await apiClients().ApiKeys.destroyApiKey(id);
        dispatch({ type: "remove", payload: id });
      } catch (err) {
        setError(getApiError(err));
      }
    },
    [dispatch],
  );

  if (apiKeys.length === 0) {
    return <Alert>No api keys have been created for this environment.</Alert>;
  }

  return (
    <>
      {error && (
        <Alert onClose={() => setError("")} mb="2" color="red">
          {error}
        </Alert>
      )}
      <Table.Root>
        <Table.Header>
          <Table.Row>
            <Table.ColumnHeaderCell>Name</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell>Key</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell>Created</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell>Expires</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell></Table.ColumnHeaderCell>
          </Table.Row>
        </Table.Header>

        <Table.Body>
          {apiKeys.map((apiKey) => (
            <Table.Row key={apiKey.id} data-testid={`ApiKeyRow_${apiKey.id}`}>
              <Table.Cell>{apiKey.name}</Table.Cell>
              <Table.Cell>
                <Badge color="gray">{apiKey.masked_secret_key}</Badge>
              </Table.Cell>
              <Table.Cell>{dates(apiKey.created_at).fromNow()}</Table.Cell>
              <Table.Cell>{apiKey.expires_at ? dates(apiKey.expires_at).fromNow() : "-"}</Table.Cell>
              <Table.Cell>
                <DropdownMenu.Root>
                  <DropdownMenu.Trigger>
                    <Button variant="ghost" size="1" color="gray">
                      <RiMoreFill size="16px" />
                    </Button>
                  </DropdownMenu.Trigger>
                  <DropdownMenu.Content>
                    <DropdownMenu.Item
                      color="red"
                      onClick={() => {
                        setError("");
                        setApiKeyToRemove(apiKey);
                        setShowConfirm(true);
                      }}
                      data-testid={`ApiKeyRow_${apiKey.id}_DestroyBtn`}
                    >
                      Destroy
                    </DropdownMenu.Item>
                  </DropdownMenu.Content>
                </DropdownMenu.Root>
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table.Root>

      <ConfirmPrompt
        title="Destroy api key"
        description={`Are you sure you want to delete the ${apiKeyToRemove?.name} api key? This can not be undone.`}
        open={showConfirm}
        onConfirm={() => destroyApiKey(apiKeyToRemove?.id as string)}
        onOpenChange={setShowConfirm}
      />
    </>
  );
}
