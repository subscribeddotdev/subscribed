import { classnames } from "@@/common/libs/classnames";
import {
  RiCheckboxMultipleFill,
  RiGitRepositoryPrivateLine,
  RiHome5Line,
  RiStackLine,
} from "@remixicon/react";
import { PropsWithChildren, useMemo } from "react";
import { Link, useLocation, useParams } from "react-router-dom";
import styles from "./LayoutDashboard.module.css";

interface MenuItemProps extends PropsWithChildren {
  href: string;
  active?: boolean;
}

export function MenuItem({ children, href, active = false }: MenuItemProps) {
  return (
    <li
      className={classnames(styles.menuItem, {
        [styles.menuItemActive]: active,
      })}
      data-active={active}
    >
      <Link className={styles.menuItemLink} to={href}>
        {children}
      </Link>
    </li>
  );
}

function getMenuItems(environment: string) {
  const basePath = `${import.meta.env.BASE_URL}/${environment}`;

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

export function MenuList() {
  const params = useParams();
  const location = useLocation();
  const menuItems = useMemo(
    () => getMenuItems(params.environment as string),
    [params]
  );

  return (
    <ul className={styles.menu}>
      {menuItems.map((item, idx) => (
        <MenuItem
          active={location.pathname === item.path}
          key={idx}
          href={item.path}
        >
          {item.icon} {item.label}
        </MenuItem>
      ))}
    </ul>
  );
}
