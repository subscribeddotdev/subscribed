import { Alert } from "@@/common/components/Alert/Alert";
import { Application } from "@@/common/libs/backendapi/client";
import { dates } from "@@/common/libs/dates";
import { usePaths } from "@@/paths";
import { Badge, Table } from "@radix-ui/themes";
import Link from "next/link";
import styles from "./ListApplications.module.css";

interface Props {
  data: Application[];
}

export function ListApplications({ data }: Props) {
  const paths = usePaths();
  if (data.length === 0) {
    return <Alert>No applications have been created for this environment.</Alert>;
  }

  return (
    <Table.Root>
      <Table.Header>
        <Table.Row>
          <Table.ColumnHeaderCell>Name</Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell>ID</Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell>Created at</Table.ColumnHeaderCell>
        </Table.Row>
      </Table.Header>

      <Table.Body>
        {data.map((app) => (
          <Table.Row key={app.id}>
            <Table.RowHeaderCell>
              <Link className={styles.link} href={paths.helpers.toApplication(app.id)}>
                {app.name}
              </Link>
            </Table.RowHeaderCell>
            <Table.RowHeaderCell>
              <Badge color="gray">{app.id}</Badge>
            </Table.RowHeaderCell>
            <Table.Cell>{dates(app.created_at).format("LL")}</Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
    </Table.Root>
  );
}
