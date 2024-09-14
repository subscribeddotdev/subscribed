import * as Label from "@radix-ui/react-label";
import { Flex, Text, TextField } from "@radix-ui/themes";
import { ComponentPropsWithoutRef, forwardRef } from "react";
import styles from "./Input.module.css";

interface Props extends ComponentPropsWithoutRef<"input"> {
  label?: string;
  error?: string;
  "data-testid"?: string;
}

type ButtonType =
  | "number"
  | "search"
  | "time"
  | "text"
  | "hidden"
  | "date"
  | "datetime-local"
  | "email"
  | "month"
  | "password"
  | "tel"
  | "url"
  | "week"
  | undefined;

export const Input = forwardRef<HTMLInputElement, Props>(
  function Input(props, ref) {
    const {
      onChange,
      onInput,
      error,
      id,
      name,
      placeholder,
      label,
      type,
      required,
    } = props;
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
          required={required}
          onChange={onChange}
          onInput={onInput}
          className={styles.input}
          placeholder={placeholder}
          type={type as ButtonType}
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
  },
);
