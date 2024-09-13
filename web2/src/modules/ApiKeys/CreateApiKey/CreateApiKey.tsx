import { Alert } from "@@/common/components/Alert/Alert";
import { Input } from "@@/common/components/Input/Input";
import {
  createApiClients,
  getApiError,
} from "@@/common/libs/backendapi/browser";
import { dates } from "@@/common/libs/dates";
import { Button, Code, Dialog, Flex } from "@radix-ui/themes";
import { useFormik } from "formik";
import { useCallback, useState } from "react";
import * as yup from "yup";
import { retrieveTokenFromTheClient } from "../../Auth/token";
import { useApiKeysContext } from "../ApiKeysContext";

interface Props {
  onSuccess(): Promise<void>;
}

export function CreateApiKey({ onSuccess }: Props) {
  const [error, setError] = useState("");
  const [open, setOpen] = useState(false);
  const [apiKey, setApiKey] = useState("");
  const state = useApiKeysContext();

  const onOpenChange = useCallback((isOpen: boolean) => {
    if (!isOpen) {
      setApiKey("");
    }

    setOpen(isOpen);
  }, []);

  const f = useFormik({
    initialValues,
    validateOnChange: false,
    validationSchema,
    onSubmit: async (values) => {
      try {
        setError("");
        const apiClient = createApiClients(retrieveTokenFromTheClient());
        const { data } = await apiClient.ApiKeys.createApiKey({
          name: values.name,
          environment_id: state.environmentId,
          expires_at: values.expires_at
            ? dates(values.expires_at).toISOString()
            : null,
        });

        setApiKey(data.unmasked_api_key);

        await onSuccess();
      } catch (error) {
        setError(getApiError(error));
      }
    },
  });

  return (
    <Dialog.Root open={open} onOpenChange={onOpenChange}>
      <Dialog.Trigger>
        <Button>Create api key</Button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth="450px">
        <form onSubmit={f.handleSubmit}>
          <Dialog.Title>New api key</Dialog.Title>

          <Flex direction="column" gap="2">
            <Input
              label="Name"
              required
              name="name"
              id="name"
              onChange={f.handleChange}
              error={f.errors.name}
            />
            <Input
              type="date"
              id="expires_at"
              label="Expires at"
              name="expires_at"
              onChange={f.handleChange}
              error={f.errors.expires_at}
            />
          </Flex>

          <Flex gap="3" mt="4" justify="end">
            <Dialog.Close>
              <Button variant="soft" color="gray">
                Cancel
              </Button>
            </Dialog.Close>
            <Button
              loading={f.isSubmitting}
              disabled={f.isSubmitting}
              type="submit"
            >
              Create
            </Button>
          </Flex>

          {error && (
            <Alert mt="4" color="red">
              {error}
            </Alert>
          )}
        </form>
      </Dialog.Content>

      {apiKey && (
        <Dialog.Content maxWidth="450px">
          <Dialog.Title>Api key successfully created</Dialog.Title>

          <Alert color="amber" mb="4">
            Please saved this api key before closing this window because we{" "}
            {"won't"} show it ever again.
          </Alert>
          <Code>{apiKey}</Code>

          <Flex gap="3" mt="4" justify="end">
            <Dialog.Close>
              <Button>Close</Button>
            </Dialog.Close>
          </Flex>
        </Dialog.Content>
      )}
    </Dialog.Root>
  );
}

function validationSchema() {
  return yup.object().shape({
    name: yup.string().required(),
    expires_at: yup.date().min(dates().add(1).toDate()).optional(),
  });
}

const initialValues = { name: "", expires_at: "" };
