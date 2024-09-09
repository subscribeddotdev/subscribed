import { Pagination } from "@@/common/libs/backendapi/client";
import { Flex, Text } from "@radix-ui/themes";
import { RiArrowLeftSLine, RiArrowRightSLine } from "@remixicon/react";
import { Button } from "../Button/Button";

interface Props {
  loading: boolean;
  pagination: Pagination;
  handler(page: number): void;
}

export function TablePaginationControl({ pagination, loading, handler }: Props) {
  return (
    <Flex mt="4" justify="between" align="center">
      <Text size="2" color="gray">
        Page {pagination.current_page} of {pagination.total_pages}
      </Text>

      <Flex gap="2">
        <Button
          size="1"
          color="gray"
          variant="outline"
          loading={loading}
          disabled={loading || pagination.current_page === 1}
          onClick={() => handler(pagination.current_page - 1)}
        >
          <RiArrowLeftSLine />
        </Button>
        <Button
          size="1"
          color="gray"
          variant="outline"
          loading={loading}
          onClick={() => handler(pagination.current_page + 1)}
          disabled={loading || pagination.current_page === pagination.total_pages}
        >
          <RiArrowRightSLine />
        </Button>
      </Flex>
    </Flex>
  );
}
