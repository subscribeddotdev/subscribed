import { Alert } from "@@/common/components/Alert/Alert";
import { TablePaginationControl } from "@@/common/components/TablePaginationControl/TablePaginationControl";
import { apiClients, getApiError } from "@@/common/libs/backendapi/browser";
import { Application, Pagination } from "@@/common/libs/backendapi/client";
import { dates } from "@@/common/libs/dates";
import { usePaths } from "@@/paths";
import { Badge, Box, Table } from "@radix-ui/themes";
import { useCallback, useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import styles from "./ListApplications.module.css";

export function ListApplications() {
  const params = useParams();
  const [data, setData] = useState<Application[]>([]);
  const [pagination, setPagination] = useState<Pagination>();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const paths = usePaths();

  const paginationHandler = useCallback(
    async (page: number) => {
      setLoading(true);
      setError("");
      try {
        const resp = await apiClients().Applications.getApplications(
          params.environment as string,
          undefined,
          page
        );
        setData(resp.data.data);
        setPagination(resp.data.pagination);
      } catch (err) {
        setError(getApiError(err));
      } finally {
        setLoading(false);
      }
    },
    [params]
  );

  useEffect(() => {
    paginationHandler(1);
  }, [paginationHandler]);

  if (data.length === 0) {
    return (
      <Alert>No applications have been created for this environment.</Alert>
    );
  }

  return (
    <Box>
      {error && (
        <Alert mb="4" color="red">
          {error}
        </Alert>
      )}

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
                <Link
                  className={styles.link}
                  to={paths.helpers.toApplication(app.id)}
                >
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

      {pagination && (
        <TablePaginationControl
          loading={loading}
          pagination={pagination}
          handler={paginationHandler}
        />
      )}
    </Box>
  );
}
