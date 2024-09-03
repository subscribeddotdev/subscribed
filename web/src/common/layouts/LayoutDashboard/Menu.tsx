import { classnames } from "@@/common/libs/classnames";
import { RiCheckboxMultipleFill, RiGitRepositoryPrivateLine, RiHome5Line, RiStackLine } from "@remixicon/react";
import Link from "next/link";
import { useRouter } from "next/router";
import { PropsWithChildren, useMemo } from "react";
import styles from "./LayoutDashboard.module.css";

interface MenuItemProps extends PropsWithChildren {
  href: string;
  active?: boolean;
}

export function MenuItem({ children, href, active = false }: MenuItemProps) {
  return (
    <li className={classnames(styles.menuItem, { [styles.menuItemActive]: active })} data-active={active}>
      <Link className={styles.menuItemLink} href={href}>
        {children}
      </Link>
    </li>
  );
}

interface Props {}

function getMenuItems(environment: string) {
  const basePath = `/dashboard/${environment}`;

  return [
    {
      path: basePath,
      label: "Getting started",
      icon: <RiHome5Line size="18" />,
    },
    {
      path: `${basePath}/applications`,
      label: "Applications",
      icon: <RiCheckboxMultipleFill size="18" />,
    },
    {
      path: `${basePath}/event-types`,
      label: "Event types",
      icon: <RiStackLine size="18" />,
    },
    {
      path: `${basePath}/api-keys`,
      label: "API keys",
      icon: <RiGitRepositoryPrivateLine size="18" />,
    },
  ];
}

export function MenuList({}: Props) {
  const router = useRouter();
  const menuItems = useMemo(() => getMenuItems(router.query.environment as string), [router.query]);

  return (
    <ul className={styles.menu}>
      {menuItems.map((item, idx) => (
        <MenuItem active={router.asPath === item.path} key={idx} href={item.path}>
          {item.icon} {item.label}
        </MenuItem>
      ))}
    </ul>
  );
}
