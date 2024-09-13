import { paths } from "@@/constants";
import { Avatar, DropdownMenu, Flex, Text } from "@radix-ui/themes";
import { useRouter } from "next/router";
import { useCallback } from "react";
import { clearTokenFromCurrentSession, useUserDetails } from "../Auth/token";
import styles from "./UserMenu.module.css";

interface Props {}

export function UserMenu({}: Props) {
  const router = useRouter();
  const user = useUserDetails();

  const onLogout = useCallback(async () => {
    clearTokenFromCurrentSession();
    await router.push(paths.signin);
  }, [router]);

  if (!user.details) {
    return null;
  }

  return (
    <>
      <DropdownMenu.Root>
        <DropdownMenu.Trigger>
          <button className={styles.menuBtn}>
            <Avatar size="1" radius="full" src="" fallback={user.details?.firstName.charAt(0)} />
          </button>
        </DropdownMenu.Trigger>
        <DropdownMenu.Content>
          <Flex pl="3" pr="3" pt="2" pb="1" direction="column">
            <Text size="2" weight="bold">
              {user.details.firstName} {user.details.lastName}
            </Text>
            <Text size="2" color="gray">
              {user.details.email}
            </Text>
          </Flex>
          <DropdownMenu.Separator />
          <DropdownMenu.Item className={styles.menuItem} onClick={onLogout}>
            Log out
          </DropdownMenu.Item>
        </DropdownMenu.Content>
      </DropdownMenu.Root>
    </>
  );
}
