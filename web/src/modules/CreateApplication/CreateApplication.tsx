import { config } from "@@/config";
import { Button, Code, Dialog, Flex, Text, TextField } from "@radix-ui/themes";
import { useState } from "react";

export function CreateApplication() {
  const [name, setName] = useState("");

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <Button>Create application</Button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth="450px">
        <Dialog.Title>New application</Dialog.Title>
        <Dialog.Description size="2" mb="4">
          Applications can only be created via our API, please copy the code
          snippet below to your terminal and specify your a valid API key in the
          header <Code>x-api-key</Code>.
        </Dialog.Description>

        <Flex direction="column" gap="3" mb="4">
          <label>
            <Text as="div" size="2" mb="1" weight="bold">
              Name
            </Text>
            <TextField.Root
              placeholder="My new app"
              onInput={(e) => setName((e.target as HTMLInputElement).value)}
            />
          </label>
        </Flex>

        <Flex direction="column" gap="3">
          <Code wrap="pretty" size="3">
            curl -X POST \<br />
            -H {`"x-api-key: ******"`} \<br />
            -H {`"Content-Type: application/json"`} \<br />
            -d {`'{"name": "${name}"}'`} \<br />
            {config.public.api}/applications
          </Code>
        </Flex>

        <Flex gap="3" mt="4" justify="end">
          <Dialog.Close>
            <Button
              onClick={() =>
                navigator.clipboard.writeText(
                  codeSnippet(name, config.public.api),
                )
              }
            >
              Copy
            </Button>
          </Dialog.Close>
        </Flex>
      </Dialog.Content>
    </Dialog.Root>
  );
}

function codeSnippet(appName: string, apiURL: string) {
  return `curl -X POST \\
-H "x-api-key: ******" \\
-H "Content-Type: application/json" \\
-d '{"name": "${appName}"}' \\
${apiURL}/applications
`;
}
