import * as Label from "@radix-ui/react-label";
import { Flex, Text, TextArea } from "@radix-ui/themes";
import { ComponentPropsWithoutRef, forwardRef } from "react";
import styles from "./InputArea.module.css";

interface Props extends ComponentPropsWithoutRef<"textarea"> {
  label?: string;
  error?: string;
  "data-testid"?: string;
}

export const InputArea = forwardRef<HTMLTextAreaElement, Props>(
  function InputArea(props, ref) {
    const { onChange, error, id, name, placeholder, label, required } = props;
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
        <TextArea
          id={id}
          ref={ref}
          name={name}
          required={required}
          onChange={onChange}
          className={styles.input}
          placeholder={placeholder}
          data-testid={props["data-testid"]}
        />
        {error && (
          <Text size="2" color="red">
            {error}
          </Text>
        )}
      </Flex>
    );
  },
);
