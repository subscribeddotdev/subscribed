import { AlertDialog, Button, Flex } from "@radix-ui/themes";
import { ReactNode, useCallback, useState } from "react";

interface Props {
  title: string;
  open: boolean;
  description?: ReactNode;
  onConfirm(): Promise<void>;
  onOpenChange(status: boolean): void;
}

export function ConfirmPrompt({ onOpenChange, onConfirm, open, title, description }: Props) {
  const [doing, setDoing] = useState(false);
  const onConfirmHandler = useCallback(async () => {
    setDoing(true);
    await onConfirm();
    onOpenChange(false);
    setDoing(false);
  }, [onConfirm, onOpenChange]);

  return (
    <AlertDialog.Root open={open} onOpenChange={onOpenChange}>
      <AlertDialog.Content maxWidth="450px">
        <AlertDialog.Title>{title}</AlertDialog.Title>
        <AlertDialog.Description size="2">{description}</AlertDialog.Description>

        <Flex gap="3" mt="4" justify="end">
          <AlertDialog.Cancel>
            <Button variant="soft" color="gray" disabled={doing}>
              Cancel
            </Button>
          </AlertDialog.Cancel>
          <Button loading={doing} disabled={doing} variant="solid" color="red" onClick={onConfirmHandler}>
            Confirm
          </Button>
        </Flex>
      </AlertDialog.Content>
    </AlertDialog.Root>
  );
}
