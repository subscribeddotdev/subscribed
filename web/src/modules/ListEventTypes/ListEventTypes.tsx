import { Alert } from "@@/common/components/Alert/Alert";
import { TablePaginationControl } from "@@/common/components/TablePaginationControl/TablePaginationControl";
import { apiClients, getApiError } from "@@/common/libs/backendapi/browser";
import { EventType, Pagination } from "@@/common/libs/backendapi/client";
import { dates } from "@@/common/libs/dates";
import { usePaths } from "@@/paths";
import { Box, Table } from "@radix-ui/themes";
import { useCallback, useEffect, useState } from "react";
import { Link } from "react-router-dom";
import styles from "./ListEventTypes.module.css";

export function ListEventTypes() {
  const [data, setData] = useState<EventType[]>([]);
  const [pagination, setPagination] = useState<Pagination>();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const paths = usePaths();

  const paginationHandler = useCallback(async (page: number) => {
    setLoading(true);
    setError("");
    try {
      const resp = await apiClients().EventTypes.getEventTypes(undefined, page);
      setData(resp.data.data);
      setPagination(resp.data.pagination);
    } catch (err) {
      setError(getApiError(err));
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    paginationHandler(1);
  }, [paginationHandler]);

  if (data.length === 0) {
    return <Alert>No event types have been created yet.</Alert>;
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
            <Table.ColumnHeaderCell>Description</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell>Created at</Table.ColumnHeaderCell>
          </Table.Row>
        </Table.Header>

        <Table.Body>
          {data.map((eventType) => (
            <Table.Row key={eventType.id}>
              <Table.RowHeaderCell>
                <Link
                  className={styles.link}
                  to={paths.helpers.toEventType(eventType.id)}
                >
                  {eventType.name}
                </Link>
              </Table.RowHeaderCell>
              <Table.RowHeaderCell>
                {eventType.description.length > 70
                  ? `${eventType.description.substring(0, 67).trimEnd()}...`
                  : eventType.description}
              </Table.RowHeaderCell>
              <Table.Cell>
                {dates(eventType.created_at).format("LL")}
              </Table.Cell>
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
