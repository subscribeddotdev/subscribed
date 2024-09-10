import { Link } from "@radix-ui/themes";
import { RiArrowRightSLine } from "@remixicon/react";
import { useRouter } from "next/router";
import { useMemo } from "react";
import styles from "./Breadcrumbs.module.css";

interface Props {
  // Used to replace variables defined in the URL
  variables?: Record<string, string>;
  renderFromIndex?: number;
}

export function Breadcrumbs({ variables = {}, renderFromIndex = 0 }: Props) {
  const router = useRouter();
  const items = useMemo(() => {
    const paths: { label: string; path: string }[] = [];

    router.asPath
      .split("/")
      .filter((item) => item)
      .forEach((item, idx) => {
        paths.push({
          label: variables[item] || item,
          path: idx > 0 ? paths[idx - 1].path + "/" + item : "/" + item,
        });
      });

    return paths;
  }, [router, variables]);

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

            {idx !== items.length - 1 && <RiArrowRightSLine size="14" className={styles.breadcrumbIcon} />}
          </li>
        );
      })}
    </ul>
  );
}
