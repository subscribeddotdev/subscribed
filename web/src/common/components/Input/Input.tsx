import * as Label from "@radix-ui/react-label";
import { Flex, Text, TextField } from "@radix-ui/themes";
import { ComponentPropsWithoutRef, forwardRef } from "react";
import styles from "./Input.module.css";

interface Props extends ComponentPropsWithoutRef<"input"> {
  label?: string;
  error?: string;
  "data-testid"?: string;
}

export const Input = forwardRef<HTMLInputElement, Props>(function Input(props, ref) {
  const { onChange, error, id, name, placeholder, label, type, required, ...rest } = props;
  return (
    <Flex direction="column" gap="1">
      {label && (
        <Label.Root htmlFor={name}>
          <Text weight="bold" size="2">
            {label}
          </Text>
          {required && (
            <Text color="red" weight="bold" size="2">
              *
            </Text>
          )}
        </Label.Root>
      )}
      <TextField.Root
        id={id}
        ref={ref}
        name={name}
        type={type as any}
        required={required}
        onChange={onChange}
        className={styles.input}
        placeholder={placeholder}
        data-testid={props["data-testid"]}
      >
        <TextField.Slot></TextField.Slot>
      </TextField.Root>
      {error && (
        <Text size="2" color="red">
          {error}
        </Text>
      )}
    </Flex>
  );
});
