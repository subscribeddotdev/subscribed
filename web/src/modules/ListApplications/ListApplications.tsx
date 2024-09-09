import { Alert } from "@@/common/components/Alert/Alert";
import { TablePaginationControl } from "@@/common/components/TablePaginationControl/TablePaginationControl";
import { apiClients } from "@@/common/libs/backendapi/browser";
import { Application, Pagination } from "@@/common/libs/backendapi/client";
import { dates } from "@@/common/libs/dates";
import { logger } from "@@/common/libs/logger";
import { usePaths } from "@@/paths";
import { Badge, Box, Table } from "@radix-ui/themes";
import Link from "next/link";
import { useRouter } from "next/router";
import { useCallback, useState } from "react";
import styles from "./ListApplications.module.css";

interface Props {
  data: Application[];
  pagination: Pagination;
}

export function ListApplications({ data: initialData, pagination: initialPagination }: Props) {
  const router = useRouter();
  const [data, setData] = useState(initialData);
  const [pagination, setPagination] = useState(initialPagination);
  const [loading, setLoading] = useState(false);

  const paginationHandler = useCallback(
    async (page: number) => {
      setLoading(true);
      try {
        const resp = await apiClients().Applications.getApplications(
          router.query.environment as string,
          undefined,
          page
        );
        setData(resp.data.data);
        setPagination(resp.data.pagination);
      } catch (error) {
        logger.error(error);
      } finally {
        setLoading(false);
      }
    },
    [router]
  );

  const paths = usePaths();
  if (data.length === 0) {
    return <Alert>No applications have been created for this environment.</Alert>;
  }

  return (
    <Box>
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

      <TablePaginationControl pagination={pagination} loading={loading} handler={paginationHandler} />
    </Box>
  );
}
