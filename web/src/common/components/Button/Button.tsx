import { classnames } from "@@/common/libs/classnames";
import { ButtonProps, Button as RadixButton } from "@radix-ui/themes";
import { forwardRef } from "react";
import styles from "./Button.module.css";

export const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  function Button({ children, ...props }, ref) {
    const { className, ...otherProps } = props;

    return (
      <RadixButton
        className={classnames(styles.root, className)}
        {...otherProps}
        ref={ref}
      >
        {children}
      </RadixButton>
    );
  },
);
