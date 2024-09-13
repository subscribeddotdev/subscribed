import { Link } from "@radix-ui/themes";
import { RiArrowRightSLine } from "@remixicon/react";
import { useMemo } from "react";
import { useLocation } from "react-router-dom";
import styles from "./Breadcrumbs.module.css";

interface Props {
  // Used to replace variables defined in the URL
  variables?: Record<string, string>;
  renderFromIndex?: number;
}

export function Breadcrumbs({ variables = {}, renderFromIndex = 0 }: Props) {
  const location = useLocation();
  const items = useMemo(() => {
    const paths: { label: string; path: string }[] = [];

    location.pathname
      .split("/")
      .filter((item) => item)
      .forEach((item, idx) => {
        paths.push({
          label: variables[item] || item,
          path: idx > 0 ? paths[idx - 1].path + "/" + item : "/" + item,
        });
      });

    return paths;
  }, [location, variables]);

  return (
    <ul className={styles.breadcrumb}>
      {items.map((item, idx) => {
        if (idx < renderFromIndex) {
          return null;
        }

        return (
          <li key={item.path} className={styles.breadcrumbItem}>
            <Link size="2" href={item.path} color="gray">
              {item.label}
            </Link>

            {idx !== items.length - 1 && (
              <RiArrowRightSLine size="14" className={styles.breadcrumbIcon} />
            )}
          </li>
        );
      })}
    </ul>
  );
}
